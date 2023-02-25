// Code generated by schema-generate. DO NOT EDIT.

package budgetsaction

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ActionThreshold 
type ActionThreshold struct {
  Type string `json:"Type"`
  Value float64 `json:"Value"`
}

// Definition 
type Definition struct {
  IamActionDefinition *IamActionDefinition `json:"IamActionDefinition,omitempty"`
  ScpActionDefinition *ScpActionDefinition `json:"ScpActionDefinition,omitempty"`
  SsmActionDefinition *SsmActionDefinition `json:"SsmActionDefinition,omitempty"`
}

// IamActionDefinition 
type IamActionDefinition struct {
  Groups []string `json:"Groups,omitempty"`
  PolicyArn string `json:"PolicyArn"`
  Roles []string `json:"Roles,omitempty"`
  Users []string `json:"Users,omitempty"`
}

// Resource An example resource schema demonstrating some basic constructs and validation rules.
type Resource struct {
  ActionId string `json:"ActionId,omitempty"`
  ActionThreshold *ActionThreshold `json:"ActionThreshold"`
  ActionType string `json:"ActionType"`
  ApprovalModel string `json:"ApprovalModel,omitempty"`
  BudgetName string `json:"BudgetName"`
  Definition *Definition `json:"Definition"`
  ExecutionRoleArn string `json:"ExecutionRoleArn"`
  NotificationType string `json:"NotificationType"`
  Subscribers []*Subscriber `json:"Subscribers"`
}

// ScpActionDefinition 
type ScpActionDefinition struct {
  PolicyId string `json:"PolicyId"`
  TargetIds []string `json:"TargetIds"`
}

// SsmActionDefinition 
type SsmActionDefinition struct {
  InstanceIds []string `json:"InstanceIds"`
  Region string `json:"Region"`
  Subtype string `json:"Subtype"`
}

// Subscriber 
type Subscriber struct {
  Address string `json:"Address"`
  Type string `json:"Type"`
}

func (strct *ActionThreshold) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *ActionThreshold) UnmarshalJSON(b []byte) error {
    TypeReceived := false
    ValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            ValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    // check if Value (a required property) was received
    if !ValueReceived {
        return errors.New("\"Value\" is required but was not present")
    }
    return nil
}

