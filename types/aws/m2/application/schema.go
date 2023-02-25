// Code generated by schema-generate. DO NOT EDIT.

package application

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Represents an application that runs on an AWS Mainframe Modernization Environment
type Resource struct {
  ApplicationArn string `json:"ApplicationArn,omitempty"`
  ApplicationId string `json:"ApplicationId,omitempty"`
  Definition interface{} `json:"Definition"`
  Description string `json:"Description,omitempty"`
  EngineType string `json:"EngineType"`

  // The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting application-related resources.
  KmsKeyId string `json:"KmsKeyId,omitempty"`
  Name string `json:"Name"`
  Tags *TagMap `json:"Tags,omitempty"`
}

// TagMap 
type TagMap struct {
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ApplicationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApplicationArn\": ")
	if tmp, err := json.Marshal(strct.ApplicationArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ApplicationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApplicationId\": ")
	if tmp, err := json.Marshal(strct.ApplicationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Definition" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Definition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Definition\": ")
	if tmp, err := json.Marshal(strct.Definition); err != nil {
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
    // "EngineType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EngineType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EngineType\": ")
	if tmp, err := json.Marshal(strct.EngineType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KmsKeyId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KmsKeyId\": ")
	if tmp, err := json.Marshal(strct.KmsKeyId); err != nil {
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
    DefinitionReceived := false
    EngineTypeReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ApplicationArn":
            if err := json.Unmarshal([]byte(v), &strct.ApplicationArn); err != nil {
                return err
             }
        case "ApplicationId":
            if err := json.Unmarshal([]byte(v), &strct.ApplicationId); err != nil {
                return err
             }
        case "Definition":
            if err := json.Unmarshal([]byte(v), &strct.Definition); err != nil {
                return err
             }
            DefinitionReceived = true
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "EngineType":
            if err := json.Unmarshal([]byte(v), &strct.EngineType); err != nil {
                return err
             }
            EngineTypeReceived = true
        case "KmsKeyId":
            if err := json.Unmarshal([]byte(v), &strct.KmsKeyId); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Definition (a required property) was received
    if !DefinitionReceived {
        return errors.New("\"Definition\" is required but was not present")
    }
    // check if EngineType (a required property) was received
    if !EngineTypeReceived {
        return errors.New("\"EngineType\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}

func (strct *TagMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

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
    for k, _ := range jsonMap {
        switch k {
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}