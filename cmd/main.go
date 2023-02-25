package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/common-fate/clio"
	"golang.org/x/sync/errgroup"

	"github.com/a-h/generate"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRetryMode(aws.RetryModeAdaptive), config.WithRetryMaxAttempts(20))
	if err != nil {
		return err
	}

	cfn := cloudformation.NewFromConfig(cfg)

	var typeSummaries []types.TypeSummary

	done := false
	var nextToken *string

	for !done {
		clio.Info("listing types...")
		res, err := cfn.ListTypes(ctx, &cloudformation.ListTypesInput{
			Type:             types.RegistryTypeResource,
			Visibility:       types.VisibilityPublic,
			ProvisioningType: types.ProvisioningTypeFullyMutable,
			DeprecatedStatus: types.DeprecatedStatusLive,
			NextToken:        nextToken,
		})
		if err != nil {
			return err
		}
		typeSummaries = append(typeSummaries, res.TypeSummaries...)

		if res.NextToken == nil {
			done = true
		} else {
			nextToken = res.NextToken
		}
	}

	var eg errgroup.Group
	eg.SetLimit(5)

	for _, t2 := range typeSummaries {
		t := t2
		eg.Go(func() error {
			got, err := cfn.DescribeType(ctx, &cloudformation.DescribeTypeInput{
				TypeName: t.TypeName,
				Type:     types.RegistryTypeResource,
			})
			if err != nil {
				clio.Infow("skipping resource due to error", "resource", *t.TypeName, "error", err)
				return nil
			}
			// clio.Infof("%s", *got.Schema)
			u, err := url.Parse("https://example.com") // unused but required to be valid by the generator
			if err != nil {
				return err
			}
			clio.Infof("generating schema for %s", *got.TypeName)

			// typename is in the format: AWS::Logs::LogGroup
			// the filename will be aws/logs/loggroup/schema.go

			filename := strings.ToLower(*got.TypeName)
			filename = strings.ReplaceAll(filename, "::", "/")
			filename += "/schema.go"
			filename = filepath.Join("types", filename)

			parts := strings.Split(*got.TypeName, "::")
			packageName := parts[len(parts)-1] // e.g. 'loggroup'
			packageName = strings.ToLower(packageName)

			dir := filepath.Dir(filename)
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}

			schema, err := generate.ParseWithSchemaKeyRequired(*got.Schema, u, false)
			if err != nil {
				return err
			}

			schema.Title = "Resource"

			w, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			defer w.Close()

			g := generate.New(schema)
			err = g.CreateTypes()
			if err != nil {
				return err
			}
			generate.Output(w, g, packageName)
			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		return err
	}

	return nil
}
