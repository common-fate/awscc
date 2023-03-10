// Code generated by schema-generate. DO NOT EDIT.

package model

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::ApiGateway::Model
type Resource struct {

  // The content type for the model.
  ContentType string `json:"ContentType,omitempty"`

  // A description that identifies this model.
  Description string `json:"Description,omitempty"`

  // A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
  Name string `json:"Name,omitempty"`

  // The ID of a REST API with which to associate this model.
  RestApiId string `json:"RestApiId"`

  // The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
  Schema interface{} `json:"Schema,omitempty"`
}

// Schema_object The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
type Schema_object struct {
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ContentType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContentType\": ")
	if tmp, err := json.Marshal(strct.ContentType); err != nil {
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
    // "RestApiId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RestApiId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RestApiId\": ")
	if tmp, err := json.Marshal(strct.RestApiId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Schema" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Schema\": ")
	if tmp, err := json.Marshal(strct.Schema); err != nil {
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
    RestApiIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContentType":
            if err := json.Unmarshal([]byte(v), &strct.ContentType); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "RestApiId":
            if err := json.Unmarshal([]byte(v), &strct.RestApiId); err != nil {
                return err
             }
            RestApiIdReceived = true
        case "Schema":
            if err := json.Unmarshal([]byte(v), &strct.Schema); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if RestApiId (a required property) was received
    if !RestApiIdReceived {
        return errors.New("\"RestApiId\" is required but was not present")
    }
    return nil
}
