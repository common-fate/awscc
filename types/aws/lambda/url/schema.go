// Code generated by schema-generate. DO NOT EDIT.

package url

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// Cors 
type Cors struct {

  // Specifies whether credentials are included in the CORS request.
  AllowCredentials bool `json:"AllowCredentials,omitempty"`

  // Represents a collection of allowed headers.
  AllowHeaders []string `json:"AllowHeaders,omitempty"`

  // Represents a collection of allowed HTTP methods.
  AllowMethods []string `json:"AllowMethods,omitempty"`

  // Represents a collection of allowed origins.
  AllowOrigins []string `json:"AllowOrigins,omitempty"`

  // Represents a collection of exposed headers.
  ExposeHeaders []string `json:"ExposeHeaders,omitempty"`
  MaxAge int `json:"MaxAge,omitempty"`
}

// Resource Resource Type definition for AWS::Lambda::Url
type Resource struct {

  // Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
  AuthType string `json:"AuthType"`
  Cors *Cors `json:"Cors,omitempty"`

  // The full Amazon Resource Name (ARN) of the function associated with the Function URL.
  FunctionArn string `json:"FunctionArn,omitempty"`

  // The generated url for this resource.
  FunctionUrl string `json:"FunctionUrl,omitempty"`

  // The invocation mode for the function?s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.
  InvokeMode string `json:"InvokeMode,omitempty"`

  // The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
  Qualifier string `json:"Qualifier,omitempty"`

  // The Amazon Resource Name (ARN) of the function associated with the Function URL.
  TargetFunctionArn string `json:"TargetFunctionArn"`
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
    // "AuthType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AuthType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthType\": ")
	if tmp, err := json.Marshal(strct.AuthType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Cors" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Cors\": ")
	if tmp, err := json.Marshal(strct.Cors); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FunctionArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionArn\": ")
	if tmp, err := json.Marshal(strct.FunctionArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FunctionUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FunctionUrl\": ")
	if tmp, err := json.Marshal(strct.FunctionUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InvokeMode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InvokeMode\": ")
	if tmp, err := json.Marshal(strct.InvokeMode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Qualifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Qualifier\": ")
	if tmp, err := json.Marshal(strct.Qualifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "TargetFunctionArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "TargetFunctionArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TargetFunctionArn\": ")
	if tmp, err := json.Marshal(strct.TargetFunctionArn); err != nil {
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
    AuthTypeReceived := false
    TargetFunctionArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AuthType":
            if err := json.Unmarshal([]byte(v), &strct.AuthType); err != nil {
                return err
             }
            AuthTypeReceived = true
        case "Cors":
            if err := json.Unmarshal([]byte(v), &strct.Cors); err != nil {
                return err
             }
        case "FunctionArn":
            if err := json.Unmarshal([]byte(v), &strct.FunctionArn); err != nil {
                return err
             }
        case "FunctionUrl":
            if err := json.Unmarshal([]byte(v), &strct.FunctionUrl); err != nil {
                return err
             }
        case "InvokeMode":
            if err := json.Unmarshal([]byte(v), &strct.InvokeMode); err != nil {
                return err
             }
        case "Qualifier":
            if err := json.Unmarshal([]byte(v), &strct.Qualifier); err != nil {
                return err
             }
        case "TargetFunctionArn":
            if err := json.Unmarshal([]byte(v), &strct.TargetFunctionArn); err != nil {
                return err
             }
            TargetFunctionArnReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AuthType (a required property) was received
    if !AuthTypeReceived {
        return errors.New("\"AuthType\" is required but was not present")
    }
    // check if TargetFunctionArn (a required property) was received
    if !TargetFunctionArnReceived {
        return errors.New("\"TargetFunctionArn\" is required but was not present")
    }
    return nil
}