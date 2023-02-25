// Code generated by schema-generate. DO NOT EDIT.

package subnetgroup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource Type definition for AWS::ElastiCache::SubnetGroup
type Resource struct {

  // The name for the cache subnet group. This value is stored as a lowercase string.
  CacheSubnetGroupName string `json:"CacheSubnetGroupName,omitempty"`

  // The description for the cache subnet group.
  Description string `json:"Description"`

  // The EC2 subnet IDs for the cache subnet group.
  SubnetIds []string `json:"SubnetIds"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A tag that can be added to an ElastiCache subnet group. Tags are composed of a Key/Value pair. You can use tags to categorize and track all your subnet groups. A tag with a null Value is permitted.
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CacheSubnetGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CacheSubnetGroupName\": ")
	if tmp, err := json.Marshal(strct.CacheSubnetGroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Description" field is required
    // only required object types supported for marshal checking (for now)
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
    // "SubnetIds" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SubnetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetIds\": ")
	if tmp, err := json.Marshal(strct.SubnetIds); err != nil {
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
    DescriptionReceived := false
    SubnetIdsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CacheSubnetGroupName":
            if err := json.Unmarshal([]byte(v), &strct.CacheSubnetGroupName); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
            DescriptionReceived = true
        case "SubnetIds":
            if err := json.Unmarshal([]byte(v), &strct.SubnetIds); err != nil {
                return err
             }
            SubnetIdsReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Description (a required property) was received
    if !DescriptionReceived {
        return errors.New("\"Description\" is required but was not present")
    }
    // check if SubnetIds (a required property) was received
    if !SubnetIdsReceived {
        return errors.New("\"SubnetIds\" is required but was not present")
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