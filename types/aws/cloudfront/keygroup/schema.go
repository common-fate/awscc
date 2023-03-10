// Code generated by schema-generate. DO NOT EDIT.

package keygroup

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// KeyGroupConfig 
type KeyGroupConfig struct {
  Comment string `json:"Comment,omitempty"`
  Items []string `json:"Items"`
  Name string `json:"Name"`
}

// Resource Resource Type definition for AWS::CloudFront::KeyGroup
type Resource struct {
  Id string `json:"Id,omitempty"`
  KeyGroupConfig *KeyGroupConfig `json:"KeyGroupConfig"`
  LastModifiedTime string `json:"LastModifiedTime,omitempty"`
}

func (strct *KeyGroupConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Comment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Comment\": ")
	if tmp, err := json.Marshal(strct.Comment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Items" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Items" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Items\": ")
	if tmp, err := json.Marshal(strct.Items); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *KeyGroupConfig) UnmarshalJSON(b []byte) error {
    ItemsReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Comment":
            if err := json.Unmarshal([]byte(v), &strct.Comment); err != nil {
                return err
             }
        case "Items":
            if err := json.Unmarshal([]byte(v), &strct.Items); err != nil {
                return err
             }
            ItemsReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Items (a required property) was received
    if !ItemsReceived {
        return errors.New("\"Items\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // "KeyGroupConfig" field is required
    if strct.KeyGroupConfig == nil {
        return nil, errors.New("KeyGroupConfig is a required field")
    }
    // Marshal the "KeyGroupConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyGroupConfig\": ")
	if tmp, err := json.Marshal(strct.KeyGroupConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LastModifiedTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LastModifiedTime\": ")
	if tmp, err := json.Marshal(strct.LastModifiedTime); err != nil {
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
    KeyGroupConfigReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "KeyGroupConfig":
            if err := json.Unmarshal([]byte(v), &strct.KeyGroupConfig); err != nil {
                return err
             }
            KeyGroupConfigReceived = true
        case "LastModifiedTime":
            if err := json.Unmarshal([]byte(v), &strct.LastModifiedTime); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if KeyGroupConfig (a required property) was received
    if !KeyGroupConfigReceived {
        return errors.New("\"KeyGroupConfig\" is required but was not present")
    }
    return nil
}
