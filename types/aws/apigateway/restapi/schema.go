// Code generated by schema-generate. DO NOT EDIT.

package restapi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Body_object 
type Body_object struct {
}

// EndpointConfiguration 
type EndpointConfiguration struct {
  Types []string `json:"Types,omitempty"`
  VpcEndpointIds []string `json:"VpcEndpointIds,omitempty"`
}

// Parameters_object 
type Parameters_object struct {
}

// Policy_object 
type Policy_object struct {
}

// Resource Resource Type definition for AWS::ApiGateway::RestApi.
type Resource struct {
  ApiKeySourceType string `json:"ApiKeySourceType,omitempty"`
  BinaryMediaTypes []string `json:"BinaryMediaTypes,omitempty"`
  Body interface{} `json:"Body,omitempty"`
  BodyS3Location *S3Location `json:"BodyS3Location,omitempty"`
  CloneFrom string `json:"CloneFrom,omitempty"`
  Description string `json:"Description,omitempty"`
  DisableExecuteApiEndpoint bool `json:"DisableExecuteApiEndpoint,omitempty"`
  EndpointConfiguration *EndpointConfiguration `json:"EndpointConfiguration,omitempty"`
  FailOnWarnings bool `json:"FailOnWarnings,omitempty"`
  MinimumCompressionSize int `json:"MinimumCompressionSize,omitempty"`
  Mode string `json:"Mode,omitempty"`
  Name string `json:"Name,omitempty"`
  Parameters interface{} `json:"Parameters,omitempty"`
  Policy interface{} `json:"Policy,omitempty"`
  RestApiId string `json:"RestApiId,omitempty"`
  RootResourceId string `json:"RootResourceId,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// S3Location 
type S3Location struct {
  Bucket string `json:"Bucket,omitempty"`
  ETag string `json:"ETag,omitempty"`
  Key string `json:"Key,omitempty"`
  Version string `json:"Version,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *EndpointConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Types" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Types\": ")
	if tmp, err := json.Marshal(strct.Types); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "VpcEndpointIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"VpcEndpointIds\": ")
	if tmp, err := json.Marshal(strct.VpcEndpointIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EndpointConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Types":
            if err := json.Unmarshal([]byte(v), &strct.Types); err != nil {
                return err
             }
        case "VpcEndpointIds":
            if err := json.Unmarshal([]byte(v), &strct.VpcEndpointIds); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Parameters_object) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Parameters_object) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, _ := range jsonMap {
        switch k {
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
    // Marshal the "ApiKeySourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApiKeySourceType\": ")
	if tmp, err := json.Marshal(strct.ApiKeySourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BinaryMediaTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BinaryMediaTypes\": ")
	if tmp, err := json.Marshal(strct.BinaryMediaTypes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Body" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Body\": ")
	if tmp, err := json.Marshal(strct.Body); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BodyS3Location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BodyS3Location\": ")
	if tmp, err := json.Marshal(strct.BodyS3Location); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CloneFrom" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CloneFrom\": ")
	if tmp, err := json.Marshal(strct.CloneFrom); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "DisableExecuteApiEndpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DisableExecuteApiEndpoint\": ")
	if tmp, err := json.Marshal(strct.DisableExecuteApiEndpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EndpointConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EndpointConfiguration\": ")
	if tmp, err := json.Marshal(strct.EndpointConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FailOnWarnings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FailOnWarnings\": ")
	if tmp, err := json.Marshal(strct.FailOnWarnings); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MinimumCompressionSize" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MinimumCompressionSize\": ")
	if tmp, err := json.Marshal(strct.MinimumCompressionSize); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Mode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Mode\": ")
	if tmp, err := json.Marshal(strct.Mode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Parameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Policy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Policy\": ")
	if tmp, err := json.Marshal(strct.Policy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RestApiId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RestApiId\": ")
	if tmp, err := json.Marshal(strct.RestApiId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RootResourceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RootResourceId\": ")
	if tmp, err := json.Marshal(strct.RootResourceId); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ApiKeySourceType":
            if err := json.Unmarshal([]byte(v), &strct.ApiKeySourceType); err != nil {
                return err
             }
        case "BinaryMediaTypes":
            if err := json.Unmarshal([]byte(v), &strct.BinaryMediaTypes); err != nil {
                return err
             }
        case "Body":
            if err := json.Unmarshal([]byte(v), &strct.Body); err != nil {
                return err
             }
        case "BodyS3Location":
            if err := json.Unmarshal([]byte(v), &strct.BodyS3Location); err != nil {
                return err
             }
        case "CloneFrom":
            if err := json.Unmarshal([]byte(v), &strct.CloneFrom); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "DisableExecuteApiEndpoint":
            if err := json.Unmarshal([]byte(v), &strct.DisableExecuteApiEndpoint); err != nil {
                return err
             }
        case "EndpointConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.EndpointConfiguration); err != nil {
                return err
             }
        case "FailOnWarnings":
            if err := json.Unmarshal([]byte(v), &strct.FailOnWarnings); err != nil {
                return err
             }
        case "MinimumCompressionSize":
            if err := json.Unmarshal([]byte(v), &strct.MinimumCompressionSize); err != nil {
                return err
             }
        case "Mode":
            if err := json.Unmarshal([]byte(v), &strct.Mode); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "Parameters":
            if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
                return err
             }
        case "Policy":
            if err := json.Unmarshal([]byte(v), &strct.Policy); err != nil {
                return err
             }
        case "RestApiId":
            if err := json.Unmarshal([]byte(v), &strct.RestApiId); err != nil {
                return err
             }
        case "RootResourceId":
            if err := json.Unmarshal([]byte(v), &strct.RootResourceId); err != nil {
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
    return nil
}

func (strct *S3Location) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Bucket" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Bucket\": ")
	if tmp, err := json.Marshal(strct.Bucket); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ETag" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ETag\": ")
	if tmp, err := json.Marshal(strct.ETag); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

func (strct *S3Location) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Bucket":
            if err := json.Unmarshal([]byte(v), &strct.Bucket); err != nil {
                return err
             }
        case "ETag":
            if err := json.Unmarshal([]byte(v), &strct.ETag); err != nil {
                return err
             }
        case "Key":
            if err := json.Unmarshal([]byte(v), &strct.Key); err != nil {
                return err
             }
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
