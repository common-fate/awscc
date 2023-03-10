// Code generated by schema-generate. DO NOT EDIT.

package policy

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// IEMap An FMS includeMap or excludeMap.
type IEMap struct {
  ACCOUNT []string `json:"ACCOUNT,omitempty"`
  ORGUNIT []string `json:"ORGUNIT,omitempty"`
}

// NetworkFirewallPolicy Network firewall policy.
type NetworkFirewallPolicy struct {
  FirewallDeploymentModel string `json:"FirewallDeploymentModel"`
}

// PolicyOption Firewall policy option.
type PolicyOption struct {
  NetworkFirewallPolicy *NetworkFirewallPolicy `json:"NetworkFirewallPolicy,omitempty"`
  ThirdPartyFirewallPolicy *ThirdPartyFirewallPolicy `json:"ThirdPartyFirewallPolicy,omitempty"`
}

// PolicyTag A policy tag.
type PolicyTag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

// Resource Creates an AWS Firewall Manager policy.
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  DeleteAllPolicyResources bool `json:"DeleteAllPolicyResources,omitempty"`
  ExcludeMap *IEMap `json:"ExcludeMap,omitempty"`
  ExcludeResourceTags bool `json:"ExcludeResourceTags"`
  Id string `json:"Id,omitempty"`
  IncludeMap *IEMap `json:"IncludeMap,omitempty"`
  PolicyDescription string `json:"PolicyDescription,omitempty"`
  PolicyName string `json:"PolicyName"`
  RemediationEnabled bool `json:"RemediationEnabled"`
  ResourceSetIds []string `json:"ResourceSetIds,omitempty"`
  ResourceTags []*ResourceTag `json:"ResourceTags,omitempty"`
  ResourceType string `json:"ResourceType,omitempty"`
  ResourceTypeList []string `json:"ResourceTypeList,omitempty"`
  ResourcesCleanUp bool `json:"ResourcesCleanUp,omitempty"`
  SecurityServicePolicyData *SecurityServicePolicyData `json:"SecurityServicePolicyData"`
  Tags []*PolicyTag `json:"Tags,omitempty"`
}

// ResourceTag A resource tag.
type ResourceTag struct {
  Key string `json:"Key"`
  Value string `json:"Value,omitempty"`
}

// SecurityServicePolicyData Firewall security service policy data.
type SecurityServicePolicyData struct {
  ManagedServiceData string `json:"ManagedServiceData,omitempty"`
  PolicyOption *PolicyOption `json:"PolicyOption,omitempty"`
  Type string `json:"Type"`
}

// ThirdPartyFirewallPolicy Third party firewall policy.
type ThirdPartyFirewallPolicy struct {
  FirewallDeploymentModel string `json:"FirewallDeploymentModel"`
}

func (strct *IEMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ACCOUNT" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ACCOUNT\": ")
	if tmp, err := json.Marshal(strct.ACCOUNT); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ORGUNIT" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ORGUNIT\": ")
	if tmp, err := json.Marshal(strct.ORGUNIT); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *IEMap) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ACCOUNT":
            if err := json.Unmarshal([]byte(v), &strct.ACCOUNT); err != nil {
                return err
             }
        case "ORGUNIT":
            if err := json.Unmarshal([]byte(v), &strct.ORGUNIT); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *NetworkFirewallPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "FirewallDeploymentModel" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FirewallDeploymentModel" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallDeploymentModel\": ")
	if tmp, err := json.Marshal(strct.FirewallDeploymentModel); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *NetworkFirewallPolicy) UnmarshalJSON(b []byte) error {
    FirewallDeploymentModelReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FirewallDeploymentModel":
            if err := json.Unmarshal([]byte(v), &strct.FirewallDeploymentModel); err != nil {
                return err
             }
            FirewallDeploymentModelReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FirewallDeploymentModel (a required property) was received
    if !FirewallDeploymentModelReceived {
        return errors.New("\"FirewallDeploymentModel\" is required but was not present")
    }
    return nil
}

