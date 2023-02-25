// Code generated by schema-generate. DO NOT EDIT.

package storedquery

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::Config::StoredQuery
type Resource struct {
  QueryArn string `json:"QueryArn,omitempty"`
  QueryDescription string `json:"QueryDescription,omitempty"`
  QueryExpression string `json:"QueryExpression"`
  QueryId string `json:"QueryId,omitempty"`
  QueryName string `json:"QueryName"`

  // The tags for the stored query.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "QueryArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueryArn\": ")
	if tmp, err := json.Marshal(strct.QueryArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "QueryDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueryDescription\": ")
	if tmp, err := json.Marshal(strct.QueryDescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "QueryExpression" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "QueryExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueryExpression\": ")
	if tmp, err := json.Marshal(strct.QueryExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "QueryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueryId\": ")
	if tmp, err := json.Marshal(strct.QueryId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "QueryName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "QueryName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueryName\": ")
	if tmp, err := json.Marshal(strct.QueryName); err != nil {
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
    QueryExpressionReceived := false
    QueryNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "QueryArn":
            if err := json.Unmarshal([]byte(v), &strct.QueryArn); err != nil {
                return err
             }
        case "QueryDescription":
            if err := json.Unmarshal([]byte(v), &strct.QueryDescription); err != nil {
                return err
             }
        case "QueryExpression":
            if err := json.Unmarshal([]byte(v), &strct.QueryExpression); err != nil {
                return err
             }
            QueryExpressionReceived = true
        case "QueryId":
            if err := json.Unmarshal([]byte(v), &strct.QueryId); err != nil {
                return err
             }
        case "QueryName":
            if err := json.Unmarshal([]byte(v), &strct.QueryName); err != nil {
                return err
             }
            QueryNameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if QueryExpression (a required property) was received
    if !QueryExpressionReceived {
        return errors.New("\"QueryExpression\" is required but was not present")
    }
    // check if QueryName (a required property) was received
    if !QueryNameReceived {
        return errors.New("\"QueryName\" is required but was not present")
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
