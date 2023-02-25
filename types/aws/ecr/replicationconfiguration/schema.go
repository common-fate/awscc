// Code generated by schema-generate. DO NOT EDIT.

package replicationconfiguration

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ReplicationConfiguration An object representing the replication configuration for a registry.
type ReplicationConfiguration struct {

  // An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.
  Rules []*ReplicationRule `json:"Rules"`
}

// ReplicationDestination An array of objects representing the details of a replication destination.
type ReplicationDestination struct {
  Region string `json:"Region"`
  RegistryId string `json:"RegistryId"`
}

// ReplicationRule An array of objects representing the details of a replication destination.
type ReplicationRule struct {

  // An array of objects representing the details of a replication destination.
  Destinations []*ReplicationDestination `json:"Destinations"`

  // An array of objects representing the details of a repository filter.
  RepositoryFilters []*RepositoryFilter `json:"RepositoryFilters,omitempty"`
}

// RepositoryFilter An array of objects representing the details of a repository filter.
type RepositoryFilter struct {
  Filter string `json:"Filter"`
  FilterType string `json:"FilterType"`
}

// Resource The AWS::ECR::ReplicationConfiguration resource configures the replication destinations for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication.html
type Resource struct {

  // The RegistryId associated with the aws account.
  RegistryId string `json:"RegistryId,omitempty"`
  ReplicationConfiguration *ReplicationConfiguration `json:"ReplicationConfiguration"`
}

func (strct *ReplicationConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Rules" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Rules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Rules\": ")
	if tmp, err := json.Marshal(strct.Rules); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReplicationConfiguration) UnmarshalJSON(b []byte) error {
    RulesReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Rules":
            if err := json.Unmarshal([]byte(v), &strct.Rules); err != nil {
                return err
             }
            RulesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Rules (a required property) was received
    if !RulesReceived {
        return errors.New("\"Rules\" is required but was not present")
    }
    return nil
}

func (strct *ReplicationDestination) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Region" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Region" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Region\": ")
	if tmp, err := json.Marshal(strct.Region); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RegistryId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RegistryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegistryId\": ")
	if tmp, err := json.Marshal(strct.RegistryId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReplicationDestination) UnmarshalJSON(b []byte) error {
    RegionReceived := false
    RegistryIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Region":
            if err := json.Unmarshal([]byte(v), &strct.Region); err != nil {
                return err
             }
            RegionReceived = true
        case "RegistryId":
            if err := json.Unmarshal([]byte(v), &strct.RegistryId); err != nil {
                return err
             }
            RegistryIdReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Region (a required property) was received
    if !RegionReceived {
        return errors.New("\"Region\" is required but was not present")
    }
    // check if RegistryId (a required property) was received
    if !RegistryIdReceived {
        return errors.New("\"RegistryId\" is required but was not present")
    }
    return nil
}

func (strct *ReplicationRule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Destinations" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Destinations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Destinations\": ")
	if tmp, err := json.Marshal(strct.Destinations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RepositoryFilters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RepositoryFilters\": ")
	if tmp, err := json.Marshal(strct.RepositoryFilters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReplicationRule) UnmarshalJSON(b []byte) error {
    DestinationsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Destinations":
            if err := json.Unmarshal([]byte(v), &strct.Destinations); err != nil {
                return err
             }
            DestinationsReceived = true
        case "RepositoryFilters":
            if err := json.Unmarshal([]byte(v), &strct.RepositoryFilters); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Destinations (a required property) was received
    if !DestinationsReceived {
        return errors.New("\"Destinations\" is required but was not present")
    }
    return nil
}

func (strct *RepositoryFilter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Filter" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Filter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Filter\": ")
	if tmp, err := json.Marshal(strct.Filter); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FilterType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FilterType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FilterType\": ")
	if tmp, err := json.Marshal(strct.FilterType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RepositoryFilter) UnmarshalJSON(b []byte) error {
    FilterReceived := false
    FilterTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Filter":
            if err := json.Unmarshal([]byte(v), &strct.Filter); err != nil {
                return err
             }
            FilterReceived = true
        case "FilterType":
            if err := json.Unmarshal([]byte(v), &strct.FilterType); err != nil {
                return err
             }
            FilterTypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Filter (a required property) was received
    if !FilterReceived {
        return errors.New("\"Filter\" is required but was not present")
    }
    // check if FilterType (a required property) was received
    if !FilterTypeReceived {
        return errors.New("\"FilterType\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "RegistryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegistryId\": ")
	if tmp, err := json.Marshal(strct.RegistryId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ReplicationConfiguration" field is required
    if strct.ReplicationConfiguration == nil {
        return nil, errors.New("ReplicationConfiguration is a required field")
    }
    // Marshal the "ReplicationConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplicationConfiguration\": ")
	if tmp, err := json.Marshal(strct.ReplicationConfiguration); err != nil {
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
    ReplicationConfigurationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "RegistryId":
            if err := json.Unmarshal([]byte(v), &strct.RegistryId); err != nil {
                return err
             }
        case "ReplicationConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.ReplicationConfiguration); err != nil {
                return err
             }
            ReplicationConfigurationReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ReplicationConfiguration (a required property) was received
    if !ReplicationConfigurationReceived {
        return errors.New("\"ReplicationConfiguration\" is required but was not present")
    }
    return nil
}
