// Code generated by schema-generate. DO NOT EDIT.

package link

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Definition of AWS::Oam::Link Resource Type
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  Label string `json:"Label,omitempty"`
  LabelTemplate string `json:"LabelTemplate,omitempty"`
  ResourceTypes []string `json:"ResourceTypes"`
  SinkIdentifier string `json:"SinkIdentifier"`

  // Tags to apply to the link
  Tags *Tags `json:"Tags,omitempty"`
}

// Tags Tags to apply to the link
type Tags struct {
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
    // Marshal the "Label" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Label\": ")
	if tmp, err := json.Marshal(strct.Label); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LabelTemplate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LabelTemplate\": ")
	if tmp, err := json.Marshal(strct.LabelTemplate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourceTypes" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ResourceTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceTypes\": ")
	if tmp, err := json.Marshal(strct.ResourceTypes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SinkIdentifier" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SinkIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SinkIdentifier\": ")
	if tmp, err := json.Marshal(strct.SinkIdentifier); err != nil {
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
    ResourceTypesReceived := false
    SinkIdentifierReceived := false
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
        case "Label":
            if err := json.Unmarshal([]byte(v), &strct.Label); err != nil {
                return err
             }
        case "LabelTemplate":
            if err := json.Unmarshal([]byte(v), &strct.LabelTemplate); err != nil {
                return err
             }
        case "ResourceTypes":
            if err := json.Unmarshal([]byte(v), &strct.ResourceTypes); err != nil {
                return err
             }
            ResourceTypesReceived = true
        case "SinkIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.SinkIdentifier); err != nil {
                return err
             }
            SinkIdentifierReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ResourceTypes (a required property) was received
    if !ResourceTypesReceived {
        return errors.New("\"ResourceTypes\" is required but was not present")
    }
    // check if SinkIdentifier (a required property) was received
    if !SinkIdentifierReceived {
        return errors.New("\"SinkIdentifier\" is required but was not present")
    }
    return nil
}

func (strct *Tags) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Tags) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, _ := range jsonMap {
        switch k {
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}