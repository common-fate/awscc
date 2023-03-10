// Code generated by schema-generate. DO NOT EDIT.

package phonenumber

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// Resource Resource Type definition for AWS::Connect::PhoneNumber
type Resource struct {

  // The phone number e164 address.
  Address string `json:"Address,omitempty"`

  // The phone number country code.
  CountryCode string `json:"CountryCode"`

  // The description of the phone number.
  Description string `json:"Description,omitempty"`

  // The phone number ARN
  PhoneNumberArn string `json:"PhoneNumberArn,omitempty"`

  // The phone number prefix.
  Prefix string `json:"Prefix,omitempty"`

  // One or more tags.
  Tags []*Tag `json:"Tags,omitempty"`

  // The ARN of the target the phone number is claimed to.
  TargetArn string `json:"TargetArn"`

  // The phone number type, either TOLL_FREE or DID.
  Type string `json:"Type"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Address" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Address\": ")
	if tmp, err := json.Marshal(strct.Address); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "CountryCode" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "CountryCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CountryCode\": ")
	if tmp, err := json.Marshal(strct.CountryCode); err != nil {
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
    // Marshal the "PhoneNumberArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PhoneNumberArn\": ")
	if tmp, err := json.Marshal(strct.PhoneNumberArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // "TargetArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetArn\": ")
	if tmp, err := json.Marshal(strct.TargetArn); err != nil {
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
    CountryCodeReceived := false
    TargetArnReceived := false
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Address":
            if err := json.Unmarshal([]byte(v), &strct.Address); err != nil {
                return err
             }
        case "CountryCode":
            if err := json.Unmarshal([]byte(v), &strct.CountryCode); err != nil {
                return err
             }
            CountryCodeReceived = true
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "PhoneNumberArn":
            if err := json.Unmarshal([]byte(v), &strct.PhoneNumberArn); err != nil {
                return err
             }
        case "Prefix":
            if err := json.Unmarshal([]byte(v), &strct.Prefix); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "TargetArn":
            if err := json.Unmarshal([]byte(v), &strct.TargetArn); err != nil {
                return err
             }
            TargetArnReceived = true
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if CountryCode (a required property) was received
    if !CountryCodeReceived {
        return errors.New("\"CountryCode\" is required but was not present")
    }
    // check if TargetArn (a required property) was received
    if !TargetArnReceived {
        return errors.New("\"TargetArn\" is required but was not present")
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
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
