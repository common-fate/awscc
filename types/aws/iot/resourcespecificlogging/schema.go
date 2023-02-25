// Code generated by schema-generate. DO NOT EDIT.

package resourcespecificlogging

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource-specific logging allows you to specify a logging level for a specific thing group.
type Resource struct {

  // The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
  LogLevel string `json:"LogLevel"`

  // Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.
  TargetId string `json:"TargetId,omitempty"`

  // The target name.
  TargetName string `json:"TargetName"`

  // The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID.
  TargetType string `json:"TargetType"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "LogLevel" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "LogLevel" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LogLevel\": ")
	if tmp, err := json.Marshal(strct.LogLevel); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TargetId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetId\": ")
	if tmp, err := json.Marshal(strct.TargetId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TargetName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetName\": ")
	if tmp, err := json.Marshal(strct.TargetName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TargetType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetType\": ")
	if tmp, err := json.Marshal(strct.TargetType); err != nil {
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
    LogLevelReceived := false
    TargetNameReceived := false
    TargetTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LogLevel":
            if err := json.Unmarshal([]byte(v), &strct.LogLevel); err != nil {
                return err
             }
            LogLevelReceived = true
        case "TargetId":
            if err := json.Unmarshal([]byte(v), &strct.TargetId); err != nil {
                return err
             }
        case "TargetName":
            if err := json.Unmarshal([]byte(v), &strct.TargetName); err != nil {
                return err
             }
            TargetNameReceived = true
        case "TargetType":
            if err := json.Unmarshal([]byte(v), &strct.TargetType); err != nil {
                return err
             }
            TargetTypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if LogLevel (a required property) was received
    if !LogLevelReceived {
        return errors.New("\"LogLevel\" is required but was not present")
    }
    // check if TargetName (a required property) was received
    if !TargetNameReceived {
        return errors.New("\"TargetName\" is required but was not present")
    }
    // check if TargetType (a required property) was received
    if !TargetTypeReceived {
        return errors.New("\"TargetType\" is required but was not present")
    }
    return nil
}
