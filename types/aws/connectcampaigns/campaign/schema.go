// Code generated by schema-generate. DO NOT EDIT.

package campaign

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AnswerMachineDetectionConfig The configuration used for answering machine detection during outbound calls
type AnswerMachineDetectionConfig struct {

  // Flag to decided whether outbound calls should have answering machine detection enabled or not
  EnableAnswerMachineDetection bool `json:"EnableAnswerMachineDetection"`
}

// DialerConfig The possible types of dialer config parameters
type DialerConfig struct {
  PredictiveDialerConfig *PredictiveDialerConfig `json:"PredictiveDialerConfig,omitempty"`
  ProgressiveDialerConfig *ProgressiveDialerConfig `json:"ProgressiveDialerConfig,omitempty"`
}

// OutboundCallConfig The configuration used for outbound calls.
type OutboundCallConfig struct {

  // The configuration used for answering machine detection during outbound calls
  AnswerMachineDetectionConfig *AnswerMachineDetectionConfig `json:"AnswerMachineDetectionConfig,omitempty"`

  // The identifier of the contact flow for the outbound call.
  ConnectContactFlowArn string `json:"ConnectContactFlowArn"`

  // The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.
  ConnectQueueArn string `json:"ConnectQueueArn"`

  // The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.
  ConnectSourcePhoneNumber string `json:"ConnectSourcePhoneNumber,omitempty"`
}

// PredictiveDialerConfig Predictive Dialer config
type PredictiveDialerConfig struct {

  // The bandwidth allocation of a queue resource.
  BandwidthAllocation float64 `json:"BandwidthAllocation"`
}

// ProgressiveDialerConfig Progressive Dialer config
type ProgressiveDialerConfig struct {

  // The bandwidth allocation of a queue resource.
  BandwidthAllocation float64 `json:"BandwidthAllocation"`
}

// Resource Definition of AWS::ConnectCampaigns::Campaign Resource Type
type Resource struct {

  // Amazon Connect Campaign Arn
  Arn string `json:"Arn,omitempty"`

  // Amazon Connect Instance Arn
  ConnectInstanceArn string `json:"ConnectInstanceArn"`
  DialerConfig *DialerConfig `json:"DialerConfig"`

  // Amazon Connect Campaign Name
  Name string `json:"Name"`
  OutboundCallConfig *OutboundCallConfig `json:"OutboundCallConfig"`

  // One or more tags.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that's 1 to 256 characters in length.
  Value string `json:"Value"`
}

func (strct *AnswerMachineDetectionConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "EnableAnswerMachineDetection" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EnableAnswerMachineDetection" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableAnswerMachineDetection\": ")
	if tmp, err := json.Marshal(strct.EnableAnswerMachineDetection); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AnswerMachineDetectionConfig) UnmarshalJSON(b []byte) error {
    EnableAnswerMachineDetectionReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EnableAnswerMachineDetection":
            if err := json.Unmarshal([]byte(v), &strct.EnableAnswerMachineDetection); err != nil {
                return err
             }
            EnableAnswerMachineDetectionReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if EnableAnswerMachineDetection (a required property) was received
    if !EnableAnswerMachineDetectionReceived {
        return errors.New("\"EnableAnswerMachineDetection\" is required but was not present")
    }
    return nil
}

