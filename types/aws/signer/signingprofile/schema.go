// Code generated by schema-generate. DO NOT EDIT.

package signingprofile

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource A signing profile is a signing template that can be used to carry out a pre-defined signing job.
type Resource struct {

  // The Amazon Resource Name (ARN) of the specified signing profile.
  Arn string `json:"Arn,omitempty"`

  // The ID of the target signing platform.
  PlatformId string `json:"PlatformId"`

  // A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. 
  ProfileName string `json:"ProfileName,omitempty"`

  // A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.
  ProfileVersion string `json:"ProfileVersion,omitempty"`

  // The Amazon Resource Name (ARN) of the specified signing profile version.
  ProfileVersionArn string `json:"ProfileVersionArn,omitempty"`

  // Signature validity period of the profile.
  SignatureValidityPeriod *SignatureValidityPeriod `json:"SignatureValidityPeriod,omitempty"`

  // A list of tags associated with the signing profile.
  Tags []*Tag `json:"Tags,omitempty"`
}

// SignatureValidityPeriod 
type SignatureValidityPeriod struct {
  Type string `json:"Type,omitempty"`
  Value int `json:"Value,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key,omitempty"`
  Value string `json:"Value,omitempty"`
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
    // "PlatformId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "PlatformId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PlatformId\": ")
	if tmp, err := json.Marshal(strct.PlatformId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProfileName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileName\": ")
	if tmp, err := json.Marshal(strct.ProfileName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProfileVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileVersion\": ")
	if tmp, err := json.Marshal(strct.ProfileVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ProfileVersionArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProfileVersionArn\": ")
	if tmp, err := json.Marshal(strct.ProfileVersionArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SignatureValidityPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SignatureValidityPeriod\": ")
	if tmp, err := json.Marshal(strct.SignatureValidityPeriod); err != nil {
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
    PlatformIdReceived := false
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
        case "PlatformId":
            if err := json.Unmarshal([]byte(v), &strct.PlatformId); err != nil {
                return err
             }
            PlatformIdReceived = true
        case "ProfileName":
            if err := json.Unmarshal([]byte(v), &strct.ProfileName); err != nil {
                return err
             }
        case "ProfileVersion":
            if err := json.Unmarshal([]byte(v), &strct.ProfileVersion); err != nil {
                return err
             }
        case "ProfileVersionArn":
            if err := json.Unmarshal([]byte(v), &strct.ProfileVersionArn); err != nil {
                return err
             }
        case "SignatureValidityPeriod":
            if err := json.Unmarshal([]byte(v), &strct.SignatureValidityPeriod); err != nil {
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
    // check if PlatformId (a required property) was received
    if !PlatformIdReceived {
        return errors.New("\"PlatformId\" is required but was not present")
    }
    return nil
}

func (strct *SignatureValidityPeriod) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *SignatureValidityPeriod) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
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
