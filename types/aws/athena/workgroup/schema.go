// Code generated by schema-generate. DO NOT EDIT.

package workgroup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// EncryptionConfiguration If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
type EncryptionConfiguration struct {
  EncryptionOption string `json:"EncryptionOption"`
  KmsKey string `json:"KmsKey,omitempty"`
}

// EngineVersion The Athena engine version for running queries.
type EngineVersion struct {
  EffectiveEngineVersion string `json:"EffectiveEngineVersion,omitempty"`
  SelectedEngineVersion string `json:"SelectedEngineVersion,omitempty"`
}

// Resource Resource schema for AWS::Athena::WorkGroup
type Resource struct {

  // The date and time the workgroup was created.
  CreationTime string `json:"CreationTime,omitempty"`

  // The workgroup description.
  Description string `json:"Description,omitempty"`

  // The workGroup name.
  Name string `json:"Name"`

  // The option to delete the workgroup and its contents even if the workgroup contains any named queries.
  RecursiveDeleteOption bool `json:"RecursiveDeleteOption,omitempty"`

  // The state of the workgroup: ENABLED or DISABLED.
  State string `json:"State,omitempty"`

  // One or more tags, separated by commas, that you want to attach to the workgroup as you create it
  Tags []*Tag `json:"Tags,omitempty"`

  // The workgroup configuration
  WorkGroupConfiguration *WorkGroupConfiguration `json:"WorkGroupConfiguration,omitempty"`

  // The workgroup configuration update object
  WorkGroupConfigurationUpdates *WorkGroupConfigurationUpdates `json:"WorkGroupConfigurationUpdates,omitempty"`
}

// ResultConfiguration The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
// 
type ResultConfiguration struct {
  EncryptionConfiguration *EncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
  OutputLocation string `json:"OutputLocation,omitempty"`
}

