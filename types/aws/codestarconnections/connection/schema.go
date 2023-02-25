// Code generated by schema-generate. DO NOT EDIT.

package connection

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Schema for AWS::CodeStarConnections::Connection resource which can be used to connect external source providers with AWS CodePipeline
type Resource struct {

  // The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.
  ConnectionArn string `json:"ConnectionArn,omitempty"`

  // The name of the connection. Connection names must be unique in an AWS user account.
  ConnectionName string `json:"ConnectionName"`

  // The current status of the connection.
  ConnectionStatus string `json:"ConnectionStatus,omitempty"`

  // The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
  HostArn string `json:"HostArn,omitempty"`

  // The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.
  OwnerAccountId string `json:"OwnerAccountId,omitempty"`

  // The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
  ProviderType string `json:"ProviderType,omitempty"`

  // Specifies the tags applied to a connection.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ConnectionArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectionArn\": ")
	if tmp, err := json.Marshal(strct.ConnectionArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ConnectionName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConnectionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectionName\": ")
	if tmp, err := json.Marshal(strct.ConnectionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ConnectionStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectionStatus\": ")
	if tmp, err := json.Marshal(strct.ConnectionStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "HostArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HostArn\": ")
	if tmp, err := json.Marshal(strct.HostArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OwnerAccountId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OwnerAccountId\": ")
	if tmp, err := json.Marshal(strct.OwnerAccountId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProviderType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProviderType\": ")
	if tmp, err := json.Marshal(strct.ProviderType); err != nil {
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
    ConnectionNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ConnectionArn":
            if err := json.Unmarshal([]byte(v), &strct.ConnectionArn); err != nil {
                return err
             }
        case "ConnectionName":
            if err := json.Unmarshal([]byte(v), &strct.ConnectionName); err != nil {
                return err
             }
            ConnectionNameReceived = true
        case "ConnectionStatus":
            if err := json.Unmarshal([]byte(v), &strct.ConnectionStatus); err != nil {
                return err
             }
        case "HostArn":
            if err := json.Unmarshal([]byte(v), &strct.HostArn); err != nil {
                return err
             }
        case "OwnerAccountId":
            if err := json.Unmarshal([]byte(v), &strct.OwnerAccountId); err != nil {
                return err
             }
        case "ProviderType":
            if err := json.Unmarshal([]byte(v), &strct.ProviderType); err != nil {
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
    // check if ConnectionName (a required property) was received
    if !ConnectionNameReceived {
        return errors.New("\"ConnectionName\" is required but was not present")
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