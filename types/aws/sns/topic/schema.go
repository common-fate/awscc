// Code generated by schema-generate. DO NOT EDIT.

package topic

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// DataProtectionPolicy The body of the policy document you want to use for this topic.
// 
// You can only add one policy per topic.
// 
// The policy must be in JSON string format.
// 
// Length Constraints: Maximum length of 30720
type DataProtectionPolicy struct {
}

// Resource Resource Type definition for AWS::SNS::Topic
type Resource struct {

  // Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
  // 
  // When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
  // 
  // (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
  // 
  // 
  ContentBasedDeduplication bool `json:"ContentBasedDeduplication,omitempty"`

  // The body of the policy document you want to use for this topic.
  // 
  // You can only add one policy per topic.
  // 
  // The policy must be in JSON string format.
  // 
  // Length Constraints: Maximum length of 30720
  DataProtectionPolicy *DataProtectionPolicy `json:"DataProtectionPolicy,omitempty"`

  // The display name to use for an Amazon SNS topic with SMS subscriptions.
  DisplayName string `json:"DisplayName,omitempty"`

  // Set to true to create a FIFO topic.
  FifoTopic bool `json:"FifoTopic,omitempty"`

  // The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
  // 
  // This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
  KmsMasterKeyId string `json:"KmsMasterKeyId,omitempty"`

  // Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
  SignatureVersion string `json:"SignatureVersion,omitempty"`

  // The SNS subscriptions (endpoints) for this topic.
  Subscription []*Subscription `json:"Subscription,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
  TopicArn string `json:"TopicArn,omitempty"`

  // The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
  // 
  // If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
  TopicName string `json:"TopicName,omitempty"`

  // Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
  TracingConfig string `json:"TracingConfig,omitempty"`
}

// Subscription 
type Subscription struct {
  Endpoint string `json:"Endpoint"`
  Protocol string `json:"Protocol"`
}

// Tag 
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 characters in length.
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ContentBasedDeduplication" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContentBasedDeduplication\": ")
	if tmp, err := json.Marshal(strct.ContentBasedDeduplication); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DataProtectionPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataProtectionPolicy\": ")
	if tmp, err := json.Marshal(strct.DataProtectionPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "FifoTopic" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FifoTopic\": ")
	if tmp, err := json.Marshal(strct.FifoTopic); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KmsMasterKeyId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KmsMasterKeyId\": ")
	if tmp, err := json.Marshal(strct.KmsMasterKeyId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SignatureVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SignatureVersion\": ")
	if tmp, err := json.Marshal(strct.SignatureVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Subscription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subscription\": ")
	if tmp, err := json.Marshal(strct.Subscription); err != nil {
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
    // Marshal the "TopicArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TopicArn\": ")
	if tmp, err := json.Marshal(strct.TopicArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TopicName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TopicName\": ")
	if tmp, err := json.Marshal(strct.TopicName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TracingConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TracingConfig\": ")
	if tmp, err := json.Marshal(strct.TracingConfig); err != nil {
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
        case "ContentBasedDeduplication":
            if err := json.Unmarshal([]byte(v), &strct.ContentBasedDeduplication); err != nil {
                return err
             }
        case "DataProtectionPolicy":
            if err := json.Unmarshal([]byte(v), &strct.DataProtectionPolicy); err != nil {
                return err
             }
        case "DisplayName":
            if err := json.Unmarshal([]byte(v), &strct.DisplayName); err != nil {
                return err
             }
        case "FifoTopic":
            if err := json.Unmarshal([]byte(v), &strct.FifoTopic); err != nil {
                return err
             }
        case "KmsMasterKeyId":
            if err := json.Unmarshal([]byte(v), &strct.KmsMasterKeyId); err != nil {
                return err
             }
        case "SignatureVersion":
            if err := json.Unmarshal([]byte(v), &strct.SignatureVersion); err != nil {
                return err
             }
        case "Subscription":
            if err := json.Unmarshal([]byte(v), &strct.Subscription); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TopicArn":
            if err := json.Unmarshal([]byte(v), &strct.TopicArn); err != nil {
                return err
             }
        case "TopicName":
            if err := json.Unmarshal([]byte(v), &strct.TopicName); err != nil {
                return err
             }
        case "TracingConfig":
            if err := json.Unmarshal([]byte(v), &strct.TracingConfig); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Subscription) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Endpoint" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Endpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Endpoint\": ")
	if tmp, err := json.Marshal(strct.Endpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Protocol" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Protocol" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Protocol\": ")
	if tmp, err := json.Marshal(strct.Protocol); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Subscription) UnmarshalJSON(b []byte) error {
    EndpointReceived := false
    ProtocolReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
            EndpointReceived = true
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
            ProtocolReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Endpoint (a required property) was received
    if !EndpointReceived {
        return errors.New("\"Endpoint\" is required but was not present")
    }
    // check if Protocol (a required property) was received
    if !ProtocolReceived {
        return errors.New("\"Protocol\" is required but was not present")
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
