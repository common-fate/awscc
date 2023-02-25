// Code generated by schema-generate. DO NOT EDIT.

package firewallpolicy

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// ActionDefinition 
type ActionDefinition struct {
  PublishMetricAction *PublishMetricAction `json:"PublishMetricAction,omitempty"`
}

// CustomAction 
type CustomAction struct {
  ActionDefinition *ActionDefinition `json:"ActionDefinition"`
  ActionName string `json:"ActionName"`
}

// Dimension 
type Dimension struct {
  Value string `json:"Value"`
}

// FirewallPolicy 
type FirewallPolicy struct {
  StatefulDefaultActions []string `json:"StatefulDefaultActions,omitempty"`
  StatefulEngineOptions *StatefulEngineOptions `json:"StatefulEngineOptions,omitempty"`
  StatefulRuleGroupReferences []*StatefulRuleGroupReference `json:"StatefulRuleGroupReferences,omitempty"`
  StatelessCustomActions []*CustomAction `json:"StatelessCustomActions,omitempty"`
  StatelessDefaultActions []string `json:"StatelessDefaultActions"`
  StatelessFragmentDefaultActions []string `json:"StatelessFragmentDefaultActions"`
  StatelessRuleGroupReferences []*StatelessRuleGroupReference `json:"StatelessRuleGroupReferences,omitempty"`
}

// PublishMetricAction 
type PublishMetricAction struct {
  Dimensions []*Dimension `json:"Dimensions"`
}

// Resource Resource type definition for AWS::NetworkFirewall::FirewallPolicy
type Resource struct {
  Description string `json:"Description,omitempty"`
  FirewallPolicy *FirewallPolicy `json:"FirewallPolicy"`
  FirewallPolicyArn string `json:"FirewallPolicyArn,omitempty"`
  FirewallPolicyId string `json:"FirewallPolicyId,omitempty"`
  FirewallPolicyName string `json:"FirewallPolicyName"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// StatefulEngineOptions 
type StatefulEngineOptions struct {
  RuleOrder string `json:"RuleOrder,omitempty"`
  StreamExceptionPolicy string `json:"StreamExceptionPolicy,omitempty"`
}

// StatefulRuleGroupOverride 
type StatefulRuleGroupOverride struct {
  Action string `json:"Action,omitempty"`
}

// StatefulRuleGroupReference 
type StatefulRuleGroupReference struct {
  Override *StatefulRuleGroupOverride `json:"Override,omitempty"`
  Priority int `json:"Priority,omitempty"`
  ResourceArn string `json:"ResourceArn"`
}

// StatelessRuleGroupReference 
type StatelessRuleGroupReference struct {
  Priority int `json:"Priority"`
  ResourceArn string `json:"ResourceArn"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *ActionDefinition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "PublishMetricAction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublishMetricAction\": ")
	if tmp, err := json.Marshal(strct.PublishMetricAction); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ActionDefinition) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PublishMetricAction":
            if err := json.Unmarshal([]byte(v), &strct.PublishMetricAction); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *CustomAction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ActionDefinition" field is required
    if strct.ActionDefinition == nil {
        return nil, errors.New("ActionDefinition is a required field")
    }
    // Marshal the "ActionDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ActionDefinition\": ")
	if tmp, err := json.Marshal(strct.ActionDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ActionName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ActionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ActionName\": ")
	if tmp, err := json.Marshal(strct.ActionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CustomAction) UnmarshalJSON(b []byte) error {
    ActionDefinitionReceived := false
    ActionNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ActionDefinition":
            if err := json.Unmarshal([]byte(v), &strct.ActionDefinition); err != nil {
                return err
             }
            ActionDefinitionReceived = true
        case "ActionName":
            if err := json.Unmarshal([]byte(v), &strct.ActionName); err != nil {
                return err
             }
            ActionNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ActionDefinition (a required property) was received
    if !ActionDefinitionReceived {
        return errors.New("\"ActionDefinition\" is required but was not present")
    }
    // check if ActionName (a required property) was received
    if !ActionNameReceived {
        return errors.New("\"ActionName\" is required but was not present")
    }
    return nil
}

