// Code generated by schema-generate. DO NOT EDIT.

package packaginggroup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Authorization 
type Authorization struct {

  // The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.
  CdnIdentifierSecret string `json:"CdnIdentifierSecret"`

  // The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.
  SecretsRoleArn string `json:"SecretsRoleArn"`
}

// LogConfiguration 
type LogConfiguration struct {

  // Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.
  LogGroupName string `json:"LogGroupName,omitempty"`
}

// Resource Resource schema for AWS::MediaPackage::PackagingGroup
type Resource struct {

  // The ARN of the PackagingGroup.
  Arn string `json:"Arn,omitempty"`

  // CDN Authorization
  Authorization *Authorization `json:"Authorization,omitempty"`

  // The fully qualified domain name for Assets in the PackagingGroup.
  DomainName string `json:"DomainName,omitempty"`

  // The configuration parameters for egress access logging.
  EgressAccessLogs *LogConfiguration `json:"EgressAccessLogs,omitempty"`

  // The ID of the PackagingGroup.
  Id string `json:"Id"`

  // A collection of tags associated with a resource
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *Authorization) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "CdnIdentifierSecret" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "CdnIdentifierSecret" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CdnIdentifierSecret\": ")
	if tmp, err := json.Marshal(strct.CdnIdentifierSecret); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SecretsRoleArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SecretsRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecretsRoleArn\": ")
	if tmp, err := json.Marshal(strct.SecretsRoleArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Authorization) UnmarshalJSON(b []byte) error {
    CdnIdentifierSecretReceived := false
    SecretsRoleArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CdnIdentifierSecret":
            if err := json.Unmarshal([]byte(v), &strct.CdnIdentifierSecret); err != nil {
                return err
             }
            CdnIdentifierSecretReceived = true
        case "SecretsRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.SecretsRoleArn); err != nil {
                return err
             }
            SecretsRoleArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if CdnIdentifierSecret (a required property) was received
    if !CdnIdentifierSecretReceived {
        return errors.New("\"CdnIdentifierSecret\" is required but was not present")
    }
    // check if SecretsRoleArn (a required property) was received
    if !SecretsRoleArnReceived {
        return errors.New("\"SecretsRoleArn\" is required but was not present")
    }
    return nil
}

func (strct *LogConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "LogGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LogGroupName\": ")
	if tmp, err := json.Marshal(strct.LogGroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LogConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LogGroupName":
            if err := json.Unmarshal([]byte(v), &strct.LogGroupName); err != nil {
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
    // Marshal the "Authorization" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Authorization\": ")
	if tmp, err := json.Marshal(strct.Authorization); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainName\": ")
	if tmp, err := json.Marshal(strct.DomainName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EgressAccessLogs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EgressAccessLogs\": ")
	if tmp, err := json.Marshal(strct.EgressAccessLogs); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Id" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
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
    IdReceived := false
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
        case "Authorization":
            if err := json.Unmarshal([]byte(v), &strct.Authorization); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
        case "EgressAccessLogs":
            if err := json.Unmarshal([]byte(v), &strct.EgressAccessLogs); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
            IdReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Id (a required property) was received
    if !IdReceived {
        return errors.New("\"Id\" is required but was not present")
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