func (strct *Definition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "IamActionDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IamActionDefinition\": ")
	if tmp, err := json.Marshal(strct.IamActionDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ScpActionDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScpActionDefinition\": ")
	if tmp, err := json.Marshal(strct.ScpActionDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SsmActionDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SsmActionDefinition\": ")
	if tmp, err := json.Marshal(strct.SsmActionDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Definition) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "IamActionDefinition":
            if err := json.Unmarshal([]byte(v), &strct.IamActionDefinition); err != nil {
                return err
             }
        case "ScpActionDefinition":
            if err := json.Unmarshal([]byte(v), &strct.ScpActionDefinition); err != nil {
                return err
             }
        case "SsmActionDefinition":
            if err := json.Unmarshal([]byte(v), &strct.SsmActionDefinition); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *IamActionDefinition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Groups" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Groups\": ")
	if tmp, err := json.Marshal(strct.Groups); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PolicyArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyArn\": ")
	if tmp, err := json.Marshal(strct.PolicyArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Roles" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Roles\": ")
	if tmp, err := json.Marshal(strct.Roles); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Users" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Users\": ")
	if tmp, err := json.Marshal(strct.Users); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *IamActionDefinition) UnmarshalJSON(b []byte) error {
    PolicyArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Groups":
            if err := json.Unmarshal([]byte(v), &strct.Groups); err != nil {
                return err
             }
        case "PolicyArn":
            if err := json.Unmarshal([]byte(v), &strct.PolicyArn); err != nil {
                return err
             }
            PolicyArnReceived = true
        case "Roles":
            if err := json.Unmarshal([]byte(v), &strct.Roles); err != nil {
                return err
             }
        case "Users":
            if err := json.Unmarshal([]byte(v), &strct.Users); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if PolicyArn (a required property) was received
    if !PolicyArnReceived {
        return errors.New("\"PolicyArn\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ActionId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ActionId\": ")
	if tmp, err := json.Marshal(strct.ActionId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ActionThreshold" field is required
    if strct.ActionThreshold == nil {
        return nil, errors.New("ActionThreshold is a required field")
    }
    // Marshal the "ActionThreshold" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ActionThreshold\": ")
	if tmp, err := json.Marshal(strct.ActionThreshold); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ActionType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ActionType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ActionType\": ")
	if tmp, err := json.Marshal(strct.ActionType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ApprovalModel" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApprovalModel\": ")
	if tmp, err := json.Marshal(strct.ApprovalModel); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "BudgetName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BudgetName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BudgetName\": ")
	if tmp, err := json.Marshal(strct.BudgetName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Definition" field is required
    if strct.Definition == nil {
        return nil, errors.New("Definition is a required field")
    }
    // Marshal the "Definition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Definition\": ")
	if tmp, err := json.Marshal(strct.Definition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ExecutionRoleArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ExecutionRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ExecutionRoleArn\": ")
	if tmp, err := json.Marshal(strct.ExecutionRoleArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "NotificationType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "NotificationType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NotificationType\": ")
	if tmp, err := json.Marshal(strct.NotificationType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Subscribers" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Subscribers" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subscribers\": ")
	if tmp, err := json.Marshal(strct.Subscribers); err != nil {
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
    ActionThresholdReceived := false
    ActionTypeReceived := false
    BudgetNameReceived := false
    DefinitionReceived := false
    ExecutionRoleArnReceived := false
    NotificationTypeReceived := false
    SubscribersReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ActionId":
            if err := json.Unmarshal([]byte(v), &strct.ActionId); err != nil {
                return err
             }
        case "ActionThreshold":
            if err := json.Unmarshal([]byte(v), &strct.ActionThreshold); err != nil {
                return err
             }
            ActionThresholdReceived = true
        case "ActionType":
            if err := json.Unmarshal([]byte(v), &strct.ActionType); err != nil {
                return err
             }
            ActionTypeReceived = true
        case "ApprovalModel":
            if err := json.Unmarshal([]byte(v), &strct.ApprovalModel); err != nil {
                return err
             }
        case "BudgetName":
            if err := json.Unmarshal([]byte(v), &strct.BudgetName); err != nil {
                return err
             }
            BudgetNameReceived = true
        case "Definition":
            if err := json.Unmarshal([]byte(v), &strct.Definition); err != nil {
                return err
             }
            DefinitionReceived = true
        case "ExecutionRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.ExecutionRoleArn); err != nil {
                return err
             }
            ExecutionRoleArnReceived = true
        case "NotificationType":
            if err := json.Unmarshal([]byte(v), &strct.NotificationType); err != nil {
                return err
             }
            NotificationTypeReceived = true
        case "Subscribers":
            if err := json.Unmarshal([]byte(v), &strct.Subscribers); err != nil {
                return err
             }
            SubscribersReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ActionThreshold (a required property) was received
    if !ActionThresholdReceived {
        return errors.New("\"ActionThreshold\" is required but was not present")
    }
    // check if ActionType (a required property) was received
    if !ActionTypeReceived {
        return errors.New("\"ActionType\" is required but was not present")
    }
    // check if BudgetName (a required property) was received
    if !BudgetNameReceived {
        return errors.New("\"BudgetName\" is required but was not present")
    }
    // check if Definition (a required property) was received
    if !DefinitionReceived {
        return errors.New("\"Definition\" is required but was not present")
    }
    // check if ExecutionRoleArn (a required property) was received
    if !ExecutionRoleArnReceived {
        return errors.New("\"ExecutionRoleArn\" is required but was not present")
    }
    // check if NotificationType (a required property) was received
    if !NotificationTypeReceived {
        return errors.New("\"NotificationType\" is required but was not present")
    }
    // check if Subscribers (a required property) was received
    if !SubscribersReceived {
        return errors.New("\"Subscribers\" is required but was not present")
    }
    return nil
}

func (strct *ScpActionDefinition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "PolicyId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyId\": ")
	if tmp, err := json.Marshal(strct.PolicyId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TargetIds" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetIds\": ")
	if tmp, err := json.Marshal(strct.TargetIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ScpActionDefinition) UnmarshalJSON(b []byte) error {
    PolicyIdReceived := false
    TargetIdsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PolicyId":
            if err := json.Unmarshal([]byte(v), &strct.PolicyId); err != nil {
                return err
             }
            PolicyIdReceived = true
        case "TargetIds":
            if err := json.Unmarshal([]byte(v), &strct.TargetIds); err != nil {
                return err
             }
            TargetIdsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if PolicyId (a required property) was received
    if !PolicyIdReceived {
        return errors.New("\"PolicyId\" is required but was not present")
    }
    // check if TargetIds (a required property) was received
    if !TargetIdsReceived {
        return errors.New("\"TargetIds\" is required but was not present")
    }
    return nil
}

func (strct *SsmActionDefinition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "InstanceIds" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "InstanceIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceIds\": ")
	if tmp, err := json.Marshal(strct.InstanceIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // "Subtype" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Subtype" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subtype\": ")
	if tmp, err := json.Marshal(strct.Subtype); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SsmActionDefinition) UnmarshalJSON(b []byte) error {
    InstanceIdsReceived := false
    RegionReceived := false
    SubtypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "InstanceIds":
            if err := json.Unmarshal([]byte(v), &strct.InstanceIds); err != nil {
                return err
             }
            InstanceIdsReceived = true
        case "Region":
            if err := json.Unmarshal([]byte(v), &strct.Region); err != nil {
                return err
             }
            RegionReceived = true
        case "Subtype":
            if err := json.Unmarshal([]byte(v), &strct.Subtype); err != nil {
                return err
             }
            SubtypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if InstanceIds (a required property) was received
    if !InstanceIdsReceived {
        return errors.New("\"InstanceIds\" is required but was not present")
    }
    // check if Region (a required property) was received
    if !RegionReceived {
        return errors.New("\"Region\" is required but was not present")
    }
    // check if Subtype (a required property) was received
    if !SubtypeReceived {
        return errors.New("\"Subtype\" is required but was not present")
    }
    return nil
}

func (strct *Subscriber) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Address" field is required
    // only required object types supported for marshal checking (for now)
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

func (strct *Subscriber) UnmarshalJSON(b []byte) error {
    AddressReceived := false
    TypeReceived := false
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
            AddressReceived = true
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Address (a required property) was received
    if !AddressReceived {
        return errors.New("\"Address\" is required but was not present")
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}