func (strct *DialerConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "PredictiveDialerConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PredictiveDialerConfig\": ")
	if tmp, err := json.Marshal(strct.PredictiveDialerConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProgressiveDialerConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProgressiveDialerConfig\": ")
	if tmp, err := json.Marshal(strct.ProgressiveDialerConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DialerConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PredictiveDialerConfig":
            if err := json.Unmarshal([]byte(v), &strct.PredictiveDialerConfig); err != nil {
                return err
             }
        case "ProgressiveDialerConfig":
            if err := json.Unmarshal([]byte(v), &strct.ProgressiveDialerConfig); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *OutboundCallConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AnswerMachineDetectionConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AnswerMachineDetectionConfig\": ")
	if tmp, err := json.Marshal(strct.AnswerMachineDetectionConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ConnectContactFlowArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConnectContactFlowArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectContactFlowArn\": ")
	if tmp, err := json.Marshal(strct.ConnectContactFlowArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ConnectQueueArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConnectQueueArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectQueueArn\": ")
	if tmp, err := json.Marshal(strct.ConnectQueueArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ConnectSourcePhoneNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectSourcePhoneNumber\": ")
	if tmp, err := json.Marshal(strct.ConnectSourcePhoneNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *OutboundCallConfig) UnmarshalJSON(b []byte) error {
    ConnectContactFlowArnReceived := false
    ConnectQueueArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AnswerMachineDetectionConfig":
            if err := json.Unmarshal([]byte(v), &strct.AnswerMachineDetectionConfig); err != nil {
                return err
             }
        case "ConnectContactFlowArn":
            if err := json.Unmarshal([]byte(v), &strct.ConnectContactFlowArn); err != nil {
                return err
             }
            ConnectContactFlowArnReceived = true
        case "ConnectQueueArn":
            if err := json.Unmarshal([]byte(v), &strct.ConnectQueueArn); err != nil {
                return err
             }
            ConnectQueueArnReceived = true
        case "ConnectSourcePhoneNumber":
            if err := json.Unmarshal([]byte(v), &strct.ConnectSourcePhoneNumber); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ConnectContactFlowArn (a required property) was received
    if !ConnectContactFlowArnReceived {
        return errors.New("\"ConnectContactFlowArn\" is required but was not present")
    }
    // check if ConnectQueueArn (a required property) was received
    if !ConnectQueueArnReceived {
        return errors.New("\"ConnectQueueArn\" is required but was not present")
    }
    return nil
}

func (strct *PredictiveDialerConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "BandwidthAllocation" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BandwidthAllocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BandwidthAllocation\": ")
	if tmp, err := json.Marshal(strct.BandwidthAllocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PredictiveDialerConfig) UnmarshalJSON(b []byte) error {
    BandwidthAllocationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BandwidthAllocation":
            if err := json.Unmarshal([]byte(v), &strct.BandwidthAllocation); err != nil {
                return err
             }
            BandwidthAllocationReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if BandwidthAllocation (a required property) was received
    if !BandwidthAllocationReceived {
        return errors.New("\"BandwidthAllocation\" is required but was not present")
    }
    return nil
}

func (strct *ProgressiveDialerConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "BandwidthAllocation" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BandwidthAllocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BandwidthAllocation\": ")
	if tmp, err := json.Marshal(strct.BandwidthAllocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ProgressiveDialerConfig) UnmarshalJSON(b []byte) error {
    BandwidthAllocationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BandwidthAllocation":
            if err := json.Unmarshal([]byte(v), &strct.BandwidthAllocation); err != nil {
                return err
             }
            BandwidthAllocationReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if BandwidthAllocation (a required property) was received
    if !BandwidthAllocationReceived {
        return errors.New("\"BandwidthAllocation\" is required but was not present")
    }
    return nil
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
    // "ConnectInstanceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConnectInstanceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConnectInstanceArn\": ")
	if tmp, err := json.Marshal(strct.ConnectInstanceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DialerConfig" field is required
    if strct.DialerConfig == nil {
        return nil, errors.New("DialerConfig is a required field")
    }
    // Marshal the "DialerConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DialerConfig\": ")
	if tmp, err := json.Marshal(strct.DialerConfig); err != nil {
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
    // "OutboundCallConfig" field is required
    if strct.OutboundCallConfig == nil {
        return nil, errors.New("OutboundCallConfig is a required field")
    }
    // Marshal the "OutboundCallConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutboundCallConfig\": ")
	if tmp, err := json.Marshal(strct.OutboundCallConfig); err != nil {
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
    ConnectInstanceArnReceived := false
    DialerConfigReceived := false
    NameReceived := false
    OutboundCallConfigReceived := false
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
        case "ConnectInstanceArn":
            if err := json.Unmarshal([]byte(v), &strct.ConnectInstanceArn); err != nil {
                return err
             }
            ConnectInstanceArnReceived = true
        case "DialerConfig":
            if err := json.Unmarshal([]byte(v), &strct.DialerConfig); err != nil {
                return err
             }
            DialerConfigReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "OutboundCallConfig":
            if err := json.Unmarshal([]byte(v), &strct.OutboundCallConfig); err != nil {
                return err
             }
            OutboundCallConfigReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ConnectInstanceArn (a required property) was received
    if !ConnectInstanceArnReceived {
        return errors.New("\"ConnectInstanceArn\" is required but was not present")
    }
    // check if DialerConfig (a required property) was received
    if !DialerConfigReceived {
        return errors.New("\"DialerConfig\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if OutboundCallConfig (a required property) was received
    if !OutboundCallConfigReceived {
        return errors.New("\"OutboundCallConfig\" is required but was not present")
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
