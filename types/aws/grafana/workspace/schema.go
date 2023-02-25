// Code generated by schema-generate. DO NOT EDIT.

package workspace

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AssertionAttributes Maps Grafana friendly names to the IdPs SAML attributes.
type AssertionAttributes struct {

  // Name of the attribute within the SAML assert to use as the users email in Grafana.
  Email string `json:"Email,omitempty"`

  // Name of the attribute within the SAML assert to use as the users groups in Grafana.
  Groups string `json:"Groups,omitempty"`

  // Name of the attribute within the SAML assert to use as the users login handle in Grafana.
  Login string `json:"Login,omitempty"`

  // Name of the attribute within the SAML assert to use as the users name in Grafana.
  Name string `json:"Name,omitempty"`

  // Name of the attribute within the SAML assert to use as the users organizations in Grafana.
  Org string `json:"Org,omitempty"`

  // Name of the attribute within the SAML assert to use as the users roles in Grafana.
  Role string `json:"Role,omitempty"`
}

// IdpMetadata IdP Metadata used to configure SAML authentication in Grafana.
type IdpMetadata struct {

  // URL that vends the IdPs metadata.
  Url string `json:"Url,omitempty"`

  // XML blob of the IdPs metadata.
  Xml string `json:"Xml,omitempty"`
}

// Resource Definition of AWS::Grafana::Workspace Resource Type
type Resource struct {
  AccountAccessType string `json:"AccountAccessType,omitempty"`

  // List of authentication providers to enable.
  AuthenticationProviders []string `json:"AuthenticationProviders,omitempty"`

  // A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
  ClientToken string `json:"ClientToken,omitempty"`

  // Timestamp when the workspace was created.
  CreationTimestamp string `json:"CreationTimestamp,omitempty"`

  // List of data sources on the service managed IAM role.
  DataSources []string `json:"DataSources,omitempty"`

  // Description of a workspace.
  Description string `json:"Description,omitempty"`

  // Endpoint for the Grafana workspace.
  Endpoint string `json:"Endpoint,omitempty"`

  // Version of Grafana the workspace is currently using.
  GrafanaVersion string `json:"GrafanaVersion,omitempty"`

  // The id that uniquely identifies a Grafana workspace.
  Id string `json:"Id,omitempty"`

  // Timestamp when the workspace was last modified
  ModificationTimestamp string `json:"ModificationTimestamp,omitempty"`

  // The user friendly name of a workspace.
  Name string `json:"Name,omitempty"`

  // List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.
  NotificationDestinations []string `json:"NotificationDestinations,omitempty"`

  // The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.
  OrganizationRoleName string `json:"OrganizationRoleName,omitempty"`

  // List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.
  OrganizationalUnits []string `json:"OrganizationalUnits,omitempty"`
  PermissionType string `json:"PermissionType,omitempty"`

  // IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.
  RoleArn string `json:"RoleArn,omitempty"`
  SamlConfiguration *SamlConfiguration `json:"SamlConfiguration,omitempty"`
  SamlConfigurationStatus string `json:"SamlConfigurationStatus,omitempty"`

  // The client ID of the AWS SSO Managed Application.
  SsoClientId string `json:"SsoClientId,omitempty"`

  // The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.
  StackSetName string `json:"StackSetName,omitempty"`
  Status string `json:"Status,omitempty"`
}

// RoleValues Maps SAML roles to the Grafana Editor and Admin roles.
type RoleValues struct {

  // List of SAML roles which will be mapped into the Grafana Admin role.
  Admin []string `json:"Admin,omitempty"`

  // List of SAML roles which will be mapped into the Grafana Editor role.
  Editor []string `json:"Editor,omitempty"`
}

