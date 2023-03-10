// Code generated by schema-generate. DO NOT EDIT.

package domain

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// Resource The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.
type Resource struct {
  Description string `json:"Description,omitempty"`
  DomainId string `json:"DomainId,omitempty"`
  Name string `json:"Name"`
  ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `json:"ServerSideEncryptionConfiguration"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// ServerSideEncryptionConfiguration 
type ServerSideEncryptionConfiguration struct {
  KmsKeyId string `json:"KmsKeyId"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Description" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainId\": ")
	if tmp, err := json.Marshal(strct.DomainId); err != nil {
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
    // "ServerSideEncryptionConfiguration" field is required
    if strct.ServerSideEncryptionConfiguration == nil {
        return nil, errors.New("ServerSideEncryptionConfiguration is a required field")
    }
    // Marshal the "ServerSideEncryptionConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerSideEncryptionConfiguration\": ")
	if tmp, err := json.Marshal(strct.ServerSideEncryptionConfiguration); err != nil {
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
    NameReceived := false
    ServerSideEncryptionConfigurationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "DomainId":
            if err := json.Unmarshal([]byte(v), &strct.DomainId); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "ServerSideEncryptionConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.ServerSideEncryptionConfiguration); err != nil {
                return err
             }
            ServerSideEncryptionConfigurationReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if ServerSideEncryptionConfiguration (a required property) was received
    if !ServerSideEncryptionConfigurationReceived {
        return errors.New("\"ServerSideEncryptionConfiguration\" is required but was not present")
    }
    return nil
}

func (strct *ServerSideEncryptionConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "KmsKeyId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "KmsKeyId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KmsKeyId\": ")
	if tmp, err := json.Marshal(strct.KmsKeyId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ServerSideEncryptionConfiguration) UnmarshalJSON(b []byte) error {
    KmsKeyIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "KmsKeyId":
            if err := json.Unmarshal([]byte(v), &strct.KmsKeyId); err != nil {
                return err
             }
            KmsKeyIdReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if KmsKeyId (a required property) was received
    if !KmsKeyIdReceived {
        return errors.New("\"KmsKeyId\" is required but was not present")
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
