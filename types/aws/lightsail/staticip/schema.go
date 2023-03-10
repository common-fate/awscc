// Code generated by schema-generate. DO NOT EDIT.

package staticip

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::Lightsail::StaticIp
type Resource struct {

  // The instance where the static IP is attached.
  AttachedTo string `json:"AttachedTo,omitempty"`

  // The static IP address.
  IpAddress string `json:"IpAddress,omitempty"`

  // A Boolean value indicating whether the static IP is attached.
  IsAttached bool `json:"IsAttached,omitempty"`
  StaticIpArn string `json:"StaticIpArn,omitempty"`

  // The name of the static IP address.
  StaticIpName string `json:"StaticIpName"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AttachedTo" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachedTo\": ")
	if tmp, err := json.Marshal(strct.AttachedTo); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpAddress\": ")
	if tmp, err := json.Marshal(strct.IpAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IsAttached" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IsAttached\": ")
	if tmp, err := json.Marshal(strct.IsAttached); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StaticIpArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StaticIpArn\": ")
	if tmp, err := json.Marshal(strct.StaticIpArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "StaticIpName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "StaticIpName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StaticIpName\": ")
	if tmp, err := json.Marshal(strct.StaticIpName); err != nil {
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
    StaticIpNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AttachedTo":
            if err := json.Unmarshal([]byte(v), &strct.AttachedTo); err != nil {
                return err
             }
        case "IpAddress":
            if err := json.Unmarshal([]byte(v), &strct.IpAddress); err != nil {
                return err
             }
        case "IsAttached":
            if err := json.Unmarshal([]byte(v), &strct.IsAttached); err != nil {
                return err
             }
        case "StaticIpArn":
            if err := json.Unmarshal([]byte(v), &strct.StaticIpArn); err != nil {
                return err
             }
        case "StaticIpName":
            if err := json.Unmarshal([]byte(v), &strct.StaticIpName); err != nil {
                return err
             }
            StaticIpNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if StaticIpName (a required property) was received
    if !StaticIpNameReceived {
        return errors.New("\"StaticIpName\" is required but was not present")
    }
    return nil
}
