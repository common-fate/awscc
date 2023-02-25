// Code generated by schema-generate. DO NOT EDIT.

package app

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// PhysicalResourceId 
type PhysicalResourceId struct {
  AwsAccountId string `json:"AwsAccountId,omitempty"`
  AwsRegion string `json:"AwsRegion,omitempty"`
  Identifier string `json:"Identifier"`
  Type string `json:"Type"`
}

// Resource Resource Type Definition for AWS::ResilienceHub::App.
type Resource struct {

  // Amazon Resource Name (ARN) of the App.
  AppArn string `json:"AppArn,omitempty"`

  // Assessment execution schedule.
  AppAssessmentSchedule string `json:"AppAssessmentSchedule,omitempty"`

  // A string containing full ResilienceHub app template body.
  AppTemplateBody string `json:"AppTemplateBody"`

  // App description.
  Description string `json:"Description,omitempty"`

  // Name of the app.
  Name string `json:"Name"`

  // Amazon Resource Name (ARN) of the Resiliency Policy.
  ResiliencyPolicyArn string `json:"ResiliencyPolicyArn,omitempty"`

  // An array of ResourceMapping objects.
  ResourceMappings []*ResourceMapping `json:"ResourceMappings"`
  Tags *TagMap `json:"Tags,omitempty"`
}

// ResourceMapping Resource mapping is used to map logical resources from template to physical resource
type ResourceMapping struct {
  LogicalStackName string `json:"LogicalStackName,omitempty"`
  MappingType string `json:"MappingType"`
  PhysicalResourceId *PhysicalResourceId `json:"PhysicalResourceId"`
  ResourceName string `json:"ResourceName,omitempty"`
  TerraformSourceName string `json:"TerraformSourceName,omitempty"`
}

// TagMap 
type TagMap struct {
}

func (strct *PhysicalResourceId) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AwsAccountId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AwsAccountId\": ")
	if tmp, err := json.Marshal(strct.AwsAccountId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AwsRegion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AwsRegion\": ")
	if tmp, err := json.Marshal(strct.AwsRegion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Identifier" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Identifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Identifier\": ")
	if tmp, err := json.Marshal(strct.Identifier); err != nil {
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

func (strct *PhysicalResourceId) UnmarshalJSON(b []byte) error {
    IdentifierReceived := false
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AwsAccountId":
            if err := json.Unmarshal([]byte(v), &strct.AwsAccountId); err != nil {
                return err
             }
        case "AwsRegion":
            if err := json.Unmarshal([]byte(v), &strct.AwsRegion); err != nil {
                return err
             }
        case "Identifier":
            if err := json.Unmarshal([]byte(v), &strct.Identifier); err != nil {
                return err
             }
            IdentifierReceived = true
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Identifier (a required property) was received
    if !IdentifierReceived {
        return errors.New("\"Identifier\" is required but was not present")
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AppArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppArn\": ")
	if tmp, err := json.Marshal(strct.AppArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AppAssessmentSchedule" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppAssessmentSchedule\": ")
	if tmp, err := json.Marshal(strct.AppAssessmentSchedule); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "AppTemplateBody" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AppTemplateBody" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppTemplateBody\": ")
	if tmp, err := json.Marshal(strct.AppTemplateBody); err != nil {
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
    // Marshal the "ResiliencyPolicyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResiliencyPolicyArn\": ")
	if tmp, err := json.Marshal(strct.ResiliencyPolicyArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourceMappings" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ResourceMappings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceMappings\": ")
	if tmp, err := json.Marshal(strct.ResourceMappings); err != nil {
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
    AppTemplateBodyReceived := false
    NameReceived := false
    ResourceMappingsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AppArn":
            if err := json.Unmarshal([]byte(v), &strct.AppArn); err != nil {
                return err
             }
        case "AppAssessmentSchedule":
            if err := json.Unmarshal([]byte(v), &strct.AppAssessmentSchedule); err != nil {
                return err
             }
        case "AppTemplateBody":
            if err := json.Unmarshal([]byte(v), &strct.AppTemplateBody); err != nil {
                return err
             }
            AppTemplateBodyReceived = true
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "ResiliencyPolicyArn":
            if err := json.Unmarshal([]byte(v), &strct.ResiliencyPolicyArn); err != nil {
                return err
             }
        case "ResourceMappings":
            if err := json.Unmarshal([]byte(v), &strct.ResourceMappings); err != nil {
                return err
             }
            ResourceMappingsReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AppTemplateBody (a required property) was received
    if !AppTemplateBodyReceived {
        return errors.New("\"AppTemplateBody\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if ResourceMappings (a required property) was received
    if !ResourceMappingsReceived {
        return errors.New("\"ResourceMappings\" is required but was not present")
    }
    return nil
}

func (strct *ResourceMapping) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "LogicalStackName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LogicalStackName\": ")
	if tmp, err := json.Marshal(strct.LogicalStackName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MappingType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "MappingType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MappingType\": ")
	if tmp, err := json.Marshal(strct.MappingType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PhysicalResourceId" field is required
    if strct.PhysicalResourceId == nil {
        return nil, errors.New("PhysicalResourceId is a required field")
    }
    // Marshal the "PhysicalResourceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PhysicalResourceId\": ")
	if tmp, err := json.Marshal(strct.PhysicalResourceId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceName\": ")
	if tmp, err := json.Marshal(strct.ResourceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TerraformSourceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TerraformSourceName\": ")
	if tmp, err := json.Marshal(strct.TerraformSourceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResourceMapping) UnmarshalJSON(b []byte) error {
    MappingTypeReceived := false
    PhysicalResourceIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LogicalStackName":
            if err := json.Unmarshal([]byte(v), &strct.LogicalStackName); err != nil {
                return err
             }
        case "MappingType":
            if err := json.Unmarshal([]byte(v), &strct.MappingType); err != nil {
                return err
             }
            MappingTypeReceived = true
        case "PhysicalResourceId":
            if err := json.Unmarshal([]byte(v), &strct.PhysicalResourceId); err != nil {
                return err
             }
            PhysicalResourceIdReceived = true
        case "ResourceName":
            if err := json.Unmarshal([]byte(v), &strct.ResourceName); err != nil {
                return err
             }
        case "TerraformSourceName":
            if err := json.Unmarshal([]byte(v), &strct.TerraformSourceName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MappingType (a required property) was received
    if !MappingTypeReceived {
        return errors.New("\"MappingType\" is required but was not present")
    }
    // check if PhysicalResourceId (a required property) was received
    if !PhysicalResourceIdReceived {
        return errors.New("\"PhysicalResourceId\" is required but was not present")
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