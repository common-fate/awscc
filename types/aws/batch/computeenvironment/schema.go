// Code generated by schema-generate. DO NOT EDIT.

package computeenvironment

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ComputeResources 
type ComputeResources struct {
  AllocationStrategy string `json:"AllocationStrategy,omitempty"`
  BidPercentage int `json:"BidPercentage,omitempty"`
  DesiredvCpus int `json:"DesiredvCpus,omitempty"`
  Ec2Configuration []*Ec2ConfigurationObject `json:"Ec2Configuration,omitempty"`
  Ec2KeyPair string `json:"Ec2KeyPair,omitempty"`
  ImageId string `json:"ImageId,omitempty"`
  InstanceRole string `json:"InstanceRole,omitempty"`
  InstanceTypes []string `json:"InstanceTypes,omitempty"`
  LaunchTemplate *LaunchTemplateSpecification `json:"LaunchTemplate,omitempty"`
  MaxvCpus int `json:"MaxvCpus"`
  MinvCpus int `json:"MinvCpus,omitempty"`
  PlacementGroup string `json:"PlacementGroup,omitempty"`
  SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
  SpotIamFleetRole string `json:"SpotIamFleetRole,omitempty"`
  Subnets []string `json:"Subnets"`

  // A key-value pair to associate with a resource.
  Tags *Tags `json:"Tags,omitempty"`
  Type string `json:"Type"`
  UpdateToLatestImageVersion bool `json:"UpdateToLatestImageVersion,omitempty"`
}

// Ec2ConfigurationObject 
type Ec2ConfigurationObject struct {
  ImageIdOverride string `json:"ImageIdOverride,omitempty"`
  ImageKubernetesVersion string `json:"ImageKubernetesVersion,omitempty"`
  ImageType string `json:"ImageType"`
}

// EksConfiguration 
type EksConfiguration struct {
  EksClusterArn string `json:"EksClusterArn"`
  KubernetesNamespace string `json:"KubernetesNamespace"`
}

// LaunchTemplateSpecification 
type LaunchTemplateSpecification struct {
  LaunchTemplateId string `json:"LaunchTemplateId,omitempty"`
  LaunchTemplateName string `json:"LaunchTemplateName,omitempty"`
  Version string `json:"Version,omitempty"`
}

// Resource Resource Type definition for AWS::Batch::ComputeEnvironment
type Resource struct {
  ComputeEnvironmentArn string `json:"ComputeEnvironmentArn,omitempty"`
  ComputeEnvironmentName string `json:"ComputeEnvironmentName,omitempty"`
  ComputeResources *ComputeResources `json:"ComputeResources,omitempty"`
  EksConfiguration *EksConfiguration `json:"EksConfiguration,omitempty"`
  ReplaceComputeEnvironment bool `json:"ReplaceComputeEnvironment,omitempty"`
  ServiceRole string `json:"ServiceRole,omitempty"`
  State string `json:"State,omitempty"`

  // A key-value pair to associate with a resource.
  Tags *Tags `json:"Tags,omitempty"`
  Type string `json:"Type"`
  UnmanagedvCpus int `json:"UnmanagedvCpus,omitempty"`
  UpdatePolicy *UpdatePolicy `json:"UpdatePolicy,omitempty"`
}

// Tags A key-value pair to associate with a resource.
type Tags struct {
}

// UpdatePolicy 
type UpdatePolicy struct {
  JobExecutionTimeoutMinutes int `json:"JobExecutionTimeoutMinutes,omitempty"`
  TerminateJobsOnUpdate bool `json:"TerminateJobsOnUpdate,omitempty"`
}

