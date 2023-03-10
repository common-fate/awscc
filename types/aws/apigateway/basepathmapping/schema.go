// Code generated by schema-generate. DO NOT EDIT.

package basepathmapping

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::ApiGateway::BasePathMapping
type Resource struct {

  // The base path name that callers of the API must provide in the URL after the domain name.
  BasePath string `json:"BasePath,omitempty"`

  // The DomainName of an AWS::ApiGateway::DomainName resource.
  DomainName string `json:"DomainName"`

  // The ID of the API.
  RestApiId string `json:"RestApiId,omitempty"`

  // The name of the API's stage.
  Stage string `json:"Stage,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "BasePath" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BasePath\": ")
	if tmp, err := json.Marshal(strct.BasePath); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DomainName" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "Stage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Stage\": ")
	if tmp, err := json.Marshal(strct.Stage); err != nil {
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
    DomainNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BasePath":
            if err := json.Unmarshal([]byte(v), &strct.BasePath); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
            DomainNameReceived = true
        case "RestApiId":
            if err := json.Unmarshal([]byte(v), &strct.RestApiId); err != nil {
                return err
             }
        case "Stage":
            if err := json.Unmarshal([]byte(v), &strct.Stage); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DomainName (a required property) was received
    if !DomainNameReceived {
        return errors.New("\"DomainName\" is required but was not present")
    }
    return nil
}
