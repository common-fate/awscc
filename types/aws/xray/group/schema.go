// Code generated by schema-generate. DO NOT EDIT.

package group

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// InsightsConfiguration 
type InsightsConfiguration struct {

  // Set the InsightsEnabled value to true to enable insights or false to disable insights.
  InsightsEnabled bool `json:"InsightsEnabled,omitempty"`

  // Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.
  NotificationsEnabled bool `json:"NotificationsEnabled,omitempty"`
}

// Resource This schema provides construct and validation rules for AWS-XRay Group resource parameters.
type Resource struct {

  // The filter expression defining criteria by which to group traces.
  FilterExpression string `json:"FilterExpression,omitempty"`

  // The ARN of the group that was generated on creation.
  GroupARN string `json:"GroupARN,omitempty"`

  // The case-sensitive name of the new group. Names must be unique.
  GroupName string `json:"GroupName,omitempty"`
  InsightsConfiguration *InsightsConfiguration `json:"InsightsConfiguration,omitempty"`
  Tags []*TagsItems `json:"Tags,omitempty"`
}

// TagsItems 
type TagsItems struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *InsightsConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "InsightsEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InsightsEnabled\": ")
	if tmp, err := json.Marshal(strct.InsightsEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NotificationsEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NotificationsEnabled\": ")
	if tmp, err := json.Marshal(strct.NotificationsEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *InsightsConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "InsightsEnabled":
            if err := json.Unmarshal([]byte(v), &strct.InsightsEnabled); err != nil {
                return err
             }
        case "NotificationsEnabled":
            if err := json.Unmarshal([]byte(v), &strct.NotificationsEnabled); err != nil {
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
    // Marshal the "FilterExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FilterExpression\": ")
	if tmp, err := json.Marshal(strct.FilterExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GroupARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GroupARN\": ")
	if tmp, err := json.Marshal(strct.GroupARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GroupName\": ")
	if tmp, err := json.Marshal(strct.GroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InsightsConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InsightsConfiguration\": ")
	if tmp, err := json.Marshal(strct.InsightsConfiguration); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FilterExpression":
            if err := json.Unmarshal([]byte(v), &strct.FilterExpression); err != nil {
                return err
             }
        case "GroupARN":
            if err := json.Unmarshal([]byte(v), &strct.GroupARN); err != nil {
                return err
             }
        case "GroupName":
            if err := json.Unmarshal([]byte(v), &strct.GroupName); err != nil {
                return err
             }
        case "InsightsConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.InsightsConfiguration); err != nil {
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
    return nil
}

func (strct *TagsItems) MarshalJSON() ([]byte, error) {
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

func (strct *TagsItems) UnmarshalJSON(b []byte) error {
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
