// Code generated by schema-generate. DO NOT EDIT.

package framework

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ControlInputParameter 
type ControlInputParameter struct {
  ParameterName string `json:"ParameterName"`
  ParameterValue string `json:"ParameterValue"`
}

// ControlScope The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.
type ControlScope struct {

  // The ID of the only AWS resource that you want your control scope to contain.
  ComplianceResourceIds []string `json:"ComplianceResourceIds,omitempty"`

  // Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS`.
  ComplianceResourceTypes []string `json:"ComplianceResourceTypes,omitempty"`

  // Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.
  Tags []*Tag `json:"Tags,omitempty"`
}

// FrameworkControl 
type FrameworkControl struct {

  // A list of ParameterName and ParameterValue pairs.
  ControlInputParameters []*ControlInputParameter `json:"ControlInputParameters,omitempty"`

  // The name of a control. This name is between 1 and 256 characters.
  ControlName string `json:"ControlName"`

  // The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.
  ControlScope *ControlScope `json:"ControlScope,omitempty"`
}

// Resource Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.
type Resource struct {

  // The date and time that a framework is created, in ISO 8601 representation. The value of CreationTime is accurate to milliseconds. For example, 2020-07-10T15:00:00.000-08:00 represents the 10th of July 2020 at 3:00 PM 8 hours behind UTC.
  CreationTime string `json:"CreationTime,omitempty"`

  // The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`
  DeploymentStatus string `json:"DeploymentStatus,omitempty"`

  // An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
  FrameworkArn string `json:"FrameworkArn,omitempty"`

  // Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
  FrameworkControls []*FrameworkControl `json:"FrameworkControls"`

  // An optional description of the framework with a maximum 1,024 characters.
  FrameworkDescription string `json:"FrameworkDescription,omitempty"`

  // The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
  FrameworkName string `json:"FrameworkName,omitempty"`

  // A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:
  // 
  // `ACTIVE` when recording is turned on for all resources governed by the framework.
  // 
  // `PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.
  // 
  // `INACTIVE` when recording is turned off for all resources governed by the framework.
  // 
  // `UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.
  FrameworkStatus string `json:"FrameworkStatus,omitempty"`

  // Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
  FrameworkTags []*Tag `json:"FrameworkTags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key,omitempty"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value,omitempty"`
}

func (strct *ControlInputParameter) MarshalJSON() ([]byte, error) {
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

func (strct *ControlInputParameter) UnmarshalJSON(b []byte) error {
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

func (strct *ControlScope) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ComplianceResourceIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComplianceResourceIds\": ")
	if tmp, err := json.Marshal(strct.ComplianceResourceIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ComplianceResourceTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComplianceResourceTypes\": ")
	if tmp, err := json.Marshal(strct.ComplianceResourceTypes); err != nil {
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

func (strct *ControlScope) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ComplianceResourceIds":
            if err := json.Unmarshal([]byte(v), &strct.ComplianceResourceIds); err != nil {
                return err
             }
        case "ComplianceResourceTypes":
            if err := json.Unmarshal([]byte(v), &strct.ComplianceResourceTypes); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *FrameworkControl) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ControlInputParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ControlInputParameters\": ")
	if tmp, err := json.Marshal(strct.ControlInputParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ControlName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ControlName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ControlName\": ")
	if tmp, err := json.Marshal(strct.ControlName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ControlScope" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ControlScope\": ")
	if tmp, err := json.Marshal(strct.ControlScope); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *FrameworkControl) UnmarshalJSON(b []byte) error {
    ControlNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ControlInputParameters":
            if err := json.Unmarshal([]byte(v), &strct.ControlInputParameters); err != nil {
                return err
             }
        case "ControlName":
            if err := json.Unmarshal([]byte(v), &strct.ControlName); err != nil {
                return err
             }
            ControlNameReceived = true
        case "ControlScope":
            if err := json.Unmarshal([]byte(v), &strct.ControlScope); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ControlName (a required property) was received
    if !ControlNameReceived {
        return errors.New("\"ControlName\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "DeploymentStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeploymentStatus\": ")
	if tmp, err := json.Marshal(strct.DeploymentStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FrameworkArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkArn\": ")
	if tmp, err := json.Marshal(strct.FrameworkArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FrameworkControls" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FrameworkControls" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkControls\": ")
	if tmp, err := json.Marshal(strct.FrameworkControls); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FrameworkDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkDescription\": ")
	if tmp, err := json.Marshal(strct.FrameworkDescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FrameworkName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkName\": ")
	if tmp, err := json.Marshal(strct.FrameworkName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FrameworkStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkStatus\": ")
	if tmp, err := json.Marshal(strct.FrameworkStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FrameworkTags" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FrameworkTags\": ")
	if tmp, err := json.Marshal(strct.FrameworkTags); err != nil {
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
    FrameworkControlsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CreationTime":
            if err := json.Unmarshal([]byte(v), &strct.CreationTime); err != nil {
                return err
             }
        case "DeploymentStatus":
            if err := json.Unmarshal([]byte(v), &strct.DeploymentStatus); err != nil {
                return err
             }
        case "FrameworkArn":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkArn); err != nil {
                return err
             }
        case "FrameworkControls":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkControls); err != nil {
                return err
             }
            FrameworkControlsReceived = true
        case "FrameworkDescription":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkDescription); err != nil {
                return err
             }
        case "FrameworkName":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkName); err != nil {
                return err
             }
        case "FrameworkStatus":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkStatus); err != nil {
                return err
             }
        case "FrameworkTags":
            if err := json.Unmarshal([]byte(v), &strct.FrameworkTags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if FrameworkControls (a required property) was received
    if !FrameworkControlsReceived {
        return errors.New("\"FrameworkControls\" is required but was not present")
    }
    return nil
}

func (strct *Tag) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
