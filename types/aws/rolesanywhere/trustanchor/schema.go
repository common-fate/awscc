// Code generated by schema-generate. DO NOT EDIT.

package trustanchor

import (
    "encoding/json"
    "fmt"
    "bytes"
    "errors"
)

// Resource Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.
type Resource struct {
  Enabled bool `json:"Enabled,omitempty"`
  Name string `json:"Name"`
  Source *Source `json:"Source"`
  Tags []*Tag `json:"Tags,omitempty"`
  TrustAnchorArn string `json:"TrustAnchorArn,omitempty"`
  TrustAnchorId string `json:"TrustAnchorId,omitempty"`
}

// Source 
type Source struct {
  SourceData interface{} `json:"SourceData,omitempty"`
  SourceType string `json:"SourceType,omitempty"`
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
    // "Source" field is required
    if strct.Source == nil {
        return nil, errors.New("Source is a required field")
    }
    // Marshal the "Source" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
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
    // Marshal the "TrustAnchorArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TrustAnchorArn\": ")
	if tmp, err := json.Marshal(strct.TrustAnchorArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TrustAnchorId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TrustAnchorId\": ")
	if tmp, err := json.Marshal(strct.TrustAnchorId); err != nil {
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
    SourceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Enabled":
            if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
            SourceReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TrustAnchorArn":
            if err := json.Unmarshal([]byte(v), &strct.TrustAnchorArn); err != nil {
                return err
             }
        case "TrustAnchorId":
            if err := json.Unmarshal([]byte(v), &strct.TrustAnchorId); err != nil {
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
    // check if Source (a required property) was received
    if !SourceReceived {
        return errors.New("\"Source\" is required but was not present")
    }
    return nil
}

func (strct *Source) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "SourceData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceData\": ")
	if tmp, err := json.Marshal(strct.SourceData); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceType\": ")
	if tmp, err := json.Marshal(strct.SourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Source) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "SourceData":
            if err := json.Unmarshal([]byte(v), &strct.SourceData); err != nil {
                return err
             }
        case "SourceType":
            if err := json.Unmarshal([]byte(v), &strct.SourceType); err != nil {
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
