// Code generated by schema-generate. DO NOT EDIT.

package globalreplicationgroup

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// GlobalReplicationGroupMember 
type GlobalReplicationGroupMember struct {

  // Regionally unique identifier for the member i.e. ReplicationGroupId.
  ReplicationGroupId string `json:"ReplicationGroupId,omitempty"`

  // The AWS region of the Global Datastore member.
  ReplicationGroupRegion string `json:"ReplicationGroupRegion,omitempty"`

  // Indicates the role of the member, primary or secondary.
  Role string `json:"Role,omitempty"`
}

// RegionalConfiguration 
type RegionalConfiguration struct {

  // The replication group id of the Global Datastore member.
  ReplicationGroupId string `json:"ReplicationGroupId,omitempty"`

  // The AWS region of the Global Datastore member.
  ReplicationGroupRegion string `json:"ReplicationGroupRegion,omitempty"`

  // A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. 
  ReshardingConfigurations []*ReshardingConfiguration `json:"ReshardingConfigurations,omitempty"`
}

// ReshardingConfiguration 
type ReshardingConfiguration struct {

  // Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.
  NodeGroupId string `json:"NodeGroupId,omitempty"`

  // A list of preferred availability zones for the nodes of new node groups.
  PreferredAvailabilityZones []string `json:"PreferredAvailabilityZones,omitempty"`
}

// Resource The AWS::ElastiCache::GlobalReplicationGroup resource creates an Amazon ElastiCache Global Replication Group.
type Resource struct {

  // AutomaticFailoverEnabled
  AutomaticFailoverEnabled bool `json:"AutomaticFailoverEnabled,omitempty"`

  // The cache node type of the Global Datastore
  CacheNodeType string `json:"CacheNodeType,omitempty"`

  // Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.
  CacheParameterGroupName string `json:"CacheParameterGroupName,omitempty"`

  // The engine version of the Global Datastore.
  EngineVersion string `json:"EngineVersion,omitempty"`

  // Indicates the number of node groups in the Global Datastore.
  GlobalNodeGroupCount int `json:"GlobalNodeGroupCount,omitempty"`

  // The optional description of the Global Datastore
  GlobalReplicationGroupDescription string `json:"GlobalReplicationGroupDescription,omitempty"`

  // The name of the Global Datastore, it is generated by ElastiCache adding a prefix to GlobalReplicationGroupIdSuffix.
  GlobalReplicationGroupId string `json:"GlobalReplicationGroupId,omitempty"`

  // The suffix name of a Global Datastore. Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix. 
  GlobalReplicationGroupIdSuffix string `json:"GlobalReplicationGroupIdSuffix,omitempty"`

  // The replication groups that comprise the Global Datastore.
  Members []*GlobalReplicationGroupMember `json:"Members"`

  // Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore 
  RegionalConfigurations []*RegionalConfiguration `json:"RegionalConfigurations,omitempty"`

  // The status of the Global Datastore
  Status string `json:"Status,omitempty"`
}

