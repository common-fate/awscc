// Code generated by schema-generate. DO NOT EDIT.

package api

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// Body 
type Body struct {
}

// BodyS3Location 
type BodyS3Location struct {
  Bucket string `json:"Bucket,omitempty"`
  Etag string `json:"Etag,omitempty"`
  Key string `json:"Key,omitempty"`
  Version string `json:"Version,omitempty"`
}

// Cors 
type Cors struct {
  AllowCredentials bool `json:"AllowCredentials,omitempty"`
  AllowHeaders []string `json:"AllowHeaders,omitempty"`
  AllowMethods []string `json:"AllowMethods,omitempty"`
  AllowOrigins []string `json:"AllowOrigins,omitempty"`
  ExposeHeaders []string `json:"ExposeHeaders,omitempty"`
  MaxAge int `json:"MaxAge,omitempty"`
}

// Resource Resource Type definition for AWS::ApiGatewayV2::Api
type Resource struct {
  ApiEndpoint string `json:"ApiEndpoint,omitempty"`
  ApiId string `json:"ApiId,omitempty"`
  ApiKeySelectionExpression string `json:"ApiKeySelectionExpression,omitempty"`
  BasePath string `json:"BasePath,omitempty"`
  Body *Body `json:"Body,omitempty"`
  BodyS3Location *BodyS3Location `json:"BodyS3Location,omitempty"`
  CorsConfiguration *Cors `json:"CorsConfiguration,omitempty"`
  CredentialsArn string `json:"CredentialsArn,omitempty"`
  Description string `json:"Description,omitempty"`
  DisableExecuteApiEndpoint bool `json:"DisableExecuteApiEndpoint,omitempty"`
  DisableSchemaValidation bool `json:"DisableSchemaValidation,omitempty"`
  FailOnWarnings bool `json:"FailOnWarnings,omitempty"`
  Name string `json:"Name,omitempty"`
  ProtocolType string `json:"ProtocolType,omitempty"`
  RouteKey string `json:"RouteKey,omitempty"`
  RouteSelectionExpression string `json:"RouteSelectionExpression,omitempty"`

  // This resource type use map for Tags, suggest to use List of Tag
  Tags *Tags `json:"Tags,omitempty"`
  Target string `json:"Target,omitempty"`
  Version string `json:"Version,omitempty"`
}

// Tags This resource type use map for Tags, suggest to use List of Tag
type Tags struct {
}

func (strct *BodyS3Location) MarshalJSON() ([]byte, error) {
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
    // Marshal the "Etag" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Etag\": ")
	if tmp, err := json.Marshal(strct.Etag); err != nil {
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

func (strct *BodyS3Location) UnmarshalJSON(b []byte) error {
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
        case "Etag":
            if err := json.Unmarshal([]byte(v), &strct.Etag); err != nil {
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

func (strct *Cors) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AllowCredentials" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowCredentials\": ")
	if tmp, err := json.Marshal(strct.AllowCredentials); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllowHeaders" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowHeaders\": ")
	if tmp, err := json.Marshal(strct.AllowHeaders); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllowMethods" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowMethods\": ")
	if tmp, err := json.Marshal(strct.AllowMethods); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AllowOrigins" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowOrigins\": ")
	if tmp, err := json.Marshal(strct.AllowOrigins); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ExposeHeaders" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ExposeHeaders\": ")
	if tmp, err := json.Marshal(strct.ExposeHeaders); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaxAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxAge\": ")
	if tmp, err := json.Marshal(strct.MaxAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Cors) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AllowCredentials":
            if err := json.Unmarshal([]byte(v), &strct.AllowCredentials); err != nil {
                return err
             }
        case "AllowHeaders":
            if err := json.Unmarshal([]byte(v), &strct.AllowHeaders); err != nil {
                return err
             }
        case "AllowMethods":
            if err := json.Unmarshal([]byte(v), &strct.AllowMethods); err != nil {
                return err
             }
        case "AllowOrigins":
            if err := json.Unmarshal([]byte(v), &strct.AllowOrigins); err != nil {
                return err
             }
        case "ExposeHeaders":
            if err := json.Unmarshal([]byte(v), &strct.ExposeHeaders); err != nil {
                return err
             }
        case "MaxAge":
            if err := json.Unmarshal([]byte(v), &strct.MaxAge); err != nil {
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
    // Marshal the "ApiEndpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApiEndpoint\": ")
	if tmp, err := json.Marshal(strct.ApiEndpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ApiId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApiId\": ")
	if tmp, err := json.Marshal(strct.ApiId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ApiKeySelectionExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApiKeySelectionExpression\": ")
	if tmp, err := json.Marshal(strct.ApiKeySelectionExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BasePath" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BasePath\": ")
	if tmp, err := json.Marshal(strct.BasePath); err != nil {
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
    // Marshal the "CorsConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CorsConfiguration\": ")
	if tmp, err := json.Marshal(strct.CorsConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CredentialsArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CredentialsArn\": ")
	if tmp, err := json.Marshal(strct.CredentialsArn); err != nil {
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
    // Marshal the "DisableSchemaValidation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DisableSchemaValidation\": ")
	if tmp, err := json.Marshal(strct.DisableSchemaValidation); err != nil {
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
    // Marshal the "ProtocolType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ProtocolType\": ")
	if tmp, err := json.Marshal(strct.ProtocolType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RouteKey" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RouteKey\": ")
	if tmp, err := json.Marshal(strct.RouteKey); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RouteSelectionExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RouteSelectionExpression\": ")
	if tmp, err := json.Marshal(strct.RouteSelectionExpression); err != nil {
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
    // Marshal the "Target" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
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

func (strct *Resource) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ApiEndpoint":
            if err := json.Unmarshal([]byte(v), &strct.ApiEndpoint); err != nil {
                return err
             }
        case "ApiId":
            if err := json.Unmarshal([]byte(v), &strct.ApiId); err != nil {
                return err
             }
        case "ApiKeySelectionExpression":
            if err := json.Unmarshal([]byte(v), &strct.ApiKeySelectionExpression); err != nil {
                return err
             }
        case "BasePath":
            if err := json.Unmarshal([]byte(v), &strct.BasePath); err != nil {
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
        case "CorsConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.CorsConfiguration); err != nil {
                return err
             }
        case "CredentialsArn":
            if err := json.Unmarshal([]byte(v), &strct.CredentialsArn); err != nil {
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
        case "DisableSchemaValidation":
            if err := json.Unmarshal([]byte(v), &strct.DisableSchemaValidation); err != nil {
                return err
             }
        case "FailOnWarnings":
            if err := json.Unmarshal([]byte(v), &strct.FailOnWarnings); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "ProtocolType":
            if err := json.Unmarshal([]byte(v), &strct.ProtocolType); err != nil {
                return err
             }
        case "RouteKey":
            if err := json.Unmarshal([]byte(v), &strct.RouteKey); err != nil {
                return err
             }
        case "RouteSelectionExpression":
            if err := json.Unmarshal([]byte(v), &strct.RouteSelectionExpression); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Target":
            if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
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

func (strct *Tags) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Tags) UnmarshalJSON(b []byte) error {
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