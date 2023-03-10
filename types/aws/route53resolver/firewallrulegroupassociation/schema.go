// Code generated by schema-generate. DO NOT EDIT.

package firewallrulegroupassociation

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.
type Resource struct {

  // Arn
  Arn string `json:"Arn,omitempty"`

  // Rfc3339TimeString
  CreationTime string `json:"CreationTime,omitempty"`

  // The id of the creator request.
  CreatorRequestId string `json:"CreatorRequestId,omitempty"`

  // FirewallRuleGroupId
  FirewallRuleGroupId string `json:"FirewallRuleGroupId"`

  // Id
  Id string `json:"Id,omitempty"`

  // ServicePrincipal
  ManagedOwnerName string `json:"ManagedOwnerName,omitempty"`

  // Rfc3339TimeString
  ModificationTime string `json:"ModificationTime,omitempty"`

  // MutationProtectionStatus
  MutationProtection string `json:"MutationProtection,omitempty"`

  // FirewallRuleGroupAssociationName
  Name string `json:"Name,omitempty"`

  // Priority
  Priority int `json:"Priority"`

  // ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
  Status string `json:"Status,omitempty"`

  // FirewallDomainListAssociationStatus
  StatusMessage string `json:"StatusMessage,omitempty"`

  // Tags
  Tags []*Tag `json:"Tags,omitempty"`

  // VpcId
  VpcId string `json:"VpcId"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
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
    // Marshal the "CreationTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreationTime\": ")
	if tmp, err := json.Marshal(strct.CreationTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CreatorRequestId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreatorRequestId\": ")
	if tmp, err := json.Marshal(strct.CreatorRequestId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FirewallRuleGroupId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FirewallRuleGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallRuleGroupId\": ")
	if tmp, err := json.Marshal(strct.FirewallRuleGroupId); err != nil {
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
    // Marshal the "ManagedOwnerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ManagedOwnerName\": ")
	if tmp, err := json.Marshal(strct.ManagedOwnerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ModificationTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ModificationTime\": ")
	if tmp, err := json.Marshal(strct.ModificationTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MutationProtection" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MutationProtection\": ")
	if tmp, err := json.Marshal(strct.MutationProtection); err != nil {
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
    // "Priority" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Priority" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Priority\": ")
	if tmp, err := json.Marshal(strct.Priority); err != nil {
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
    // Marshal the "StatusMessage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatusMessage\": ")
	if tmp, err := json.Marshal(strct.StatusMessage); err != nil {
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
    // "VpcId" field is required
    // only required object types supported for marshal checking (for now)
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
    FirewallRuleGroupIdReceived := false
    PriorityReceived := false
    VpcIdReceived := false
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
        case "CreationTime":
            if err := json.Unmarshal([]byte(v), &strct.CreationTime); err != nil {
                return err
             }
        case "CreatorRequestId":
            if err := json.Unmarshal([]byte(v), &strct.CreatorRequestId); err != nil {
                return err
             }
        case "FirewallRuleGroupId":
            if err := json.Unmarshal([]byte(v), &strct.FirewallRuleGroupId); err != nil {
                return err
             }
            FirewallRuleGroupIdReceived = true
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "ManagedOwnerName":
            if err := json.Unmarshal([]byte(v), &strct.ManagedOwnerName); err != nil {
                return err
             }
        case "ModificationTime":
            if err := json.Unmarshal([]byte(v), &strct.ModificationTime); err != nil {
                return err
             }
        case "MutationProtection":
            if err := json.Unmarshal([]byte(v), &strct.MutationProtection); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "Priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
            PriorityReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "StatusMessage":
            if err := json.Unmarshal([]byte(v), &strct.StatusMessage); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "VpcId":
            if err := json.Unmarshal([]byte(v), &strct.VpcId); err != nil {
                return err
             }
            VpcIdReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FirewallRuleGroupId (a required property) was received
    if !FirewallRuleGroupIdReceived {
        return errors.New("\"FirewallRuleGroupId\" is required but was not present")
    }
    // check if Priority (a required property) was received
    if !PriorityReceived {
        return errors.New("\"Priority\" is required but was not present")
    }
    // check if VpcId (a required property) was received
    if !VpcIdReceived {
        return errors.New("\"VpcId\" is required but was not present")
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
