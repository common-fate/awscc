// Code generated by schema-generate. DO NOT EDIT.

package oidcprovider

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// Resource Resource Type definition for AWS::IAM::OIDCProvider
type Resource struct {

  // Amazon Resource Name (ARN) of the OIDC provider
  Arn string `json:"Arn,omitempty"`
  ClientIdList []string `json:"ClientIdList,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
  ThumbprintList []string `json:"ThumbprintList"`
  Url string `json:"Url,omitempty"`
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
    // Marshal the "ClientIdList" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientIdList\": ")
	if tmp, err := json.Marshal(strct.ClientIdList); err != nil {
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
    // "ThumbprintList" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ThumbprintList" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ThumbprintList\": ")
	if tmp, err := json.Marshal(strct.ThumbprintList); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Url" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
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
    ThumbprintListReceived := false
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
        case "ClientIdList":
            if err := json.Unmarshal([]byte(v), &strct.ClientIdList); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "ThumbprintList":
            if err := json.Unmarshal([]byte(v), &strct.ThumbprintList); err != nil {
                return err
             }
            ThumbprintListReceived = true
        case "Url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ThumbprintList (a required property) was received
    if !ThumbprintListReceived {
        return errors.New("\"ThumbprintList\" is required but was not present")
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
