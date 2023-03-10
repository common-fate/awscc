// Code generated by schema-generate. DO NOT EDIT.

package identityproviderconfig

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// OidcIdentityProviderConfig An object representing an OpenID Connect (OIDC) configuration.
type OidcIdentityProviderConfig struct {

  // This is also known as audience. The ID for the client application that makes authentication requests to the OpenID identity provider.
  ClientId string `json:"ClientId"`

  // The JWT claim that the provider uses to return your groups.
  GroupsClaim string `json:"GroupsClaim,omitempty"`

  // The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).
  GroupsPrefix string `json:"GroupsPrefix,omitempty"`

  // The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.
  IssuerUrl string `json:"IssuerUrl"`
  RequiredClaims []*RequiredClaim `json:"RequiredClaims,omitempty"`

  // The JSON Web Token (JWT) claim to use as the username. The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.
  UsernameClaim string `json:"UsernameClaim,omitempty"`

  // The prefix that is prepended to username claims to prevent clashes with existing names. If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.
  UsernamePrefix string `json:"UsernamePrefix,omitempty"`
}

// RequiredClaim The key value pairs that describe required claims in the identity token. If set, each claim is verified to be present in the token with a matching value.
type RequiredClaim struct {

  // The key of the requiredClaims.
  Key string `json:"Key"`

  // The value for the requiredClaims.
  Value string `json:"Value"`
}

// Resource An object representing an Amazon EKS IdentityProviderConfig.
type Resource struct {

  // The name of the identity provider configuration.
  ClusterName string `json:"ClusterName"`

  // The ARN of the configuration.
  IdentityProviderConfigArn string `json:"IdentityProviderConfigArn,omitempty"`

  // The name of the OIDC provider configuration.
  IdentityProviderConfigName string `json:"IdentityProviderConfigName,omitempty"`
  Oidc *OidcIdentityProviderConfig `json:"Oidc,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // The type of the identity provider configuration.
  Type string `json:"Type"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *OidcIdentityProviderConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ClientId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ClientId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientId\": ")
	if tmp, err := json.Marshal(strct.ClientId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GroupsClaim" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GroupsClaim\": ")
	if tmp, err := json.Marshal(strct.GroupsClaim); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GroupsPrefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GroupsPrefix\": ")
	if tmp, err := json.Marshal(strct.GroupsPrefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IssuerUrl" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "IssuerUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IssuerUrl\": ")
	if tmp, err := json.Marshal(strct.IssuerUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RequiredClaims" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RequiredClaims\": ")
	if tmp, err := json.Marshal(strct.RequiredClaims); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UsernameClaim" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UsernameClaim\": ")
	if tmp, err := json.Marshal(strct.UsernameClaim); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UsernamePrefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UsernamePrefix\": ")
	if tmp, err := json.Marshal(strct.UsernamePrefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *OidcIdentityProviderConfig) UnmarshalJSON(b []byte) error {
    ClientIdReceived := false
    IssuerUrlReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ClientId":
            if err := json.Unmarshal([]byte(v), &strct.ClientId); err != nil {
                return err
             }
            ClientIdReceived = true
        case "GroupsClaim":
            if err := json.Unmarshal([]byte(v), &strct.GroupsClaim); err != nil {
                return err
             }
        case "GroupsPrefix":
            if err := json.Unmarshal([]byte(v), &strct.GroupsPrefix); err != nil {
                return err
             }
        case "IssuerUrl":
            if err := json.Unmarshal([]byte(v), &strct.IssuerUrl); err != nil {
                return err
             }
            IssuerUrlReceived = true
        case "RequiredClaims":
            if err := json.Unmarshal([]byte(v), &strct.RequiredClaims); err != nil {
                return err
             }
        case "UsernameClaim":
            if err := json.Unmarshal([]byte(v), &strct.UsernameClaim); err != nil {
                return err
             }
        case "UsernamePrefix":
            if err := json.Unmarshal([]byte(v), &strct.UsernamePrefix); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ClientId (a required property) was received
    if !ClientIdReceived {
        return errors.New("\"ClientId\" is required but was not present")
    }
    // check if IssuerUrl (a required property) was received
    if !IssuerUrlReceived {
        return errors.New("\"IssuerUrl\" is required but was not present")
    }
    return nil
}

func (strct *RequiredClaim) MarshalJSON() ([]byte, error) {
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

func (strct *RequiredClaim) UnmarshalJSON(b []byte) error {
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

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ClusterName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ClusterName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClusterName\": ")
	if tmp, err := json.Marshal(strct.ClusterName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IdentityProviderConfigArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdentityProviderConfigArn\": ")
	if tmp, err := json.Marshal(strct.IdentityProviderConfigArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IdentityProviderConfigName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdentityProviderConfigName\": ")
	if tmp, err := json.Marshal(strct.IdentityProviderConfigName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Oidc" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Oidc\": ")
	if tmp, err := json.Marshal(strct.Oidc); err != nil {
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
    // "Type" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
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
    ClusterNameReceived := false
    TypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ClusterName":
            if err := json.Unmarshal([]byte(v), &strct.ClusterName); err != nil {
                return err
             }
            ClusterNameReceived = true
        case "IdentityProviderConfigArn":
            if err := json.Unmarshal([]byte(v), &strct.IdentityProviderConfigArn); err != nil {
                return err
             }
        case "IdentityProviderConfigName":
            if err := json.Unmarshal([]byte(v), &strct.IdentityProviderConfigName); err != nil {
                return err
             }
        case "Oidc":
            if err := json.Unmarshal([]byte(v), &strct.Oidc); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            TypeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ClusterName (a required property) was received
    if !ClusterNameReceived {
        return errors.New("\"ClusterName\" is required but was not present")
    }
    // check if Type (a required property) was received
    if !TypeReceived {
        return errors.New("\"Type\" is required but was not present")
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
