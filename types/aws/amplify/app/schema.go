// Code generated by schema-generate. DO NOT EDIT.

package app

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// AutoBranchCreationConfig 
type AutoBranchCreationConfig struct {
  AutoBranchCreationPatterns []string `json:"AutoBranchCreationPatterns,omitempty"`
  BasicAuthConfig *BasicAuthConfig `json:"BasicAuthConfig,omitempty"`
  BuildSpec string `json:"BuildSpec,omitempty"`
  EnableAutoBranchCreation bool `json:"EnableAutoBranchCreation,omitempty"`
  EnableAutoBuild bool `json:"EnableAutoBuild,omitempty"`
  EnablePerformanceMode bool `json:"EnablePerformanceMode,omitempty"`
  EnablePullRequestPreview bool `json:"EnablePullRequestPreview,omitempty"`
  EnvironmentVariables []*EnvironmentVariable `json:"EnvironmentVariables,omitempty"`
  Framework string `json:"Framework,omitempty"`
  PullRequestEnvironmentName string `json:"PullRequestEnvironmentName,omitempty"`
  Stage string `json:"Stage,omitempty"`
}

// BasicAuthConfig 
type BasicAuthConfig struct {
  EnableBasicAuth bool `json:"EnableBasicAuth,omitempty"`
  Password string `json:"Password,omitempty"`
  Username string `json:"Username,omitempty"`
}

// CustomRule 
type CustomRule struct {
  Condition string `json:"Condition,omitempty"`
  Source string `json:"Source"`
  Status string `json:"Status,omitempty"`
  Target string `json:"Target"`
}

// EnvironmentVariable 
type EnvironmentVariable struct {
  Name string `json:"Name"`
  Value string `json:"Value"`
}

