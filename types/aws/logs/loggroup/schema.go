// Code generated by schema-generate. DO NOT EDIT.

package loggroup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// DataProtectionPolicy The body of the policy document you want to use for this topic.
// 
// You can only add one policy per topic.
// 
// The policy must be in JSON string format.
// 
// Length Constraints: Maximum length of 30720
type DataProtectionPolicy struct {
}

// Resource Resource schema for AWS::Logs::LogGroup
type Resource struct {

  // The CloudWatch log group ARN.
  Arn string `json:"Arn,omitempty"`

  // The body of the policy document you want to use for this topic.
  // 
  // You can only add one policy per topic.
  // 
  // The policy must be in JSON string format.
  // 
  // Length Constraints: Maximum length of 30720
  DataProtectionPolicy *DataProtectionPolicy `json:"DataProtectionPolicy,omitempty"`

  // The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
  KmsKeyId string `json:"KmsKeyId,omitempty"`

  // The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
  LogGroupName string `json:"LogGroupName,omitempty"`

  // The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
  RetentionInDays int `json:"RetentionInDays,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.
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
    // Marshal the "DataProtectionPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataProtectionPolicy\": ")
	if tmp, err := json.Marshal(strct.DataProtectionPolicy); err != nil {
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
    // Marshal the "LogGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LogGroupName\": ")
	if tmp, err := json.Marshal(strct.LogGroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RetentionInDays" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RetentionInDays\": ")
	if tmp, err := json.Marshal(strct.RetentionInDays); err != nil {
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
        case "DataProtectionPolicy":
            if err := json.Unmarshal([]byte(v), &strct.DataProtectionPolicy); err != nil {
                return err
             }
        case "KmsKeyId":
            if err := json.Unmarshal([]byte(v), &strct.KmsKeyId); err != nil {
                return err
             }
        case "LogGroupName":
            if err := json.Unmarshal([]byte(v), &strct.LogGroupName); err != nil {
                return err
             }
        case "RetentionInDays":
            if err := json.Unmarshal([]byte(v), &strct.RetentionInDays); err != nil {
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