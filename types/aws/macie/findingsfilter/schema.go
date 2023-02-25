// Code generated by schema-generate. DO NOT EDIT.

package findingsfilter

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// Criterion Map of filter criteria.
type Criterion struct {
}

// CriterionAdditionalProperties 
type CriterionAdditionalProperties struct {
  Eq []string `json:"eq,omitempty"`
  Gt int `json:"gt,omitempty"`
  Gte int `json:"gte,omitempty"`
  Lt int `json:"lt,omitempty"`
  Lte int `json:"lte,omitempty"`
  Neq []string `json:"neq,omitempty"`
}

// FindingCriteria 
type FindingCriteria struct {
  Criterion *Criterion `json:"Criterion,omitempty"`
}

// FindingsFilterListItem Returned by ListHandler representing filter name and ID.
type FindingsFilterListItem struct {
  Id string `json:"Id,omitempty"`
  Name string `json:"Name,omitempty"`
}

// Resource Macie FindingsFilter resource schema.
type Resource struct {

  // Findings filter action.
  Action string `json:"Action,omitempty"`

  // Findings filter ARN.
  Arn string `json:"Arn,omitempty"`

  // Findings filter description
  Description string `json:"Description,omitempty"`

  // Findings filter criteria.
  FindingCriteria *FindingCriteria `json:"FindingCriteria"`

  // Findings filters list.
  FindingsFilterListItems []*FindingsFilterListItem `json:"FindingsFilterListItems,omitempty"`

  // Findings filter ID.
  Id string `json:"Id,omitempty"`

  // Findings filter name
  Name string `json:"Name"`

  // Findings filter position.
  Position int `json:"Position,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
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
    // "FindingCriteria" field is required
    if strct.FindingCriteria == nil {
        return nil, errors.New("FindingCriteria is a required field")
    }
    // Marshal the "FindingCriteria" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FindingCriteria\": ")
	if tmp, err := json.Marshal(strct.FindingCriteria); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FindingsFilterListItems" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FindingsFilterListItems\": ")
	if tmp, err := json.Marshal(strct.FindingsFilterListItems); err != nil {
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
    // Marshal the "Position" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Position\": ")
	if tmp, err := json.Marshal(strct.Position); err != nil {
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
    FindingCriteriaReceived := false
    NameReceived := false
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
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "FindingCriteria":
            if err := json.Unmarshal([]byte(v), &strct.FindingCriteria); err != nil {
                return err
             }
            FindingCriteriaReceived = true
        case "FindingsFilterListItems":
            if err := json.Unmarshal([]byte(v), &strct.FindingsFilterListItems); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Position":
            if err := json.Unmarshal([]byte(v), &strct.Position); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FindingCriteria (a required property) was received
    if !FindingCriteriaReceived {
        return errors.New("\"FindingCriteria\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}
