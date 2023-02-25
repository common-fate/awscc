// Code generated by schema-generate. DO NOT EDIT.

package ipamresourcediscovery

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// IpamOperatingRegion The regions IPAM Resource Discovery is enabled for. Allows for monitoring.
type IpamOperatingRegion struct {

  // The name of the region.
  RegionName string `json:"RegionName"`
}

// Resource Resource Schema of AWS::EC2::IPAMResourceDiscovery Type
type Resource struct {
  Description string `json:"Description,omitempty"`

  // Amazon Resource Name (Arn) for the Resource Discovery.
  IpamResourceDiscoveryArn string `json:"IpamResourceDiscoveryArn,omitempty"`

  // Id of the IPAM Pool.
  IpamResourceDiscoveryId string `json:"IpamResourceDiscoveryId,omitempty"`

  // The region the resource discovery is setup in. 
  IpamResourceDiscoveryRegion string `json:"IpamResourceDiscoveryRegion,omitempty"`

  // Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
  IsDefault bool `json:"IsDefault,omitempty"`

  // The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring
  OperatingRegions []*IpamOperatingRegion `json:"OperatingRegions,omitempty"`

  // Owner Account ID of the Resource Discovery
  OwnerId string `json:"OwnerId,omitempty"`

  // The state of this Resource Discovery.
  State string `json:"State,omitempty"`

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
    // Marshal the "IpamResourceDiscoveryArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamResourceDiscoveryArn\": ")
	if tmp, err := json.Marshal(strct.IpamResourceDiscoveryArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamResourceDiscoveryId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamResourceDiscoveryId\": ")
	if tmp, err := json.Marshal(strct.IpamResourceDiscoveryId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamResourceDiscoveryRegion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamResourceDiscoveryRegion\": ")
	if tmp, err := json.Marshal(strct.IpamResourceDiscoveryRegion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IsDefault" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IsDefault\": ")
	if tmp, err := json.Marshal(strct.IsDefault); err != nil {
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
    // Marshal the "OwnerId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OwnerId\": ")
	if tmp, err := json.Marshal(strct.OwnerId); err != nil {
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
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "IpamResourceDiscoveryArn":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryArn); err != nil {
                return err
             }
        case "IpamResourceDiscoveryId":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryId); err != nil {
                return err
             }
        case "IpamResourceDiscoveryRegion":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryRegion); err != nil {
                return err
             }
        case "IsDefault":
            if err := json.Unmarshal([]byte(v), &strct.IsDefault); err != nil {
                return err
             }
        case "OperatingRegions":
            if err := json.Unmarshal([]byte(v), &strct.OperatingRegions); err != nil {
                return err
             }
        case "OwnerId":
            if err := json.Unmarshal([]byte(v), &strct.OwnerId); err != nil {
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