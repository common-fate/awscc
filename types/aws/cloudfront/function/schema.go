// Code generated by schema-generate. DO NOT EDIT.

package function

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// FunctionConfig 
type FunctionConfig struct {
  Comment string `json:"Comment"`
  Runtime string `json:"Runtime"`
}

// FunctionMetadata 
type FunctionMetadata struct {
  FunctionARN string `json:"FunctionARN,omitempty"`
}

// Resource Resource Type definition for AWS::CloudFront::Function
type Resource struct {
  AutoPublish bool `json:"AutoPublish,omitempty"`
  FunctionARN string `json:"FunctionARN,omitempty"`
  FunctionCode string `json:"FunctionCode"`
  FunctionConfig *FunctionConfig `json:"FunctionConfig"`
  FunctionMetadata *FunctionMetadata `json:"FunctionMetadata,omitempty"`
  Name string `json:"Name"`
  Stage string `json:"Stage,omitempty"`
}

func (strct *FunctionConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Comment" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Comment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Comment\": ")
	if tmp, err := json.Marshal(strct.Comment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Runtime" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Runtime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Runtime\": ")
	if tmp, err := json.Marshal(strct.Runtime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *FunctionConfig) UnmarshalJSON(b []byte) error {
    CommentReceived := false
    RuntimeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Comment":
            if err := json.Unmarshal([]byte(v), &strct.Comment); err != nil {
                return err
             }
            CommentReceived = true
        case "Runtime":
            if err := json.Unmarshal([]byte(v), &strct.Runtime); err != nil {
                return err
             }
            RuntimeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Comment (a required property) was received
    if !CommentReceived {
        return errors.New("\"Comment\" is required but was not present")
    }
    // check if Runtime (a required property) was received
    if !RuntimeReceived {
        return errors.New("\"Runtime\" is required but was not present")
    }
    return nil
}

func (strct *FunctionMetadata) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "FunctionARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionARN\": ")
	if tmp, err := json.Marshal(strct.FunctionARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *FunctionMetadata) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FunctionARN":
            if err := json.Unmarshal([]byte(v), &strct.FunctionARN); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AutoPublish" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoPublish\": ")
	if tmp, err := json.Marshal(strct.AutoPublish); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FunctionARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionARN\": ")
	if tmp, err := json.Marshal(strct.FunctionARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FunctionCode" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FunctionCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionCode\": ")
	if tmp, err := json.Marshal(strct.FunctionCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FunctionConfig" field is required
    if strct.FunctionConfig == nil {
        return nil, errors.New("FunctionConfig is a required field")
    }
    // Marshal the "FunctionConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionConfig\": ")
	if tmp, err := json.Marshal(strct.FunctionConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FunctionMetadata" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionMetadata\": ")
	if tmp, err := json.Marshal(strct.FunctionMetadata); err != nil {
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
    // Marshal the "Stage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Stage\": ")
	if tmp, err := json.Marshal(strct.Stage); err != nil {
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
    FunctionCodeReceived := false
    FunctionConfigReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AutoPublish":
            if err := json.Unmarshal([]byte(v), &strct.AutoPublish); err != nil {
                return err
             }
        case "FunctionARN":
            if err := json.Unmarshal([]byte(v), &strct.FunctionARN); err != nil {
                return err
             }
        case "FunctionCode":
            if err := json.Unmarshal([]byte(v), &strct.FunctionCode); err != nil {
                return err
             }
            FunctionCodeReceived = true
        case "FunctionConfig":
            if err := json.Unmarshal([]byte(v), &strct.FunctionConfig); err != nil {
                return err
             }
            FunctionConfigReceived = true
        case "FunctionMetadata":
            if err := json.Unmarshal([]byte(v), &strct.FunctionMetadata); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Stage":
            if err := json.Unmarshal([]byte(v), &strct.Stage); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FunctionCode (a required property) was received
    if !FunctionCodeReceived {
        return errors.New("\"FunctionCode\" is required but was not present")
    }
    // check if FunctionConfig (a required property) was received
    if !FunctionConfigReceived {
        return errors.New("\"FunctionConfig\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}
