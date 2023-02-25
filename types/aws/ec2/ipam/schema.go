// Code generated by schema-generate. DO NOT EDIT.

package ipam

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// IpamOperatingRegion The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
type IpamOperatingRegion struct {

  // The name of the region.
  RegionName string `json:"RegionName"`
}

// Resource Resource Schema of AWS::EC2::IPAM Type
type Resource struct {

  // The Amazon Resource Name (ARN) of the IPAM.
  Arn string `json:"Arn,omitempty"`

  // The Id of the default association to the default resource discovery, created with this IPAM.
  DefaultResourceDiscoveryAssociationId string `json:"DefaultResourceDiscoveryAssociationId,omitempty"`

  // The Id of the default resource discovery, created with this IPAM.
  DefaultResourceDiscoveryId string `json:"DefaultResourceDiscoveryId,omitempty"`
  Description string `json:"Description,omitempty"`

  // Id of the IPAM.
  IpamId string `json:"IpamId,omitempty"`

  // The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
  OperatingRegions []*IpamOperatingRegion `json:"OperatingRegions,omitempty"`

  // The Id of the default scope for publicly routable IP space, created with this IPAM.
  PrivateDefaultScopeId string `json:"PrivateDefaultScopeId,omitempty"`

  // The Id of the default scope for publicly routable IP space, created with this IPAM.
  PublicDefaultScopeId string `json:"PublicDefaultScopeId,omitempty"`

  // The count of resource discoveries associated with this IPAM.
  ResourceDiscoveryAssociationCount int `json:"ResourceDiscoveryAssociationCount,omitempty"`

  // The number of scopes that currently exist in this IPAM.
  ScopeCount int `json:"ScopeCount,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *IpamOperatingRegion) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "RegionName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RegionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionName\": ")
	if tmp, err := json.Marshal(strct.RegionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *IpamOperatingRegion) UnmarshalJSON(b []byte) error {
    RegionNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "RegionName":
            if err := json.Unmarshal([]byte(v), &strct.RegionName); err != nil {
                return err
             }
            RegionNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if RegionName (a required property) was received
    if !RegionNameReceived {
        return errors.New("\"RegionName\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Arn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Arn\": ")
	if tmp, err := json.Marshal(strct.Arn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultResourceDiscoveryAssociationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultResourceDiscoveryAssociationId\": ")
	if tmp, err := json.Marshal(strct.DefaultResourceDiscoveryAssociationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultResourceDiscoveryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultResourceDiscoveryId\": ")
	if tmp, err := json.Marshal(strct.DefaultResourceDiscoveryId); err != nil {
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
    // Marshal the "IpamId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamId\": ")
	if tmp, err := json.Marshal(strct.IpamId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OperatingRegions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OperatingRegions\": ")
	if tmp, err := json.Marshal(strct.OperatingRegions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PrivateDefaultScopeId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PrivateDefaultScopeId\": ")
	if tmp, err := json.Marshal(strct.PrivateDefaultScopeId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicDefaultScopeId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicDefaultScopeId\": ")
	if tmp, err := json.Marshal(strct.PublicDefaultScopeId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceDiscoveryAssociationCount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceDiscoveryAssociationCount\": ")
	if tmp, err := json.Marshal(strct.ResourceDiscoveryAssociationCount); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ScopeCount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScopeCount\": ")
	if tmp, err := json.Marshal(strct.ScopeCount); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "DefaultResourceDiscoveryAssociationId":
            if err := json.Unmarshal([]byte(v), &strct.DefaultResourceDiscoveryAssociationId); err != nil {
                return err
             }
        case "DefaultResourceDiscoveryId":
            if err := json.Unmarshal([]byte(v), &strct.DefaultResourceDiscoveryId); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "IpamId":
            if err := json.Unmarshal([]byte(v), &strct.IpamId); err != nil {
                return err
             }
        case "OperatingRegions":
            if err := json.Unmarshal([]byte(v), &strct.OperatingRegions); err != nil {
                return err
             }
        case "PrivateDefaultScopeId":
            if err := json.Unmarshal([]byte(v), &strct.PrivateDefaultScopeId); err != nil {
                return err
             }
        case "PublicDefaultScopeId":
            if err := json.Unmarshal([]byte(v), &strct.PublicDefaultScopeId); err != nil {
                return err
             }
        case "ResourceDiscoveryAssociationCount":
            if err := json.Unmarshal([]byte(v), &strct.ResourceDiscoveryAssociationCount); err != nil {
                return err
             }
        case "ScopeCount":
            if err := json.Unmarshal([]byte(v), &strct.ScopeCount); err != nil {
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