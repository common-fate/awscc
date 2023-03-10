// Code generated by schema-generate. DO NOT EDIT.

package account

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource You can use AWS::Organizations::Account to manage accounts in organization.
type Resource struct {

  // If the account was created successfully, the unique identifier (ID) of the new account.
  AccountId string `json:"AccountId,omitempty"`

  // The friendly name of the member account.
  AccountName string `json:"AccountName"`

  // The Amazon Resource Name (ARN) of the account.
  Arn string `json:"Arn,omitempty"`

  // The email address of the owner to assign to the new member account.
  Email string `json:"Email"`

  // The method by which the account joined the organization.
  JoinedMethod string `json:"JoinedMethod,omitempty"`

  // The date the account became a part of the organization.
  JoinedTimestamp string `json:"JoinedTimestamp,omitempty"`

  // List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root.
  ParentIds []string `json:"ParentIds,omitempty"`

  // The name of an IAM role that AWS Organizations automatically preconfigures in the new member account. Default name is OrganizationAccountAccessRole if not specified.
  RoleName string `json:"RoleName,omitempty"`

  // The status of the account in the organization.
  Status string `json:"Status,omitempty"`

  // A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A custom key-value pair associated with a resource within your organization.
type Tag struct {

  // The key identifier, or name, of the tag.
  Key string `json:"Key"`

  // The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AccountId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccountId\": ")
	if tmp, err := json.Marshal(strct.AccountId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "AccountName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AccountName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccountName\": ")
	if tmp, err := json.Marshal(strct.AccountName); err != nil {
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
    // "Email" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Email" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Email\": ")
	if tmp, err := json.Marshal(strct.Email); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "JoinedMethod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"JoinedMethod\": ")
	if tmp, err := json.Marshal(strct.JoinedMethod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "JoinedTimestamp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"JoinedTimestamp\": ")
	if tmp, err := json.Marshal(strct.JoinedTimestamp); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ParentIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ParentIds\": ")
	if tmp, err := json.Marshal(strct.ParentIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RoleName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleName\": ")
	if tmp, err := json.Marshal(strct.RoleName); err != nil {
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
    AccountNameReceived := false
    EmailReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AccountId":
            if err := json.Unmarshal([]byte(v), &strct.AccountId); err != nil {
                return err
             }
        case "AccountName":
            if err := json.Unmarshal([]byte(v), &strct.AccountName); err != nil {
                return err
             }
            AccountNameReceived = true
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "Email":
            if err := json.Unmarshal([]byte(v), &strct.Email); err != nil {
                return err
             }
            EmailReceived = true
        case "JoinedMethod":
            if err := json.Unmarshal([]byte(v), &strct.JoinedMethod); err != nil {
                return err
             }
        case "JoinedTimestamp":
            if err := json.Unmarshal([]byte(v), &strct.JoinedTimestamp); err != nil {
                return err
             }
        case "ParentIds":
            if err := json.Unmarshal([]byte(v), &strct.ParentIds); err != nil {
                return err
             }
        case "RoleName":
            if err := json.Unmarshal([]byte(v), &strct.RoleName); err != nil {
                return err
             }
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
    // check if AccountName (a required property) was received
    if !AccountNameReceived {
        return errors.New("\"AccountName\" is required but was not present")
    }
    // check if Email (a required property) was received
    if !EmailReceived {
        return errors.New("\"Email\" is required but was not present")
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
