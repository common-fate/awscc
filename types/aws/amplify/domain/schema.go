// Code generated by schema-generate. DO NOT EDIT.

package domain

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.
type Resource struct {
  AppId string `json:"AppId"`
  Arn string `json:"Arn,omitempty"`
  AutoSubDomainCreationPatterns []string `json:"AutoSubDomainCreationPatterns,omitempty"`
  AutoSubDomainIAMRole string `json:"AutoSubDomainIAMRole,omitempty"`
  CertificateRecord string `json:"CertificateRecord,omitempty"`
  DomainName string `json:"DomainName"`
  DomainStatus string `json:"DomainStatus,omitempty"`
  EnableAutoSubDomain bool `json:"EnableAutoSubDomain,omitempty"`
  StatusReason string `json:"StatusReason,omitempty"`
  SubDomainSettings []*SubDomainSetting `json:"SubDomainSettings"`
}

// SubDomainSetting 
type SubDomainSetting struct {
  BranchName string `json:"BranchName"`
  Prefix string `json:"Prefix"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AppId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AppId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppId\": ")
	if tmp, err := json.Marshal(strct.AppId); err != nil {
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
    // Marshal the "AutoSubDomainCreationPatterns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoSubDomainCreationPatterns\": ")
	if tmp, err := json.Marshal(strct.AutoSubDomainCreationPatterns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AutoSubDomainIAMRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoSubDomainIAMRole\": ")
	if tmp, err := json.Marshal(strct.AutoSubDomainIAMRole); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CertificateRecord" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CertificateRecord\": ")
	if tmp, err := json.Marshal(strct.CertificateRecord); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DomainName" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "DomainStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainStatus\": ")
	if tmp, err := json.Marshal(strct.DomainStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnableAutoSubDomain" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableAutoSubDomain\": ")
	if tmp, err := json.Marshal(strct.EnableAutoSubDomain); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StatusReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StatusReason\": ")
	if tmp, err := json.Marshal(strct.StatusReason); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SubDomainSettings" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SubDomainSettings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubDomainSettings\": ")
	if tmp, err := json.Marshal(strct.SubDomainSettings); err != nil {
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
    AppIdReceived := false
    DomainNameReceived := false
    SubDomainSettingsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AppId":
            if err := json.Unmarshal([]byte(v), &strct.AppId); err != nil {
                return err
             }
            AppIdReceived = true
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AutoSubDomainCreationPatterns":
            if err := json.Unmarshal([]byte(v), &strct.AutoSubDomainCreationPatterns); err != nil {
                return err
             }
        case "AutoSubDomainIAMRole":
            if err := json.Unmarshal([]byte(v), &strct.AutoSubDomainIAMRole); err != nil {
                return err
             }
        case "CertificateRecord":
            if err := json.Unmarshal([]byte(v), &strct.CertificateRecord); err != nil {
                return err
             }
        case "DomainName":
            if err := json.Unmarshal([]byte(v), &strct.DomainName); err != nil {
                return err
             }
            DomainNameReceived = true
        case "DomainStatus":
            if err := json.Unmarshal([]byte(v), &strct.DomainStatus); err != nil {
                return err
             }
        case "EnableAutoSubDomain":
            if err := json.Unmarshal([]byte(v), &strct.EnableAutoSubDomain); err != nil {
                return err
             }
        case "StatusReason":
            if err := json.Unmarshal([]byte(v), &strct.StatusReason); err != nil {
                return err
             }
        case "SubDomainSettings":
            if err := json.Unmarshal([]byte(v), &strct.SubDomainSettings); err != nil {
                return err
             }
            SubDomainSettingsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AppId (a required property) was received
    if !AppIdReceived {
        return errors.New("\"AppId\" is required but was not present")
    }
    // check if DomainName (a required property) was received
    if !DomainNameReceived {
        return errors.New("\"DomainName\" is required but was not present")
    }
    // check if SubDomainSettings (a required property) was received
    if !SubDomainSettingsReceived {
        return errors.New("\"SubDomainSettings\" is required but was not present")
    }
    return nil
}

func (strct *SubDomainSetting) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "BranchName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BranchName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BranchName\": ")
	if tmp, err := json.Marshal(strct.BranchName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Prefix" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Prefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Prefix\": ")
	if tmp, err := json.Marshal(strct.Prefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SubDomainSetting) UnmarshalJSON(b []byte) error {
    BranchNameReceived := false
    PrefixReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BranchName":
            if err := json.Unmarshal([]byte(v), &strct.BranchName); err != nil {
                return err
             }
            BranchNameReceived = true
        case "Prefix":
            if err := json.Unmarshal([]byte(v), &strct.Prefix); err != nil {
                return err
             }
            PrefixReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if BranchName (a required property) was received
    if !BranchNameReceived {
        return errors.New("\"BranchName\" is required but was not present")
    }
    // check if Prefix (a required property) was received
    if !PrefixReceived {
        return errors.New("\"Prefix\" is required but was not present")
    }
    return nil
}
