// Code generated by schema-generate. DO NOT EDIT.

package variantstore

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ReferenceItem 
type ReferenceItem struct {
  ReferenceArn string `json:"ReferenceArn"`
}

// Resource Definition of AWS::Omics::VariantStore Resource Type
type Resource struct {
  CreationTime string `json:"CreationTime,omitempty"`
  Description string `json:"Description,omitempty"`
  Id string `json:"Id,omitempty"`
  Name string `json:"Name"`
  Reference *ReferenceItem `json:"Reference"`
  SseConfig *SseConfig `json:"SseConfig,omitempty"`
  Status string `json:"Status,omitempty"`
  StatusMessage string `json:"StatusMessage,omitempty"`
  StoreArn string `json:"StoreArn,omitempty"`
  StoreSizeBytes float64 `json:"StoreSizeBytes,omitempty"`
  Tags *TagMap `json:"Tags,omitempty"`
  UpdateTime string `json:"UpdateTime,omitempty"`
}

// SseConfig 
type SseConfig struct {
  KeyArn string `json:"KeyArn,omitempty"`
  Type string `json:"Type"`
}

// TagMap 
type TagMap struct {
}

func (strct *ReferenceItem) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ReferenceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ReferenceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReferenceArn\": ")
	if tmp, err := json.Marshal(strct.ReferenceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReferenceItem) UnmarshalJSON(b []byte) error {
    ReferenceArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ReferenceArn":
            if err := json.Unmarshal([]byte(v), &strct.ReferenceArn); err != nil {
                return err
             }
            ReferenceArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ReferenceArn (a required property) was received
    if !ReferenceArnReceived {
        return errors.New("\"ReferenceArn\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CreationTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreationTime\": ")
	if tmp, err := json.Marshal(strct.CreationTime); err != nil {
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
    // "Reference" field is required
    if strct.Reference == nil {
        return nil, errors.New("Reference is a required field")
    }
    // Marshal the "Reference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Reference\": ")
	if tmp, err := json.Marshal(strct.Reference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SseConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SseConfig\": ")
	if tmp, err := json.Marshal(strct.SseConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatusMessage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatusMessage\": ")
	if tmp, err := json.Marshal(strct.StatusMessage); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StoreArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StoreArn\": ")
	if tmp, err := json.Marshal(strct.StoreArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StoreSizeBytes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StoreSizeBytes\": ")
	if tmp, err := json.Marshal(strct.StoreSizeBytes); err != nil {
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
    // Marshal the "UpdateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UpdateTime\": ")
	if tmp, err := json.Marshal(strct.UpdateTime); err != nil {
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
    ReferenceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CreationTime":
            if err := json.Unmarshal([]byte(v), &strct.CreationTime); err != nil {
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
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Reference":
            if err := json.Unmarshal([]byte(v), &strct.Reference); err != nil {
                return err
             }
            ReferenceReceived = true
        case "SseConfig":
            if err := json.Unmarshal([]byte(v), &strct.SseConfig); err != nil {
                return err
             }
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "StatusMessage":
            if err := json.Unmarshal([]byte(v), &strct.StatusMessage); err != nil {
                return err
             }
        case "StoreArn":
            if err := json.Unmarshal([]byte(v), &strct.StoreArn); err != nil {
                return err
             }
        case "StoreSizeBytes":
            if err := json.Unmarshal([]byte(v), &strct.StoreSizeBytes); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "UpdateTime":
            if err := json.Unmarshal([]byte(v), &strct.UpdateTime); err != nil {
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
    // check if Reference (a required property) was received
    if !ReferenceReceived {
        return errors.New("\"Reference\" is required but was not present")
    }
    return nil
}

func (strct *SseConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "KeyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyArn\": ")
	if tmp, err := json.Marshal(strct.KeyArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Type" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SseConfig) UnmarshalJSON(b []byte) error {
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "KeyArn":
            if err := json.Unmarshal([]byte(v), &strct.KeyArn); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}

func (strct *TagMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TagMap) UnmarshalJSON(b []byte) error {
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
