// Code generated by schema-generate. DO NOT EDIT.

package locationfsxontap

import (
    "encoding/json"
    "fmt"
    "bytes"
    "errors"
)

// NFS NFS protocol configuration for FSx ONTAP file system.
type NFS struct {
  MountOptions *NfsMountOptions `json:"MountOptions"`
}

// NfsMountOptions The NFS mount options that DataSync can use to mount your NFS share.
type NfsMountOptions struct {

  // The specific NFS version that you want DataSync to use to mount your NFS share.
  Version string `json:"Version,omitempty"`
}

// Protocol Configuration settings for NFS or SMB protocol.
type Protocol struct {
  NFS *NFS `json:"NFS,omitempty"`
  SMB *SMB `json:"SMB,omitempty"`
}

// Resource Resource schema for AWS::DataSync::LocationFSxONTAP.
type Resource struct {

  // The Amazon Resource Name (ARN) for the FSx ONAP file system.
  FsxFilesystemArn string `json:"FsxFilesystemArn,omitempty"`

  // The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.
  LocationArn string `json:"LocationArn,omitempty"`

  // The URL of the FSx ONTAP file system that was described.
  LocationUri string `json:"LocationUri,omitempty"`
  Protocol *Protocol `json:"Protocol,omitempty"`

  // The ARNs of the security groups that are to use to configure the FSx ONTAP file system.
  SecurityGroupArns []string `json:"SecurityGroupArns"`

  // The Amazon Resource Name (ARN) for the FSx ONTAP SVM.
  StorageVirtualMachineArn string `json:"StorageVirtualMachineArn"`

  // A subdirectory in the location's path.
  Subdirectory string `json:"Subdirectory,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`
}

// SMB SMB protocol configuration for FSx ONTAP file system.
type SMB struct {

  // The name of the Windows domain that the SMB server belongs to.
  Domain string `json:"Domain,omitempty"`
  MountOptions *SmbMountOptions `json:"MountOptions"`

  // The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
  Password string `json:"Password"`

  // The user who can mount the share, has the permissions to access files and folders in the SMB share.
  User string `json:"User"`
}

// SmbMountOptions The mount options used by DataSync to access the SMB server.
type SmbMountOptions struct {

  // The specific SMB version that you want DataSync to use to mount your SMB share.
  Version string `json:"Version,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key for an AWS resource tag.
  Key string `json:"Key"`

  // The value for an AWS resource tag.
  Value string `json:"Value"`
}

func (strct *NFS) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "MountOptions" field is required
    if strct.MountOptions == nil {
        return nil, errors.New("MountOptions is a required field")
    }
    // Marshal the "MountOptions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MountOptions\": ")
	if tmp, err := json.Marshal(strct.MountOptions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *NFS) UnmarshalJSON(b []byte) error {
    MountOptionsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "MountOptions":
            if err := json.Unmarshal([]byte(v), &strct.MountOptions); err != nil {
                return err
             }
            MountOptionsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MountOptions (a required property) was received
    if !MountOptionsReceived {
        return errors.New("\"MountOptions\" is required but was not present")
    }
    return nil
}

func (strct *NfsMountOptions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Version" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *NfsMountOptions) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Version":
            if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Protocol) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "NFS" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NFS\": ")
	if tmp, err := json.Marshal(strct.NFS); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SMB" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SMB\": ")
	if tmp, err := json.Marshal(strct.SMB); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Protocol) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "NFS":
            if err := json.Unmarshal([]byte(v), &strct.NFS); err != nil {
                return err
             }
        case "SMB":
            if err := json.Unmarshal([]byte(v), &strct.SMB); err != nil {
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
    // Marshal the "FsxFilesystemArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FsxFilesystemArn\": ")
	if tmp, err := json.Marshal(strct.FsxFilesystemArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LocationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LocationArn\": ")
	if tmp, err := json.Marshal(strct.LocationArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LocationUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LocationUri\": ")
	if tmp, err := json.Marshal(strct.LocationUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Protocol" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Protocol\": ")
	if tmp, err := json.Marshal(strct.Protocol); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SecurityGroupArns" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SecurityGroupArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityGroupArns\": ")
	if tmp, err := json.Marshal(strct.SecurityGroupArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "StorageVirtualMachineArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "StorageVirtualMachineArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StorageVirtualMachineArn\": ")
	if tmp, err := json.Marshal(strct.StorageVirtualMachineArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Subdirectory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Subdirectory\": ")
	if tmp, err := json.Marshal(strct.Subdirectory); err != nil {
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
    SecurityGroupArnsReceived := false
    StorageVirtualMachineArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "FsxFilesystemArn":
            if err := json.Unmarshal([]byte(v), &strct.FsxFilesystemArn); err != nil {
                return err
             }
        case "LocationArn":
            if err := json.Unmarshal([]byte(v), &strct.LocationArn); err != nil {
                return err
             }
        case "LocationUri":
            if err := json.Unmarshal([]byte(v), &strct.LocationUri); err != nil {
                return err
             }
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
        case "SecurityGroupArns":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupArns); err != nil {
                return err
             }
            SecurityGroupArnsReceived = true
        case "StorageVirtualMachineArn":
            if err := json.Unmarshal([]byte(v), &strct.StorageVirtualMachineArn); err != nil {
                return err
             }
            StorageVirtualMachineArnReceived = true
        case "Subdirectory":
            if err := json.Unmarshal([]byte(v), &strct.Subdirectory); err != nil {
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
    // check if SecurityGroupArns (a required property) was received
    if !SecurityGroupArnsReceived {
        return errors.New("\"SecurityGroupArns\" is required but was not present")
    }
    // check if StorageVirtualMachineArn (a required property) was received
    if !StorageVirtualMachineArnReceived {
        return errors.New("\"StorageVirtualMachineArn\" is required but was not present")
    }
    return nil
}

func (strct *SMB) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Domain" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Domain\": ")
	if tmp, err := json.Marshal(strct.Domain); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "MountOptions" field is required
    if strct.MountOptions == nil {
        return nil, errors.New("MountOptions is a required field")
    }
    // Marshal the "MountOptions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MountOptions\": ")
	if tmp, err := json.Marshal(strct.MountOptions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Password" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Password" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Password\": ")
	if tmp, err := json.Marshal(strct.Password); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "User" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "User" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"User\": ")
	if tmp, err := json.Marshal(strct.User); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SMB) UnmarshalJSON(b []byte) error {
    MountOptionsReceived := false
    PasswordReceived := false
    UserReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Domain":
            if err := json.Unmarshal([]byte(v), &strct.Domain); err != nil {
                return err
             }
        case "MountOptions":
            if err := json.Unmarshal([]byte(v), &strct.MountOptions); err != nil {
                return err
             }
            MountOptionsReceived = true
        case "Password":
            if err := json.Unmarshal([]byte(v), &strct.Password); err != nil {
                return err
             }
            PasswordReceived = true
        case "User":
            if err := json.Unmarshal([]byte(v), &strct.User); err != nil {
                return err
             }
            UserReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if MountOptions (a required property) was received
    if !MountOptionsReceived {
        return errors.New("\"MountOptions\" is required but was not present")
    }
    // check if Password (a required property) was received
    if !PasswordReceived {
        return errors.New("\"Password\" is required but was not present")
    }
    // check if User (a required property) was received
    if !UserReceived {
        return errors.New("\"User\" is required but was not present")
    }
    return nil
}

func (strct *SmbMountOptions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Version" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SmbMountOptions) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Version":
            if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
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
