// Code generated by schema-generate. DO NOT EDIT.

package connectattachment

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ConnectAttachmentOptions Connect attachment options for protocol
type ConnectAttachmentOptions struct {

  // Tunnel protocol for connect attachment
  Protocol string `json:"Protocol,omitempty"`
}

// ProposedSegmentChange A key-value pair to associate with a resource.
type ProposedSegmentChange struct {

  // New policy rule number of the attachment
  AttachmentPolicyRuleNumber int `json:"AttachmentPolicyRuleNumber,omitempty"`

  // Proposed segment name
  SegmentName string `json:"SegmentName,omitempty"`

  // Proposed tags for the Segment.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Resource AWS::NetworkManager::ConnectAttachment Resource Type Definition
type Resource struct {

  // The ID of the attachment.
  AttachmentId string `json:"AttachmentId,omitempty"`

  // The policy rule number associated with the attachment.
  AttachmentPolicyRuleNumber int `json:"AttachmentPolicyRuleNumber,omitempty"`

  // The type of attachment.
  AttachmentType string `json:"AttachmentType,omitempty"`

  // The ARN of a core network for the VPC attachment.
  CoreNetworkArn string `json:"CoreNetworkArn,omitempty"`

  // ID of the CoreNetwork that the attachment will be attached to.
  CoreNetworkId string `json:"CoreNetworkId"`

  // Creation time of the attachment.
  CreatedAt string `json:"CreatedAt,omitempty"`

  // Edge location of the attachment.
  EdgeLocation string `json:"EdgeLocation"`

  // Protocol options for connect attachment
  Options *ConnectAttachmentOptions `json:"Options"`

  // The ID of the attachment account owner.
  OwnerAccountId string `json:"OwnerAccountId,omitempty"`

  // The attachment to move from one segment to another.
  ProposedSegmentChange *ProposedSegmentChange `json:"ProposedSegmentChange,omitempty"`

  // The attachment resource ARN.
  ResourceArn string `json:"ResourceArn,omitempty"`

  // The name of the segment attachment.
  SegmentName string `json:"SegmentName,omitempty"`

  // State of the attachment.
  State string `json:"State,omitempty"`

  // Tags for the attachment.
  Tags []*Tag `json:"Tags,omitempty"`

  // Id of transport attachment
  TransportAttachmentId string `json:"TransportAttachmentId"`

  // Last update time of the attachment.
  UpdatedAt string `json:"UpdatedAt,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *ConnectAttachmentOptions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *ConnectAttachmentOptions) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ProposedSegmentChange) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AttachmentPolicyRuleNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachmentPolicyRuleNumber\": ")
	if tmp, err := json.Marshal(strct.AttachmentPolicyRuleNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SegmentName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SegmentName\": ")
	if tmp, err := json.Marshal(strct.SegmentName); err != nil {
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

func (strct *ProposedSegmentChange) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AttachmentPolicyRuleNumber":
            if err := json.Unmarshal([]byte(v), &strct.AttachmentPolicyRuleNumber); err != nil {
                return err
             }
        case "SegmentName":
            if err := json.Unmarshal([]byte(v), &strct.SegmentName); err != nil {
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

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AttachmentId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachmentId\": ")
	if tmp, err := json.Marshal(strct.AttachmentId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AttachmentPolicyRuleNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachmentPolicyRuleNumber\": ")
	if tmp, err := json.Marshal(strct.AttachmentPolicyRuleNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AttachmentType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachmentType\": ")
	if tmp, err := json.Marshal(strct.AttachmentType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CoreNetworkArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CoreNetworkArn\": ")
	if tmp, err := json.Marshal(strct.CoreNetworkArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "CoreNetworkId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "CoreNetworkId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CoreNetworkId\": ")
	if tmp, err := json.Marshal(strct.CoreNetworkId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CreatedAt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreatedAt\": ")
	if tmp, err := json.Marshal(strct.CreatedAt); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "EdgeLocation" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EdgeLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EdgeLocation\": ")
	if tmp, err := json.Marshal(strct.EdgeLocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Options" field is required
    if strct.Options == nil {
        return nil, errors.New("Options is a required field")
    }
    // Marshal the "Options" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Options\": ")
	if tmp, err := json.Marshal(strct.Options); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OwnerAccountId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OwnerAccountId\": ")
	if tmp, err := json.Marshal(strct.OwnerAccountId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProposedSegmentChange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProposedSegmentChange\": ")
	if tmp, err := json.Marshal(strct.ProposedSegmentChange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceArn\": ")
	if tmp, err := json.Marshal(strct.ResourceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SegmentName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SegmentName\": ")
	if tmp, err := json.Marshal(strct.SegmentName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "State" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"State\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
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
    // "TransportAttachmentId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TransportAttachmentId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TransportAttachmentId\": ")
	if tmp, err := json.Marshal(strct.TransportAttachmentId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UpdatedAt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UpdatedAt\": ")
	if tmp, err := json.Marshal(strct.UpdatedAt); err != nil {
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
    CoreNetworkIdReceived := false
    EdgeLocationReceived := false
    OptionsReceived := false
    TransportAttachmentIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AttachmentId":
            if err := json.Unmarshal([]byte(v), &strct.AttachmentId); err != nil {
                return err
             }
        case "AttachmentPolicyRuleNumber":
            if err := json.Unmarshal([]byte(v), &strct.AttachmentPolicyRuleNumber); err != nil {
                return err
             }
        case "AttachmentType":
            if err := json.Unmarshal([]byte(v), &strct.AttachmentType); err != nil {
                return err
             }
        case "CoreNetworkArn":
            if err := json.Unmarshal([]byte(v), &strct.CoreNetworkArn); err != nil {
                return err
             }
        case "CoreNetworkId":
            if err := json.Unmarshal([]byte(v), &strct.CoreNetworkId); err != nil {
                return err
             }
            CoreNetworkIdReceived = true
        case "CreatedAt":
            if err := json.Unmarshal([]byte(v), &strct.CreatedAt); err != nil {
                return err
             }
        case "EdgeLocation":
            if err := json.Unmarshal([]byte(v), &strct.EdgeLocation); err != nil {
                return err
             }
            EdgeLocationReceived = true
        case "Options":
            if err := json.Unmarshal([]byte(v), &strct.Options); err != nil {
                return err
             }
            OptionsReceived = true
        case "OwnerAccountId":
            if err := json.Unmarshal([]byte(v), &strct.OwnerAccountId); err != nil {
                return err
             }
        case "ProposedSegmentChange":
            if err := json.Unmarshal([]byte(v), &strct.ProposedSegmentChange); err != nil {
                return err
             }
        case "ResourceArn":
            if err := json.Unmarshal([]byte(v), &strct.ResourceArn); err != nil {
                return err
             }
        case "SegmentName":
            if err := json.Unmarshal([]byte(v), &strct.SegmentName); err != nil {
                return err
             }
        case "State":
            if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TransportAttachmentId":
            if err := json.Unmarshal([]byte(v), &strct.TransportAttachmentId); err != nil {
                return err
             }
            TransportAttachmentIdReceived = true
        case "UpdatedAt":
            if err := json.Unmarshal([]byte(v), &strct.UpdatedAt); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if CoreNetworkId (a required property) was received
    if !CoreNetworkIdReceived {
        return errors.New("\"CoreNetworkId\" is required but was not present")
    }
    // check if EdgeLocation (a required property) was received
    if !EdgeLocationReceived {
        return errors.New("\"EdgeLocation\" is required but was not present")
    }
    // check if Options (a required property) was received
    if !OptionsReceived {
        return errors.New("\"Options\" is required but was not present")
    }
    // check if TransportAttachmentId (a required property) was received
    if !TransportAttachmentIdReceived {
        return errors.New("\"TransportAttachmentId\" is required but was not present")
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
