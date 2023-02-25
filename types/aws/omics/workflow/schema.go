// Code generated by schema-generate. DO NOT EDIT.

package workflow

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Definition of AWS::Omics::Workflow Resource Type
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  CreationTime string `json:"CreationTime,omitempty"`
  DefinitionUri string `json:"DefinitionUri,omitempty"`
  Description string `json:"Description,omitempty"`
  Engine string `json:"Engine,omitempty"`
  Id string `json:"Id,omitempty"`
  Main string `json:"Main,omitempty"`
  Name string `json:"Name,omitempty"`
  ParameterTemplate *WorkflowParameterTemplate `json:"ParameterTemplate,omitempty"`
  Status string `json:"Status,omitempty"`
  StorageCapacity float64 `json:"StorageCapacity,omitempty"`
  Tags *TagMap `json:"Tags,omitempty"`
  Type string `json:"Type,omitempty"`
}

// TagMap A map of resource tags
type TagMap struct {
}

// WorkflowParameter 
type WorkflowParameter struct {
  Description string `json:"Description,omitempty"`
  Optional bool `json:"Optional,omitempty"`
}

// WorkflowParameterTemplate 
type WorkflowParameterTemplate struct {
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
    // Marshal the "DefinitionUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefinitionUri\": ")
	if tmp, err := json.Marshal(strct.DefinitionUri); err != nil {
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
    // Marshal the "Engine" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Engine\": ")
	if tmp, err := json.Marshal(strct.Engine); err != nil {
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
    // Marshal the "Main" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Main\": ")
	if tmp, err := json.Marshal(strct.Main); err != nil {
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
    // Marshal the "ParameterTemplate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParameterTemplate\": ")
	if tmp, err := json.Marshal(strct.ParameterTemplate); err != nil {
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
    // Marshal the "StorageCapacity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StorageCapacity\": ")
	if tmp, err := json.Marshal(strct.StorageCapacity); err != nil {
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
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
        case "DefinitionUri":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionUri); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Engine":
            if err := json.Unmarshal([]byte(v), &strct.Engine); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "Main":
            if err := json.Unmarshal([]byte(v), &strct.Main); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "ParameterTemplate":
            if err := json.Unmarshal([]byte(v), &strct.ParameterTemplate); err != nil {
                return err
             }
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "StorageCapacity":
            if err := json.Unmarshal([]byte(v), &strct.StorageCapacity); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *TagMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TagMap) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, _ := range jsonMap {
        switch k {
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *WorkflowParameter) MarshalJSON() ([]byte, error) {
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
    // Marshal the "Optional" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Optional\": ")
	if tmp, err := json.Marshal(strct.Optional); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WorkflowParameter) UnmarshalJSON(b []byte) error {
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
        case "Optional":
            if err := json.Unmarshal([]byte(v), &strct.Optional); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *WorkflowParameterTemplate) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WorkflowParameterTemplate) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, _ := range jsonMap {
        switch k {
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
