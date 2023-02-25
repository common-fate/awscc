// Code generated by schema-generate. DO NOT EDIT.

package configurationseteventdestination

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// CloudWatchDestination An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
type CloudWatchDestination struct {

  // A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
  DimensionConfigurations []*DimensionConfiguration `json:"DimensionConfigurations,omitempty"`
}

// DimensionConfiguration A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
type DimensionConfiguration struct {

  // The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.
  DefaultDimensionValue string `json:"DefaultDimensionValue"`

  // The name of an Amazon CloudWatch dimension associated with an email sending metric.
  DimensionName string `json:"DimensionName"`

  // The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.
  DimensionValueSource string `json:"DimensionValueSource"`
}

// EventDestination 
type EventDestination struct {

  // An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
  CloudWatchDestination *CloudWatchDestination `json:"CloudWatchDestination,omitempty"`

  // Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   
  Enabled bool `json:"Enabled,omitempty"`

  // An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
  KinesisFirehoseDestination *KinesisFirehoseDestination `json:"KinesisFirehoseDestination,omitempty"`

  // The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure, deliveryDelay, and subscription.
  MatchingEventTypes []string `json:"MatchingEventTypes"`

  // The name of the event destination set.
  Name string `json:"Name,omitempty"`

  // An object that contains SNS topic ARN associated event destination.
  SnsDestination *SnsDestination `json:"SnsDestination,omitempty"`
}

// KinesisFirehoseDestination An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
type KinesisFirehoseDestination struct {

  // The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.
  DeliveryStreamARN string `json:"DeliveryStreamARN"`

  // The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.
  IAMRoleARN string `json:"IAMRoleARN"`
}

// Resource Resource Type definition for AWS::SES::ConfigurationSetEventDestination
type Resource struct {

  // The name of the configuration set that contains the event destination.
  ConfigurationSetName string `json:"ConfigurationSetName"`

  // The event destination object.
  EventDestination *EventDestination `json:"EventDestination"`
  Id string `json:"Id,omitempty"`
}

// SnsDestination An object that contains SNS topic ARN associated event destination.
type SnsDestination struct {
  TopicARN string `json:"TopicARN"`
}

func (strct *CloudWatchDestination) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DimensionConfigurations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DimensionConfigurations\": ")
	if tmp, err := json.Marshal(strct.DimensionConfigurations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CloudWatchDestination) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DimensionConfigurations":
            if err := json.Unmarshal([]byte(v), &strct.DimensionConfigurations); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *DimensionConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "DefaultDimensionValue" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DefaultDimensionValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultDimensionValue\": ")
	if tmp, err := json.Marshal(strct.DefaultDimensionValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DimensionName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DimensionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DimensionName\": ")
	if tmp, err := json.Marshal(strct.DimensionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DimensionValueSource" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DimensionValueSource" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DimensionValueSource\": ")
	if tmp, err := json.Marshal(strct.DimensionValueSource); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DimensionConfiguration) UnmarshalJSON(b []byte) error {
    DefaultDimensionValueReceived := false
    DimensionNameReceived := false
    DimensionValueSourceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DefaultDimensionValue":
            if err := json.Unmarshal([]byte(v), &strct.DefaultDimensionValue); err != nil {
                return err
             }
            DefaultDimensionValueReceived = true
        case "DimensionName":
            if err := json.Unmarshal([]byte(v), &strct.DimensionName); err != nil {
                return err
             }
            DimensionNameReceived = true
        case "DimensionValueSource":
            if err := json.Unmarshal([]byte(v), &strct.DimensionValueSource); err != nil {
                return err
             }
            DimensionValueSourceReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DefaultDimensionValue (a required property) was received
    if !DefaultDimensionValueReceived {
        return errors.New("\"DefaultDimensionValue\" is required but was not present")
    }
    // check if DimensionName (a required property) was received
    if !DimensionNameReceived {
        return errors.New("\"DimensionName\" is required but was not present")
    }
    // check if DimensionValueSource (a required property) was received
    if !DimensionValueSourceReceived {
        return errors.New("\"DimensionValueSource\" is required but was not present")
    }
    return nil
}

