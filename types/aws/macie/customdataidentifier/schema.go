// Code generated by schema-generate. DO NOT EDIT.

package customdataidentifier

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// Resource Macie CustomDataIdentifier resource schema
type Resource struct {

  // Custom data identifier ARN.
  Arn string `json:"Arn,omitempty"`

  // Description of custom data identifier.
  Description string `json:"Description,omitempty"`

  // Custom data identifier ID.
  Id string `json:"Id,omitempty"`

  // Words to be ignored.
  IgnoreWords []string `json:"IgnoreWords,omitempty"`

  // Keywords to be matched against.
  Keywords []string `json:"Keywords,omitempty"`

  // Maximum match distance.
  MaximumMatchDistance int `json:"MaximumMatchDistance,omitempty"`

  // Name of custom data identifier.
  Name string `json:"Name"`

  // Regular expression for custom data identifier.
  Regex string `json:"Regex"`
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
    // Marshal the "IgnoreWords" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IgnoreWords\": ")
	if tmp, err := json.Marshal(strct.IgnoreWords); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Keywords" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Keywords\": ")
	if tmp, err := json.Marshal(strct.Keywords); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaximumMatchDistance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaximumMatchDistance\": ")
	if tmp, err := json.Marshal(strct.MaximumMatchDistance); err != nil {
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
    // "Regex" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Regex" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Regex\": ")
	if tmp, err := json.Marshal(strct.Regex); err != nil {
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
    NameReceived := false
    RegexReceived := false
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
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "IgnoreWords":
            if err := json.Unmarshal([]byte(v), &strct.IgnoreWords); err != nil {
                return err
             }
        case "Keywords":
            if err := json.Unmarshal([]byte(v), &strct.Keywords); err != nil {
                return err
             }
        case "MaximumMatchDistance":
            if err := json.Unmarshal([]byte(v), &strct.MaximumMatchDistance); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Regex":
            if err := json.Unmarshal([]byte(v), &strct.Regex); err != nil {
                return err
             }
            RegexReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Regex (a required property) was received
    if !RegexReceived {
        return errors.New("\"Regex\" is required but was not present")
    }
    return nil
}
