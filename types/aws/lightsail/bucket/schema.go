// Code generated by schema-generate. DO NOT EDIT.

package bucket

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AccessRules An object that sets the public accessibility of objects in the specified bucket.
type AccessRules struct {

  // A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.
  AllowPublicOverrides bool `json:"AllowPublicOverrides,omitempty"`

  // Specifies the anonymous access to all objects in a bucket.
  GetObject string `json:"GetObject,omitempty"`
}

// Resource Resource Type definition for AWS::Lightsail::Bucket
type Resource struct {

  // Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.
  AbleToUpdateBundle bool `json:"AbleToUpdateBundle,omitempty"`
  AccessRules *AccessRules `json:"AccessRules,omitempty"`
  BucketArn string `json:"BucketArn,omitempty"`

  // The name for the bucket.
  BucketName string `json:"BucketName"`

  // The ID of the bundle to use for the bucket.
  BundleId string `json:"BundleId"`

  // Specifies whether to enable or disable versioning of objects in the bucket.
  ObjectVersioning bool `json:"ObjectVersioning,omitempty"`

  // An array of strings to specify the AWS account IDs that can access the bucket.
  ReadOnlyAccessAccounts []string `json:"ReadOnlyAccessAccounts,omitempty"`

  // The names of the Lightsail resources for which to set bucket access.
  ResourcesReceivingAccess []string `json:"ResourcesReceivingAccess,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // The URL of the bucket.
  Url string `json:"Url,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value,omitempty"`
}

func (strct *AccessRules) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AllowPublicOverrides" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowPublicOverrides\": ")
	if tmp, err := json.Marshal(strct.AllowPublicOverrides); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GetObject" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GetObject\": ")
	if tmp, err := json.Marshal(strct.GetObject); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AccessRules) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AllowPublicOverrides":
            if err := json.Unmarshal([]byte(v), &strct.AllowPublicOverrides); err != nil {
                return err
             }
        case "GetObject":
            if err := json.Unmarshal([]byte(v), &strct.GetObject); err != nil {
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
    // Marshal the "AbleToUpdateBundle" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AbleToUpdateBundle\": ")
	if tmp, err := json.Marshal(strct.AbleToUpdateBundle); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AccessRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessRules\": ")
	if tmp, err := json.Marshal(strct.AccessRules); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BucketArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BucketArn\": ")
	if tmp, err := json.Marshal(strct.BucketArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "BucketName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BucketName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BucketName\": ")
	if tmp, err := json.Marshal(strct.BucketName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "BundleId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BundleId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BundleId\": ")
	if tmp, err := json.Marshal(strct.BundleId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ObjectVersioning" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ObjectVersioning\": ")
	if tmp, err := json.Marshal(strct.ObjectVersioning); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ReadOnlyAccessAccounts" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReadOnlyAccessAccounts\": ")
	if tmp, err := json.Marshal(strct.ReadOnlyAccessAccounts); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourcesReceivingAccess" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourcesReceivingAccess\": ")
	if tmp, err := json.Marshal(strct.ResourcesReceivingAccess); err != nil {
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
    // Marshal the "Url" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
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
    BucketNameReceived := false
    BundleIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AbleToUpdateBundle":
            if err := json.Unmarshal([]byte(v), &strct.AbleToUpdateBundle); err != nil {
                return err
             }
        case "AccessRules":
            if err := json.Unmarshal([]byte(v), &strct.AccessRules); err != nil {
                return err
             }
        case "BucketArn":
            if err := json.Unmarshal([]byte(v), &strct.BucketArn); err != nil {
                return err
             }
        case "BucketName":
            if err := json.Unmarshal([]byte(v), &strct.BucketName); err != nil {
                return err
             }
            BucketNameReceived = true
        case "BundleId":
            if err := json.Unmarshal([]byte(v), &strct.BundleId); err != nil {
                return err
             }
            BundleIdReceived = true
        case "ObjectVersioning":
            if err := json.Unmarshal([]byte(v), &strct.ObjectVersioning); err != nil {
                return err
             }
        case "ReadOnlyAccessAccounts":
            if err := json.Unmarshal([]byte(v), &strct.ReadOnlyAccessAccounts); err != nil {
                return err
             }
        case "ResourcesReceivingAccess":
            if err := json.Unmarshal([]byte(v), &strct.ResourcesReceivingAccess); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if BucketName (a required property) was received
    if !BucketNameReceived {
        return errors.New("\"BucketName\" is required but was not present")
    }
    // check if BundleId (a required property) was received
    if !BundleIdReceived {
        return errors.New("\"BundleId\" is required but was not present")
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
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Key (a required property) was received
    if !KeyReceived {
        return errors.New("\"Key\" is required but was not present")
    }
    return nil
}
