// Code generated by schema-generate. DO NOT EDIT.

package account

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// Resource Resource Type definition for AWS::ApiGateway::Account
type Resource struct {

  // The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.
  CloudWatchRoleArn string `json:"CloudWatchRoleArn,omitempty"`

  // Primary identifier which is manually generated.
  Id string `json:"Id,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CloudWatchRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CloudWatchRoleArn\": ")
	if tmp, err := json.Marshal(strct.CloudWatchRoleArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CloudWatchRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.CloudWatchRoleArn); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
