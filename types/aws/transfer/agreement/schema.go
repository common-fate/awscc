// Code generated by schema-generate. DO NOT EDIT.

package agreement

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::Transfer::Agreement
type Resource struct {

  // Specifies the access role for the agreement.
  AccessRole string `json:"AccessRole"`

  // A unique identifier for the agreement.
  AgreementId string `json:"AgreementId,omitempty"`

  // Specifies the unique Amazon Resource Name (ARN) for the agreement.
  Arn string `json:"Arn,omitempty"`

  // Specifies the base directory for the agreement.
  BaseDirectory string `json:"BaseDirectory"`

  // A textual description for the agreement.
  Description string `json:"Description,omitempty"`

  // A unique identifier for the local profile.
  LocalProfileId string `json:"LocalProfileId"`

  // A unique identifier for the partner profile.
  PartnerProfileId string `json:"PartnerProfileId"`

  // A unique identifier for the server.
  ServerId string `json:"ServerId"`

  // Specifies the status of the agreement.
  Status string `json:"Status,omitempty"`

  // Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag Creates a key-value pair for a specific resource.
type Tag struct {

  // The name assigned to the tag that you create.
  Key string `json:"Key"`

  // Contains one or more values that you assigned to the key name you create.
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AccessRole" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AccessRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessRole\": ")
	if tmp, err := json.Marshal(strct.AccessRole); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AgreementId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AgreementId\": ")
	if tmp, err := json.Marshal(strct.AgreementId); err != nil {
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
    // "BaseDirectory" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BaseDirectory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BaseDirectory\": ")
	if tmp, err := json.Marshal(strct.BaseDirectory); err != nil {
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
    // "LocalProfileId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "LocalProfileId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LocalProfileId\": ")
	if tmp, err := json.Marshal(strct.LocalProfileId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "PartnerProfileId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PartnerProfileId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PartnerProfileId\": ")
	if tmp, err := json.Marshal(strct.PartnerProfileId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ServerId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ServerId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServerId\": ")
	if tmp, err := json.Marshal(strct.ServerId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
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
    AccessRoleReceived := false
    BaseDirectoryReceived := false
    LocalProfileIdReceived := false
    PartnerProfileIdReceived := false
    ServerIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AccessRole":
            if err := json.Unmarshal([]byte(v), &strct.AccessRole); err != nil {
                return err
             }
            AccessRoleReceived = true
        case "AgreementId":
            if err := json.Unmarshal([]byte(v), &strct.AgreementId); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "BaseDirectory":
            if err := json.Unmarshal([]byte(v), &strct.BaseDirectory); err != nil {
                return err
             }
            BaseDirectoryReceived = true
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "LocalProfileId":
            if err := json.Unmarshal([]byte(v), &strct.LocalProfileId); err != nil {
                return err
             }
            LocalProfileIdReceived = true
        case "PartnerProfileId":
            if err := json.Unmarshal([]byte(v), &strct.PartnerProfileId); err != nil {
                return err
             }
            PartnerProfileIdReceived = true
        case "ServerId":
            if err := json.Unmarshal([]byte(v), &strct.ServerId); err != nil {
                return err
             }
            ServerIdReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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
    // check if AccessRole (a required property) was received
    if !AccessRoleReceived {
        return errors.New("\"AccessRole\" is required but was not present")
    }
    // check if BaseDirectory (a required property) was received
    if !BaseDirectoryReceived {
        return errors.New("\"BaseDirectory\" is required but was not present")
    }
    // check if LocalProfileId (a required property) was received
    if !LocalProfileIdReceived {
        return errors.New("\"LocalProfileId\" is required but was not present")
    }
    // check if PartnerProfileId (a required property) was received
    if !PartnerProfileIdReceived {
        return errors.New("\"PartnerProfileId\" is required but was not present")
    }
    // check if ServerId (a required property) was received
    if !ServerIdReceived {
        return errors.New("\"ServerId\" is required but was not present")
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