// SamlConfiguration SAML configuration data associated with an AMG workspace.
type SamlConfiguration struct {

  // List of SAML organizations allowed to access Grafana.
  AllowedOrganizations []string `json:"AllowedOrganizations,omitempty"`
  AssertionAttributes *AssertionAttributes `json:"AssertionAttributes,omitempty"`
  IdpMetadata *IdpMetadata `json:"IdpMetadata"`

  // The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.
  LoginValidityDuration float64 `json:"LoginValidityDuration,omitempty"`
  RoleValues *RoleValues `json:"RoleValues,omitempty"`
}

func (strct *AssertionAttributes) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Email" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Email\": ")
	if tmp, err := json.Marshal(strct.Email); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Groups" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Groups\": ")
	if tmp, err := json.Marshal(strct.Groups); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Login" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Login\": ")
	if tmp, err := json.Marshal(strct.Login); err != nil {
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
    // Marshal the "Org" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Org\": ")
	if tmp, err := json.Marshal(strct.Org); err != nil {
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

func (strct *AssertionAttributes) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Email":
            if err := json.Unmarshal([]byte(v), &strct.Email); err != nil {
                return err
             }
        case "Groups":
            if err := json.Unmarshal([]byte(v), &strct.Groups); err != nil {
                return err
             }
        case "Login":
            if err := json.Unmarshal([]byte(v), &strct.Login); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "Org":
            if err := json.Unmarshal([]byte(v), &strct.Org); err != nil {
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

func (strct *IdpMetadata) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Url" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Xml" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Xml\": ")
	if tmp, err := json.Marshal(strct.Xml); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *IdpMetadata) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        case "Xml":
            if err := json.Unmarshal([]byte(v), &strct.Xml); err != nil {
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
    // Marshal the "AccountAccessType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccountAccessType\": ")
	if tmp, err := json.Marshal(strct.AccountAccessType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AuthenticationProviders" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AuthenticationProviders\": ")
	if tmp, err := json.Marshal(strct.AuthenticationProviders); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ClientToken" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClientToken\": ")
	if tmp, err := json.Marshal(strct.ClientToken); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CreationTimestamp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreationTimestamp\": ")
	if tmp, err := json.Marshal(strct.CreationTimestamp); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DataSources" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataSources\": ")
	if tmp, err := json.Marshal(strct.DataSources); err != nil {
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
    // Marshal the "Endpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Endpoint\": ")
	if tmp, err := json.Marshal(strct.Endpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "GrafanaVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GrafanaVersion\": ")
	if tmp, err := json.Marshal(strct.GrafanaVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ModificationTimestamp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ModificationTimestamp\": ")
	if tmp, err := json.Marshal(strct.ModificationTimestamp); err != nil {
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
    // Marshal the "NotificationDestinations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"NotificationDestinations\": ")
	if tmp, err := json.Marshal(strct.NotificationDestinations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OrganizationRoleName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OrganizationRoleName\": ")
	if tmp, err := json.Marshal(strct.OrganizationRoleName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OrganizationalUnits" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OrganizationalUnits\": ")
	if tmp, err := json.Marshal(strct.OrganizationalUnits); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PermissionType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PermissionType\": ")
	if tmp, err := json.Marshal(strct.PermissionType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RoleArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleArn\": ")
	if tmp, err := json.Marshal(strct.RoleArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SamlConfiguration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SamlConfiguration\": ")
	if tmp, err := json.Marshal(strct.SamlConfiguration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SamlConfigurationStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SamlConfigurationStatus\": ")
	if tmp, err := json.Marshal(strct.SamlConfigurationStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SsoClientId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SsoClientId\": ")
	if tmp, err := json.Marshal(strct.SsoClientId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StackSetName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StackSetName\": ")
	if tmp, err := json.Marshal(strct.StackSetName); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AccountAccessType":
            if err := json.Unmarshal([]byte(v), &strct.AccountAccessType); err != nil {
                return err
             }
        case "AuthenticationProviders":
            if err := json.Unmarshal([]byte(v), &strct.AuthenticationProviders); err != nil {
                return err
             }
        case "ClientToken":
            if err := json.Unmarshal([]byte(v), &strct.ClientToken); err != nil {
                return err
             }
        case "CreationTimestamp":
            if err := json.Unmarshal([]byte(v), &strct.CreationTimestamp); err != nil {
                return err
             }
        case "DataSources":
            if err := json.Unmarshal([]byte(v), &strct.DataSources); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "Endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
        case "GrafanaVersion":
            if err := json.Unmarshal([]byte(v), &strct.GrafanaVersion); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "ModificationTimestamp":
            if err := json.Unmarshal([]byte(v), &strct.ModificationTimestamp); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "NotificationDestinations":
            if err := json.Unmarshal([]byte(v), &strct.NotificationDestinations); err != nil {
                return err
             }
        case "OrganizationRoleName":
            if err := json.Unmarshal([]byte(v), &strct.OrganizationRoleName); err != nil {
                return err
             }
        case "OrganizationalUnits":
            if err := json.Unmarshal([]byte(v), &strct.OrganizationalUnits); err != nil {
                return err
             }
        case "PermissionType":
            if err := json.Unmarshal([]byte(v), &strct.PermissionType); err != nil {
                return err
             }
        case "RoleArn":
            if err := json.Unmarshal([]byte(v), &strct.RoleArn); err != nil {
                return err
             }
        case "SamlConfiguration":
            if err := json.Unmarshal([]byte(v), &strct.SamlConfiguration); err != nil {
                return err
             }
        case "SamlConfigurationStatus":
            if err := json.Unmarshal([]byte(v), &strct.SamlConfigurationStatus); err != nil {
                return err
             }
        case "SsoClientId":
            if err := json.Unmarshal([]byte(v), &strct.SsoClientId); err != nil {
                return err
             }
        case "StackSetName":
            if err := json.Unmarshal([]byte(v), &strct.StackSetName); err != nil {
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
    return nil
}

func (strct *RoleValues) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Admin" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Admin\": ")
	if tmp, err := json.Marshal(strct.Admin); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Editor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Editor\": ")
	if tmp, err := json.Marshal(strct.Editor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RoleValues) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Admin":
            if err := json.Unmarshal([]byte(v), &strct.Admin); err != nil {
                return err
             }
        case "Editor":
            if err := json.Unmarshal([]byte(v), &strct.Editor); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *SamlConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AllowedOrganizations" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AllowedOrganizations\": ")
	if tmp, err := json.Marshal(strct.AllowedOrganizations); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AssertionAttributes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssertionAttributes\": ")
	if tmp, err := json.Marshal(strct.AssertionAttributes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IdpMetadata" field is required
    if strct.IdpMetadata == nil {
        return nil, errors.New("IdpMetadata is a required field")
    }
    // Marshal the "IdpMetadata" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IdpMetadata\": ")
	if tmp, err := json.Marshal(strct.IdpMetadata); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LoginValidityDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LoginValidityDuration\": ")
	if tmp, err := json.Marshal(strct.LoginValidityDuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RoleValues" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RoleValues\": ")
	if tmp, err := json.Marshal(strct.RoleValues); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SamlConfiguration) UnmarshalJSON(b []byte) error {
    IdpMetadataReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AllowedOrganizations":
            if err := json.Unmarshal([]byte(v), &strct.AllowedOrganizations); err != nil {
                return err
             }
        case "AssertionAttributes":
            if err := json.Unmarshal([]byte(v), &strct.AssertionAttributes); err != nil {
                return err
             }
        case "IdpMetadata":
            if err := json.Unmarshal([]byte(v), &strct.IdpMetadata); err != nil {
                return err
             }
            IdpMetadataReceived = true
        case "LoginValidityDuration":
            if err := json.Unmarshal([]byte(v), &strct.LoginValidityDuration); err != nil {
                return err
             }
        case "RoleValues":
            if err := json.Unmarshal([]byte(v), &strct.RoleValues); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if IdpMetadata (a required property) was received
    if !IdpMetadataReceived {
        return errors.New("\"IdpMetadata\" is required but was not present")
    }
    return nil
}