func (strct *GlobalReplicationGroupMember) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ReplicationGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplicationGroupId\": ")
	if tmp, err := json.Marshal(strct.ReplicationGroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ReplicationGroupRegion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplicationGroupRegion\": ")
	if tmp, err := json.Marshal(strct.ReplicationGroupRegion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Role" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Role\": ")
	if tmp, err := json.Marshal(strct.Role); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GlobalReplicationGroupMember) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ReplicationGroupId":
            if err := json.Unmarshal([]byte(v), &strct.ReplicationGroupId); err != nil {
                return err
             }
        case "ReplicationGroupRegion":
            if err := json.Unmarshal([]byte(v), &strct.ReplicationGroupRegion); err != nil {
                return err
             }
        case "Role":
            if err := json.Unmarshal([]byte(v), &strct.Role); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *RegionalConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ReplicationGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplicationGroupId\": ")
	if tmp, err := json.Marshal(strct.ReplicationGroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ReplicationGroupRegion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReplicationGroupRegion\": ")
	if tmp, err := json.Marshal(strct.ReplicationGroupRegion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ReshardingConfigurations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ReshardingConfigurations\": ")
	if tmp, err := json.Marshal(strct.ReshardingConfigurations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RegionalConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ReplicationGroupId":
            if err := json.Unmarshal([]byte(v), &strct.ReplicationGroupId); err != nil {
                return err
             }
        case "ReplicationGroupRegion":
            if err := json.Unmarshal([]byte(v), &strct.ReplicationGroupRegion); err != nil {
                return err
             }
        case "ReshardingConfigurations":
            if err := json.Unmarshal([]byte(v), &strct.ReshardingConfigurations); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ReshardingConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "NodeGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NodeGroupId\": ")
	if tmp, err := json.Marshal(strct.NodeGroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PreferredAvailabilityZones" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PreferredAvailabilityZones\": ")
	if tmp, err := json.Marshal(strct.PreferredAvailabilityZones); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReshardingConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "NodeGroupId":
            if err := json.Unmarshal([]byte(v), &strct.NodeGroupId); err != nil {
                return err
             }
        case "PreferredAvailabilityZones":
            if err := json.Unmarshal([]byte(v), &strct.PreferredAvailabilityZones); err != nil {
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
    // Marshal the "AutomaticFailoverEnabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutomaticFailoverEnabled\": ")
	if tmp, err := json.Marshal(strct.AutomaticFailoverEnabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CacheNodeType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CacheNodeType\": ")
	if tmp, err := json.Marshal(strct.CacheNodeType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CacheParameterGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CacheParameterGroupName\": ")
	if tmp, err := json.Marshal(strct.CacheParameterGroupName); err != nil {
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
    // Marshal the "GlobalNodeGroupCount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalNodeGroupCount\": ")
	if tmp, err := json.Marshal(strct.GlobalNodeGroupCount); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GlobalReplicationGroupDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalReplicationGroupDescription\": ")
	if tmp, err := json.Marshal(strct.GlobalReplicationGroupDescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GlobalReplicationGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalReplicationGroupId\": ")
	if tmp, err := json.Marshal(strct.GlobalReplicationGroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GlobalReplicationGroupIdSuffix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalReplicationGroupIdSuffix\": ")
	if tmp, err := json.Marshal(strct.GlobalReplicationGroupIdSuffix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Members" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Members" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Members\": ")
	if tmp, err := json.Marshal(strct.Members); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegionalConfigurations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionalConfigurations\": ")
	if tmp, err := json.Marshal(strct.RegionalConfigurations); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    MembersReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AutomaticFailoverEnabled":
            if err := json.Unmarshal([]byte(v), &strct.AutomaticFailoverEnabled); err != nil {
                return err
             }
        case "CacheNodeType":
            if err := json.Unmarshal([]byte(v), &strct.CacheNodeType); err != nil {
                return err
             }
        case "CacheParameterGroupName":
            if err := json.Unmarshal([]byte(v), &strct.CacheParameterGroupName); err != nil {
                return err
             }
        case "EngineVersion":
            if err := json.Unmarshal([]byte(v), &strct.EngineVersion); err != nil {
                return err
             }
        case "GlobalNodeGroupCount":
            if err := json.Unmarshal([]byte(v), &strct.GlobalNodeGroupCount); err != nil {
                return err
             }
        case "GlobalReplicationGroupDescription":
            if err := json.Unmarshal([]byte(v), &strct.GlobalReplicationGroupDescription); err != nil {
                return err
             }
        case "GlobalReplicationGroupId":
            if err := json.Unmarshal([]byte(v), &strct.GlobalReplicationGroupId); err != nil {
                return err
             }
        case "GlobalReplicationGroupIdSuffix":
            if err := json.Unmarshal([]byte(v), &strct.GlobalReplicationGroupIdSuffix); err != nil {
                return err
             }
        case "Members":
            if err := json.Unmarshal([]byte(v), &strct.Members); err != nil {
                return err
             }
            MembersReceived = true
        case "RegionalConfigurations":
            if err := json.Unmarshal([]byte(v), &strct.RegionalConfigurations); err != nil {
                return err
             }
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Members (a required property) was received
    if !MembersReceived {
        return errors.New("\"Members\" is required but was not present")
    }
    return nil
}