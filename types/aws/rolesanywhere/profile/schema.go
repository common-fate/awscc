// Code generated by schema-generate. DO NOT EDIT.

package profile

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Definition of AWS::RolesAnywhere::Profile Resource Type
type Resource struct {
  DurationSeconds float64 `json:"DurationSeconds,omitempty"`
  Enabled bool `json:"Enabled,omitempty"`
  ManagedPolicyArns []string `json:"ManagedPolicyArns,omitempty"`
  Name string `json:"Name"`
  ProfileArn string `json:"ProfileArn,omitempty"`
  ProfileId string `json:"ProfileId,omitempty"`
  RequireInstanceProperties bool `json:"RequireInstanceProperties,omitempty"`
  RoleArns []string `json:"RoleArns"`
  SessionPolicy string `json:"SessionPolicy,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DurationSeconds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DurationSeconds\": ")
	if tmp, err := json.Marshal(strct.DurationSeconds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Enabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Enabled\": ")
	if tmp, err := json.Marshal(strct.Enabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ManagedPolicyArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ManagedPolicyArns\": ")
	if tmp, err := json.Marshal(strct.ManagedPolicyArns); err != nil {
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
    // Marshal the "ProfileArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileArn\": ")
	if tmp, err := json.Marshal(strct.ProfileArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProfileId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileId\": ")
	if tmp, err := json.Marshal(strct.ProfileId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RequireInstanceProperties" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RequireInstanceProperties\": ")
	if tmp, err := json.Marshal(strct.RequireInstanceProperties); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RoleArns" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RoleArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleArns\": ")
	if tmp, err := json.Marshal(strct.RoleArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SessionPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SessionPolicy\": ")
	if tmp, err := json.Marshal(strct.SessionPolicy); err != nil {
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
    NameReceived := false
    RoleArnsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DurationSeconds":
            if err := json.Unmarshal([]byte(v), &strct.DurationSeconds); err != nil {
                return err
             }
        case "Enabled":
            if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
                return err
             }
        case "ManagedPolicyArns":
            if err := json.Unmarshal([]byte(v), &strct.ManagedPolicyArns); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "ProfileArn":
            if err := json.Unmarshal([]byte(v), &strct.ProfileArn); err != nil {
                return err
             }
        case "ProfileId":
            if err := json.Unmarshal([]byte(v), &strct.ProfileId); err != nil {
                return err
             }
        case "RequireInstanceProperties":
            if err := json.Unmarshal([]byte(v), &strct.RequireInstanceProperties); err != nil {
                return err
             }
        case "RoleArns":
            if err := json.Unmarshal([]byte(v), &strct.RoleArns); err != nil {
                return err
             }
            RoleArnsReceived = true
        case "SessionPolicy":
            if err := json.Unmarshal([]byte(v), &strct.SessionPolicy); err != nil {
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
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if RoleArns (a required property) was received
    if !RoleArnsReceived {
        return errors.New("\"RoleArns\" is required but was not present")
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