// Code generated by schema-generate. DO NOT EDIT.

package authorizer

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// JWTConfiguration 
type JWTConfiguration struct {
  Audience []string `json:"Audience,omitempty"`
  Issuer string `json:"Issuer,omitempty"`
}

// Resource Resource Type definition for AWS::ApiGatewayV2::Authorizer
type Resource struct {
  ApiId string `json:"ApiId"`
  AuthorizerCredentialsArn string `json:"AuthorizerCredentialsArn,omitempty"`
  AuthorizerId string `json:"AuthorizerId,omitempty"`
  AuthorizerPayloadFormatVersion string `json:"AuthorizerPayloadFormatVersion,omitempty"`
  AuthorizerResultTtlInSeconds int `json:"AuthorizerResultTtlInSeconds,omitempty"`
  AuthorizerType string `json:"AuthorizerType"`
  AuthorizerUri string `json:"AuthorizerUri,omitempty"`
  EnableSimpleResponses bool `json:"EnableSimpleResponses,omitempty"`
  IdentitySource []string `json:"IdentitySource,omitempty"`
  IdentityValidationExpression string `json:"IdentityValidationExpression,omitempty"`
  JwtConfiguration *JWTConfiguration `json:"JwtConfiguration,omitempty"`
  Name string `json:"Name"`
}

func (strct *JWTConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Audience" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Audience\": ")
	if tmp, err := json.Marshal(strct.Audience); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Issuer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Issuer\": ")
	if tmp, err := json.Marshal(strct.Issuer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *JWTConfiguration) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Audience":
            if err := json.Unmarshal([]byte(v), &strct.Audience); err != nil {
                return err
             }
        case "Issuer":
            if err := json.Unmarshal([]byte(v), &strct.Issuer); err != nil {
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
    // "ApiId" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "AuthorizerCredentialsArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerCredentialsArn\": ")
	if tmp, err := json.Marshal(strct.AuthorizerCredentialsArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AuthorizerId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerId\": ")
	if tmp, err := json.Marshal(strct.AuthorizerId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AuthorizerPayloadFormatVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerPayloadFormatVersion\": ")
	if tmp, err := json.Marshal(strct.AuthorizerPayloadFormatVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AuthorizerResultTtlInSeconds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerResultTtlInSeconds\": ")
	if tmp, err := json.Marshal(strct.AuthorizerResultTtlInSeconds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "AuthorizerType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AuthorizerType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerType\": ")
	if tmp, err := json.Marshal(strct.AuthorizerType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AuthorizerUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthorizerUri\": ")
	if tmp, err := json.Marshal(strct.AuthorizerUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnableSimpleResponses" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableSimpleResponses\": ")
	if tmp, err := json.Marshal(strct.EnableSimpleResponses); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IdentitySource" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdentitySource\": ")
	if tmp, err := json.Marshal(strct.IdentitySource); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IdentityValidationExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdentityValidationExpression\": ")
	if tmp, err := json.Marshal(strct.IdentityValidationExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "JwtConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"JwtConfiguration\": ")
	if tmp, err := json.Marshal(strct.JwtConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    ApiIdReceived := false
    AuthorizerTypeReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ApiId":
            if err := json.Unmarshal([]byte(v), &strct.ApiId); err != nil {
                return err
             }
            ApiIdReceived = true
        case "AuthorizerCredentialsArn":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerCredentialsArn); err != nil {
                return err
             }
        case "AuthorizerId":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerId); err != nil {
                return err
             }
        case "AuthorizerPayloadFormatVersion":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerPayloadFormatVersion); err != nil {
                return err
             }
        case "AuthorizerResultTtlInSeconds":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerResultTtlInSeconds); err != nil {
                return err
             }
        case "AuthorizerType":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerType); err != nil {
                return err
             }
            AuthorizerTypeReceived = true
        case "AuthorizerUri":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizerUri); err != nil {
                return err
             }
        case "EnableSimpleResponses":
            if err := json.Unmarshal([]byte(v), &strct.EnableSimpleResponses); err != nil {
                return err
             }
        case "IdentitySource":
            if err := json.Unmarshal([]byte(v), &strct.IdentitySource); err != nil {
                return err
             }
        case "IdentityValidationExpression":
            if err := json.Unmarshal([]byte(v), &strct.IdentityValidationExpression); err != nil {
                return err
             }
        case "JwtConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.JwtConfiguration); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ApiId (a required property) was received
    if !ApiIdReceived {
        return errors.New("\"ApiId\" is required but was not present")
    }
    // check if AuthorizerType (a required property) was received
    if !AuthorizerTypeReceived {
        return errors.New("\"AuthorizerType\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}