// Code generated by schema-generate. DO NOT EDIT.

package repository

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// EncryptionConfiguration The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
// 
// By default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.
// 
// For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html
type EncryptionConfiguration struct {
  EncryptionType string `json:"EncryptionType"`
  KmsKey string `json:"KmsKey,omitempty"`
}

// ImageScanningConfiguration The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.
type ImageScanningConfiguration struct {
  ScanOnPush bool `json:"ScanOnPush,omitempty"`
}

// LifecyclePolicy The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
type LifecyclePolicy struct {
  LifecyclePolicyText string `json:"LifecyclePolicyText,omitempty"`
  RegistryId string `json:"RegistryId,omitempty"`
}

// RepositoryPolicyText_object The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
type RepositoryPolicyText_object struct {
}

// Resource The AWS::ECR::Repository resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  EncryptionConfiguration *EncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
  ImageScanningConfiguration *ImageScanningConfiguration `json:"ImageScanningConfiguration,omitempty"`

  // The image tag mutability setting for the repository.
  ImageTagMutability string `json:"ImageTagMutability,omitempty"`
  LifecyclePolicy *LifecyclePolicy `json:"LifecyclePolicy,omitempty"`

  // The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
  RepositoryName string `json:"RepositoryName,omitempty"`

  // The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
  RepositoryPolicyText interface{} `json:"RepositoryPolicyText,omitempty"`
  RepositoryUri string `json:"RepositoryUri,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *EncryptionConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "EncryptionType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EncryptionType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionType\": ")
	if tmp, err := json.Marshal(strct.EncryptionType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KmsKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KmsKey\": ")
	if tmp, err := json.Marshal(strct.KmsKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EncryptionConfiguration) UnmarshalJSON(b []byte) error {
    EncryptionTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EncryptionType":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionType); err != nil {
                return err
             }
            EncryptionTypeReceived = true
        case "KmsKey":
            if err := json.Unmarshal([]byte(v), &strct.KmsKey); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if EncryptionType (a required property) was received
    if !EncryptionTypeReceived {
        return errors.New("\"EncryptionType\" is required but was not present")
    }
    return nil
}

func (strct *ImageScanningConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ScanOnPush" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScanOnPush\": ")
	if tmp, err := json.Marshal(strct.ScanOnPush); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ImageScanningConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ScanOnPush":
            if err := json.Unmarshal([]byte(v), &strct.ScanOnPush); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *LifecyclePolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "LifecyclePolicyText" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LifecyclePolicyText\": ")
	if tmp, err := json.Marshal(strct.LifecyclePolicyText); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegistryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegistryId\": ")
	if tmp, err := json.Marshal(strct.RegistryId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LifecyclePolicy) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LifecyclePolicyText":
            if err := json.Unmarshal([]byte(v), &strct.LifecyclePolicyText); err != nil {
                return err
             }
        case "RegistryId":
            if err := json.Unmarshal([]byte(v), &strct.RegistryId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Arn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Arn\": ")
	if tmp, err := json.Marshal(strct.Arn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EncryptionConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionConfiguration\": ")
	if tmp, err := json.Marshal(strct.EncryptionConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImageScanningConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageScanningConfiguration\": ")
	if tmp, err := json.Marshal(strct.ImageScanningConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImageTagMutability" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageTagMutability\": ")
	if tmp, err := json.Marshal(strct.ImageTagMutability); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LifecyclePolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LifecyclePolicy\": ")
	if tmp, err := json.Marshal(strct.LifecyclePolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RepositoryName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RepositoryName\": ")
	if tmp, err := json.Marshal(strct.RepositoryName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RepositoryPolicyText" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RepositoryPolicyText\": ")
	if tmp, err := json.Marshal(strct.RepositoryPolicyText); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RepositoryUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RepositoryUri\": ")
	if tmp, err := json.Marshal(strct.RepositoryUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Tags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Tags\": ")
	if tmp, err := json.Marshal(strct.Tags); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "EncryptionConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionConfiguration); err != nil {
                return err
             }
        case "ImageScanningConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.ImageScanningConfiguration); err != nil {
                return err
             }
        case "ImageTagMutability":
            if err := json.Unmarshal([]byte(v), &strct.ImageTagMutability); err != nil {
                return err
             }
        case "LifecyclePolicy":
            if err := json.Unmarshal([]byte(v), &strct.LifecyclePolicy); err != nil {
                return err
             }
        case "RepositoryName":
            if err := json.Unmarshal([]byte(v), &strct.RepositoryName); err != nil {
                return err
             }
        case "RepositoryPolicyText":
            if err := json.Unmarshal([]byte(v), &strct.RepositoryPolicyText); err != nil {
                return err
             }
        case "RepositoryUri":
            if err := json.Unmarshal([]byte(v), &strct.RepositoryUri); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Tag) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Key" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Key" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Key\": ")
	if tmp, err := json.Marshal(strct.Key); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Value" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Value" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Value\": ")
	if tmp, err := json.Marshal(strct.Value); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Tag) UnmarshalJSON(b []byte) error {
    KeyReceived := false
    ValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Key":
            if err := json.Unmarshal([]byte(v), &strct.Key); err != nil {
                return err
             }
            KeyReceived = true
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            ValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Key (a required property) was received
    if !KeyReceived {
        return errors.New("\"Key\" is required but was not present")
    }
    // check if Value (a required property) was received
    if !ValueReceived {
        return errors.New("\"Value\" is required but was not present")
    }
    return nil
}
