// Code generated by schema-generate. DO NOT EDIT.

package listener

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// PortRange A port range to support for connections from  clients to your accelerator.
type PortRange struct {
  FromPort int `json:"FromPort"`
  ToPort int `json:"ToPort"`
}

// Resource Resource Type definition for AWS::GlobalAccelerator::Listener
type Resource struct {

  // The Amazon Resource Name (ARN) of the accelerator.
  AcceleratorArn string `json:"AcceleratorArn"`

  // Client affinity lets you direct all requests from a user to the same endpoint.
  ClientAffinity string `json:"ClientAffinity,omitempty"`

  // The Amazon Resource Name (ARN) of the listener.
  ListenerArn string `json:"ListenerArn,omitempty"`
  PortRanges []*PortRange `json:"PortRanges"`

  // The protocol for the listener.
  Protocol string `json:"Protocol"`
}

func (strct *PortRange) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "FromPort" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FromPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FromPort\": ")
	if tmp, err := json.Marshal(strct.FromPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ToPort" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ToPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ToPort\": ")
	if tmp, err := json.Marshal(strct.ToPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PortRange) UnmarshalJSON(b []byte) error {
    FromPortReceived := false
    ToPortReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FromPort":
            if err := json.Unmarshal([]byte(v), &strct.FromPort); err != nil {
                return err
             }
            FromPortReceived = true
        case "ToPort":
            if err := json.Unmarshal([]byte(v), &strct.ToPort); err != nil {
                return err
             }
            ToPortReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FromPort (a required property) was received
    if !FromPortReceived {
        return errors.New("\"FromPort\" is required but was not present")
    }
    // check if ToPort (a required property) was received
    if !ToPortReceived {
        return errors.New("\"ToPort\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AcceleratorArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AcceleratorArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AcceleratorArn\": ")
	if tmp, err := json.Marshal(strct.AcceleratorArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ClientAffinity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientAffinity\": ")
	if tmp, err := json.Marshal(strct.ClientAffinity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ListenerArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ListenerArn\": ")
	if tmp, err := json.Marshal(strct.ListenerArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PortRanges" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PortRanges" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PortRanges\": ")
	if tmp, err := json.Marshal(strct.PortRanges); err != nil {
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
    AcceleratorArnReceived := false
    PortRangesReceived := false
    ProtocolReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AcceleratorArn":
            if err := json.Unmarshal([]byte(v), &strct.AcceleratorArn); err != nil {
                return err
             }
            AcceleratorArnReceived = true
        case "ClientAffinity":
            if err := json.Unmarshal([]byte(v), &strct.ClientAffinity); err != nil {
                return err
             }
        case "ListenerArn":
            if err := json.Unmarshal([]byte(v), &strct.ListenerArn); err != nil {
                return err
             }
        case "PortRanges":
            if err := json.Unmarshal([]byte(v), &strct.PortRanges); err != nil {
                return err
             }
            PortRangesReceived = true
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
            ProtocolReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AcceleratorArn (a required property) was received
    if !AcceleratorArnReceived {
        return errors.New("\"AcceleratorArn\" is required but was not present")
    }
    // check if PortRanges (a required property) was received
    if !PortRangesReceived {
        return errors.New("\"PortRanges\" is required but was not present")
    }
    // check if Protocol (a required property) was received
    if !ProtocolReceived {
        return errors.New("\"Protocol\" is required but was not present")
    }
    return nil
}
