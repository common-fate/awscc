// Code generated by schema-generate. DO NOT EDIT.

package site

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Location The location of the site
type Location struct {

  // The physical address.
  Address string `json:"Address,omitempty"`

  // The latitude.
  Latitude string `json:"Latitude,omitempty"`

  // The longitude.
  Longitude string `json:"Longitude,omitempty"`
}

// Resource The AWS::NetworkManager::Site type describes a site.
type Resource struct {

  // The description of the site.
  Description string `json:"Description,omitempty"`

  // The ID of the global network.
  GlobalNetworkId string `json:"GlobalNetworkId"`

  // The location of the site.
  Location *Location `json:"Location,omitempty"`

  // The Amazon Resource Name (ARN) of the site.
  SiteArn string `json:"SiteArn,omitempty"`

  // The ID of the site.
  SiteId string `json:"SiteId,omitempty"`

  // The tags for the site.
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag A key-value pair to associate with a site resource.
type Tag struct {
  Key string `json:"Key,omitempty"`
  Value string `json:"Value,omitempty"`
}

func (strct *Location) MarshalJSON() ([]byte, error) {
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
    // Marshal the "Latitude" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Latitude\": ")
	if tmp, err := json.Marshal(strct.Latitude); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Longitude" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Longitude\": ")
	if tmp, err := json.Marshal(strct.Longitude); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Location) UnmarshalJSON(b []byte) error {
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
        case "Latitude":
            if err := json.Unmarshal([]byte(v), &strct.Latitude); err != nil {
                return err
             }
        case "Longitude":
            if err := json.Unmarshal([]byte(v), &strct.Longitude); err != nil {
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
    // "GlobalNetworkId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "GlobalNetworkId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GlobalNetworkId\": ")
	if tmp, err := json.Marshal(strct.GlobalNetworkId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SiteArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SiteArn\": ")
	if tmp, err := json.Marshal(strct.SiteArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SiteId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SiteId\": ")
	if tmp, err := json.Marshal(strct.SiteId); err != nil {
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
    GlobalNetworkIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "GlobalNetworkId":
            if err := json.Unmarshal([]byte(v), &strct.GlobalNetworkId); err != nil {
                return err
             }
            GlobalNetworkIdReceived = true
        case "Location":
            if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
                return err
             }
        case "SiteArn":
            if err := json.Unmarshal([]byte(v), &strct.SiteArn); err != nil {
                return err
             }
        case "SiteId":
            if err := json.Unmarshal([]byte(v), &strct.SiteId); err != nil {
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
    // check if GlobalNetworkId (a required property) was received
    if !GlobalNetworkIdReceived {
        return errors.New("\"GlobalNetworkId\" is required but was not present")
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