func (strct *Dimension) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *Dimension) UnmarshalJSON(b []byte) error {
    ValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            ValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Value (a required property) was received
    if !ValueReceived {
        return errors.New("\"Value\" is required but was not present")
    }
    return nil
}

func (strct *FirewallPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "StatefulDefaultActions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatefulDefaultActions\": ")
	if tmp, err := json.Marshal(strct.StatefulDefaultActions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatefulEngineOptions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatefulEngineOptions\": ")
	if tmp, err := json.Marshal(strct.StatefulEngineOptions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatefulRuleGroupReferences" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatefulRuleGroupReferences\": ")
	if tmp, err := json.Marshal(strct.StatefulRuleGroupReferences); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatelessCustomActions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatelessCustomActions\": ")
	if tmp, err := json.Marshal(strct.StatelessCustomActions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "StatelessDefaultActions" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "StatelessDefaultActions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatelessDefaultActions\": ")
	if tmp, err := json.Marshal(strct.StatelessDefaultActions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "StatelessFragmentDefaultActions" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "StatelessFragmentDefaultActions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatelessFragmentDefaultActions\": ")
	if tmp, err := json.Marshal(strct.StatelessFragmentDefaultActions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatelessRuleGroupReferences" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatelessRuleGroupReferences\": ")
	if tmp, err := json.Marshal(strct.StatelessRuleGroupReferences); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *FirewallPolicy) UnmarshalJSON(b []byte) error {
    StatelessDefaultActionsReceived := false
    StatelessFragmentDefaultActionsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "StatefulDefaultActions":
            if err := json.Unmarshal([]byte(v), &strct.StatefulDefaultActions); err != nil {
                return err
             }
        case "StatefulEngineOptions":
            if err := json.Unmarshal([]byte(v), &strct.StatefulEngineOptions); err != nil {
                return err
             }
        case "StatefulRuleGroupReferences":
            if err := json.Unmarshal([]byte(v), &strct.StatefulRuleGroupReferences); err != nil {
                return err
             }
        case "StatelessCustomActions":
            if err := json.Unmarshal([]byte(v), &strct.StatelessCustomActions); err != nil {
                return err
             }
        case "StatelessDefaultActions":
            if err := json.Unmarshal([]byte(v), &strct.StatelessDefaultActions); err != nil {
                return err
             }
            StatelessDefaultActionsReceived = true
        case "StatelessFragmentDefaultActions":
            if err := json.Unmarshal([]byte(v), &strct.StatelessFragmentDefaultActions); err != nil {
                return err
             }
            StatelessFragmentDefaultActionsReceived = true
        case "StatelessRuleGroupReferences":
            if err := json.Unmarshal([]byte(v), &strct.StatelessRuleGroupReferences); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if StatelessDefaultActions (a required property) was received
    if !StatelessDefaultActionsReceived {
        return errors.New("\"StatelessDefaultActions\" is required but was not present")
    }
    // check if StatelessFragmentDefaultActions (a required property) was received
    if !StatelessFragmentDefaultActionsReceived {
        return errors.New("\"StatelessFragmentDefaultActions\" is required but was not present")
    }
    return nil
}