func (strct *EventDestination) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CloudWatchDestination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CloudWatchDestination\": ")
	if tmp, err := json.Marshal(strct.CloudWatchDestination); err != nil {
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
    // Marshal the "KinesisFirehoseDestination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KinesisFirehoseDestination\": ")
	if tmp, err := json.Marshal(strct.KinesisFirehoseDestination); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MatchingEventTypes" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "MatchingEventTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MatchingEventTypes\": ")
	if tmp, err := json.Marshal(strct.MatchingEventTypes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "SnsDestination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SnsDestination\": ")
	if tmp, err := json.Marshal(strct.SnsDestination); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EventDestination) UnmarshalJSON(b []byte) error {
    MatchingEventTypesReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CloudWatchDestination":
            if err := json.Unmarshal([]byte(v), &strct.CloudWatchDestination); err != nil {
                return err
             }
        case "Enabled":
            if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
                return err
             }
        case "KinesisFirehoseDestination":
            if err := json.Unmarshal([]byte(v), &strct.KinesisFirehoseDestination); err != nil {
                return err
             }
        case "MatchingEventTypes":
            if err := json.Unmarshal([]byte(v), &strct.MatchingEventTypes); err != nil {
                return err
             }
            MatchingEventTypesReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "SnsDestination":
            if err := json.Unmarshal([]byte(v), &strct.SnsDestination); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MatchingEventTypes (a required property) was received
    if !MatchingEventTypesReceived {
        return errors.New("\"MatchingEventTypes\" is required but was not present")
    }
    return nil
}

func (strct *KinesisFirehoseDestination) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "DeliveryStreamARN" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DeliveryStreamARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeliveryStreamARN\": ")
	if tmp, err := json.Marshal(strct.DeliveryStreamARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IAMRoleARN" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "IAMRoleARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IAMRoleARN\": ")
	if tmp, err := json.Marshal(strct.IAMRoleARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *KinesisFirehoseDestination) UnmarshalJSON(b []byte) error {
    DeliveryStreamARNReceived := false
    IAMRoleARNReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DeliveryStreamARN":
            if err := json.Unmarshal([]byte(v), &strct.DeliveryStreamARN); err != nil {
                return err
             }
            DeliveryStreamARNReceived = true
        case "IAMRoleARN":
            if err := json.Unmarshal([]byte(v), &strct.IAMRoleARN); err != nil {
                return err
             }
            IAMRoleARNReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DeliveryStreamARN (a required property) was received
    if !DeliveryStreamARNReceived {
        return errors.New("\"DeliveryStreamARN\" is required but was not present")
    }
    // check if IAMRoleARN (a required property) was received
    if !IAMRoleARNReceived {
        return errors.New("\"IAMRoleARN\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ConfigurationSetName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConfigurationSetName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConfigurationSetName\": ")
	if tmp, err := json.Marshal(strct.ConfigurationSetName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "EventDestination" field is required
    if strct.EventDestination == nil {
        return nil, errors.New("EventDestination is a required field")
    }
    // Marshal the "EventDestination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventDestination\": ")
	if tmp, err := json.Marshal(strct.EventDestination); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    ConfigurationSetNameReceived := false
    EventDestinationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ConfigurationSetName":
            if err := json.Unmarshal([]byte(v), &strct.ConfigurationSetName); err != nil {
                return err
             }
            ConfigurationSetNameReceived = true
        case "EventDestination":
            if err := json.Unmarshal([]byte(v), &strct.EventDestination); err != nil {
                return err
             }
            EventDestinationReceived = true
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ConfigurationSetName (a required property) was received
    if !ConfigurationSetNameReceived {
        return errors.New("\"ConfigurationSetName\" is required but was not present")
    }
    // check if EventDestination (a required property) was received
    if !EventDestinationReceived {
        return errors.New("\"EventDestination\" is required but was not present")
    }
    return nil
}

func (strct *SnsDestination) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "TopicARN" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TopicARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TopicARN\": ")
	if tmp, err := json.Marshal(strct.TopicARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SnsDestination) UnmarshalJSON(b []byte) error {
    TopicARNReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "TopicARN":
            if err := json.Unmarshal([]byte(v), &strct.TopicARN); err != nil {
                return err
             }
            TopicARNReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if TopicARN (a required property) was received
    if !TopicARNReceived {
        return errors.New("\"TopicARN\" is required but was not present")
    }
    return nil
}
