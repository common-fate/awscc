// Code generated by schema-generate. DO NOT EDIT.

package domain

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// Resource A domain defined for 3rd party data source in Profile Service
type Resource struct {

  // The time of this integration got created
  CreatedAt string `json:"CreatedAt,omitempty"`

  // The URL of the SQS dead letter queue
  DeadLetterQueueUrl string `json:"DeadLetterQueueUrl,omitempty"`

  // The default encryption key
  DefaultEncryptionKey string `json:"DefaultEncryptionKey,omitempty"`

  // The default number of days until the data within the domain expires.
  DefaultExpirationDays int `json:"DefaultExpirationDays,omitempty"`

  // The unique name of the domain.
  DomainName string `json:"DomainName"`

  // The time of this integration got last updated at
  LastUpdatedAt string `json:"LastUpdatedAt,omitempty"`

  // The tags (keys and values) associated with the domain
  Tags []*Tag `json:"Tags,omitempty"`
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
    // Marshal the "DeadLetterQueueUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeadLetterQueueUrl\": ")
	if tmp, err := json.Marshal(strct.DeadLetterQueueUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultEncryptionKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultEncryptionKey\": ")
	if tmp, err := json.Marshal(strct.DefaultEncryptionKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultExpirationDays" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultExpirationDays\": ")
	if tmp, err := json.Marshal(strct.DefaultExpirationDays); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DomainName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DomainName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainName\": ")
	if tmp, err := json.Marshal(strct.DomainName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LastUpdatedAt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LastUpdatedAt\": ")
	if tmp, err := json.Marshal(strct.LastUpdatedAt); err != nil {
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
    DomainNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CreatedAt":
            if err := json.Unmarshal([]byte(v), &strct.CreatedAt); err != nil {
                return err
             }
        case "DeadLetterQueueUrl":
            if err := json.Unmarshal([]byte(v), &strct.DeadLetterQueueUrl); err != nil {
                return err
             }
        case "DefaultEncryptionKey":
            if err := json.Unmarshal([]byte(v), &strct.DefaultEncryptionKey); err != nil {
                return err
             }
        case "DefaultExpirationDays":
            if err := json.Unmarshal([]byte(v), &strct.DefaultExpirationDays); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
            DomainNameReceived = true
        case "LastUpdatedAt":
            if err := json.Unmarshal([]byte(v), &strct.LastUpdatedAt); err != nil {
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
    // check if DomainName (a required property) was received
    if !DomainNameReceived {
        return errors.New("\"DomainName\" is required but was not present")
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
