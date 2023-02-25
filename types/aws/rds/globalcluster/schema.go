// Code generated by schema-generate. DO NOT EDIT.

package globalcluster

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// Resource Resource Type definition for AWS::RDS::GlobalCluster
type Resource struct {

  // The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.
  DeletionProtection bool `json:"DeletionProtection,omitempty"`

  // The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).
  // If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
  Engine string `json:"Engine,omitempty"`

  // The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
  EngineVersion string `json:"EngineVersion,omitempty"`

  // The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.
  GlobalClusterIdentifier string `json:"GlobalClusterIdentifier,omitempty"`

  // The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.
  SourceDBClusterIdentifier string `json:"SourceDBClusterIdentifier,omitempty"`

  //  The storage encryption setting for the new global database cluster.
  // If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
  StorageEncrypted bool `json:"StorageEncrypted,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DeletionProtection" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DeletionProtection\": ")
	if tmp, err := json.Marshal(strct.DeletionProtection); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Engine" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Engine\": ")
	if tmp, err := json.Marshal(strct.Engine); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EngineVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EngineVersion\": ")
	if tmp, err := json.Marshal(strct.EngineVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GlobalClusterIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalClusterIdentifier\": ")
	if tmp, err := json.Marshal(strct.GlobalClusterIdentifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SourceDBClusterIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceDBClusterIdentifier\": ")
	if tmp, err := json.Marshal(strct.SourceDBClusterIdentifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StorageEncrypted" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StorageEncrypted\": ")
	if tmp, err := json.Marshal(strct.StorageEncrypted); err != nil {
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
        case "DeletionProtection":
            if err := json.Unmarshal([]byte(v), &strct.DeletionProtection); err != nil {
                return err
             }
        case "Engine":
            if err := json.Unmarshal([]byte(v), &strct.Engine); err != nil {
                return err
             }
        case "EngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.EngineVersion); err != nil {
                return err
             }
        case "GlobalClusterIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.GlobalClusterIdentifier); err != nil {
                return err
             }
        case "SourceDBClusterIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.SourceDBClusterIdentifier); err != nil {
                return err
             }
        case "StorageEncrypted":
            if err := json.Unmarshal([]byte(v), &strct.StorageEncrypted); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}