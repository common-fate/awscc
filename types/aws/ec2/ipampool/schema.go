// Code generated by schema-generate. DO NOT EDIT.

package ipampool

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ProvisionedCidr An address space to be inserted into this pool. All allocations must be made from this address space.
type ProvisionedCidr struct {
  Cidr string `json:"Cidr"`
}

// Resource Resource Schema of AWS::EC2::IPAMPool Type
type Resource struct {

  // The address family of the address space in this pool. Either IPv4 or IPv6.
  AddressFamily string `json:"AddressFamily"`

  // The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
  AllocationDefaultNetmaskLength int `json:"AllocationDefaultNetmaskLength,omitempty"`

  // The maximum allowed netmask length for allocations made from this pool.
  AllocationMaxNetmaskLength int `json:"AllocationMaxNetmaskLength,omitempty"`

  // The minimum allowed netmask length for allocations made from this pool.
  AllocationMinNetmaskLength int `json:"AllocationMinNetmaskLength,omitempty"`

  // When specified, an allocation will not be allowed unless a resource has a matching set of tags.
  AllocationResourceTags []*Tag `json:"AllocationResourceTags,omitempty"`

  // The Amazon Resource Name (ARN) of the IPAM Pool.
  Arn string `json:"Arn,omitempty"`

  // Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
  AutoImport bool `json:"AutoImport,omitempty"`

  // Limits which service in Amazon Web Services that the pool can be used in.
  AwsService string `json:"AwsService,omitempty"`
  Description string `json:"Description,omitempty"`

  // The Amazon Resource Name (ARN) of the IPAM this pool is a part of.
  IpamArn string `json:"IpamArn,omitempty"`

  // Id of the IPAM Pool.
  IpamPoolId string `json:"IpamPoolId,omitempty"`

  // The Amazon Resource Name (ARN) of the scope this pool is a part of.
  IpamScopeArn string `json:"IpamScopeArn,omitempty"`

  // The Id of the scope this pool is a part of.
  IpamScopeId string `json:"IpamScopeId"`

  // Determines whether this scope contains publicly routable space or space for a private network
  IpamScopeType string `json:"IpamScopeType,omitempty"`

  // The region of this pool. If not set, this will default to "None" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.
  Locale string `json:"Locale,omitempty"`

  // The depth of this pool in the source pool hierarchy.
  PoolDepth int `json:"PoolDepth,omitempty"`

  // A list of cidrs representing the address space available for allocation in this pool.
  ProvisionedCidrs []*ProvisionedCidr `json:"ProvisionedCidrs,omitempty"`

  // The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.
  PublicIpSource string `json:"PublicIpSource,omitempty"`

  // Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
  PubliclyAdvertisable bool `json:"PubliclyAdvertisable,omitempty"`

  // The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.
  SourceIpamPoolId string `json:"SourceIpamPoolId,omitempty"`

  // The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete"
  State string `json:"State,omitempty"`

  // An explanation of how the pool arrived at it current state.
  StateMessage string `json:"StateMessage,omitempty"`

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

func (strct *ProvisionedCidr) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Cidr" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Cidr" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Cidr\": ")
	if tmp, err := json.Marshal(strct.Cidr); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ProvisionedCidr) UnmarshalJSON(b []byte) error {
    CidrReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Cidr":
            if err := json.Unmarshal([]byte(v), &strct.Cidr); err != nil {
                return err
             }
            CidrReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Cidr (a required property) was received
    if !CidrReceived {
        return errors.New("\"Cidr\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AddressFamily" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AddressFamily" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AddressFamily\": ")
	if tmp, err := json.Marshal(strct.AddressFamily); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllocationDefaultNetmaskLength" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllocationDefaultNetmaskLength\": ")
	if tmp, err := json.Marshal(strct.AllocationDefaultNetmaskLength); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllocationMaxNetmaskLength" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllocationMaxNetmaskLength\": ")
	if tmp, err := json.Marshal(strct.AllocationMaxNetmaskLength); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllocationMinNetmaskLength" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllocationMinNetmaskLength\": ")
	if tmp, err := json.Marshal(strct.AllocationMinNetmaskLength); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllocationResourceTags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllocationResourceTags\": ")
	if tmp, err := json.Marshal(strct.AllocationResourceTags); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "AutoImport" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoImport\": ")
	if tmp, err := json.Marshal(strct.AutoImport); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AwsService" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AwsService\": ")
	if tmp, err := json.Marshal(strct.AwsService); err != nil {
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
    // Marshal the "IpamPoolId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamPoolId\": ")
	if tmp, err := json.Marshal(strct.IpamPoolId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamScopeArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamScopeArn\": ")
	if tmp, err := json.Marshal(strct.IpamScopeArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IpamScopeId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "IpamScopeId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamScopeId\": ")
	if tmp, err := json.Marshal(strct.IpamScopeId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IpamScopeType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpamScopeType\": ")
	if tmp, err := json.Marshal(strct.IpamScopeType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Locale" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Locale\": ")
	if tmp, err := json.Marshal(strct.Locale); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PoolDepth" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PoolDepth\": ")
	if tmp, err := json.Marshal(strct.PoolDepth); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProvisionedCidrs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProvisionedCidrs\": ")
	if tmp, err := json.Marshal(strct.ProvisionedCidrs); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicIpSource" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicIpSource\": ")
	if tmp, err := json.Marshal(strct.PublicIpSource); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PubliclyAdvertisable" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PubliclyAdvertisable\": ")
	if tmp, err := json.Marshal(strct.PubliclyAdvertisable); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SourceIpamPoolId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceIpamPoolId\": ")
	if tmp, err := json.Marshal(strct.SourceIpamPoolId); err != nil {
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
    // Marshal the "StateMessage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StateMessage\": ")
	if tmp, err := json.Marshal(strct.StateMessage); err != nil {
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
    AddressFamilyReceived := false
    IpamScopeIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AddressFamily":
            if err := json.Unmarshal([]byte(v), &strct.AddressFamily); err != nil {
                return err
             }
            AddressFamilyReceived = true
        case "AllocationDefaultNetmaskLength":
            if err := json.Unmarshal([]byte(v), &strct.AllocationDefaultNetmaskLength); err != nil {
                return err
             }
        case "AllocationMaxNetmaskLength":
            if err := json.Unmarshal([]byte(v), &strct.AllocationMaxNetmaskLength); err != nil {
                return err
             }
        case "AllocationMinNetmaskLength":
            if err := json.Unmarshal([]byte(v), &strct.AllocationMinNetmaskLength); err != nil {
                return err
             }
        case "AllocationResourceTags":
            if err := json.Unmarshal([]byte(v), &strct.AllocationResourceTags); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AutoImport":
            if err := json.Unmarshal([]byte(v), &strct.AutoImport); err != nil {
                return err
             }
        case "AwsService":
            if err := json.Unmarshal([]byte(v), &strct.AwsService); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "IpamArn":
            if err := json.Unmarshal([]byte(v), &strct.IpamArn); err != nil {
                return err
             }
        case "IpamPoolId":
            if err := json.Unmarshal([]byte(v), &strct.IpamPoolId); err != nil {
                return err
             }
        case "IpamScopeArn":
            if err := json.Unmarshal([]byte(v), &strct.IpamScopeArn); err != nil {
                return err
             }
        case "IpamScopeId":
            if err := json.Unmarshal([]byte(v), &strct.IpamScopeId); err != nil {
                return err
             }
            IpamScopeIdReceived = true
        case "IpamScopeType":
            if err := json.Unmarshal([]byte(v), &strct.IpamScopeType); err != nil {
                return err
             }
        case "Locale":
            if err := json.Unmarshal([]byte(v), &strct.Locale); err != nil {
                return err
             }
        case "PoolDepth":
            if err := json.Unmarshal([]byte(v), &strct.PoolDepth); err != nil {
                return err
             }
        case "ProvisionedCidrs":
            if err := json.Unmarshal([]byte(v), &strct.ProvisionedCidrs); err != nil {
                return err
             }
        case "PublicIpSource":
            if err := json.Unmarshal([]byte(v), &strct.PublicIpSource); err != nil {
                return err
             }
        case "PubliclyAdvertisable":
            if err := json.Unmarshal([]byte(v), &strct.PubliclyAdvertisable); err != nil {
                return err
             }
        case "SourceIpamPoolId":
            if err := json.Unmarshal([]byte(v), &strct.SourceIpamPoolId); err != nil {
                return err
             }
        case "State":
            if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
                return err
             }
        case "StateMessage":
            if err := json.Unmarshal([]byte(v), &strct.StateMessage); err != nil {
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
    // check if AddressFamily (a required property) was received
    if !AddressFamilyReceived {
        return errors.New("\"AddressFamily\" is required but was not present")
    }
    // check if IpamScopeId (a required property) was received
    if !IpamScopeIdReceived {
        return errors.New("\"IpamScopeId\" is required but was not present")
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