// Code generated by schema-generate. DO NOT EDIT.

package workgroup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ConfigParameter 
type ConfigParameter struct {
  ParameterKey string `json:"ParameterKey,omitempty"`
  ParameterValue string `json:"ParameterValue,omitempty"`
}

// Endpoint 
type Endpoint struct {
  Address string `json:"Address,omitempty"`
  Port int `json:"Port,omitempty"`
  VpcEndpoints []*VpcEndpoint `json:"VpcEndpoints,omitempty"`
}

// NetworkInterface 
type NetworkInterface struct {
  AvailabilityZone string `json:"AvailabilityZone,omitempty"`
  NetworkInterfaceId string `json:"NetworkInterfaceId,omitempty"`
  PrivateIpAddress string `json:"PrivateIpAddress,omitempty"`
  SubnetId string `json:"SubnetId,omitempty"`
}

// Resource Definition of AWS::RedshiftServerless::Workgroup Resource Type
type Resource struct {
  BaseCapacity int `json:"BaseCapacity,omitempty"`
  ConfigParameters []*ConfigParameter `json:"ConfigParameters,omitempty"`
  EnhancedVpcRouting bool `json:"EnhancedVpcRouting,omitempty"`
  NamespaceName string `json:"NamespaceName,omitempty"`
  Port int `json:"Port,omitempty"`
  PubliclyAccessible bool `json:"PubliclyAccessible,omitempty"`
  SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
  SubnetIds []string `json:"SubnetIds,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
  Workgroup *Workgroup `json:"Workgroup,omitempty"`
  WorkgroupName string `json:"WorkgroupName"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

// VpcEndpoint 
type VpcEndpoint struct {
  NetworkInterfaces []*NetworkInterface `json:"NetworkInterfaces,omitempty"`
  VpcEndpointId string `json:"VpcEndpointId,omitempty"`
  VpcId string `json:"VpcId,omitempty"`
}

// Workgroup 
type Workgroup struct {
  BaseCapacity int `json:"BaseCapacity,omitempty"`
  ConfigParameters []*ConfigParameter `json:"ConfigParameters,omitempty"`
  CreationDate string `json:"CreationDate,omitempty"`
  Endpoint *Endpoint `json:"Endpoint,omitempty"`
  EnhancedVpcRouting bool `json:"EnhancedVpcRouting,omitempty"`
  NamespaceName string `json:"NamespaceName,omitempty"`
  PubliclyAccessible bool `json:"PubliclyAccessible,omitempty"`
  SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
  Status string `json:"Status,omitempty"`
  SubnetIds []string `json:"SubnetIds,omitempty"`
  WorkgroupArn string `json:"WorkgroupArn,omitempty"`
  WorkgroupId string `json:"WorkgroupId,omitempty"`
  WorkgroupName string `json:"WorkgroupName,omitempty"`
}

func (strct *ConfigParameter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ParameterKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParameterKey\": ")
	if tmp, err := json.Marshal(strct.ParameterKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ParameterValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParameterValue\": ")
	if tmp, err := json.Marshal(strct.ParameterValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ConfigParameter) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ParameterKey":
            if err := json.Unmarshal([]byte(v), &strct.ParameterKey); err != nil {
                return err
             }
        case "ParameterValue":
            if err := json.Unmarshal([]byte(v), &strct.ParameterValue); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Endpoint) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Address" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Address\": ")
	if tmp, err := json.Marshal(strct.Address); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Port" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Port\": ")
	if tmp, err := json.Marshal(strct.Port); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "VpcEndpoints" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"VpcEndpoints\": ")
	if tmp, err := json.Marshal(strct.VpcEndpoints); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Endpoint) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Address":
            if err := json.Unmarshal([]byte(v), &strct.Address); err != nil {
                return err
             }
        case "Port":
            if err := json.Unmarshal([]byte(v), &strct.Port); err != nil {
                return err
             }
        case "VpcEndpoints":
            if err := json.Unmarshal([]byte(v), &strct.VpcEndpoints); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *NetworkInterface) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AvailabilityZone" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AvailabilityZone\": ")
	if tmp, err := json.Marshal(strct.AvailabilityZone); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NetworkInterfaceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NetworkInterfaceId\": ")
	if tmp, err := json.Marshal(strct.NetworkInterfaceId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PrivateIpAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PrivateIpAddress\": ")
	if tmp, err := json.Marshal(strct.PrivateIpAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SubnetId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetId\": ")
	if tmp, err := json.Marshal(strct.SubnetId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *NetworkInterface) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AvailabilityZone":
            if err := json.Unmarshal([]byte(v), &strct.AvailabilityZone); err != nil {
                return err
             }
        case "NetworkInterfaceId":
            if err := json.Unmarshal([]byte(v), &strct.NetworkInterfaceId); err != nil {
                return err
             }
        case "PrivateIpAddress":
            if err := json.Unmarshal([]byte(v), &strct.PrivateIpAddress); err != nil {
                return err
             }
        case "SubnetId":
            if err := json.Unmarshal([]byte(v), &strct.SubnetId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "BaseCapacity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BaseCapacity\": ")
	if tmp, err := json.Marshal(strct.BaseCapacity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ConfigParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConfigParameters\": ")
	if tmp, err := json.Marshal(strct.ConfigParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnhancedVpcRouting" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnhancedVpcRouting\": ")
	if tmp, err := json.Marshal(strct.EnhancedVpcRouting); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NamespaceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NamespaceName\": ")
	if tmp, err := json.Marshal(strct.NamespaceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Port" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Port\": ")
	if tmp, err := json.Marshal(strct.Port); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PubliclyAccessible" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PubliclyAccessible\": ")
	if tmp, err := json.Marshal(strct.PubliclyAccessible); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecurityGroupIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityGroupIds\": ")
	if tmp, err := json.Marshal(strct.SecurityGroupIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SubnetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetIds\": ")
	if tmp, err := json.Marshal(strct.SubnetIds); err != nil {
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
    // Marshal the "Workgroup" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Workgroup\": ")
	if tmp, err := json.Marshal(strct.Workgroup); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "WorkgroupName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "WorkgroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkgroupName\": ")
	if tmp, err := json.Marshal(strct.WorkgroupName); err != nil {
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
    WorkgroupNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BaseCapacity":
            if err := json.Unmarshal([]byte(v), &strct.BaseCapacity); err != nil {
                return err
             }
        case "ConfigParameters":
            if err := json.Unmarshal([]byte(v), &strct.ConfigParameters); err != nil {
                return err
             }
        case "EnhancedVpcRouting":
            if err := json.Unmarshal([]byte(v), &strct.EnhancedVpcRouting); err != nil {
                return err
             }
        case "NamespaceName":
            if err := json.Unmarshal([]byte(v), &strct.NamespaceName); err != nil {
                return err
             }
        case "Port":
            if err := json.Unmarshal([]byte(v), &strct.Port); err != nil {
                return err
             }
        case "PubliclyAccessible":
            if err := json.Unmarshal([]byte(v), &strct.PubliclyAccessible); err != nil {
                return err
             }
        case "SecurityGroupIds":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupIds); err != nil {
                return err
             }
        case "SubnetIds":
            if err := json.Unmarshal([]byte(v), &strct.SubnetIds); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Workgroup":
            if err := json.Unmarshal([]byte(v), &strct.Workgroup); err != nil {
                return err
             }
        case "WorkgroupName":
            if err := json.Unmarshal([]byte(v), &strct.WorkgroupName); err != nil {
                return err
             }
            WorkgroupNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if WorkgroupName (a required property) was received
    if !WorkgroupNameReceived {
        return errors.New("\"WorkgroupName\" is required but was not present")
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

func (strct *VpcEndpoint) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "NetworkInterfaces" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NetworkInterfaces\": ")
	if tmp, err := json.Marshal(strct.NetworkInterfaces); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "VpcEndpointId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"VpcEndpointId\": ")
	if tmp, err := json.Marshal(strct.VpcEndpointId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "VpcId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"VpcId\": ")
	if tmp, err := json.Marshal(strct.VpcId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *VpcEndpoint) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "NetworkInterfaces":
            if err := json.Unmarshal([]byte(v), &strct.NetworkInterfaces); err != nil {
                return err
             }
        case "VpcEndpointId":
            if err := json.Unmarshal([]byte(v), &strct.VpcEndpointId); err != nil {
                return err
             }
        case "VpcId":
            if err := json.Unmarshal([]byte(v), &strct.VpcId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Workgroup) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "BaseCapacity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BaseCapacity\": ")
	if tmp, err := json.Marshal(strct.BaseCapacity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ConfigParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConfigParameters\": ")
	if tmp, err := json.Marshal(strct.ConfigParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CreationDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreationDate\": ")
	if tmp, err := json.Marshal(strct.CreationDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Endpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Endpoint\": ")
	if tmp, err := json.Marshal(strct.Endpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnhancedVpcRouting" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnhancedVpcRouting\": ")
	if tmp, err := json.Marshal(strct.EnhancedVpcRouting); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NamespaceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NamespaceName\": ")
	if tmp, err := json.Marshal(strct.NamespaceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PubliclyAccessible" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PubliclyAccessible\": ")
	if tmp, err := json.Marshal(strct.PubliclyAccessible); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecurityGroupIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityGroupIds\": ")
	if tmp, err := json.Marshal(strct.SecurityGroupIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "SubnetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetIds\": ")
	if tmp, err := json.Marshal(strct.SubnetIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "WorkgroupArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkgroupArn\": ")
	if tmp, err := json.Marshal(strct.WorkgroupArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "WorkgroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkgroupId\": ")
	if tmp, err := json.Marshal(strct.WorkgroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "WorkgroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkgroupName\": ")
	if tmp, err := json.Marshal(strct.WorkgroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Workgroup) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BaseCapacity":
            if err := json.Unmarshal([]byte(v), &strct.BaseCapacity); err != nil {
                return err
             }
        case "ConfigParameters":
            if err := json.Unmarshal([]byte(v), &strct.ConfigParameters); err != nil {
                return err
             }
        case "CreationDate":
            if err := json.Unmarshal([]byte(v), &strct.CreationDate); err != nil {
                return err
             }
        case "Endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
        case "EnhancedVpcRouting":
            if err := json.Unmarshal([]byte(v), &strct.EnhancedVpcRouting); err != nil {
                return err
             }
        case "NamespaceName":
            if err := json.Unmarshal([]byte(v), &strct.NamespaceName); err != nil {
                return err
             }
        case "PubliclyAccessible":
            if err := json.Unmarshal([]byte(v), &strct.PubliclyAccessible); err != nil {
                return err
             }
        case "SecurityGroupIds":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupIds); err != nil {
                return err
             }
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "SubnetIds":
            if err := json.Unmarshal([]byte(v), &strct.SubnetIds); err != nil {
                return err
             }
        case "WorkgroupArn":
            if err := json.Unmarshal([]byte(v), &strct.WorkgroupArn); err != nil {
                return err
             }
        case "WorkgroupId":
            if err := json.Unmarshal([]byte(v), &strct.WorkgroupId); err != nil {
                return err
             }
        case "WorkgroupName":
            if err := json.Unmarshal([]byte(v), &strct.WorkgroupName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