func (strct *PublishMetricAction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Dimensions" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Dimensions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Dimensions\": ")
	if tmp, err := json.Marshal(strct.Dimensions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PublishMetricAction) UnmarshalJSON(b []byte) error {
    DimensionsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Dimensions":
            if err := json.Unmarshal([]byte(v), &strct.Dimensions); err != nil {
                return err
             }
            DimensionsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Dimensions (a required property) was received
    if !DimensionsReceived {
        return errors.New("\"Dimensions\" is required but was not present")
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
    // "FirewallPolicy" field is required
    if strct.FirewallPolicy == nil {
        return nil, errors.New("FirewallPolicy is a required field")
    }
    // Marshal the "FirewallPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallPolicy\": ")
	if tmp, err := json.Marshal(strct.FirewallPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FirewallPolicyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallPolicyArn\": ")
	if tmp, err := json.Marshal(strct.FirewallPolicyArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FirewallPolicyId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallPolicyId\": ")
	if tmp, err := json.Marshal(strct.FirewallPolicyId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FirewallPolicyName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FirewallPolicyName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FirewallPolicyName\": ")
	if tmp, err := json.Marshal(strct.FirewallPolicyName); err != nil {
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
    FirewallPolicyReceived := false
    FirewallPolicyNameReceived := false
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
        case "FirewallPolicy":
            if err := json.Unmarshal([]byte(v), &strct.FirewallPolicy); err != nil {
                return err
             }
            FirewallPolicyReceived = true
        case "FirewallPolicyArn":
            if err := json.Unmarshal([]byte(v), &strct.FirewallPolicyArn); err != nil {
                return err
             }
        case "FirewallPolicyId":
            if err := json.Unmarshal([]byte(v), &strct.FirewallPolicyId); err != nil {
                return err
             }
        case "FirewallPolicyName":
            if err := json.Unmarshal([]byte(v), &strct.FirewallPolicyName); err != nil {
                return err
             }
            FirewallPolicyNameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FirewallPolicy (a required property) was received
    if !FirewallPolicyReceived {
        return errors.New("\"FirewallPolicy\" is required but was not present")
    }
    // check if FirewallPolicyName (a required property) was received
    if !FirewallPolicyNameReceived {
        return errors.New("\"FirewallPolicyName\" is required but was not present")
    }
    return nil
}

func (strct *StatefulEngineOptions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "RuleOrder" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RuleOrder\": ")
	if tmp, err := json.Marshal(strct.RuleOrder); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StreamExceptionPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StreamExceptionPolicy\": ")
	if tmp, err := json.Marshal(strct.StreamExceptionPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StatefulEngineOptions) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "RuleOrder":
            if err := json.Unmarshal([]byte(v), &strct.RuleOrder); err != nil {
                return err
             }
        case "StreamExceptionPolicy":
            if err := json.Unmarshal([]byte(v), &strct.StreamExceptionPolicy); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *StatefulRuleGroupOverride) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Action" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Action\": ")
	if tmp, err := json.Marshal(strct.Action); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StatefulRuleGroupOverride) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Action":
            if err := json.Unmarshal([]byte(v), &strct.Action); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *StatefulRuleGroupReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Override" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Override\": ")
	if tmp, err := json.Marshal(strct.Override); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // "ResourceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ResourceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceArn\": ")
	if tmp, err := json.Marshal(strct.ResourceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StatefulRuleGroupReference) UnmarshalJSON(b []byte) error {
    ResourceArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Override":
            if err := json.Unmarshal([]byte(v), &strct.Override); err != nil {
                return err
             }
        case "Priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
        case "ResourceArn":
            if err := json.Unmarshal([]byte(v), &strct.ResourceArn); err != nil {
                return err
             }
            ResourceArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ResourceArn (a required property) was received
    if !ResourceArnReceived {
        return errors.New("\"ResourceArn\" is required but was not present")
    }
    return nil
}

func (strct *StatelessRuleGroupReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // "ResourceArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ResourceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceArn\": ")
	if tmp, err := json.Marshal(strct.ResourceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StatelessRuleGroupReference) UnmarshalJSON(b []byte) error {
    PriorityReceived := false
    ResourceArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
            PriorityReceived = true
        case "ResourceArn":
            if err := json.Unmarshal([]byte(v), &strct.ResourceArn); err != nil {
                return err
             }
            ResourceArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Priority (a required property) was received
    if !PriorityReceived {
        return errors.New("\"Priority\" is required but was not present")
    }
    // check if ResourceArn (a required property) was received
    if !ResourceArnReceived {
        return errors.New("\"ResourceArn\" is required but was not present")
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