// Code generated by schema-generate. DO NOT EDIT.

package resolverrule

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::Route53Resolver::ResolverRule
type Resource struct {

  // The Amazon Resource Name (ARN) of the resolver rule.
  Arn string `json:"Arn,omitempty"`

  // DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
  DomainName string `json:"DomainName"`

  // The name for the Resolver rule
  Name string `json:"Name,omitempty"`

  // The ID of the endpoint that the rule is associated with.
  ResolverEndpointId string `json:"ResolverEndpointId,omitempty"`

  // The ID of the endpoint that the rule is associated with.
  ResolverRuleId string `json:"ResolverRuleId,omitempty"`

  // When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
  RuleType string `json:"RuleType"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
  TargetIps []*TargetAddress `json:"TargetIps,omitempty"`
}

// Tag 
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

// TargetAddress 
type TargetAddress struct {

  // One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. 
  Ip string `json:"Ip,omitempty"`

  // One IPv6 address that you want to forward DNS queries to. You can specify only IPv6 addresses. 
  Ipv6 string `json:"Ipv6,omitempty"`

  // The port at Ip that you want to forward DNS queries to. 
  Port string `json:"Port,omitempty"`
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
    // Marshal the "ResolverEndpointId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResolverEndpointId\": ")
	if tmp, err := json.Marshal(strct.ResolverEndpointId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResolverRuleId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResolverRuleId\": ")
	if tmp, err := json.Marshal(strct.ResolverRuleId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RuleType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RuleType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RuleType\": ")
	if tmp, err := json.Marshal(strct.RuleType); err != nil {
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
    // Marshal the "TargetIps" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetIps\": ")
	if tmp, err := json.Marshal(strct.TargetIps); err != nil {
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
    RuleTypeReceived := false
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
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
            DomainNameReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "ResolverEndpointId":
            if err := json.Unmarshal([]byte(v), &strct.ResolverEndpointId); err != nil {
                return err
             }
        case "ResolverRuleId":
            if err := json.Unmarshal([]byte(v), &strct.ResolverRuleId); err != nil {
                return err
             }
        case "RuleType":
            if err := json.Unmarshal([]byte(v), &strct.RuleType); err != nil {
                return err
             }
            RuleTypeReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TargetIps":
            if err := json.Unmarshal([]byte(v), &strct.TargetIps); err != nil {
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
    // check if RuleType (a required property) was received
    if !RuleTypeReceived {
        return errors.New("\"RuleType\" is required but was not present")
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

func (strct *TargetAddress) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Ip" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ip\": ")
	if tmp, err := json.Marshal(strct.Ip); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Ipv6" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ipv6\": ")
	if tmp, err := json.Marshal(strct.Ipv6); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TargetAddress) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Ip":
            if err := json.Unmarshal([]byte(v), &strct.Ip); err != nil {
                return err
             }
        case "Ipv6":
            if err := json.Unmarshal([]byte(v), &strct.Ipv6); err != nil {
                return err
             }
        case "Port":
            if err := json.Unmarshal([]byte(v), &strct.Port); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
