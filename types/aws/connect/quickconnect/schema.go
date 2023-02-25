// Code generated by schema-generate. DO NOT EDIT.

package quickconnect

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// PhoneNumberQuickConnectConfig The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.
type PhoneNumberQuickConnectConfig struct {
  PhoneNumber string `json:"PhoneNumber"`
}

// QueueQuickConnectConfig The queue configuration. This is required only if QuickConnectType is QUEUE.
type QueueQuickConnectConfig struct {
  ContactFlowArn string `json:"ContactFlowArn"`
  QueueArn string `json:"QueueArn"`
}

// QuickConnectConfig Configuration settings for the quick connect.
type QuickConnectConfig struct {
  PhoneConfig *PhoneNumberQuickConnectConfig `json:"PhoneConfig,omitempty"`
  QueueConfig *QueueQuickConnectConfig `json:"QueueConfig,omitempty"`
  QuickConnectType string `json:"QuickConnectType"`
  UserConfig *UserQuickConnectConfig `json:"UserConfig,omitempty"`
}

// Resource Resource Type definition for AWS::Connect::QuickConnect
type Resource struct {

  // The description of the quick connect.
  Description string `json:"Description,omitempty"`

  // The identifier of the Amazon Connect instance.
  InstanceArn string `json:"InstanceArn"`

  // The name of the quick connect.
  Name string `json:"Name"`

  // The Amazon Resource Name (ARN) for the quick connect.
  QuickConnectArn string `json:"QuickConnectArn,omitempty"`

  // Configuration settings for the quick connect.
  QuickConnectConfig *QuickConnectConfig `json:"QuickConnectConfig"`

  // One or more tags.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

// UserQuickConnectConfig The user configuration. This is required only if QuickConnectType is USER.
type UserQuickConnectConfig struct {
  ContactFlowArn string `json:"ContactFlowArn"`
  UserArn string `json:"UserArn"`
}

func (strct *PhoneNumberQuickConnectConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "PhoneNumber" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PhoneNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PhoneNumber\": ")
	if tmp, err := json.Marshal(strct.PhoneNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PhoneNumberQuickConnectConfig) UnmarshalJSON(b []byte) error {
    PhoneNumberReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PhoneNumber":
            if err := json.Unmarshal([]byte(v), &strct.PhoneNumber); err != nil {
                return err
             }
            PhoneNumberReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if PhoneNumber (a required property) was received
    if !PhoneNumberReceived {
        return errors.New("\"PhoneNumber\" is required but was not present")
    }
    return nil
}

func (strct *QueueQuickConnectConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ContactFlowArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ContactFlowArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContactFlowArn\": ")
	if tmp, err := json.Marshal(strct.ContactFlowArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "QueueArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "QueueArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueueArn\": ")
	if tmp, err := json.Marshal(strct.QueueArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *QueueQuickConnectConfig) UnmarshalJSON(b []byte) error {
    ContactFlowArnReceived := false
    QueueArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContactFlowArn":
            if err := json.Unmarshal([]byte(v), &strct.ContactFlowArn); err != nil {
                return err
             }
            ContactFlowArnReceived = true
        case "QueueArn":
            if err := json.Unmarshal([]byte(v), &strct.QueueArn); err != nil {
                return err
             }
            QueueArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ContactFlowArn (a required property) was received
    if !ContactFlowArnReceived {
        return errors.New("\"ContactFlowArn\" is required but was not present")
    }
    // check if QueueArn (a required property) was received
    if !QueueArnReceived {
        return errors.New("\"QueueArn\" is required but was not present")
    }
    return nil
}

func (strct *QuickConnectConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "PhoneConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PhoneConfig\": ")
	if tmp, err := json.Marshal(strct.PhoneConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "QueueConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QueueConfig\": ")
	if tmp, err := json.Marshal(strct.QueueConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "QuickConnectType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "QuickConnectType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QuickConnectType\": ")
	if tmp, err := json.Marshal(strct.QuickConnectType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UserConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UserConfig\": ")
	if tmp, err := json.Marshal(strct.UserConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *QuickConnectConfig) UnmarshalJSON(b []byte) error {
    QuickConnectTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PhoneConfig":
            if err := json.Unmarshal([]byte(v), &strct.PhoneConfig); err != nil {
                return err
             }
        case "QueueConfig":
            if err := json.Unmarshal([]byte(v), &strct.QueueConfig); err != nil {
                return err
             }
        case "QuickConnectType":
            if err := json.Unmarshal([]byte(v), &strct.QuickConnectType); err != nil {
                return err
             }
            QuickConnectTypeReceived = true
        case "UserConfig":
            if err := json.Unmarshal([]byte(v), &strct.UserConfig); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if QuickConnectType (a required property) was received
    if !QuickConnectTypeReceived {
        return errors.New("\"QuickConnectType\" is required but was not present")
    }
    return nil
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
    // "InstanceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "InstanceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceArn\": ")
	if tmp, err := json.Marshal(strct.InstanceArn); err != nil {
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
    // Marshal the "QuickConnectArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QuickConnectArn\": ")
	if tmp, err := json.Marshal(strct.QuickConnectArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "QuickConnectConfig" field is required
    if strct.QuickConnectConfig == nil {
        return nil, errors.New("QuickConnectConfig is a required field")
    }
    // Marshal the "QuickConnectConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"QuickConnectConfig\": ")
	if tmp, err := json.Marshal(strct.QuickConnectConfig); err != nil {
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
    InstanceArnReceived := false
    NameReceived := false
    QuickConnectConfigReceived := false
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
        case "InstanceArn":
            if err := json.Unmarshal([]byte(v), &strct.InstanceArn); err != nil {
                return err
             }
            InstanceArnReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "QuickConnectArn":
            if err := json.Unmarshal([]byte(v), &strct.QuickConnectArn); err != nil {
                return err
             }
        case "QuickConnectConfig":
            if err := json.Unmarshal([]byte(v), &strct.QuickConnectConfig); err != nil {
                return err
             }
            QuickConnectConfigReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if InstanceArn (a required property) was received
    if !InstanceArnReceived {
        return errors.New("\"InstanceArn\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if QuickConnectConfig (a required property) was received
    if !QuickConnectConfigReceived {
        return errors.New("\"QuickConnectConfig\" is required but was not present")
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

func (strct *UserQuickConnectConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ContactFlowArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ContactFlowArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContactFlowArn\": ")
	if tmp, err := json.Marshal(strct.ContactFlowArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "UserArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "UserArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UserArn\": ")
	if tmp, err := json.Marshal(strct.UserArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UserQuickConnectConfig) UnmarshalJSON(b []byte) error {
    ContactFlowArnReceived := false
    UserArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContactFlowArn":
            if err := json.Unmarshal([]byte(v), &strct.ContactFlowArn); err != nil {
                return err
             }
            ContactFlowArnReceived = true
        case "UserArn":
            if err := json.Unmarshal([]byte(v), &strct.UserArn); err != nil {
                return err
             }
            UserArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ContactFlowArn (a required property) was received
    if !ContactFlowArnReceived {
        return errors.New("\"ContactFlowArn\" is required but was not present")
    }
    // check if UserArn (a required property) was received
    if !UserArnReceived {
        return errors.New("\"UserArn\" is required but was not present")
    }
    return nil
}
