// Code generated by schema-generate. DO NOT EDIT.

package workspace

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// LoggingConfiguration Logging configuration
type LoggingConfiguration struct {

  // CloudWatch log group ARN
  LogGroupArn string `json:"LogGroupArn,omitempty"`
}

// Resource Resource Type definition for AWS::APS::Workspace
type Resource struct {

  // The AMP Workspace alert manager definition data
  AlertManagerDefinition string `json:"AlertManagerDefinition,omitempty"`

  // AMP Workspace alias.
  Alias string `json:"Alias,omitempty"`

  // Workspace arn.
  Arn string `json:"Arn,omitempty"`
  LoggingConfiguration *LoggingConfiguration `json:"LoggingConfiguration,omitempty"`

  // AMP Workspace prometheus endpoint
  PrometheusEndpoint string `json:"PrometheusEndpoint,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // Required to identify a specific APS Workspace.
  WorkspaceId string `json:"WorkspaceId,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *LoggingConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "LogGroupArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LogGroupArn\": ")
	if tmp, err := json.Marshal(strct.LogGroupArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LoggingConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "LogGroupArn":
            if err := json.Unmarshal([]byte(v), &strct.LogGroupArn); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AlertManagerDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AlertManagerDefinition\": ")
	if tmp, err := json.Marshal(strct.AlertManagerDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Alias" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Alias\": ")
	if tmp, err := json.Marshal(strct.Alias); err != nil {
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
    // Marshal the "LoggingConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LoggingConfiguration\": ")
	if tmp, err := json.Marshal(strct.LoggingConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PrometheusEndpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PrometheusEndpoint\": ")
	if tmp, err := json.Marshal(strct.PrometheusEndpoint); err != nil {
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
    // Marshal the "WorkspaceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkspaceId\": ")
	if tmp, err := json.Marshal(strct.WorkspaceId); err != nil {
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
        case "AlertManagerDefinition":
            if err := json.Unmarshal([]byte(v), &strct.AlertManagerDefinition); err != nil {
                return err
             }
        case "Alias":
            if err := json.Unmarshal([]byte(v), &strct.Alias); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "LoggingConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.LoggingConfiguration); err != nil {
                return err
             }
        case "PrometheusEndpoint":
            if err := json.Unmarshal([]byte(v), &strct.PrometheusEndpoint); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "WorkspaceId":
            if err := json.Unmarshal([]byte(v), &strct.WorkspaceId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
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
