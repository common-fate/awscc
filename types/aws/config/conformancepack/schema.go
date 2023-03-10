// Code generated by schema-generate. DO NOT EDIT.

package conformancepack

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ConformancePackInputParameter Input parameters in the form of key-value pairs for the conformance pack.
type ConformancePackInputParameter struct {
  ParameterName string `json:"ParameterName"`
  ParameterValue string `json:"ParameterValue"`
}

// Resource A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.
type Resource struct {

  // A list of ConformancePackInputParameter objects.
  ConformancePackInputParameters []*ConformancePackInputParameter `json:"ConformancePackInputParameters,omitempty"`

  // Name of the conformance pack which will be assigned as the unique identifier.
  ConformancePackName string `json:"ConformancePackName"`

  // AWS Config stores intermediate files while processing conformance pack template.
  DeliveryS3Bucket string `json:"DeliveryS3Bucket,omitempty"`

  // The prefix for delivery S3 bucket.
  DeliveryS3KeyPrefix string `json:"DeliveryS3KeyPrefix,omitempty"`

  // A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
  TemplateBody string `json:"TemplateBody,omitempty"`

  // Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
  TemplateS3Uri string `json:"TemplateS3Uri,omitempty"`

  // The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
  TemplateSSMDocumentDetails *TemplateSSMDocumentDetails `json:"TemplateSSMDocumentDetails,omitempty"`
}

// TemplateSSMDocumentDetails The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
type TemplateSSMDocumentDetails struct {
  DocumentName string `json:"DocumentName,omitempty"`
  DocumentVersion string `json:"DocumentVersion,omitempty"`
}

func (strct *ConformancePackInputParameter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ParameterName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ParameterName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParameterName\": ")
	if tmp, err := json.Marshal(strct.ParameterName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ParameterValue" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ParameterValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParameterValue\": ")
	if tmp, err := json.Marshal(strct.ParameterValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ConformancePackInputParameter) UnmarshalJSON(b []byte) error {
    ParameterNameReceived := false
    ParameterValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ParameterName":
            if err := json.Unmarshal([]byte(v), &strct.ParameterName); err != nil {
                return err
             }
            ParameterNameReceived = true
        case "ParameterValue":
            if err := json.Unmarshal([]byte(v), &strct.ParameterValue); err != nil {
                return err
             }
            ParameterValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ParameterName (a required property) was received
    if !ParameterNameReceived {
        return errors.New("\"ParameterName\" is required but was not present")
    }
    // check if ParameterValue (a required property) was received
    if !ParameterValueReceived {
        return errors.New("\"ParameterValue\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ConformancePackInputParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConformancePackInputParameters\": ")
	if tmp, err := json.Marshal(strct.ConformancePackInputParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ConformancePackName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ConformancePackName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConformancePackName\": ")
	if tmp, err := json.Marshal(strct.ConformancePackName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DeliveryS3Bucket" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeliveryS3Bucket\": ")
	if tmp, err := json.Marshal(strct.DeliveryS3Bucket); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DeliveryS3KeyPrefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeliveryS3KeyPrefix\": ")
	if tmp, err := json.Marshal(strct.DeliveryS3KeyPrefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "TemplateS3Uri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateS3Uri\": ")
	if tmp, err := json.Marshal(strct.TemplateS3Uri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TemplateSSMDocumentDetails" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TemplateSSMDocumentDetails\": ")
	if tmp, err := json.Marshal(strct.TemplateSSMDocumentDetails); err != nil {
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
    ConformancePackNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ConformancePackInputParameters":
            if err := json.Unmarshal([]byte(v), &strct.ConformancePackInputParameters); err != nil {
                return err
             }
        case "ConformancePackName":
            if err := json.Unmarshal([]byte(v), &strct.ConformancePackName); err != nil {
                return err
             }
            ConformancePackNameReceived = true
        case "DeliveryS3Bucket":
            if err := json.Unmarshal([]byte(v), &strct.DeliveryS3Bucket); err != nil {
                return err
             }
        case "DeliveryS3KeyPrefix":
            if err := json.Unmarshal([]byte(v), &strct.DeliveryS3KeyPrefix); err != nil {
                return err
             }
        case "TemplateBody":
            if err := json.Unmarshal([]byte(v), &strct.TemplateBody); err != nil {
                return err
             }
        case "TemplateS3Uri":
            if err := json.Unmarshal([]byte(v), &strct.TemplateS3Uri); err != nil {
                return err
             }
        case "TemplateSSMDocumentDetails":
            if err := json.Unmarshal([]byte(v), &strct.TemplateSSMDocumentDetails); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ConformancePackName (a required property) was received
    if !ConformancePackNameReceived {
        return errors.New("\"ConformancePackName\" is required but was not present")
    }
    return nil
}

func (strct *TemplateSSMDocumentDetails) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DocumentName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DocumentName\": ")
	if tmp, err := json.Marshal(strct.DocumentName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DocumentVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DocumentVersion\": ")
	if tmp, err := json.Marshal(strct.DocumentVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TemplateSSMDocumentDetails) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DocumentName":
            if err := json.Unmarshal([]byte(v), &strct.DocumentName); err != nil {
                return err
             }
        case "DocumentVersion":
            if err := json.Unmarshal([]byte(v), &strct.DocumentVersion); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
