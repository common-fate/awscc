// Code generated by schema-generate. DO NOT EDIT.

package custommetric

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource A custom metric published by your devices to Device Defender.
type Resource struct {

  // Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
  DisplayName string `json:"DisplayName,omitempty"`

  // The Amazon Resource Number (ARN) of the custom metric.
  MetricArn string `json:"MetricArn,omitempty"`

  // The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
  MetricName string `json:"MetricName,omitempty"`

  // The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
  MetricType string `json:"MetricType"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The tag's key.
  Key string `json:"Key"`

  // The tag's value.
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DisplayName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DisplayName\": ")
	if tmp, err := json.Marshal(strct.DisplayName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MetricArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MetricArn\": ")
	if tmp, err := json.Marshal(strct.MetricArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MetricName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MetricName\": ")
	if tmp, err := json.Marshal(strct.MetricName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MetricType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "MetricType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MetricType\": ")
	if tmp, err := json.Marshal(strct.MetricType); err != nil {
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
    MetricTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DisplayName":
            if err := json.Unmarshal([]byte(v), &strct.DisplayName); err != nil {
                return err
             }
        case "MetricArn":
            if err := json.Unmarshal([]byte(v), &strct.MetricArn); err != nil {
                return err
             }
        case "MetricName":
            if err := json.Unmarshal([]byte(v), &strct.MetricName); err != nil {
                return err
             }
        case "MetricType":
            if err := json.Unmarshal([]byte(v), &strct.MetricType); err != nil {
                return err
             }
            MetricTypeReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MetricType (a required property) was received
    if !MetricTypeReceived {
        return errors.New("\"MetricType\" is required but was not present")
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