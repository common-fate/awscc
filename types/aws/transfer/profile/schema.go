// Code generated by schema-generate. DO NOT EDIT.

package profile

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::Transfer::Profile
type Resource struct {

  // Specifies the unique Amazon Resource Name (ARN) for the profile.
  Arn string `json:"Arn,omitempty"`

  // AS2 identifier agreed with a trading partner.
  As2Id string `json:"As2Id"`

  // List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.
  CertificateIds []string `json:"CertificateIds,omitempty"`

  // A unique identifier for the profile
  ProfileId string `json:"ProfileId,omitempty"`

  // Enum specifying whether the profile is local or associated with a trading partner.
  ProfileType string `json:"ProfileType"`

  // An array of key-value pairs to apply to this resource.
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
    // "As2Id" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "As2Id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"As2Id\": ")
	if tmp, err := json.Marshal(strct.As2Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CertificateIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CertificateIds\": ")
	if tmp, err := json.Marshal(strct.CertificateIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProfileId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileId\": ")
	if tmp, err := json.Marshal(strct.ProfileId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ProfileType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ProfileType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileType\": ")
	if tmp, err := json.Marshal(strct.ProfileType); err != nil {
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
    As2IdReceived := false
    ProfileTypeReceived := false
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
        case "As2Id":
            if err := json.Unmarshal([]byte(v), &strct.As2Id); err != nil {
                return err
             }
            As2IdReceived = true
        case "CertificateIds":
            if err := json.Unmarshal([]byte(v), &strct.CertificateIds); err != nil {
                return err
             }
        case "ProfileId":
            if err := json.Unmarshal([]byte(v), &strct.ProfileId); err != nil {
                return err
             }
        case "ProfileType":
            if err := json.Unmarshal([]byte(v), &strct.ProfileType); err != nil {
                return err
             }
            ProfileTypeReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if As2Id (a required property) was received
    if !As2IdReceived {
        return errors.New("\"As2Id\" is required but was not present")
    }
    // check if ProfileType (a required property) was received
    if !ProfileTypeReceived {
        return errors.New("\"ProfileType\" is required but was not present")
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
