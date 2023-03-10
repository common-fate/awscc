// Code generated by schema-generate. DO NOT EDIT.

package resourcepolicy

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource The resource schema for AWSLogs ResourcePolicy
type Resource struct {

  // The policy document
  PolicyDocument string `json:"PolicyDocument"`

  // A name for resource policy
  PolicyName string `json:"PolicyName"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "PolicyDocument" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyDocument" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyDocument\": ")
	if tmp, err := json.Marshal(strct.PolicyDocument); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PolicyName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyName\": ")
	if tmp, err := json.Marshal(strct.PolicyName); err != nil {
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
    PolicyDocumentReceived := false
    PolicyNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PolicyDocument":
            if err := json.Unmarshal([]byte(v), &strct.PolicyDocument); err != nil {
                return err
             }
            PolicyDocumentReceived = true
        case "PolicyName":
            if err := json.Unmarshal([]byte(v), &strct.PolicyName); err != nil {
                return err
             }
            PolicyNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if PolicyDocument (a required property) was received
    if !PolicyDocumentReceived {
        return errors.New("\"PolicyDocument\" is required but was not present")
    }
    // check if PolicyName (a required property) was received
    if !PolicyNameReceived {
        return errors.New("\"PolicyName\" is required but was not present")
    }
    return nil
}