func (strct *PolicyOption) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "NetworkFirewallPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NetworkFirewallPolicy\": ")
	if tmp, err := json.Marshal(strct.NetworkFirewallPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ThirdPartyFirewallPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ThirdPartyFirewallPolicy\": ")
	if tmp, err := json.Marshal(strct.ThirdPartyFirewallPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PolicyOption) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "NetworkFirewallPolicy":
            if err := json.Unmarshal([]byte(v), &strct.NetworkFirewallPolicy); err != nil {
                return err
             }
        case "ThirdPartyFirewallPolicy":
            if err := json.Unmarshal([]byte(v), &strct.ThirdPartyFirewallPolicy); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *PolicyTag) MarshalJSON() ([]byte, error) {
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

func (strct *PolicyTag) UnmarshalJSON(b []byte) error {
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
    // Marshal the "DeleteAllPolicyResources" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeleteAllPolicyResources\": ")
	if tmp, err := json.Marshal(strct.DeleteAllPolicyResources); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ExcludeMap" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ExcludeMap\": ")
	if tmp, err := json.Marshal(strct.ExcludeMap); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ExcludeResourceTags" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ExcludeResourceTags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ExcludeResourceTags\": ")
	if tmp, err := json.Marshal(strct.ExcludeResourceTags); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IncludeMap" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IncludeMap\": ")
	if tmp, err := json.Marshal(strct.IncludeMap); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PolicyDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyDescription\": ")
	if tmp, err := json.Marshal(strct.PolicyDescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PolicyName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyName\": ")
	if tmp, err := json.Marshal(strct.PolicyName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RemediationEnabled" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RemediationEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RemediationEnabled\": ")
	if tmp, err := json.Marshal(strct.RemediationEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceSetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceSetIds\": ")
	if tmp, err := json.Marshal(strct.ResourceSetIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceTags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceTags\": ")
	if tmp, err := json.Marshal(strct.ResourceTags); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceType\": ")
	if tmp, err := json.Marshal(strct.ResourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceTypeList" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceTypeList\": ")
	if tmp, err := json.Marshal(strct.ResourceTypeList); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourcesCleanUp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourcesCleanUp\": ")
	if tmp, err := json.Marshal(strct.ResourcesCleanUp); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SecurityServicePolicyData" field is required
    if strct.SecurityServicePolicyData == nil {
        return nil, errors.New("SecurityServicePolicyData is a required field")
    }
    // Marshal the "SecurityServicePolicyData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityServicePolicyData\": ")
	if tmp, err := json.Marshal(strct.SecurityServicePolicyData); err != nil {
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
    ExcludeResourceTagsReceived := false
    PolicyNameReceived := false
    RemediationEnabledReceived := false
    SecurityServicePolicyDataReceived := false
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
        case "DeleteAllPolicyResources":
            if err := json.Unmarshal([]byte(v), &strct.DeleteAllPolicyResources); err != nil {
                return err
             }
        case "ExcludeMap":
            if err := json.Unmarshal([]byte(v), &strct.ExcludeMap); err != nil {
                return err
             }
        case "ExcludeResourceTags":
            if err := json.Unmarshal([]byte(v), &strct.ExcludeResourceTags); err != nil {
                return err
             }
            ExcludeResourceTagsReceived = true
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "IncludeMap":
            if err := json.Unmarshal([]byte(v), &strct.IncludeMap); err != nil {
                return err
             }
        case "PolicyDescription":
            if err := json.Unmarshal([]byte(v), &strct.PolicyDescription); err != nil {
                return err
             }
        case "PolicyName":
            if err := json.Unmarshal([]byte(v), &strct.PolicyName); err != nil {
                return err
             }
            PolicyNameReceived = true
        case "RemediationEnabled":
            if err := json.Unmarshal([]byte(v), &strct.RemediationEnabled); err != nil {
                return err
             }
            RemediationEnabledReceived = true
        case "ResourceSetIds":
            if err := json.Unmarshal([]byte(v), &strct.ResourceSetIds); err != nil {
                return err
             }
        case "ResourceTags":
            if err := json.Unmarshal([]byte(v), &strct.ResourceTags); err != nil {
                return err
             }
        case "ResourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
        case "ResourceTypeList":
            if err := json.Unmarshal([]byte(v), &strct.ResourceTypeList); err != nil {
                return err
             }
        case "ResourcesCleanUp":
            if err := json.Unmarshal([]byte(v), &strct.ResourcesCleanUp); err != nil {
                return err
             }
        case "SecurityServicePolicyData":
            if err := json.Unmarshal([]byte(v), &strct.SecurityServicePolicyData); err != nil {
                return err
             }
            SecurityServicePolicyDataReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ExcludeResourceTags (a required property) was received
    if !ExcludeResourceTagsReceived {
        return errors.New("\"ExcludeResourceTags\" is required but was not present")
    }
    // check if PolicyName (a required property) was received
    if !PolicyNameReceived {
        return errors.New("\"PolicyName\" is required but was not present")
    }
    // check if RemediationEnabled (a required property) was received
    if !RemediationEnabledReceived {
        return errors.New("\"RemediationEnabled\" is required but was not present")
    }
    // check if SecurityServicePolicyData (a required property) was received
    if !SecurityServicePolicyDataReceived {
        return errors.New("\"SecurityServicePolicyData\" is required but was not present")
    }
    return nil
}

func (strct *ResourceTag) MarshalJSON() ([]byte, error) {
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

func (strct *ResourceTag) UnmarshalJSON(b []byte) error {
    KeyReceived := false
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
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Key (a required property) was received
    if !KeyReceived {
        return errors.New("\"Key\" is required but was not present")
    }
    return nil
}

func (strct *SecurityServicePolicyData) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ManagedServiceData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ManagedServiceData\": ")
	if tmp, err := json.Marshal(strct.ManagedServiceData); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PolicyOption" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyOption\": ")
	if tmp, err := json.Marshal(strct.PolicyOption); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Type" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SecurityServicePolicyData) UnmarshalJSON(b []byte) error {
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ManagedServiceData":
            if err := json.Unmarshal([]byte(v), &strct.ManagedServiceData); err != nil {
                return err
             }
        case "PolicyOption":
            if err := json.Unmarshal([]byte(v), &strct.PolicyOption); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}

func (strct *ThirdPartyFirewallPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "FirewallDeploymentModel" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FirewallDeploymentModel" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallDeploymentModel\": ")
	if tmp, err := json.Marshal(strct.FirewallDeploymentModel); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ThirdPartyFirewallPolicy) UnmarshalJSON(b []byte) error {
    FirewallDeploymentModelReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FirewallDeploymentModel":
            if err := json.Unmarshal([]byte(v), &strct.FirewallDeploymentModel); err != nil {
                return err
             }
            FirewallDeploymentModelReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FirewallDeploymentModel (a required property) was received
    if !FirewallDeploymentModelReceived {
        return errors.New("\"FirewallDeploymentModel\" is required but was not present")
    }
    return nil
}