func (strct *ComputeResources) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AllocationStrategy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllocationStrategy\": ")
	if tmp, err := json.Marshal(strct.AllocationStrategy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BidPercentage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BidPercentage\": ")
	if tmp, err := json.Marshal(strct.BidPercentage); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DesiredvCpus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DesiredvCpus\": ")
	if tmp, err := json.Marshal(strct.DesiredvCpus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Ec2Configuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ec2Configuration\": ")
	if tmp, err := json.Marshal(strct.Ec2Configuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Ec2KeyPair" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ec2KeyPair\": ")
	if tmp, err := json.Marshal(strct.Ec2KeyPair); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImageId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageId\": ")
	if tmp, err := json.Marshal(strct.ImageId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InstanceRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceRole\": ")
	if tmp, err := json.Marshal(strct.InstanceRole); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InstanceTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceTypes\": ")
	if tmp, err := json.Marshal(strct.InstanceTypes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LaunchTemplate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LaunchTemplate\": ")
	if tmp, err := json.Marshal(strct.LaunchTemplate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MaxvCpus" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "MaxvCpus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxvCpus\": ")
	if tmp, err := json.Marshal(strct.MaxvCpus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MinvCpus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MinvCpus\": ")
	if tmp, err := json.Marshal(strct.MinvCpus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PlacementGroup" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PlacementGroup\": ")
	if tmp, err := json.Marshal(strct.PlacementGroup); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecurityGroupIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityGroupIds\": ")
	if tmp, err := json.Marshal(strct.SecurityGroupIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SpotIamFleetRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SpotIamFleetRole\": ")
	if tmp, err := json.Marshal(strct.SpotIamFleetRole); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Subnets" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Subnets" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subnets\": ")
	if tmp, err := json.Marshal(strct.Subnets); err != nil {
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
    // "Type" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UpdateToLatestImageVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UpdateToLatestImageVersion\": ")
	if tmp, err := json.Marshal(strct.UpdateToLatestImageVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ComputeResources) UnmarshalJSON(b []byte) error {
    MaxvCpusReceived := false
    SubnetsReceived := false
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AllocationStrategy":
            if err := json.Unmarshal([]byte(v), &strct.AllocationStrategy); err != nil {
                return err
             }
        case "BidPercentage":
            if err := json.Unmarshal([]byte(v), &strct.BidPercentage); err != nil {
                return err
             }
        case "DesiredvCpus":
            if err := json.Unmarshal([]byte(v), &strct.DesiredvCpus); err != nil {
                return err
             }
        case "Ec2Configuration":
            if err := json.Unmarshal([]byte(v), &strct.Ec2Configuration); err != nil {
                return err
             }
        case "Ec2KeyPair":
            if err := json.Unmarshal([]byte(v), &strct.Ec2KeyPair); err != nil {
                return err
             }
        case "ImageId":
            if err := json.Unmarshal([]byte(v), &strct.ImageId); err != nil {
                return err
             }
        case "InstanceRole":
            if err := json.Unmarshal([]byte(v), &strct.InstanceRole); err != nil {
                return err
             }
        case "InstanceTypes":
            if err := json.Unmarshal([]byte(v), &strct.InstanceTypes); err != nil {
                return err
             }
        case "LaunchTemplate":
            if err := json.Unmarshal([]byte(v), &strct.LaunchTemplate); err != nil {
                return err
             }
        case "MaxvCpus":
            if err := json.Unmarshal([]byte(v), &strct.MaxvCpus); err != nil {
                return err
             }
            MaxvCpusReceived = true
        case "MinvCpus":
            if err := json.Unmarshal([]byte(v), &strct.MinvCpus); err != nil {
                return err
             }
        case "PlacementGroup":
            if err := json.Unmarshal([]byte(v), &strct.PlacementGroup); err != nil {
                return err
             }
        case "SecurityGroupIds":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupIds); err != nil {
                return err
             }
        case "SpotIamFleetRole":
            if err := json.Unmarshal([]byte(v), &strct.SpotIamFleetRole); err != nil {
                return err
             }
        case "Subnets":
            if err := json.Unmarshal([]byte(v), &strct.Subnets); err != nil {
                return err
             }
            SubnetsReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        case "UpdateToLatestImageVersion":
            if err := json.Unmarshal([]byte(v), &strct.UpdateToLatestImageVersion); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MaxvCpus (a required property) was received
    if !MaxvCpusReceived {
        return errors.New("\"MaxvCpus\" is required but was not present")
    }
    // check if Subnets (a required property) was received
    if !SubnetsReceived {
        return errors.New("\"Subnets\" is required but was not present")
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}

func (strct *Ec2ConfigurationObject) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ImageIdOverride" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageIdOverride\": ")
	if tmp, err := json.Marshal(strct.ImageIdOverride); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImageKubernetesVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageKubernetesVersion\": ")
	if tmp, err := json.Marshal(strct.ImageKubernetesVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ImageType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ImageType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageType\": ")
	if tmp, err := json.Marshal(strct.ImageType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Ec2ConfigurationObject) UnmarshalJSON(b []byte) error {
    ImageTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ImageIdOverride":
            if err := json.Unmarshal([]byte(v), &strct.ImageIdOverride); err != nil {
                return err
             }
        case "ImageKubernetesVersion":
            if err := json.Unmarshal([]byte(v), &strct.ImageKubernetesVersion); err != nil {
                return err
             }
        case "ImageType":
            if err := json.Unmarshal([]byte(v), &strct.ImageType); err != nil {
                return err
             }
            ImageTypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ImageType (a required property) was received
    if !ImageTypeReceived {
        return errors.New("\"ImageType\" is required but was not present")
    }
    return nil
}

