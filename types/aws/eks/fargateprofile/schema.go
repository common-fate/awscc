// Code generated by schema-generate. DO NOT EDIT.

package fargateprofile

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Label A key-value pair to associate with a pod.
type Label struct {

  // The key name of the label.
  Key string `json:"Key"`

  // The value for the label. 
  Value string `json:"Value"`
}

// Resource Resource Schema for AWS::EKS::FargateProfile
type Resource struct {
  Arn string `json:"Arn,omitempty"`

  // Name of the Cluster
  ClusterName string `json:"ClusterName"`

  // Name of FargateProfile
  FargateProfileName string `json:"FargateProfileName,omitempty"`

  // The IAM policy arn for pods
  PodExecutionRoleArn string `json:"PodExecutionRoleArn"`
  Selectors []*Selector `json:"Selectors"`
  Subnets []string `json:"Subnets,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Selector 
type Selector struct {
  Labels []*Label `json:"Labels,omitempty"`
  Namespace string `json:"Namespace"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *Label) MarshalJSON() ([]byte, error) {
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

func (strct *Label) UnmarshalJSON(b []byte) error {
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
    // "ClusterName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ClusterName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClusterName\": ")
	if tmp, err := json.Marshal(strct.ClusterName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FargateProfileName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FargateProfileName\": ")
	if tmp, err := json.Marshal(strct.FargateProfileName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PodExecutionRoleArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PodExecutionRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PodExecutionRoleArn\": ")
	if tmp, err := json.Marshal(strct.PodExecutionRoleArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Selectors" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Selectors" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Selectors\": ")
	if tmp, err := json.Marshal(strct.Selectors); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Subnets" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subnets\": ")
	if tmp, err := json.Marshal(strct.Subnets); err != nil {
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
    ClusterNameReceived := false
    PodExecutionRoleArnReceived := false
    SelectorsReceived := false
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
        case "ClusterName":
            if err := json.Unmarshal([]byte(v), &strct.ClusterName); err != nil {
                return err
             }
            ClusterNameReceived = true
        case "FargateProfileName":
            if err := json.Unmarshal([]byte(v), &strct.FargateProfileName); err != nil {
                return err
             }
        case "PodExecutionRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.PodExecutionRoleArn); err != nil {
                return err
             }
            PodExecutionRoleArnReceived = true
        case "Selectors":
            if err := json.Unmarshal([]byte(v), &strct.Selectors); err != nil {
                return err
             }
            SelectorsReceived = true
        case "Subnets":
            if err := json.Unmarshal([]byte(v), &strct.Subnets); err != nil {
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
    // check if ClusterName (a required property) was received
    if !ClusterNameReceived {
        return errors.New("\"ClusterName\" is required but was not present")
    }
    // check if PodExecutionRoleArn (a required property) was received
    if !PodExecutionRoleArnReceived {
        return errors.New("\"PodExecutionRoleArn\" is required but was not present")
    }
    // check if Selectors (a required property) was received
    if !SelectorsReceived {
        return errors.New("\"Selectors\" is required but was not present")
    }
    return nil
}

func (strct *Selector) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Labels" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Namespace" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Namespace" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Namespace\": ")
	if tmp, err := json.Marshal(strct.Namespace); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Selector) UnmarshalJSON(b []byte) error {
    NamespaceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Labels":
            if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
                return err
             }
        case "Namespace":
            if err := json.Unmarshal([]byte(v), &strct.Namespace); err != nil {
                return err
             }
            NamespaceReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Namespace (a required property) was received
    if !NamespaceReceived {
        return errors.New("\"Namespace\" is required but was not present")
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