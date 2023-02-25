// Code generated by schema-generate. DO NOT EDIT.

package keysigningkey

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
type Resource struct {

  // The unique string (ID) used to identify a hosted zone.
  HostedZoneId string `json:"HostedZoneId"`

  // The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
  KeyManagementServiceArn string `json:"KeyManagementServiceArn"`

  // An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
  Name string `json:"Name"`

  // A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
  Status string `json:"Status"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "HostedZoneId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "HostedZoneId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HostedZoneId\": ")
	if tmp, err := json.Marshal(strct.HostedZoneId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "KeyManagementServiceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "KeyManagementServiceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyManagementServiceArn\": ")
	if tmp, err := json.Marshal(strct.KeyManagementServiceArn); err != nil {
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
    // "Status" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
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
    HostedZoneIdReceived := false
    KeyManagementServiceArnReceived := false
    NameReceived := false
    StatusReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "HostedZoneId":
            if err := json.Unmarshal([]byte(v), &strct.HostedZoneId); err != nil {
                return err
             }
            HostedZoneIdReceived = true
        case "KeyManagementServiceArn":
            if err := json.Unmarshal([]byte(v), &strct.KeyManagementServiceArn); err != nil {
                return err
             }
            KeyManagementServiceArnReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
            StatusReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if HostedZoneId (a required property) was received
    if !HostedZoneIdReceived {
        return errors.New("\"HostedZoneId\" is required but was not present")
    }
    // check if KeyManagementServiceArn (a required property) was received
    if !KeyManagementServiceArnReceived {
        return errors.New("\"KeyManagementServiceArn\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Status (a required property) was received
    if !StatusReceived {
        return errors.New("\"Status\" is required but was not present")
    }
    return nil
}