func (strct *EksConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "EksClusterArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EksClusterArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EksClusterArn\": ")
	if tmp, err := json.Marshal(strct.EksClusterArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "KubernetesNamespace" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "KubernetesNamespace" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KubernetesNamespace\": ")
	if tmp, err := json.Marshal(strct.KubernetesNamespace); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EksConfiguration) UnmarshalJSON(b []byte) error {
    EksClusterArnReceived := false
    KubernetesNamespaceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EksClusterArn":
            if err := json.Unmarshal([]byte(v), &strct.EksClusterArn); err != nil {
                return err
             }
            EksClusterArnReceived = true
        case "KubernetesNamespace":
            if err := json.Unmarshal([]byte(v), &strct.KubernetesNamespace); err != nil {
                return err
             }
            KubernetesNamespaceReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if EksClusterArn (a required property) was received
    if !EksClusterArnReceived {
        return errors.New("\"EksClusterArn\" is required but was not present")
    }
    // check if KubernetesNamespace (a required property) was received
    if !KubernetesNamespaceReceived {
        return errors.New("\"KubernetesNamespace\" is required but was not present")
    }
    return nil
}

func (strct *LaunchTemplateSpecification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "LaunchTemplateId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LaunchTemplateId\": ")
	if tmp, err := json.Marshal(strct.LaunchTemplateId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LaunchTemplateName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LaunchTemplateName\": ")
	if tmp, err := json.Marshal(strct.LaunchTemplateName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Version" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LaunchTemplateSpecification) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LaunchTemplateId":
            if err := json.Unmarshal([]byte(v), &strct.LaunchTemplateId); err != nil {
                return err
             }
        case "LaunchTemplateName":
            if err := json.Unmarshal([]byte(v), &strct.LaunchTemplateName); err != nil {
                return err
             }
        case "Version":
            if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
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
    // Marshal the "ComputeEnvironmentArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComputeEnvironmentArn\": ")
	if tmp, err := json.Marshal(strct.ComputeEnvironmentArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ComputeEnvironmentName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComputeEnvironmentName\": ")
	if tmp, err := json.Marshal(strct.ComputeEnvironmentName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ComputeResources" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComputeResources\": ")
	if tmp, err := json.Marshal(strct.ComputeResources); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EksConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EksConfiguration\": ")
	if tmp, err := json.Marshal(strct.EksConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ReplaceComputeEnvironment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplaceComputeEnvironment\": ")
	if tmp, err := json.Marshal(strct.ReplaceComputeEnvironment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServiceRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceRole\": ")
	if tmp, err := json.Marshal(strct.ServiceRole); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "State" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"State\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
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
    // "Type" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UnmanagedvCpus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UnmanagedvCpus\": ")
	if tmp, err := json.Marshal(strct.UnmanagedvCpus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UpdatePolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UpdatePolicy\": ")
	if tmp, err := json.Marshal(strct.UpdatePolicy); err != nil {
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
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ComputeEnvironmentArn":
            if err := json.Unmarshal([]byte(v), &strct.ComputeEnvironmentArn); err != nil {
                return err
             }
        case "ComputeEnvironmentName":
            if err := json.Unmarshal([]byte(v), &strct.ComputeEnvironmentName); err != nil {
                return err
             }
        case "ComputeResources":
            if err := json.Unmarshal([]byte(v), &strct.ComputeResources); err != nil {
                return err
             }
        case "EksConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EksConfiguration); err != nil {
                return err
             }
        case "ReplaceComputeEnvironment":
            if err := json.Unmarshal([]byte(v), &strct.ReplaceComputeEnvironment); err != nil {
                return err
             }
        case "ServiceRole":
            if err := json.Unmarshal([]byte(v), &strct.ServiceRole); err != nil {
                return err
             }
        case "State":
            if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        case "UnmanagedvCpus":
            if err := json.Unmarshal([]byte(v), &strct.UnmanagedvCpus); err != nil {
                return err
             }
        case "UpdatePolicy":
            if err := json.Unmarshal([]byte(v), &strct.UpdatePolicy); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}

func (strct *Tags) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Tags) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, _ := range jsonMap {
        switch k {
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *UpdatePolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "JobExecutionTimeoutMinutes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"JobExecutionTimeoutMinutes\": ")
	if tmp, err := json.Marshal(strct.JobExecutionTimeoutMinutes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TerminateJobsOnUpdate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TerminateJobsOnUpdate\": ")
	if tmp, err := json.Marshal(strct.TerminateJobsOnUpdate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UpdatePolicy) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "JobExecutionTimeoutMinutes":
            if err := json.Unmarshal([]byte(v), &strct.JobExecutionTimeoutMinutes); err != nil {
                return err
             }
        case "TerminateJobsOnUpdate":
            if err := json.Unmarshal([]byte(v), &strct.TerminateJobsOnUpdate); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
