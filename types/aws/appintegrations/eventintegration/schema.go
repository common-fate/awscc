// Code generated by schema-generate. DO NOT EDIT.

package eventintegration

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// EventFilter 
type EventFilter struct {

  // The source of the events.
  Source string `json:"Source"`
}

// EventIntegrationAssociation 
type EventIntegrationAssociation struct {

  // The metadata associated with the client.
  ClientAssociationMetadata []*Metadata `json:"ClientAssociationMetadata,omitempty"`

  // The identifier for the client that is associated with the event integration.
  ClientId string `json:"ClientId,omitempty"`

  // The name of the Eventbridge rule.
  EventBridgeRuleName string `json:"EventBridgeRuleName,omitempty"`

  // The Amazon Resource Name (ARN) for the event integration association.
  EventIntegrationAssociationArn string `json:"EventIntegrationAssociationArn,omitempty"`

  // The identifier for the event integration association.
  EventIntegrationAssociationId string `json:"EventIntegrationAssociationId,omitempty"`
}

// Metadata 
type Metadata struct {

  // A key to identify the metadata.
  Key string `json:"Key"`

  // Corresponding metadata value for the key.
  Value string `json:"Value"`
}

// Resource Resource Type definition for AWS::AppIntegrations::EventIntegration
type Resource struct {

  // The associations with the event integration.
  Associations []*EventIntegrationAssociation `json:"Associations,omitempty"`

  // The event integration description.
  Description string `json:"Description,omitempty"`

  // The Amazon Eventbridge bus for the event integration.
  EventBridgeBus string `json:"EventBridgeBus"`

  // The EventFilter (source) associated with the event integration.
  EventFilter *EventFilter `json:"EventFilter"`

  // The Amazon Resource Name (ARN) of the event integration.
  EventIntegrationArn string `json:"EventIntegrationArn,omitempty"`

  // The name of the event integration.
  Name string `json:"Name"`

  // The tags (keys and values) associated with the event integration.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {

  // A key to identify the tag.
  Key string `json:"Key"`

  // Corresponding tag value for the key.
  Value string `json:"Value"`
}

func (strct *EventFilter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Source" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Source" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EventFilter) UnmarshalJSON(b []byte) error {
    SourceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
            SourceReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Source (a required property) was received
    if !SourceReceived {
        return errors.New("\"Source\" is required but was not present")
    }
    return nil
}

func (strct *EventIntegrationAssociation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ClientAssociationMetadata" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientAssociationMetadata\": ")
	if tmp, err := json.Marshal(strct.ClientAssociationMetadata); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ClientId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientId\": ")
	if tmp, err := json.Marshal(strct.ClientId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EventBridgeRuleName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventBridgeRuleName\": ")
	if tmp, err := json.Marshal(strct.EventBridgeRuleName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EventIntegrationAssociationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventIntegrationAssociationArn\": ")
	if tmp, err := json.Marshal(strct.EventIntegrationAssociationArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EventIntegrationAssociationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventIntegrationAssociationId\": ")
	if tmp, err := json.Marshal(strct.EventIntegrationAssociationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EventIntegrationAssociation) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ClientAssociationMetadata":
            if err := json.Unmarshal([]byte(v), &strct.ClientAssociationMetadata); err != nil {
                return err
             }
        case "ClientId":
            if err := json.Unmarshal([]byte(v), &strct.ClientId); err != nil {
                return err
             }
        case "EventBridgeRuleName":
            if err := json.Unmarshal([]byte(v), &strct.EventBridgeRuleName); err != nil {
                return err
             }
        case "EventIntegrationAssociationArn":
            if err := json.Unmarshal([]byte(v), &strct.EventIntegrationAssociationArn); err != nil {
                return err
             }
        case "EventIntegrationAssociationId":
            if err := json.Unmarshal([]byte(v), &strct.EventIntegrationAssociationId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Metadata) MarshalJSON() ([]byte, error) {
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

func (strct *Metadata) UnmarshalJSON(b []byte) error {
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

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Associations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Associations\": ")
	if tmp, err := json.Marshal(strct.Associations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // "EventBridgeBus" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EventBridgeBus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventBridgeBus\": ")
	if tmp, err := json.Marshal(strct.EventBridgeBus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "EventFilter" field is required
    if strct.EventFilter == nil {
        return nil, errors.New("EventFilter is a required field")
    }
    // Marshal the "EventFilter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventFilter\": ")
	if tmp, err := json.Marshal(strct.EventFilter); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EventIntegrationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventIntegrationArn\": ")
	if tmp, err := json.Marshal(strct.EventIntegrationArn); err != nil {
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
    EventBridgeBusReceived := false
    EventFilterReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Associations":
            if err := json.Unmarshal([]byte(v), &strct.Associations); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "EventBridgeBus":
            if err := json.Unmarshal([]byte(v), &strct.EventBridgeBus); err != nil {
                return err
             }
            EventBridgeBusReceived = true
        case "EventFilter":
            if err := json.Unmarshal([]byte(v), &strct.EventFilter); err != nil {
                return err
             }
            EventFilterReceived = true
        case "EventIntegrationArn":
            if err := json.Unmarshal([]byte(v), &strct.EventIntegrationArn); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if EventBridgeBus (a required property) was received
    if !EventBridgeBusReceived {
        return errors.New("\"EventBridgeBus\" is required but was not present")
    }
    // check if EventFilter (a required property) was received
    if !EventFilterReceived {
        return errors.New("\"EventFilter\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
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
