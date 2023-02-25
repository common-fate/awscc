// Code generated by schema-generate. DO NOT EDIT.

package vdmattributes

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// DashboardAttributes Preferences regarding the Dashboard feature.
type DashboardAttributes struct {

  // Whether emails sent from this account have engagement tracking enabled.
  EngagementMetrics string `json:"EngagementMetrics,omitempty"`
}

// GuardianAttributes Preferences regarding the Guardian feature.
type GuardianAttributes struct {

  // Whether emails sent from this account have optimized delivery algorithm enabled.
  OptimizedSharedDelivery string `json:"OptimizedSharedDelivery,omitempty"`
}

// Resource Resource Type definition for AWS::SES::VdmAttributes
type Resource struct {
  DashboardAttributes *DashboardAttributes `json:"DashboardAttributes,omitempty"`
  GuardianAttributes *GuardianAttributes `json:"GuardianAttributes,omitempty"`

  // Unique identifier for this resource
  VdmAttributesResourceId string `json:"VdmAttributesResourceId,omitempty"`
}

func (strct *DashboardAttributes) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EngagementMetrics" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EngagementMetrics\": ")
	if tmp, err := json.Marshal(strct.EngagementMetrics); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DashboardAttributes) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EngagementMetrics":
            if err := json.Unmarshal([]byte(v), &strct.EngagementMetrics); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *GuardianAttributes) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "OptimizedSharedDelivery" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OptimizedSharedDelivery\": ")
	if tmp, err := json.Marshal(strct.OptimizedSharedDelivery); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GuardianAttributes) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "OptimizedSharedDelivery":
            if err := json.Unmarshal([]byte(v), &strct.OptimizedSharedDelivery); err != nil {
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
    // Marshal the "DashboardAttributes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DashboardAttributes\": ")
	if tmp, err := json.Marshal(strct.DashboardAttributes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GuardianAttributes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GuardianAttributes\": ")
	if tmp, err := json.Marshal(strct.GuardianAttributes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "VdmAttributesResourceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"VdmAttributesResourceId\": ")
	if tmp, err := json.Marshal(strct.VdmAttributesResourceId); err != nil {
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
        case "DashboardAttributes":
            if err := json.Unmarshal([]byte(v), &strct.DashboardAttributes); err != nil {
                return err
             }
        case "GuardianAttributes":
            if err := json.Unmarshal([]byte(v), &strct.GuardianAttributes); err != nil {
                return err
             }
        case "VdmAttributesResourceId":
            if err := json.Unmarshal([]byte(v), &strct.VdmAttributesResourceId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}