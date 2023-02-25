// Code generated by schema-generate. DO NOT EDIT.

package provisioningtemplate

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ProvisioningHook 
type ProvisioningHook struct {
  PayloadVersion string `json:"PayloadVersion,omitempty"`
  TargetArn string `json:"TargetArn,omitempty"`
}

// Resource Creates a fleet provisioning template.
type Resource struct {
  Description string `json:"Description,omitempty"`
  Enabled bool `json:"Enabled,omitempty"`
  PreProvisioningHook *ProvisioningHook `json:"PreProvisioningHook,omitempty"`
  ProvisioningRoleArn string `json:"ProvisioningRoleArn"`
  Tags []*Tag `json:"Tags,omitempty"`
  TemplateArn string `json:"TemplateArn,omitempty"`
  TemplateBody string `json:"TemplateBody"`
  TemplateName string `json:"TemplateName,omitempty"`
  TemplateType string `json:"TemplateType,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
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
    // Marshal the "Enabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Enabled\": ")
	if tmp, err := json.Marshal(strct.Enabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PreProvisioningHook" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PreProvisioningHook\": ")
	if tmp, err := json.Marshal(strct.PreProvisioningHook); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ProvisioningRoleArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ProvisioningRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProvisioningRoleArn\": ")
	if tmp, err := json.Marshal(strct.ProvisioningRoleArn); err != nil {
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
    // Marshal the "TemplateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateArn\": ")
	if tmp, err := json.Marshal(strct.TemplateArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TemplateBody" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TemplateBody" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateBody\": ")
	if tmp, err := json.Marshal(strct.TemplateBody); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TemplateName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateName\": ")
	if tmp, err := json.Marshal(strct.TemplateName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TemplateType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateType\": ")
	if tmp, err := json.Marshal(strct.TemplateType); err != nil {
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
    ProvisioningRoleArnReceived := false
    TemplateBodyReceived := false
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
        case "Enabled":
            if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
                return err
             }
        case "PreProvisioningHook":
            if err := json.Unmarshal([]byte(v), &strct.PreProvisioningHook); err != nil {
                return err
             }
        case "ProvisioningRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.ProvisioningRoleArn); err != nil {
                return err
             }
            ProvisioningRoleArnReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TemplateArn":
            if err := json.Unmarshal([]byte(v), &strct.TemplateArn); err != nil {
                return err
             }
        case "TemplateBody":
            if err := json.Unmarshal([]byte(v), &strct.TemplateBody); err != nil {
                return err
             }
            TemplateBodyReceived = true
        case "TemplateName":
            if err := json.Unmarshal([]byte(v), &strct.TemplateName); err != nil {
                return err
             }
        case "TemplateType":
            if err := json.Unmarshal([]byte(v), &strct.TemplateType); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ProvisioningRoleArn (a required property) was received
    if !ProvisioningRoleArnReceived {
        return errors.New("\"ProvisioningRoleArn\" is required but was not present")
    }
    // check if TemplateBody (a required property) was received
    if !TemplateBodyReceived {
        return errors.New("\"TemplateBody\" is required but was not present")
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