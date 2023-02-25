// Code generated by schema-generate. DO NOT EDIT.

package organizationconformancepack

import (
    "errors"
    "fmt"
    "bytes"
    "encoding/json"
)

// ConformancePackInputParameter Input parameters in the form of key-value pairs for the conformance pack.
type ConformancePackInputParameter struct {
  ParameterName string `json:"ParameterName"`
  ParameterValue string `json:"ParameterValue"`
}

// Resource Resource schema for AWS::Config::OrganizationConformancePack.
type Resource struct {

  // A list of ConformancePackInputParameter objects.
  ConformancePackInputParameters []*ConformancePackInputParameter `json:"ConformancePackInputParameters,omitempty"`

  // AWS Config stores intermediate files while processing conformance pack template.
  DeliveryS3Bucket string `json:"DeliveryS3Bucket,omitempty"`

  // The prefix for the delivery S3 bucket.
  DeliveryS3KeyPrefix string `json:"DeliveryS3KeyPrefix,omitempty"`

  // A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.
  ExcludedAccounts []string `json:"ExcludedAccounts,omitempty"`

  // The name of the organization conformance pack.
  OrganizationConformancePackName string `json:"OrganizationConformancePackName"`

  // A string containing full conformance pack template body.
  TemplateBody string `json:"TemplateBody,omitempty"`

  // Location of file containing the template body.
  TemplateS3Uri string `json:"TemplateS3Uri,omitempty"`
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
    // Marshal the "ExcludedAccounts" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ExcludedAccounts\": ")
	if tmp, err := json.Marshal(strct.ExcludedAccounts); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "OrganizationConformancePackName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "OrganizationConformancePackName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OrganizationConformancePackName\": ")
	if tmp, err := json.Marshal(strct.OrganizationConformancePackName); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    OrganizationConformancePackNameReceived := false
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
        case "DeliveryS3Bucket":
            if err := json.Unmarshal([]byte(v), &strct.DeliveryS3Bucket); err != nil {
                return err
             }
        case "DeliveryS3KeyPrefix":
            if err := json.Unmarshal([]byte(v), &strct.DeliveryS3KeyPrefix); err != nil {
                return err
             }
        case "ExcludedAccounts":
            if err := json.Unmarshal([]byte(v), &strct.ExcludedAccounts); err != nil {
                return err
             }
        case "OrganizationConformancePackName":
            if err := json.Unmarshal([]byte(v), &strct.OrganizationConformancePackName); err != nil {
                return err
             }
            OrganizationConformancePackNameReceived = true
        case "TemplateBody":
            if err := json.Unmarshal([]byte(v), &strct.TemplateBody); err != nil {
                return err
             }
        case "TemplateS3Uri":
            if err := json.Unmarshal([]byte(v), &strct.TemplateS3Uri); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if OrganizationConformancePackName (a required property) was received
    if !OrganizationConformancePackNameReceived {
        return errors.New("\"OrganizationConformancePackName\" is required but was not present")
    }
    return nil
}
