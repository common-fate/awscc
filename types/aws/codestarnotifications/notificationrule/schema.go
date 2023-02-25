// Code generated by schema-generate. DO NOT EDIT.

package notificationrule

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::CodeStarNotifications::NotificationRule
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  CreatedBy string `json:"CreatedBy,omitempty"`
  DetailType string `json:"DetailType"`
  EventTypeId string `json:"EventTypeId,omitempty"`
  EventTypeIds []string `json:"EventTypeIds"`
  Name string `json:"Name"`
  Resource string `json:"Resource"`
  Status string `json:"Status,omitempty"`
  Tags *Tags `json:"Tags,omitempty"`
  TargetAddress string `json:"TargetAddress,omitempty"`
  Targets []*Target `json:"Targets"`
}

// Tags 
type Tags struct {
}

// Target 
type Target struct {
  TargetAddress string `json:"TargetAddress"`
  TargetType string `json:"TargetType"`
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
    // Marshal the "CreatedBy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreatedBy\": ")
	if tmp, err := json.Marshal(strct.CreatedBy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DetailType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DetailType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DetailType\": ")
	if tmp, err := json.Marshal(strct.DetailType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EventTypeId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventTypeId\": ")
	if tmp, err := json.Marshal(strct.EventTypeId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "EventTypeIds" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EventTypeIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EventTypeIds\": ")
	if tmp, err := json.Marshal(strct.EventTypeIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
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
    // "Resource" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Resource" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Resource\": ")
	if tmp, err := json.Marshal(strct.Resource); err != nil {
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
    // Marshal the "TargetAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetAddress\": ")
	if tmp, err := json.Marshal(strct.TargetAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Targets" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Targets" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Targets\": ")
	if tmp, err := json.Marshal(strct.Targets); err != nil {
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
    DetailTypeReceived := false
    EventTypeIdsReceived := false
    NameReceived := false
    ResourceReceived := false
    TargetsReceived := false
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
        case "CreatedBy":
            if err := json.Unmarshal([]byte(v), &strct.CreatedBy); err != nil {
                return err
             }
        case "DetailType":
            if err := json.Unmarshal([]byte(v), &strct.DetailType); err != nil {
                return err
             }
            DetailTypeReceived = true
        case "EventTypeId":
            if err := json.Unmarshal([]byte(v), &strct.EventTypeId); err != nil {
                return err
             }
        case "EventTypeIds":
            if err := json.Unmarshal([]byte(v), &strct.EventTypeIds); err != nil {
                return err
             }
            EventTypeIdsReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Resource":
            if err := json.Unmarshal([]byte(v), &strct.Resource); err != nil {
                return err
             }
            ResourceReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TargetAddress":
            if err := json.Unmarshal([]byte(v), &strct.TargetAddress); err != nil {
                return err
             }
        case "Targets":
            if err := json.Unmarshal([]byte(v), &strct.Targets); err != nil {
                return err
             }
            TargetsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DetailType (a required property) was received
    if !DetailTypeReceived {
        return errors.New("\"DetailType\" is required but was not present")
    }
    // check if EventTypeIds (a required property) was received
    if !EventTypeIdsReceived {
        return errors.New("\"EventTypeIds\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Resource (a required property) was received
    if !ResourceReceived {
        return errors.New("\"Resource\" is required but was not present")
    }
    // check if Targets (a required property) was received
    if !TargetsReceived {
        return errors.New("\"Targets\" is required but was not present")
    }
    return nil
}

func (strct *Target) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "TargetAddress" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetAddress\": ")
	if tmp, err := json.Marshal(strct.TargetAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TargetType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetType\": ")
	if tmp, err := json.Marshal(strct.TargetType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Target) UnmarshalJSON(b []byte) error {
    TargetAddressReceived := false
    TargetTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "TargetAddress":
            if err := json.Unmarshal([]byte(v), &strct.TargetAddress); err != nil {
                return err
             }
            TargetAddressReceived = true
        case "TargetType":
            if err := json.Unmarshal([]byte(v), &strct.TargetType); err != nil {
                return err
             }
            TargetTypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if TargetAddress (a required property) was received
    if !TargetAddressReceived {
        return errors.New("\"TargetAddress\" is required but was not present")
    }
    // check if TargetType (a required property) was received
    if !TargetTypeReceived {
        return errors.New("\"TargetType\" is required but was not present")
    }
    return nil
}