// Code generated by schema-generate. DO NOT EDIT.

package domainconfiguration

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// AuthorizerConfig 
type AuthorizerConfig struct {
  AllowAuthorizerOverride bool `json:"AllowAuthorizerOverride,omitempty"`
  DefaultAuthorizerName string `json:"DefaultAuthorizerName,omitempty"`
}

// Resource Create and manage a Domain Configuration
type Resource struct {
  Arn string `json:"Arn,omitempty"`
  AuthorizerConfig *AuthorizerConfig `json:"AuthorizerConfig,omitempty"`
  DomainConfigurationName string `json:"DomainConfigurationName,omitempty"`
  DomainConfigurationStatus string `json:"DomainConfigurationStatus,omitempty"`
  DomainName string `json:"DomainName,omitempty"`
  DomainType string `json:"DomainType,omitempty"`
  ServerCertificateArns []string `json:"ServerCertificateArns,omitempty"`
  ServerCertificates []*ServerCertificateSummary `json:"ServerCertificates,omitempty"`
  ServiceType string `json:"ServiceType,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
  ValidationCertificateArn string `json:"ValidationCertificateArn,omitempty"`
}

// ServerCertificateSummary 
type ServerCertificateSummary struct {
  ServerCertificateArn string `json:"ServerCertificateArn,omitempty"`
  ServerCertificateStatus string `json:"ServerCertificateStatus,omitempty"`
  ServerCertificateStatusDetail string `json:"ServerCertificateStatusDetail,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *AuthorizerConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AllowAuthorizerOverride" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowAuthorizerOverride\": ")
	if tmp, err := json.Marshal(strct.AllowAuthorizerOverride); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultAuthorizerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultAuthorizerName\": ")
	if tmp, err := json.Marshal(strct.DefaultAuthorizerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AuthorizerConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AllowAuthorizerOverride":
            if err := json.Unmarshal([]byte(v), &strct.AllowAuthorizerOverride); err != nil {
                return err
             }
        case "DefaultAuthorizerName":
            if err := json.Unmarshal([]byte(v), &strct.DefaultAuthorizerName); err != nil {
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
    // Marshal the "AuthorizerConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerConfig\": ")
	if tmp, err := json.Marshal(strct.AuthorizerConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainConfigurationName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainConfigurationName\": ")
	if tmp, err := json.Marshal(strct.DomainConfigurationName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainConfigurationStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainConfigurationStatus\": ")
	if tmp, err := json.Marshal(strct.DomainConfigurationStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainName\": ")
	if tmp, err := json.Marshal(strct.DomainName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainType\": ")
	if tmp, err := json.Marshal(strct.DomainType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServerCertificateArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerCertificateArns\": ")
	if tmp, err := json.Marshal(strct.ServerCertificateArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServerCertificates" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerCertificates\": ")
	if tmp, err := json.Marshal(strct.ServerCertificates); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServiceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceType\": ")
	if tmp, err := json.Marshal(strct.ServiceType); err != nil {
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
    // Marshal the "ValidationCertificateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ValidationCertificateArn\": ")
	if tmp, err := json.Marshal(strct.ValidationCertificateArn); err != nil {
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
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AuthorizerConfig":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerConfig); err != nil {
                return err
             }
        case "DomainConfigurationName":
            if err := json.Unmarshal([]byte(v), &strct.DomainConfigurationName); err != nil {
                return err
             }
        case "DomainConfigurationStatus":
            if err := json.Unmarshal([]byte(v), &strct.DomainConfigurationStatus); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
        case "DomainType":
            if err := json.Unmarshal([]byte(v), &strct.DomainType); err != nil {
                return err
             }
        case "ServerCertificateArns":
            if err := json.Unmarshal([]byte(v), &strct.ServerCertificateArns); err != nil {
                return err
             }
        case "ServerCertificates":
            if err := json.Unmarshal([]byte(v), &strct.ServerCertificates); err != nil {
                return err
             }
        case "ServiceType":
            if err := json.Unmarshal([]byte(v), &strct.ServiceType); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "ValidationCertificateArn":
            if err := json.Unmarshal([]byte(v), &strct.ValidationCertificateArn); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ServerCertificateSummary) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ServerCertificateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerCertificateArn\": ")
	if tmp, err := json.Marshal(strct.ServerCertificateArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServerCertificateStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerCertificateStatus\": ")
	if tmp, err := json.Marshal(strct.ServerCertificateStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServerCertificateStatusDetail" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerCertificateStatusDetail\": ")
	if tmp, err := json.Marshal(strct.ServerCertificateStatusDetail); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ServerCertificateSummary) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ServerCertificateArn":
            if err := json.Unmarshal([]byte(v), &strct.ServerCertificateArn); err != nil {
                return err
             }
        case "ServerCertificateStatus":
            if err := json.Unmarshal([]byte(v), &strct.ServerCertificateStatus); err != nil {
                return err
             }
        case "ServerCertificateStatusDetail":
            if err := json.Unmarshal([]byte(v), &strct.ServerCertificateStatusDetail); err != nil {
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
