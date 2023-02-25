// Code generated by schema-generate. DO NOT EDIT.

package profilinggroup

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// AgentPermissions The agent permissions attached to this profiling group.
type AgentPermissions struct {

  // The principals for the agent permissions.
  Principals []string `json:"Principals"`
}

// Channel Notification medium for users to get alerted for events that occur in application profile. We support SNS topic as a notification channel.
type Channel struct {
  ChannelId string `json:"channelId,omitempty"`
  ChannelUri string `json:"channelUri"`
}

// Resource This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.
type Resource struct {

  // The agent permissions attached to this profiling group.
  AgentPermissions *AgentPermissions `json:"AgentPermissions,omitempty"`

  // Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency
  AnomalyDetectionNotificationConfiguration []*Channel `json:"AnomalyDetectionNotificationConfiguration,omitempty"`

  // The Amazon Resource Name (ARN) of the specified profiling group.
  Arn string `json:"Arn,omitempty"`

  // The compute platform of the profiling group.
  ComputePlatform string `json:"ComputePlatform,omitempty"`

  // The name of the profiling group.
  ProfilingGroupName string `json:"ProfilingGroupName"`

  // The tags associated with a profiling group.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
  Value string `json:"Value"`
}

func (strct *AgentPermissions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Principals" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Principals" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Principals\": ")
	if tmp, err := json.Marshal(strct.Principals); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AgentPermissions) UnmarshalJSON(b []byte) error {
    PrincipalsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Principals":
            if err := json.Unmarshal([]byte(v), &strct.Principals); err != nil {
                return err
             }
            PrincipalsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Principals (a required property) was received
    if !PrincipalsReceived {
        return errors.New("\"Principals\" is required but was not present")
    }
    return nil
}

func (strct *Channel) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "channelId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"channelId\": ")
	if tmp, err := json.Marshal(strct.ChannelId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ChannelUri" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "channelUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"channelUri\": ")
	if tmp, err := json.Marshal(strct.ChannelUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Channel) UnmarshalJSON(b []byte) error {
    channelUriReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "channelId":
            if err := json.Unmarshal([]byte(v), &strct.ChannelId); err != nil {
                return err
             }
        case "channelUri":
            if err := json.Unmarshal([]byte(v), &strct.ChannelUri); err != nil {
                return err
             }
            channelUriReceived = true
        }
    }
    // check if channelUri (a required property) was received
    if !channelUriReceived {
        return errors.New("\"channelUri\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AgentPermissions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AgentPermissions\": ")
	if tmp, err := json.Marshal(strct.AgentPermissions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AnomalyDetectionNotificationConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AnomalyDetectionNotificationConfiguration\": ")
	if tmp, err := json.Marshal(strct.AnomalyDetectionNotificationConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "ComputePlatform" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComputePlatform\": ")
	if tmp, err := json.Marshal(strct.ComputePlatform); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ProfilingGroupName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ProfilingGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfilingGroupName\": ")
	if tmp, err := json.Marshal(strct.ProfilingGroupName); err != nil {
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
    ProfilingGroupNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AgentPermissions":
            if err := json.Unmarshal([]byte(v), &strct.AgentPermissions); err != nil {
                return err
             }
        case "AnomalyDetectionNotificationConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.AnomalyDetectionNotificationConfiguration); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "ComputePlatform":
            if err := json.Unmarshal([]byte(v), &strct.ComputePlatform); err != nil {
                return err
             }
        case "ProfilingGroupName":
            if err := json.Unmarshal([]byte(v), &strct.ProfilingGroupName); err != nil {
                return err
             }
            ProfilingGroupNameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ProfilingGroupName (a required property) was received
    if !ProfilingGroupNameReceived {
        return errors.New("\"ProfilingGroupName\" is required but was not present")
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
