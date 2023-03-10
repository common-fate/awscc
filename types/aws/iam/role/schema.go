// Code generated by schema-generate. DO NOT EDIT.

package role

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AssumeRolePolicyDocument_object The trust policy that is associated with this role.
type AssumeRolePolicyDocument_object struct {
}

// Policy The inline policy document that is embedded in the specified IAM role.
type Policy struct {

  // The policy document.
  PolicyDocument interface{} `json:"PolicyDocument"`

  // The friendly name (not ARN) identifying the policy.
  PolicyName string `json:"PolicyName"`
}

// PolicyDocument_object The policy document.
type PolicyDocument_object struct {
}

// Resource Resource Type definition for AWS::IAM::Role
type Resource struct {

  // The Amazon Resource Name (ARN) for the role.
  Arn string `json:"Arn,omitempty"`

  // The trust policy that is associated with this role.
  AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument"`

  // A description of the role that you provide.
  Description string `json:"Description,omitempty"`

  // A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. 
  ManagedPolicyArns []string `json:"ManagedPolicyArns,omitempty"`

  // The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. 
  MaxSessionDuration int `json:"MaxSessionDuration,omitempty"`

  // The path to the role.
  Path string `json:"Path,omitempty"`

  // The ARN of the policy used to set the permissions boundary for the role.
  PermissionsBoundary string `json:"PermissionsBoundary,omitempty"`

  // Adds or updates an inline policy document that is embedded in the specified IAM role. 
  Policies []*Policy `json:"Policies,omitempty"`

  // The stable and unique string identifying the role.
  RoleId string `json:"RoleId,omitempty"`

  // A name for the IAM role, up to 64 characters in length.
  RoleName string `json:"RoleName,omitempty"`

  // A list of tags that are attached to the role.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *Policy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "PolicyDocument" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyDocument" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyDocument\": ")
	if tmp, err := json.Marshal(strct.PolicyDocument); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PolicyName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PolicyName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PolicyName\": ")
	if tmp, err := json.Marshal(strct.PolicyName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Policy) UnmarshalJSON(b []byte) error {
    PolicyDocumentReceived := false
    PolicyNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "PolicyDocument":
            if err := json.Unmarshal([]byte(v), &strct.PolicyDocument); err != nil {
                return err
             }
            PolicyDocumentReceived = true
        case "PolicyName":
            if err := json.Unmarshal([]byte(v), &strct.PolicyName); err != nil {
                return err
             }
            PolicyNameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if PolicyDocument (a required property) was received
    if !PolicyDocumentReceived {
        return errors.New("\"PolicyDocument\" is required but was not present")
    }
    // check if PolicyName (a required property) was received
    if !PolicyNameReceived {
        return errors.New("\"PolicyName\" is required but was not present")
    }
    return nil
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
    // "AssumeRolePolicyDocument" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AssumeRolePolicyDocument" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssumeRolePolicyDocument\": ")
	if tmp, err := json.Marshal(strct.AssumeRolePolicyDocument); err != nil {
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
    // Marshal the "ManagedPolicyArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ManagedPolicyArns\": ")
	if tmp, err := json.Marshal(strct.ManagedPolicyArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaxSessionDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxSessionDuration\": ")
	if tmp, err := json.Marshal(strct.MaxSessionDuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Path" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Path\": ")
	if tmp, err := json.Marshal(strct.Path); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PermissionsBoundary" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PermissionsBoundary\": ")
	if tmp, err := json.Marshal(strct.PermissionsBoundary); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Policies" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Policies\": ")
	if tmp, err := json.Marshal(strct.Policies); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RoleId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleId\": ")
	if tmp, err := json.Marshal(strct.RoleId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RoleName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleName\": ")
	if tmp, err := json.Marshal(strct.RoleName); err != nil {
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
    AssumeRolePolicyDocumentReceived := false
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
        case "AssumeRolePolicyDocument":
            if err := json.Unmarshal([]byte(v), &strct.AssumeRolePolicyDocument); err != nil {
                return err
             }
            AssumeRolePolicyDocumentReceived = true
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "ManagedPolicyArns":
            if err := json.Unmarshal([]byte(v), &strct.ManagedPolicyArns); err != nil {
                return err
             }
        case "MaxSessionDuration":
            if err := json.Unmarshal([]byte(v), &strct.MaxSessionDuration); err != nil {
                return err
             }
        case "Path":
            if err := json.Unmarshal([]byte(v), &strct.Path); err != nil {
                return err
             }
        case "PermissionsBoundary":
            if err := json.Unmarshal([]byte(v), &strct.PermissionsBoundary); err != nil {
                return err
             }
        case "Policies":
            if err := json.Unmarshal([]byte(v), &strct.Policies); err != nil {
                return err
             }
        case "RoleId":
            if err := json.Unmarshal([]byte(v), &strct.RoleId); err != nil {
                return err
             }
        case "RoleName":
            if err := json.Unmarshal([]byte(v), &strct.RoleName); err != nil {
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
    // check if AssumeRolePolicyDocument (a required property) was received
    if !AssumeRolePolicyDocumentReceived {
        return errors.New("\"AssumeRolePolicyDocument\" is required but was not present")
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
