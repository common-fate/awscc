// Code generated by schema-generate. DO NOT EDIT.

package addon

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Schema for AWS::EKS::Addon
type Resource struct {

  // Name of Addon
  AddonName string `json:"AddonName"`

  // Version of Addon
  AddonVersion string `json:"AddonVersion,omitempty"`

  // Amazon Resource Name (ARN) of the add-on
  Arn string `json:"Arn,omitempty"`

  // Name of Cluster
  ClusterName string `json:"ClusterName"`

  // The configuration values to use with the add-on
  ConfigurationValues string `json:"ConfigurationValues,omitempty"`

  // PreserveOnDelete parameter value
  PreserveOnDelete bool `json:"PreserveOnDelete,omitempty"`

  // Resolve parameter value conflicts
  ResolveConflicts string `json:"ResolveConflicts,omitempty"`

  // IAM role to bind to the add-on's service account
  ServiceAccountRoleArn string `json:"ServiceAccountRoleArn,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AddonName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AddonName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AddonName\": ")
	if tmp, err := json.Marshal(strct.AddonName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AddonVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AddonVersion\": ")
	if tmp, err := json.Marshal(strct.AddonVersion); err != nil {
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
    // Marshal the "ConfigurationValues" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ConfigurationValues\": ")
	if tmp, err := json.Marshal(strct.ConfigurationValues); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PreserveOnDelete" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PreserveOnDelete\": ")
	if tmp, err := json.Marshal(strct.PreserveOnDelete); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResolveConflicts" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResolveConflicts\": ")
	if tmp, err := json.Marshal(strct.ResolveConflicts); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServiceAccountRoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceAccountRoleArn\": ")
	if tmp, err := json.Marshal(strct.ServiceAccountRoleArn); err != nil {
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
    AddonNameReceived := false
    ClusterNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AddonName":
            if err := json.Unmarshal([]byte(v), &strct.AddonName); err != nil {
                return err
             }
            AddonNameReceived = true
        case "AddonVersion":
            if err := json.Unmarshal([]byte(v), &strct.AddonVersion); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "ClusterName":
            if err := json.Unmarshal([]byte(v), &strct.ClusterName); err != nil {
                return err
             }
            ClusterNameReceived = true
        case "ConfigurationValues":
            if err := json.Unmarshal([]byte(v), &strct.ConfigurationValues); err != nil {
                return err
             }
        case "PreserveOnDelete":
            if err := json.Unmarshal([]byte(v), &strct.PreserveOnDelete); err != nil {
                return err
             }
        case "ResolveConflicts":
            if err := json.Unmarshal([]byte(v), &strct.ResolveConflicts); err != nil {
                return err
             }
        case "ServiceAccountRoleArn":
            if err := json.Unmarshal([]byte(v), &strct.ServiceAccountRoleArn); err != nil {
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
    // check if AddonName (a required property) was received
    if !AddonNameReceived {
        return errors.New("\"AddonName\" is required but was not present")
    }
    // check if ClusterName (a required property) was received
    if !ClusterNameReceived {
        return errors.New("\"ClusterName\" is required but was not present")
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