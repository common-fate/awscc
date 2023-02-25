// Code generated by schema-generate. DO NOT EDIT.

package datarepositoryassociation

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// AutoExportPolicy Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.
type AutoExportPolicy struct {
  Events []string `json:"Events"`
}

// AutoImportPolicy Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.
type AutoImportPolicy struct {
  Events []string `json:"Events"`
}

// Resource Resource Type definition for AWS::FSx::DataRepositoryAssociation
type Resource struct {

  // The system-generated, unique ID of the data repository association.
  AssociationId string `json:"AssociationId,omitempty"`

  // A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.
  BatchImportMetaDataOnCreate bool `json:"BatchImportMetaDataOnCreate,omitempty"`

  // The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.
  DataRepositoryPath string `json:"DataRepositoryPath"`

  // The globally unique ID of the file system, assigned by Amazon FSx.
  FileSystemId string `json:"FileSystemId"`

  // This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
  FileSystemPath string `json:"FileSystemPath"`

  // For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
  ImportedFileChunkSize int `json:"ImportedFileChunkSize,omitempty"`

  // The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.
  ResourceARN string `json:"ResourceARN,omitempty"`

  // The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
  S3 *S3 `json:"S3,omitempty"`

  // A list of Tag values, with a maximum of 50 elements.
  Tags []*Tag `json:"Tags,omitempty"`
}

// S3 The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
type S3 struct {
  AutoExportPolicy *AutoExportPolicy `json:"AutoExportPolicy,omitempty"`
  AutoImportPolicy *AutoImportPolicy `json:"AutoImportPolicy,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
  Value string `json:"Value"`
}

func (strct *AutoExportPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Events" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Events" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Events\": ")
	if tmp, err := json.Marshal(strct.Events); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AutoExportPolicy) UnmarshalJSON(b []byte) error {
    EventsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Events":
            if err := json.Unmarshal([]byte(v), &strct.Events); err != nil {
                return err
             }
            EventsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Events (a required property) was received
    if !EventsReceived {
        return errors.New("\"Events\" is required but was not present")
    }
    return nil
}

func (strct *AutoImportPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Events" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Events" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Events\": ")
	if tmp, err := json.Marshal(strct.Events); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AutoImportPolicy) UnmarshalJSON(b []byte) error {
    EventsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Events":
            if err := json.Unmarshal([]byte(v), &strct.Events); err != nil {
                return err
             }
            EventsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Events (a required property) was received
    if !EventsReceived {
        return errors.New("\"Events\" is required but was not present")
    }
    return nil
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AssociationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssociationId\": ")
	if tmp, err := json.Marshal(strct.AssociationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BatchImportMetaDataOnCreate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BatchImportMetaDataOnCreate\": ")
	if tmp, err := json.Marshal(strct.BatchImportMetaDataOnCreate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DataRepositoryPath" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DataRepositoryPath" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataRepositoryPath\": ")
	if tmp, err := json.Marshal(strct.DataRepositoryPath); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FileSystemId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FileSystemId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FileSystemId\": ")
	if tmp, err := json.Marshal(strct.FileSystemId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "FileSystemPath" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "FileSystemPath" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FileSystemPath\": ")
	if tmp, err := json.Marshal(strct.FileSystemPath); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImportedFileChunkSize" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImportedFileChunkSize\": ")
	if tmp, err := json.Marshal(strct.ImportedFileChunkSize); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceARN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceARN\": ")
	if tmp, err := json.Marshal(strct.ResourceARN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "S3" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"S3\": ")
	if tmp, err := json.Marshal(strct.S3); err != nil {
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
    DataRepositoryPathReceived := false
    FileSystemIdReceived := false
    FileSystemPathReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AssociationId":
            if err := json.Unmarshal([]byte(v), &strct.AssociationId); err != nil {
                return err
             }
        case "BatchImportMetaDataOnCreate":
            if err := json.Unmarshal([]byte(v), &strct.BatchImportMetaDataOnCreate); err != nil {
                return err
             }
        case "DataRepositoryPath":
            if err := json.Unmarshal([]byte(v), &strct.DataRepositoryPath); err != nil {
                return err
             }
            DataRepositoryPathReceived = true
        case "FileSystemId":
            if err := json.Unmarshal([]byte(v), &strct.FileSystemId); err != nil {
                return err
             }
            FileSystemIdReceived = true
        case "FileSystemPath":
            if err := json.Unmarshal([]byte(v), &strct.FileSystemPath); err != nil {
                return err
             }
            FileSystemPathReceived = true
        case "ImportedFileChunkSize":
            if err := json.Unmarshal([]byte(v), &strct.ImportedFileChunkSize); err != nil {
                return err
             }
        case "ResourceARN":
            if err := json.Unmarshal([]byte(v), &strct.ResourceARN); err != nil {
                return err
             }
        case "S3":
            if err := json.Unmarshal([]byte(v), &strct.S3); err != nil {
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
    // check if DataRepositoryPath (a required property) was received
    if !DataRepositoryPathReceived {
        return errors.New("\"DataRepositoryPath\" is required but was not present")
    }
    // check if FileSystemId (a required property) was received
    if !FileSystemIdReceived {
        return errors.New("\"FileSystemId\" is required but was not present")
    }
    // check if FileSystemPath (a required property) was received
    if !FileSystemPathReceived {
        return errors.New("\"FileSystemPath\" is required but was not present")
    }
    return nil
}

func (strct *S3) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AutoExportPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoExportPolicy\": ")
	if tmp, err := json.Marshal(strct.AutoExportPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AutoImportPolicy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoImportPolicy\": ")
	if tmp, err := json.Marshal(strct.AutoImportPolicy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *S3) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AutoExportPolicy":
            if err := json.Unmarshal([]byte(v), &strct.AutoExportPolicy); err != nil {
                return err
             }
        case "AutoImportPolicy":
            if err := json.Unmarshal([]byte(v), &strct.AutoImportPolicy); err != nil {
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