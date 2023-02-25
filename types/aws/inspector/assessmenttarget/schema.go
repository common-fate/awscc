// Code generated by schema-generate. DO NOT EDIT.

package assessmenttarget

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::Inspector::AssessmentTarget
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  AssessmentTargetName string `json:"AssessmentTargetName,omitempty"`
  ResourceGroupArn string `json:"ResourceGroupArn,omitempty"`
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
    // Marshal the "AssessmentTargetName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssessmentTargetName\": ")
	if tmp, err := json.Marshal(strct.AssessmentTargetName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceGroupArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceGroupArn\": ")
	if tmp, err := json.Marshal(strct.ResourceGroupArn); err != nil {
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
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AssessmentTargetName":
            if err := json.Unmarshal([]byte(v), &strct.AssessmentTargetName); err != nil {
                return err
             }
        case "ResourceGroupArn":
            if err := json.Unmarshal([]byte(v), &strct.ResourceGroupArn); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