// Resource The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.
type Resource struct {
  AccessToken string `json:"AccessToken,omitempty"`
  AppId string `json:"AppId,omitempty"`
  AppName string `json:"AppName,omitempty"`
  Arn string `json:"Arn,omitempty"`
  AutoBranchCreationConfig *AutoBranchCreationConfig `json:"AutoBranchCreationConfig,omitempty"`
  BasicAuthConfig *BasicAuthConfig `json:"BasicAuthConfig,omitempty"`
  BuildSpec string `json:"BuildSpec,omitempty"`
  CustomHeaders string `json:"CustomHeaders,omitempty"`
  CustomRules []*CustomRule `json:"CustomRules,omitempty"`
  DefaultDomain string `json:"DefaultDomain,omitempty"`
  Description string `json:"Description,omitempty"`
  EnableBranchAutoDeletion bool `json:"EnableBranchAutoDeletion,omitempty"`
  EnvironmentVariables []*EnvironmentVariable `json:"EnvironmentVariables,omitempty"`
  IAMServiceRole string `json:"IAMServiceRole,omitempty"`
  Name string `json:"Name"`
  OauthToken string `json:"OauthToken,omitempty"`
  Platform string `json:"Platform,omitempty"`
  Repository string `json:"Repository,omitempty"`
  Tags []*Tag `json:"Tags,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *AutoBranchCreationConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AutoBranchCreationPatterns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoBranchCreationPatterns\": ")
	if tmp, err := json.Marshal(strct.AutoBranchCreationPatterns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BasicAuthConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BasicAuthConfig\": ")
	if tmp, err := json.Marshal(strct.BasicAuthConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BuildSpec" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BuildSpec\": ")
	if tmp, err := json.Marshal(strct.BuildSpec); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnableAutoBranchCreation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableAutoBranchCreation\": ")
	if tmp, err := json.Marshal(strct.EnableAutoBranchCreation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnableAutoBuild" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableAutoBuild\": ")
	if tmp, err := json.Marshal(strct.EnableAutoBuild); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnablePerformanceMode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnablePerformanceMode\": ")
	if tmp, err := json.Marshal(strct.EnablePerformanceMode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnablePullRequestPreview" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnablePullRequestPreview\": ")
	if tmp, err := json.Marshal(strct.EnablePullRequestPreview); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnvironmentVariables" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnvironmentVariables\": ")
	if tmp, err := json.Marshal(strct.EnvironmentVariables); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Framework" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Framework\": ")
	if tmp, err := json.Marshal(strct.Framework); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PullRequestEnvironmentName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PullRequestEnvironmentName\": ")
	if tmp, err := json.Marshal(strct.PullRequestEnvironmentName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Stage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Stage\": ")
	if tmp, err := json.Marshal(strct.Stage); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AutoBranchCreationConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AutoBranchCreationPatterns":
            if err := json.Unmarshal([]byte(v), &strct.AutoBranchCreationPatterns); err != nil {
                return err
             }
        case "BasicAuthConfig":
            if err := json.Unmarshal([]byte(v), &strct.BasicAuthConfig); err != nil {
                return err
             }
        case "BuildSpec":
            if err := json.Unmarshal([]byte(v), &strct.BuildSpec); err != nil {
                return err
             }
        case "EnableAutoBranchCreation":
            if err := json.Unmarshal([]byte(v), &strct.EnableAutoBranchCreation); err != nil {
                return err
             }
        case "EnableAutoBuild":
            if err := json.Unmarshal([]byte(v), &strct.EnableAutoBuild); err != nil {
                return err
             }
        case "EnablePerformanceMode":
            if err := json.Unmarshal([]byte(v), &strct.EnablePerformanceMode); err != nil {
                return err
             }
        case "EnablePullRequestPreview":
            if err := json.Unmarshal([]byte(v), &strct.EnablePullRequestPreview); err != nil {
                return err
             }
        case "EnvironmentVariables":
            if err := json.Unmarshal([]byte(v), &strct.EnvironmentVariables); err != nil {
                return err
             }
        case "Framework":
            if err := json.Unmarshal([]byte(v), &strct.Framework); err != nil {
                return err
             }
        case "PullRequestEnvironmentName":
            if err := json.Unmarshal([]byte(v), &strct.PullRequestEnvironmentName); err != nil {
                return err
             }
        case "Stage":
            if err := json.Unmarshal([]byte(v), &strct.Stage); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *BasicAuthConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EnableBasicAuth" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableBasicAuth\": ")
	if tmp, err := json.Marshal(strct.EnableBasicAuth); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Password" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Password\": ")
	if tmp, err := json.Marshal(strct.Password); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Username" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Username\": ")
	if tmp, err := json.Marshal(strct.Username); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *BasicAuthConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EnableBasicAuth":
            if err := json.Unmarshal([]byte(v), &strct.EnableBasicAuth); err != nil {
                return err
             }
        case "Password":
            if err := json.Unmarshal([]byte(v), &strct.Password); err != nil {
                return err
             }
        case "Username":
            if err := json.Unmarshal([]byte(v), &strct.Username); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *CustomRule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Condition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Condition\": ")
	if tmp, err := json.Marshal(strct.Condition); err != nil {
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
    // "Target" field is required
    // only required object types supported for marshal checking (for now)
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CustomRule) UnmarshalJSON(b []byte) error {
    SourceReceived := false
    TargetReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Condition":
            if err := json.Unmarshal([]byte(v), &strct.Condition); err != nil {
                return err
             }
        case "Source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
            SourceReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "Target":
            if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
                return err
             }
            TargetReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Source (a required property) was received
    if !SourceReceived {
        return errors.New("\"Source\" is required but was not present")
    }
    // check if Target (a required property) was received
    if !TargetReceived {
        return errors.New("\"Target\" is required but was not present")
    }
    return nil
}

func (strct *EnvironmentVariable) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *EnvironmentVariable) UnmarshalJSON(b []byte) error {
    NameReceived := false
    ValueReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            ValueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
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
    // Marshal the "AccessToken" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessToken\": ")
	if tmp, err := json.Marshal(strct.AccessToken); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AppId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppId\": ")
	if tmp, err := json.Marshal(strct.AppId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AppName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppName\": ")
	if tmp, err := json.Marshal(strct.AppName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Arn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Arn\": ")
	if tmp, err := json.Marshal(strct.Arn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AutoBranchCreationConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoBranchCreationConfig\": ")
	if tmp, err := json.Marshal(strct.AutoBranchCreationConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BasicAuthConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BasicAuthConfig\": ")
	if tmp, err := json.Marshal(strct.BasicAuthConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "BuildSpec" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BuildSpec\": ")
	if tmp, err := json.Marshal(strct.BuildSpec); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CustomHeaders" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CustomHeaders\": ")
	if tmp, err := json.Marshal(strct.CustomHeaders); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CustomRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CustomRules\": ")
	if tmp, err := json.Marshal(strct.CustomRules); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultDomain" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultDomain\": ")
	if tmp, err := json.Marshal(strct.DefaultDomain); err != nil {
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
    // Marshal the "EnableBranchAutoDeletion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnableBranchAutoDeletion\": ")
	if tmp, err := json.Marshal(strct.EnableBranchAutoDeletion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EnvironmentVariables" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnvironmentVariables\": ")
	if tmp, err := json.Marshal(strct.EnvironmentVariables); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IAMServiceRole" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IAMServiceRole\": ")
	if tmp, err := json.Marshal(strct.IAMServiceRole); err != nil {
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
    // Marshal the "OauthToken" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OauthToken\": ")
	if tmp, err := json.Marshal(strct.OauthToken); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Platform" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Platform\": ")
	if tmp, err := json.Marshal(strct.Platform); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Repository" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Repository\": ")
	if tmp, err := json.Marshal(strct.Repository); err != nil {
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
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AccessToken":
            if err := json.Unmarshal([]byte(v), &strct.AccessToken); err != nil {
                return err
             }
        case "AppId":
            if err := json.Unmarshal([]byte(v), &strct.AppId); err != nil {
                return err
             }
        case "AppName":
            if err := json.Unmarshal([]byte(v), &strct.AppName); err != nil {
                return err
             }
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AutoBranchCreationConfig":
            if err := json.Unmarshal([]byte(v), &strct.AutoBranchCreationConfig); err != nil {
                return err
             }
        case "BasicAuthConfig":
            if err := json.Unmarshal([]byte(v), &strct.BasicAuthConfig); err != nil {
                return err
             }
        case "BuildSpec":
            if err := json.Unmarshal([]byte(v), &strct.BuildSpec); err != nil {
                return err
             }
        case "CustomHeaders":
            if err := json.Unmarshal([]byte(v), &strct.CustomHeaders); err != nil {
                return err
             }
        case "CustomRules":
            if err := json.Unmarshal([]byte(v), &strct.CustomRules); err != nil {
                return err
             }
        case "DefaultDomain":
            if err := json.Unmarshal([]byte(v), &strct.DefaultDomain); err != nil {
                return err
             }
        case "Description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "EnableBranchAutoDeletion":
            if err := json.Unmarshal([]byte(v), &strct.EnableBranchAutoDeletion); err != nil {
                return err
             }
        case "EnvironmentVariables":
            if err := json.Unmarshal([]byte(v), &strct.EnvironmentVariables); err != nil {
                return err
             }
        case "IAMServiceRole":
            if err := json.Unmarshal([]byte(v), &strct.IAMServiceRole); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "OauthToken":
            if err := json.Unmarshal([]byte(v), &strct.OauthToken); err != nil {
                return err
             }
        case "Platform":
            if err := json.Unmarshal([]byte(v), &strct.Platform); err != nil {
                return err
             }
        case "Repository":
            if err := json.Unmarshal([]byte(v), &strct.Repository); err != nil {
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
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
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
