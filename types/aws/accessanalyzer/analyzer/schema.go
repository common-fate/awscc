// Code generated by schema-generate. DO NOT EDIT.

package analyzer

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// ArchiveRule An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.
type ArchiveRule struct {
  Filter []*Filter `json:"Filter"`

  // The archive rule name
  RuleName string `json:"RuleName"`
}

// Filter 
type Filter struct {
  Contains []string `json:"Contains,omitempty"`
  Eq []string `json:"Eq,omitempty"`
  Exists bool `json:"Exists,omitempty"`
  Neq []string `json:"Neq,omitempty"`
  Property string `json:"Property"`
}

// Resource The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
type Resource struct {

  // Analyzer name
  AnalyzerName string `json:"AnalyzerName,omitempty"`
  ArchiveRules []*ArchiveRule `json:"ArchiveRules,omitempty"`

  // Amazon Resource Name (ARN) of the analyzer
  Arn string `json:"Arn,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // The type of the analyzer, must be ACCOUNT or ORGANIZATION
  Type string `json:"Type"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *ArchiveRule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Filter" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Filter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Filter\": ")
	if tmp, err := json.Marshal(strct.Filter); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RuleName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RuleName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RuleName\": ")
	if tmp, err := json.Marshal(strct.RuleName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ArchiveRule) UnmarshalJSON(b []byte) error {
    FilterReceived := false
    RuleNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Filter":
            if err := json.Unmarshal([]byte(v), &strct.Filter); err != nil {
                return err
             }
            FilterReceived = true
        case "RuleName":
            if err := json.Unmarshal([]byte(v), &strct.RuleName); err != nil {
                return err
             }
            RuleNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Filter (a required property) was received
    if !FilterReceived {
        return errors.New("\"Filter\" is required but was not present")
    }
    // check if RuleName (a required property) was received
    if !RuleNameReceived {
        return errors.New("\"RuleName\" is required but was not present")
    }
    return nil
}

func (strct *Filter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Contains" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Contains\": ")
	if tmp, err := json.Marshal(strct.Contains); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Eq" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Eq\": ")
	if tmp, err := json.Marshal(strct.Eq); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Exists" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Exists\": ")
	if tmp, err := json.Marshal(strct.Exists); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Neq" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Neq\": ")
	if tmp, err := json.Marshal(strct.Neq); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Property" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Property" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Property\": ")
	if tmp, err := json.Marshal(strct.Property); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Filter) UnmarshalJSON(b []byte) error {
    PropertyReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Contains":
            if err := json.Unmarshal([]byte(v), &strct.Contains); err != nil {
                return err
             }
        case "Eq":
            if err := json.Unmarshal([]byte(v), &strct.Eq); err != nil {
                return err
             }
        case "Exists":
            if err := json.Unmarshal([]byte(v), &strct.Exists); err != nil {
                return err
             }
        case "Neq":
            if err := json.Unmarshal([]byte(v), &strct.Neq); err != nil {
                return err
             }
        case "Property":
            if err := json.Unmarshal([]byte(v), &strct.Property); err != nil {
                return err
             }
            PropertyReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Property (a required property) was received
    if !PropertyReceived {
        return errors.New("\"Property\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AnalyzerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AnalyzerName\": ")
	if tmp, err := json.Marshal(strct.AnalyzerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ArchiveRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ArchiveRules\": ")
	if tmp, err := json.Marshal(strct.ArchiveRules); err != nil {
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AnalyzerName":
            if err := json.Unmarshal([]byte(v), &strct.AnalyzerName); err != nil {
                return err
             }
        case "ArchiveRules":
            if err := json.Unmarshal([]byte(v), &strct.ArchiveRules); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
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
