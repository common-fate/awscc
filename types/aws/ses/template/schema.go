// Code generated by schema-generate. DO NOT EDIT.

package template

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::SES::Template
type Resource struct {
  Id string `json:"Id,omitempty"`
  Template *Template `json:"Template,omitempty"`
}

// Template The content of the email, composed of a subject line, an HTML part, and a text-only part
type Template struct {

  // The HTML body of the email.
  HtmlPart string `json:"HtmlPart,omitempty"`

  // The subject line of the email.
  SubjectPart string `json:"SubjectPart"`

  // The name of the template.
  TemplateName string `json:"TemplateName,omitempty"`

  // The email body that is visible to recipients whose email clients do not display HTML content.
  TextPart string `json:"TextPart,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "Template" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Template\": ")
	if tmp, err := json.Marshal(strct.Template); err != nil {
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
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "Template":
            if err := json.Unmarshal([]byte(v), &strct.Template); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Template) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "HtmlPart" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HtmlPart\": ")
	if tmp, err := json.Marshal(strct.HtmlPart); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SubjectPart" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SubjectPart" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubjectPart\": ")
	if tmp, err := json.Marshal(strct.SubjectPart); err != nil {
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
    // Marshal the "TextPart" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TextPart\": ")
	if tmp, err := json.Marshal(strct.TextPart); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Template) UnmarshalJSON(b []byte) error {
    SubjectPartReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "HtmlPart":
            if err := json.Unmarshal([]byte(v), &strct.HtmlPart); err != nil {
                return err
             }
        case "SubjectPart":
            if err := json.Unmarshal([]byte(v), &strct.SubjectPart); err != nil {
                return err
             }
            SubjectPartReceived = true
        case "TemplateName":
            if err := json.Unmarshal([]byte(v), &strct.TemplateName); err != nil {
                return err
             }
        case "TextPart":
            if err := json.Unmarshal([]byte(v), &strct.TextPart); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if SubjectPart (a required property) was received
    if !SubjectPartReceived {
        return errors.New("\"SubjectPart\" is required but was not present")
    }
    return nil
}
