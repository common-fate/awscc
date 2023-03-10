// Code generated by schema-generate. DO NOT EDIT.

package batchscramsecret

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::MSK::BatchScramSecret
type Resource struct {
  ClusterArn string `json:"ClusterArn"`
  SecretArnList []string `json:"SecretArnList,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ClusterArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ClusterArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClusterArn\": ")
	if tmp, err := json.Marshal(strct.ClusterArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecretArnList" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecretArnList\": ")
	if tmp, err := json.Marshal(strct.SecretArnList); err != nil {
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
    ClusterArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ClusterArn":
            if err := json.Unmarshal([]byte(v), &strct.ClusterArn); err != nil {
                return err
             }
            ClusterArnReceived = true
        case "SecretArnList":
            if err := json.Unmarshal([]byte(v), &strct.SecretArnList); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ClusterArn (a required property) was received
    if !ClusterArnReceived {
        return errors.New("\"ClusterArn\" is required but was not present")
    }
    return nil
}