// ResultConfigurationUpdates The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. 
type ResultConfigurationUpdates struct {
  EncryptionConfiguration *EncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
  OutputLocation string `json:"OutputLocation,omitempty"`
  RemoveEncryptionConfiguration bool `json:"RemoveEncryptionConfiguration,omitempty"`
  RemoveOutputLocation bool `json:"RemoveOutputLocation,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

// WorkGroupConfiguration 
type WorkGroupConfiguration struct {
  BytesScannedCutoffPerQuery int `json:"BytesScannedCutoffPerQuery,omitempty"`
  EnforceWorkGroupConfiguration bool `json:"EnforceWorkGroupConfiguration,omitempty"`
  EngineVersion *EngineVersion `json:"EngineVersion,omitempty"`
  PublishCloudWatchMetricsEnabled bool `json:"PublishCloudWatchMetricsEnabled,omitempty"`
  RequesterPaysEnabled bool `json:"RequesterPaysEnabled,omitempty"`
  ResultConfiguration *ResultConfiguration `json:"ResultConfiguration,omitempty"`
}

// WorkGroupConfigurationUpdates The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. 
type WorkGroupConfigurationUpdates struct {
  BytesScannedCutoffPerQuery int `json:"BytesScannedCutoffPerQuery,omitempty"`
  EnforceWorkGroupConfiguration bool `json:"EnforceWorkGroupConfiguration,omitempty"`
  EngineVersion *EngineVersion `json:"EngineVersion,omitempty"`
  PublishCloudWatchMetricsEnabled bool `json:"PublishCloudWatchMetricsEnabled,omitempty"`
  RemoveBytesScannedCutoffPerQuery bool `json:"RemoveBytesScannedCutoffPerQuery,omitempty"`
  RequesterPaysEnabled bool `json:"RequesterPaysEnabled,omitempty"`
  ResultConfigurationUpdates *ResultConfigurationUpdates `json:"ResultConfigurationUpdates,omitempty"`
}

func (strct *EncryptionConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "EncryptionOption" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "EncryptionOption" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionOption\": ")
	if tmp, err := json.Marshal(strct.EncryptionOption); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KmsKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KmsKey\": ")
	if tmp, err := json.Marshal(strct.KmsKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EncryptionConfiguration) UnmarshalJSON(b []byte) error {
    EncryptionOptionReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EncryptionOption":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionOption); err != nil {
                return err
             }
            EncryptionOptionReceived = true
        case "KmsKey":
            if err := json.Unmarshal([]byte(v), &strct.KmsKey); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if EncryptionOption (a required property) was received
    if !EncryptionOptionReceived {
        return errors.New("\"EncryptionOption\" is required but was not present")
    }
    return nil
}

func (strct *EngineVersion) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EffectiveEngineVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EffectiveEngineVersion\": ")
	if tmp, err := json.Marshal(strct.EffectiveEngineVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SelectedEngineVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SelectedEngineVersion\": ")
	if tmp, err := json.Marshal(strct.SelectedEngineVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EngineVersion) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EffectiveEngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.EffectiveEngineVersion); err != nil {
                return err
             }
        case "SelectedEngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.SelectedEngineVersion); err != nil {
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
    // Marshal the "RecursiveDeleteOption" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RecursiveDeleteOption\": ")
	if tmp, err := json.Marshal(strct.RecursiveDeleteOption); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "State" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"State\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
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
    // Marshal the "WorkGroupConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkGroupConfiguration\": ")
	if tmp, err := json.Marshal(strct.WorkGroupConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "WorkGroupConfigurationUpdates" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WorkGroupConfigurationUpdates\": ")
	if tmp, err := json.Marshal(strct.WorkGroupConfigurationUpdates); err != nil {
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
    NameReceived := false
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
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "RecursiveDeleteOption":
            if err := json.Unmarshal([]byte(v), &strct.RecursiveDeleteOption); err != nil {
                return err
             }
        case "State":
            if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "WorkGroupConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.WorkGroupConfiguration); err != nil {
                return err
             }
        case "WorkGroupConfigurationUpdates":
            if err := json.Unmarshal([]byte(v), &strct.WorkGroupConfigurationUpdates); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}

func (strct *ResultConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EncryptionConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionConfiguration\": ")
	if tmp, err := json.Marshal(strct.EncryptionConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OutputLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputLocation\": ")
	if tmp, err := json.Marshal(strct.OutputLocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResultConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EncryptionConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionConfiguration); err != nil {
                return err
             }
        case "OutputLocation":
            if err := json.Unmarshal([]byte(v), &strct.OutputLocation); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ResultConfigurationUpdates) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EncryptionConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionConfiguration\": ")
	if tmp, err := json.Marshal(strct.EncryptionConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OutputLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputLocation\": ")
	if tmp, err := json.Marshal(strct.OutputLocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RemoveEncryptionConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RemoveEncryptionConfiguration\": ")
	if tmp, err := json.Marshal(strct.RemoveEncryptionConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RemoveOutputLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RemoveOutputLocation\": ")
	if tmp, err := json.Marshal(strct.RemoveOutputLocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResultConfigurationUpdates) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EncryptionConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionConfiguration); err != nil {
                return err
             }
        case "OutputLocation":
            if err := json.Unmarshal([]byte(v), &strct.OutputLocation); err != nil {
                return err
             }
        case "RemoveEncryptionConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.RemoveEncryptionConfiguration); err != nil {
                return err
             }
        case "RemoveOutputLocation":
            if err := json.Unmarshal([]byte(v), &strct.RemoveOutputLocation); err != nil {
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

func (strct *WorkGroupConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "BytesScannedCutoffPerQuery" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BytesScannedCutoffPerQuery\": ")
	if tmp, err := json.Marshal(strct.BytesScannedCutoffPerQuery); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnforceWorkGroupConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnforceWorkGroupConfiguration\": ")
	if tmp, err := json.Marshal(strct.EnforceWorkGroupConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EngineVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EngineVersion\": ")
	if tmp, err := json.Marshal(strct.EngineVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublishCloudWatchMetricsEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublishCloudWatchMetricsEnabled\": ")
	if tmp, err := json.Marshal(strct.PublishCloudWatchMetricsEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RequesterPaysEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RequesterPaysEnabled\": ")
	if tmp, err := json.Marshal(strct.RequesterPaysEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResultConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResultConfiguration\": ")
	if tmp, err := json.Marshal(strct.ResultConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WorkGroupConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BytesScannedCutoffPerQuery":
            if err := json.Unmarshal([]byte(v), &strct.BytesScannedCutoffPerQuery); err != nil {
                return err
             }
        case "EnforceWorkGroupConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EnforceWorkGroupConfiguration); err != nil {
                return err
             }
        case "EngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.EngineVersion); err != nil {
                return err
             }
        case "PublishCloudWatchMetricsEnabled":
            if err := json.Unmarshal([]byte(v), &strct.PublishCloudWatchMetricsEnabled); err != nil {
                return err
             }
        case "RequesterPaysEnabled":
            if err := json.Unmarshal([]byte(v), &strct.RequesterPaysEnabled); err != nil {
                return err
             }
        case "ResultConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.ResultConfiguration); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *WorkGroupConfigurationUpdates) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "BytesScannedCutoffPerQuery" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BytesScannedCutoffPerQuery\": ")
	if tmp, err := json.Marshal(strct.BytesScannedCutoffPerQuery); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnforceWorkGroupConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnforceWorkGroupConfiguration\": ")
	if tmp, err := json.Marshal(strct.EnforceWorkGroupConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EngineVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EngineVersion\": ")
	if tmp, err := json.Marshal(strct.EngineVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublishCloudWatchMetricsEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublishCloudWatchMetricsEnabled\": ")
	if tmp, err := json.Marshal(strct.PublishCloudWatchMetricsEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RemoveBytesScannedCutoffPerQuery" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RemoveBytesScannedCutoffPerQuery\": ")
	if tmp, err := json.Marshal(strct.RemoveBytesScannedCutoffPerQuery); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RequesterPaysEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RequesterPaysEnabled\": ")
	if tmp, err := json.Marshal(strct.RequesterPaysEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResultConfigurationUpdates" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResultConfigurationUpdates\": ")
	if tmp, err := json.Marshal(strct.ResultConfigurationUpdates); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WorkGroupConfigurationUpdates) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BytesScannedCutoffPerQuery":
            if err := json.Unmarshal([]byte(v), &strct.BytesScannedCutoffPerQuery); err != nil {
                return err
             }
        case "EnforceWorkGroupConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EnforceWorkGroupConfiguration); err != nil {
                return err
             }
        case "EngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.EngineVersion); err != nil {
                return err
             }
        case "PublishCloudWatchMetricsEnabled":
            if err := json.Unmarshal([]byte(v), &strct.PublishCloudWatchMetricsEnabled); err != nil {
                return err
             }
        case "RemoveBytesScannedCutoffPerQuery":
            if err := json.Unmarshal([]byte(v), &strct.RemoveBytesScannedCutoffPerQuery); err != nil {
                return err
             }
        case "RequesterPaysEnabled":
            if err := json.Unmarshal([]byte(v), &strct.RequesterPaysEnabled); err != nil {
                return err
             }
        case "ResultConfigurationUpdates":
            if err := json.Unmarshal([]byte(v), &strct.ResultConfigurationUpdates); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
