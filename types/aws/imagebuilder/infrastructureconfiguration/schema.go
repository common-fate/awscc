// Code generated by schema-generate. DO NOT EDIT.

package infrastructureconfiguration

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// InstanceMetadataOptions The instance metadata option settings for the infrastructure configuration.
type InstanceMetadataOptions struct {

  // Limit the number of hops that an instance metadata request can traverse to reach its destination.
  HttpPutResponseHopLimit int `json:"HttpPutResponseHopLimit,omitempty"`

  // Indicates whether a signed token header is required for instance metadata retrieval requests. The values affect the response as follows: 
  HttpTokens string `json:"HttpTokens,omitempty"`
}

// Logging The logging configuration of the infrastructure configuration.
type Logging struct {
  S3Logs *S3Logs `json:"S3Logs,omitempty"`
}

// Resource Resource schema for AWS::ImageBuilder::InfrastructureConfiguration
type Resource struct {

  // The Amazon Resource Name (ARN) of the infrastructure configuration.
  Arn string `json:"Arn,omitempty"`

  // The description of the infrastructure configuration.
  Description string `json:"Description,omitempty"`

  // The instance metadata option settings for the infrastructure configuration.
  InstanceMetadataOptions *InstanceMetadataOptions `json:"InstanceMetadataOptions,omitempty"`

  // The instance profile of the infrastructure configuration.
  InstanceProfileName string `json:"InstanceProfileName"`

  // The instance types of the infrastructure configuration.
  InstanceTypes []string `json:"InstanceTypes,omitempty"`

  // The EC2 key pair of the infrastructure configuration..
  KeyPair string `json:"KeyPair,omitempty"`

  // The logging configuration of the infrastructure configuration.
  Logging *Logging `json:"Logging,omitempty"`

  // The name of the infrastructure configuration.
  Name string `json:"Name"`

  // The tags attached to the resource created by Image Builder.
  ResourceTags *ResourceTags `json:"ResourceTags,omitempty"`

  // The security group IDs of the infrastructure configuration.
  SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`

  // The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
  SnsTopicArn string `json:"SnsTopicArn,omitempty"`

  // The subnet ID of the infrastructure configuration.
  SubnetId string `json:"SubnetId,omitempty"`

  // The tags associated with the component.
  Tags *Tags `json:"Tags,omitempty"`

  // The terminate instance on failure configuration of the infrastructure configuration.
  TerminateInstanceOnFailure bool `json:"TerminateInstanceOnFailure,omitempty"`
}

// ResourceTags The tags attached to the resource created by Image Builder.
type ResourceTags struct {
}

// S3Logs The S3 path in which to store the logs.
type S3Logs struct {

  // S3BucketName
  S3BucketName string `json:"S3BucketName,omitempty"`

  // S3KeyPrefix
  S3KeyPrefix string `json:"S3KeyPrefix,omitempty"`
}

// TagMap TagMap
type TagMap struct {

  // TagKey
  TagKey string `json:"TagKey,omitempty"`

  // TagValue
  TagValue string `json:"TagValue,omitempty"`
}

// Tags The tags associated with the component.
type Tags struct {
}

func (strct *InstanceMetadataOptions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "HttpPutResponseHopLimit" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HttpPutResponseHopLimit\": ")
	if tmp, err := json.Marshal(strct.HttpPutResponseHopLimit); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "HttpTokens" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HttpTokens\": ")
	if tmp, err := json.Marshal(strct.HttpTokens); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *InstanceMetadataOptions) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "HttpPutResponseHopLimit":
            if err := json.Unmarshal([]byte(v), &strct.HttpPutResponseHopLimit); err != nil {
                return err
             }
        case "HttpTokens":
            if err := json.Unmarshal([]byte(v), &strct.HttpTokens); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Logging) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "S3Logs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"S3Logs\": ")
	if tmp, err := json.Marshal(strct.S3Logs); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Logging) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "S3Logs":
            if err := json.Unmarshal([]byte(v), &strct.S3Logs); err != nil {
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
    // Marshal the "Description" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InstanceMetadataOptions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceMetadataOptions\": ")
	if tmp, err := json.Marshal(strct.InstanceMetadataOptions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "InstanceProfileName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "InstanceProfileName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceProfileName\": ")
	if tmp, err := json.Marshal(strct.InstanceProfileName); err != nil {
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
    // Marshal the "KeyPair" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyPair\": ")
	if tmp, err := json.Marshal(strct.KeyPair); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Logging" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Logging\": ")
	if tmp, err := json.Marshal(strct.Logging); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceTags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceTags\": ")
	if tmp, err := json.Marshal(strct.ResourceTags); err != nil {
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
    // Marshal the "SnsTopicArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SnsTopicArn\": ")
	if tmp, err := json.Marshal(strct.SnsTopicArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SubnetId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetId\": ")
	if tmp, err := json.Marshal(strct.SubnetId); err != nil {
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
    // Marshal the "TerminateInstanceOnFailure" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TerminateInstanceOnFailure\": ")
	if tmp, err := json.Marshal(strct.TerminateInstanceOnFailure); err != nil {
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
    InstanceProfileNameReceived := false
    NameReceived := false
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
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "InstanceMetadataOptions":
            if err := json.Unmarshal([]byte(v), &strct.InstanceMetadataOptions); err != nil {
                return err
             }
        case "InstanceProfileName":
            if err := json.Unmarshal([]byte(v), &strct.InstanceProfileName); err != nil {
                return err
             }
            InstanceProfileNameReceived = true
        case "InstanceTypes":
            if err := json.Unmarshal([]byte(v), &strct.InstanceTypes); err != nil {
                return err
             }
        case "KeyPair":
            if err := json.Unmarshal([]byte(v), &strct.KeyPair); err != nil {
                return err
             }
        case "Logging":
            if err := json.Unmarshal([]byte(v), &strct.Logging); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "ResourceTags":
            if err := json.Unmarshal([]byte(v), &strct.ResourceTags); err != nil {
                return err
             }
        case "SecurityGroupIds":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupIds); err != nil {
                return err
             }
        case "SnsTopicArn":
            if err := json.Unmarshal([]byte(v), &strct.SnsTopicArn); err != nil {
                return err
             }
        case "SubnetId":
            if err := json.Unmarshal([]byte(v), &strct.SubnetId); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TerminateInstanceOnFailure":
            if err := json.Unmarshal([]byte(v), &strct.TerminateInstanceOnFailure); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if InstanceProfileName (a required property) was received
    if !InstanceProfileNameReceived {
        return errors.New("\"InstanceProfileName\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}

func (strct *ResourceTags) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResourceTags) UnmarshalJSON(b []byte) error {
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

func (strct *S3Logs) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "S3BucketName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"S3BucketName\": ")
	if tmp, err := json.Marshal(strct.S3BucketName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "S3KeyPrefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"S3KeyPrefix\": ")
	if tmp, err := json.Marshal(strct.S3KeyPrefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *S3Logs) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "S3BucketName":
            if err := json.Unmarshal([]byte(v), &strct.S3BucketName); err != nil {
                return err
             }
        case "S3KeyPrefix":
            if err := json.Unmarshal([]byte(v), &strct.S3KeyPrefix); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *TagMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "TagKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TagKey\": ")
	if tmp, err := json.Marshal(strct.TagKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TagValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TagValue\": ")
	if tmp, err := json.Marshal(strct.TagValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TagMap) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "TagKey":
            if err := json.Unmarshal([]byte(v), &strct.TagKey); err != nil {
                return err
             }
        case "TagValue":
            if err := json.Unmarshal([]byte(v), &strct.TagValue); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
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
