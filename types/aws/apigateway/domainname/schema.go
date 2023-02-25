// Code generated by schema-generate. DO NOT EDIT.

package domainname

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// EndpointConfiguration 
type EndpointConfiguration struct {
  Types []string `json:"Types,omitempty"`
}

// MutualTlsAuthentication 
type MutualTlsAuthentication struct {
  TruststoreUri string `json:"TruststoreUri,omitempty"`
  TruststoreVersion string `json:"TruststoreVersion,omitempty"`
}

// Resource Resource Type definition for AWS::ApiGateway::DomainName.
type Resource struct {
  CertificateArn string `json:"CertificateArn,omitempty"`
  DistributionDomainName string `json:"DistributionDomainName,omitempty"`
  DistributionHostedZoneId string `json:"DistributionHostedZoneId,omitempty"`
  DomainName string `json:"DomainName,omitempty"`
  EndpointConfiguration *EndpointConfiguration `json:"EndpointConfiguration,omitempty"`
  MutualTlsAuthentication *MutualTlsAuthentication `json:"MutualTlsAuthentication,omitempty"`
  OwnershipVerificationCertificateArn string `json:"OwnershipVerificationCertificateArn,omitempty"`
  RegionalCertificateArn string `json:"RegionalCertificateArn,omitempty"`
  RegionalDomainName string `json:"RegionalDomainName,omitempty"`
  RegionalHostedZoneId string `json:"RegionalHostedZoneId,omitempty"`
  SecurityPolicy string `json:"SecurityPolicy,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key,omitempty"`
  Value string `json:"Value,omitempty"`
}

func (strct *EndpointConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Types" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Types\": ")
	if tmp, err := json.Marshal(strct.Types); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EndpointConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Types":
            if err := json.Unmarshal([]byte(v), &strct.Types); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *MutualTlsAuthentication) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "TruststoreUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TruststoreUri\": ")
	if tmp, err := json.Marshal(strct.TruststoreUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TruststoreVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TruststoreVersion\": ")
	if tmp, err := json.Marshal(strct.TruststoreVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MutualTlsAuthentication) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "TruststoreUri":
            if err := json.Unmarshal([]byte(v), &strct.TruststoreUri); err != nil {
                return err
             }
        case "TruststoreVersion":
            if err := json.Unmarshal([]byte(v), &strct.TruststoreVersion); err != nil {
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
    // Marshal the "CertificateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CertificateArn\": ")
	if tmp, err := json.Marshal(strct.CertificateArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DistributionDomainName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DistributionDomainName\": ")
	if tmp, err := json.Marshal(strct.DistributionDomainName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DistributionHostedZoneId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DistributionHostedZoneId\": ")
	if tmp, err := json.Marshal(strct.DistributionHostedZoneId); err != nil {
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
    // Marshal the "EndpointConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EndpointConfiguration\": ")
	if tmp, err := json.Marshal(strct.EndpointConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MutualTlsAuthentication" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MutualTlsAuthentication\": ")
	if tmp, err := json.Marshal(strct.MutualTlsAuthentication); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OwnershipVerificationCertificateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OwnershipVerificationCertificateArn\": ")
	if tmp, err := json.Marshal(strct.OwnershipVerificationCertificateArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegionalCertificateArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionalCertificateArn\": ")
	if tmp, err := json.Marshal(strct.RegionalCertificateArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegionalDomainName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionalDomainName\": ")
	if tmp, err := json.Marshal(strct.RegionalDomainName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegionalHostedZoneId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionalHostedZoneId\": ")
	if tmp, err := json.Marshal(strct.RegionalHostedZoneId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecurityPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityPolicy\": ")
	if tmp, err := json.Marshal(strct.SecurityPolicy); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CertificateArn":
            if err := json.Unmarshal([]byte(v), &strct.CertificateArn); err != nil {
                return err
             }
        case "DistributionDomainName":
            if err := json.Unmarshal([]byte(v), &strct.DistributionDomainName); err != nil {
                return err
             }
        case "DistributionHostedZoneId":
            if err := json.Unmarshal([]byte(v), &strct.DistributionHostedZoneId); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
        case "EndpointConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EndpointConfiguration); err != nil {
                return err
             }
        case "MutualTlsAuthentication":
            if err := json.Unmarshal([]byte(v), &strct.MutualTlsAuthentication); err != nil {
                return err
             }
        case "OwnershipVerificationCertificateArn":
            if err := json.Unmarshal([]byte(v), &strct.OwnershipVerificationCertificateArn); err != nil {
                return err
             }
        case "RegionalCertificateArn":
            if err := json.Unmarshal([]byte(v), &strct.RegionalCertificateArn); err != nil {
                return err
             }
        case "RegionalDomainName":
            if err := json.Unmarshal([]byte(v), &strct.RegionalDomainName); err != nil {
                return err
             }
        case "RegionalHostedZoneId":
            if err := json.Unmarshal([]byte(v), &strct.RegionalHostedZoneId); err != nil {
                return err
             }
        case "SecurityPolicy":
            if err := json.Unmarshal([]byte(v), &strct.SecurityPolicy); err != nil {
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
