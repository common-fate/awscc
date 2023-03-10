// Code generated by schema-generate. DO NOT EDIT.

package continuousdeploymentpolicy

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ContinuousDeploymentPolicyConfig 
type ContinuousDeploymentPolicyConfig struct {
  Enabled bool `json:"Enabled"`
  StagingDistributionDnsNames []string `json:"StagingDistributionDnsNames"`
  TrafficConfig *TrafficConfig `json:"TrafficConfig,omitempty"`
}

// Resource Resource Type definition for AWS::CloudFront::ContinuousDeploymentPolicy
type Resource struct {
  ContinuousDeploymentPolicyConfig *ContinuousDeploymentPolicyConfig `json:"ContinuousDeploymentPolicyConfig"`
  Id string `json:"Id,omitempty"`
  LastModifiedTime string `json:"LastModifiedTime,omitempty"`
}

// SessionStickinessConfig 
type SessionStickinessConfig struct {
  IdleTTL int `json:"IdleTTL"`
  MaximumTTL int `json:"MaximumTTL"`
}

// SingleHeaderConfig 
type SingleHeaderConfig struct {
  Header string `json:"Header"`
  Value string `json:"Value"`
}

// SingleWeightConfig 
type SingleWeightConfig struct {
  SessionStickinessConfig *SessionStickinessConfig `json:"SessionStickinessConfig,omitempty"`
  Weight float64 `json:"Weight"`
}

// TrafficConfig 
type TrafficConfig struct {
  SingleHeaderConfig *SingleHeaderConfig `json:"SingleHeaderConfig,omitempty"`
  SingleWeightConfig *SingleWeightConfig `json:"SingleWeightConfig,omitempty"`
  Type string `json:"Type"`
}

func (strct *ContinuousDeploymentPolicyConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Enabled" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Enabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Enabled\": ")
	if tmp, err := json.Marshal(strct.Enabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "StagingDistributionDnsNames" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "StagingDistributionDnsNames" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StagingDistributionDnsNames\": ")
	if tmp, err := json.Marshal(strct.StagingDistributionDnsNames); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TrafficConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TrafficConfig\": ")
	if tmp, err := json.Marshal(strct.TrafficConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ContinuousDeploymentPolicyConfig) UnmarshalJSON(b []byte) error {
    EnabledReceived := false
    StagingDistributionDnsNamesReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Enabled":
            if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
                return err
             }
            EnabledReceived = true
        case "StagingDistributionDnsNames":
            if err := json.Unmarshal([]byte(v), &strct.StagingDistributionDnsNames); err != nil {
                return err
             }
            StagingDistributionDnsNamesReceived = true
        case "TrafficConfig":
            if err := json.Unmarshal([]byte(v), &strct.TrafficConfig); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Enabled (a required property) was received
    if !EnabledReceived {
        return errors.New("\"Enabled\" is required but was not present")
    }
    // check if StagingDistributionDnsNames (a required property) was received
    if !StagingDistributionDnsNamesReceived {
        return errors.New("\"StagingDistributionDnsNames\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ContinuousDeploymentPolicyConfig" field is required
    if strct.ContinuousDeploymentPolicyConfig == nil {
        return nil, errors.New("ContinuousDeploymentPolicyConfig is a required field")
    }
    // Marshal the "ContinuousDeploymentPolicyConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContinuousDeploymentPolicyConfig\": ")
	if tmp, err := json.Marshal(strct.ContinuousDeploymentPolicyConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "LastModifiedTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LastModifiedTime\": ")
	if tmp, err := json.Marshal(strct.LastModifiedTime); err != nil {
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
    ContinuousDeploymentPolicyConfigReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContinuousDeploymentPolicyConfig":
            if err := json.Unmarshal([]byte(v), &strct.ContinuousDeploymentPolicyConfig); err != nil {
                return err
             }
            ContinuousDeploymentPolicyConfigReceived = true
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "LastModifiedTime":
            if err := json.Unmarshal([]byte(v), &strct.LastModifiedTime); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ContinuousDeploymentPolicyConfig (a required property) was received
    if !ContinuousDeploymentPolicyConfigReceived {
        return errors.New("\"ContinuousDeploymentPolicyConfig\" is required but was not present")
    }
    return nil
}

func (strct *SessionStickinessConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "IdleTTL" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "IdleTTL" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdleTTL\": ")
	if tmp, err := json.Marshal(strct.IdleTTL); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MaximumTTL" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "MaximumTTL" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaximumTTL\": ")
	if tmp, err := json.Marshal(strct.MaximumTTL); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SessionStickinessConfig) UnmarshalJSON(b []byte) error {
    IdleTTLReceived := false
    MaximumTTLReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "IdleTTL":
            if err := json.Unmarshal([]byte(v), &strct.IdleTTL); err != nil {
                return err
             }
            IdleTTLReceived = true
        case "MaximumTTL":
            if err := json.Unmarshal([]byte(v), &strct.MaximumTTL); err != nil {
                return err
             }
            MaximumTTLReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if IdleTTL (a required property) was received
    if !IdleTTLReceived {
        return errors.New("\"IdleTTL\" is required but was not present")
    }
    // check if MaximumTTL (a required property) was received
    if !MaximumTTLReceived {
        return errors.New("\"MaximumTTL\" is required but was not present")
    }
    return nil
}

func (strct *SingleHeaderConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Header" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Header" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Header\": ")
	if tmp, err := json.Marshal(strct.Header); err != nil {
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

func (strct *SingleHeaderConfig) UnmarshalJSON(b []byte) error {
    HeaderReceived := false
    ValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Header":
            if err := json.Unmarshal([]byte(v), &strct.Header); err != nil {
                return err
             }
            HeaderReceived = true
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            ValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Header (a required property) was received
    if !HeaderReceived {
        return errors.New("\"Header\" is required but was not present")
    }
    // check if Value (a required property) was received
    if !ValueReceived {
        return errors.New("\"Value\" is required but was not present")
    }
    return nil
}

func (strct *SingleWeightConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "SessionStickinessConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SessionStickinessConfig\": ")
	if tmp, err := json.Marshal(strct.SessionStickinessConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Weight" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Weight" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Weight\": ")
	if tmp, err := json.Marshal(strct.Weight); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SingleWeightConfig) UnmarshalJSON(b []byte) error {
    WeightReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "SessionStickinessConfig":
            if err := json.Unmarshal([]byte(v), &strct.SessionStickinessConfig); err != nil {
                return err
             }
        case "Weight":
            if err := json.Unmarshal([]byte(v), &strct.Weight); err != nil {
                return err
             }
            WeightReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Weight (a required property) was received
    if !WeightReceived {
        return errors.New("\"Weight\" is required but was not present")
    }
    return nil
}

func (strct *TrafficConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "SingleHeaderConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SingleHeaderConfig\": ")
	if tmp, err := json.Marshal(strct.SingleHeaderConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SingleWeightConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SingleWeightConfig\": ")
	if tmp, err := json.Marshal(strct.SingleWeightConfig); err != nil {
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

func (strct *TrafficConfig) UnmarshalJSON(b []byte) error {
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "SingleHeaderConfig":
            if err := json.Unmarshal([]byte(v), &strct.SingleHeaderConfig); err != nil {
                return err
             }
        case "SingleWeightConfig":
            if err := json.Unmarshal([]byte(v), &strct.SingleWeightConfig); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
    }
    return nil
}
