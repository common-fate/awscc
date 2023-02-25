// Code generated by schema-generate. DO NOT EDIT.

package dbparametergroup

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// Parameters An array of parameter names and values for the parameter update.
type Parameters struct {
}

// Resource The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family
type Resource struct {

  // Specifies the name of the DB parameter group
  DBParameterGroupName string `json:"DBParameterGroupName,omitempty"`

  // Provides the customer-specified description for this DB parameter group.
  Description string `json:"Description"`

  // The DB parameter group family name.
  Family string `json:"Family"`

  // An array of parameter names and values for the parameter update.
  Parameters *Parameters `json:"Parameters,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DBParameterGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DBParameterGroupName\": ")
	if tmp, err := json.Marshal(strct.DBParameterGroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Description" field is required
    // only required object types supported for marshal checking (for now)
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
    // "Family" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Family" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Family\": ")
	if tmp, err := json.Marshal(strct.Family); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Parameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
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
    DescriptionReceived := false
    FamilyReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DBParameterGroupName":
            if err := json.Unmarshal([]byte(v), &strct.DBParameterGroupName); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
            DescriptionReceived = true
        case "Family":
            if err := json.Unmarshal([]byte(v), &strct.Family); err != nil {
                return err
             }
            FamilyReceived = true
        case "Parameters":
            if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
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
    // check if Description (a required property) was received
    if !DescriptionReceived {
        return errors.New("\"Description\" is required but was not present")
    }
    // check if Family (a required property) was received
    if !FamilyReceived {
        return errors.New("\"Family\" is required but was not present")
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
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Key (a required property) was received
    if !KeyReceived {
        return errors.New("\"Key\" is required but was not present")
    }
    return nil
}
