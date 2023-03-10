// Code generated by schema-generate. DO NOT EDIT.

package ipamresourcediscoveryassociation

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// Resource Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type
type Resource struct {

  // Arn of the IPAM.
  IpamArn string `json:"IpamArn,omitempty"`

  // The Id of the IPAM this Resource Discovery is associated to.
  IpamId string `json:"IpamId"`

  // The home region of the IPAM.
  IpamRegion string `json:"IpamRegion,omitempty"`

  // The Amazon Resource Name (ARN) of the resource discovery association is a part of.
  IpamResourceDiscoveryAssociationArn string `json:"IpamResourceDiscoveryAssociationArn,omitempty"`

  // Id of the IPAM Resource Discovery Association.
  IpamResourceDiscoveryAssociationId string `json:"IpamResourceDiscoveryAssociationId,omitempty"`

  // The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
  IpamResourceDiscoveryId string `json:"IpamResourceDiscoveryId"`

  // If the Resource Discovery Association exists due as part of CreateIpam.
  IsDefault bool `json:"IsDefault,omitempty"`

  // The AWS Account ID for the account where the shared IPAM exists.
  OwnerId string `json:"OwnerId,omitempty"`

  // The status of the resource discovery.
  ResourceDiscoveryStatus string `json:"ResourceDiscoveryStatus,omitempty"`

  // The operational state of the Resource Discovery Association. Related to Create/Delete activities.
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

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "IpamArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamArn\": ")
	if tmp, err := json.Marshal(strct.IpamArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IpamId" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "IpamRegion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamRegion\": ")
	if tmp, err := json.Marshal(strct.IpamRegion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamResourceDiscoveryAssociationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamResourceDiscoveryAssociationArn\": ")
	if tmp, err := json.Marshal(strct.IpamResourceDiscoveryAssociationArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamResourceDiscoveryAssociationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamResourceDiscoveryAssociationId\": ")
	if tmp, err := json.Marshal(strct.IpamResourceDiscoveryAssociationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IpamResourceDiscoveryId" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "ResourceDiscoveryStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceDiscoveryStatus\": ")
	if tmp, err := json.Marshal(strct.ResourceDiscoveryStatus); err != nil {
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
    IpamIdReceived := false
    IpamResourceDiscoveryIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "IpamArn":
            if err := json.Unmarshal([]byte(v), &strct.IpamArn); err != nil {
                return err
             }
        case "IpamId":
            if err := json.Unmarshal([]byte(v), &strct.IpamId); err != nil {
                return err
             }
            IpamIdReceived = true
        case "IpamRegion":
            if err := json.Unmarshal([]byte(v), &strct.IpamRegion); err != nil {
                return err
             }
        case "IpamResourceDiscoveryAssociationArn":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryAssociationArn); err != nil {
                return err
             }
        case "IpamResourceDiscoveryAssociationId":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryAssociationId); err != nil {
                return err
             }
        case "IpamResourceDiscoveryId":
            if err := json.Unmarshal([]byte(v), &strct.IpamResourceDiscoveryId); err != nil {
                return err
             }
            IpamResourceDiscoveryIdReceived = true
        case "IsDefault":
            if err := json.Unmarshal([]byte(v), &strct.IsDefault); err != nil {
                return err
             }
        case "OwnerId":
            if err := json.Unmarshal([]byte(v), &strct.OwnerId); err != nil {
                return err
             }
        case "ResourceDiscoveryStatus":
            if err := json.Unmarshal([]byte(v), &strct.ResourceDiscoveryStatus); err != nil {
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
    // check if IpamId (a required property) was received
    if !IpamIdReceived {
        return errors.New("\"IpamId\" is required but was not present")
    }
    // check if IpamResourceDiscoveryId (a required property) was received
    if !IpamResourceDiscoveryIdReceived {
        return errors.New("\"IpamResourceDiscoveryId\" is required but was not present")
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
