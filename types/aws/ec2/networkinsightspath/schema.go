// Code generated by schema-generate. DO NOT EDIT.

package networkinsightspath

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource Resource schema for AWS::EC2::NetworkInsightsPath
type Resource struct {
  CreatedDate string `json:"CreatedDate,omitempty"`
  Destination string `json:"Destination,omitempty"`
  DestinationArn string `json:"DestinationArn,omitempty"`
  DestinationIp string `json:"DestinationIp,omitempty"`
  DestinationPort int `json:"DestinationPort,omitempty"`
  NetworkInsightsPathArn string `json:"NetworkInsightsPathArn,omitempty"`
  NetworkInsightsPathId string `json:"NetworkInsightsPathId,omitempty"`
  Protocol string `json:"Protocol"`
  Source string `json:"Source"`
  SourceArn string `json:"SourceArn,omitempty"`
  SourceIp string `json:"SourceIp,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CreatedDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreatedDate\": ")
	if tmp, err := json.Marshal(strct.CreatedDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Destination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Destination\": ")
	if tmp, err := json.Marshal(strct.Destination); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DestinationArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DestinationArn\": ")
	if tmp, err := json.Marshal(strct.DestinationArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DestinationIp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DestinationIp\": ")
	if tmp, err := json.Marshal(strct.DestinationIp); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DestinationPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DestinationPort\": ")
	if tmp, err := json.Marshal(strct.DestinationPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NetworkInsightsPathArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NetworkInsightsPathArn\": ")
	if tmp, err := json.Marshal(strct.NetworkInsightsPathArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "NetworkInsightsPathId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NetworkInsightsPathId\": ")
	if tmp, err := json.Marshal(strct.NetworkInsightsPathId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Protocol" field is required
    // only required object types supported for marshal checking (for now)
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
    // "Source" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Source" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SourceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceArn\": ")
	if tmp, err := json.Marshal(strct.SourceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SourceIp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceIp\": ")
	if tmp, err := json.Marshal(strct.SourceIp); err != nil {
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
    ProtocolReceived := false
    SourceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CreatedDate":
            if err := json.Unmarshal([]byte(v), &strct.CreatedDate); err != nil {
                return err
             }
        case "Destination":
            if err := json.Unmarshal([]byte(v), &strct.Destination); err != nil {
                return err
             }
        case "DestinationArn":
            if err := json.Unmarshal([]byte(v), &strct.DestinationArn); err != nil {
                return err
             }
        case "DestinationIp":
            if err := json.Unmarshal([]byte(v), &strct.DestinationIp); err != nil {
                return err
             }
        case "DestinationPort":
            if err := json.Unmarshal([]byte(v), &strct.DestinationPort); err != nil {
                return err
             }
        case "NetworkInsightsPathArn":
            if err := json.Unmarshal([]byte(v), &strct.NetworkInsightsPathArn); err != nil {
                return err
             }
        case "NetworkInsightsPathId":
            if err := json.Unmarshal([]byte(v), &strct.NetworkInsightsPathId); err != nil {
                return err
             }
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
            ProtocolReceived = true
        case "Source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
            SourceReceived = true
        case "SourceArn":
            if err := json.Unmarshal([]byte(v), &strct.SourceArn); err != nil {
                return err
             }
        case "SourceIp":
            if err := json.Unmarshal([]byte(v), &strct.SourceIp); err != nil {
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
    // check if Protocol (a required property) was received
    if !ProtocolReceived {
        return errors.New("\"Protocol\" is required but was not present")
    }
    // check if Source (a required property) was received
    if !SourceReceived {
        return errors.New("\"Source\" is required but was not present")
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
