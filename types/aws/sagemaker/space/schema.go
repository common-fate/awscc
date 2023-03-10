// Code generated by schema-generate. DO NOT EDIT.

package space

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// CustomImage A custom SageMaker image.
type CustomImage struct {

  // The Name of the AppImageConfig.
  AppImageConfigName string `json:"AppImageConfigName"`

  // The name of the CustomImage. Must be unique to your account.
  ImageName string `json:"ImageName"`

  // The version number of the CustomImage.
  ImageVersionNumber int `json:"ImageVersionNumber,omitempty"`
}

// JupyterServerAppSettings The JupyterServer app settings.
type JupyterServerAppSettings struct {
  DefaultResourceSpec *ResourceSpec `json:"DefaultResourceSpec,omitempty"`
}

// KernelGatewayAppSettings The kernel gateway app settings.
type KernelGatewayAppSettings struct {

  // A list of custom SageMaker images that are configured to run as a KernelGateway app.
  CustomImages []*CustomImage `json:"CustomImages,omitempty"`

  // The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.
  DefaultResourceSpec *ResourceSpec `json:"DefaultResourceSpec,omitempty"`
}

// Resource Resource Type definition for AWS::SageMaker::Space
type Resource struct {

  // The ID of the associated Domain.
  DomainId string `json:"DomainId"`

  // The space Amazon Resource Name (ARN).
  SpaceArn string `json:"SpaceArn,omitempty"`

  // A name for the Space.
  SpaceName string `json:"SpaceName"`

  // A collection of settings.
  SpaceSettings *SpaceSettings `json:"SpaceSettings,omitempty"`

  // A list of tags to apply to the space.
  Tags []*Tag `json:"Tags,omitempty"`
}

// ResourceSpec 
type ResourceSpec struct {

  // The instance type that the image version runs on.
  InstanceType string `json:"InstanceType,omitempty"`

  // The ARN of the SageMaker image that the image version belongs to.
  SageMakerImageArn string `json:"SageMakerImageArn,omitempty"`

  // The ARN of the image version created on the instance.
  SageMakerImageVersionArn string `json:"SageMakerImageVersionArn,omitempty"`
}

// SpaceSettings A collection of settings that apply to spaces of Amazon SageMaker Studio. These settings are specified when the CreateSpace API is called.
type SpaceSettings struct {

  // The Jupyter server's app settings.
  JupyterServerAppSettings *JupyterServerAppSettings `json:"JupyterServerAppSettings,omitempty"`

  // The kernel gateway app settings.
  KernelGatewayAppSettings *KernelGatewayAppSettings `json:"KernelGatewayAppSettings,omitempty"`
}

// Tag 
type Tag struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

func (strct *CustomImage) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AppImageConfigName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AppImageConfigName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AppImageConfigName\": ")
	if tmp, err := json.Marshal(strct.AppImageConfigName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ImageName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ImageName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageName\": ")
	if tmp, err := json.Marshal(strct.ImageName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ImageVersionNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ImageVersionNumber\": ")
	if tmp, err := json.Marshal(strct.ImageVersionNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CustomImage) UnmarshalJSON(b []byte) error {
    AppImageConfigNameReceived := false
    ImageNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AppImageConfigName":
            if err := json.Unmarshal([]byte(v), &strct.AppImageConfigName); err != nil {
                return err
             }
            AppImageConfigNameReceived = true
        case "ImageName":
            if err := json.Unmarshal([]byte(v), &strct.ImageName); err != nil {
                return err
             }
            ImageNameReceived = true
        case "ImageVersionNumber":
            if err := json.Unmarshal([]byte(v), &strct.ImageVersionNumber); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AppImageConfigName (a required property) was received
    if !AppImageConfigNameReceived {
        return errors.New("\"AppImageConfigName\" is required but was not present")
    }
    // check if ImageName (a required property) was received
    if !ImageNameReceived {
        return errors.New("\"ImageName\" is required but was not present")
    }
    return nil
}

func (strct *JupyterServerAppSettings) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DefaultResourceSpec" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultResourceSpec\": ")
	if tmp, err := json.Marshal(strct.DefaultResourceSpec); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *JupyterServerAppSettings) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DefaultResourceSpec":
            if err := json.Unmarshal([]byte(v), &strct.DefaultResourceSpec); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *KernelGatewayAppSettings) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CustomImages" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CustomImages\": ")
	if tmp, err := json.Marshal(strct.CustomImages); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DefaultResourceSpec" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DefaultResourceSpec\": ")
	if tmp, err := json.Marshal(strct.DefaultResourceSpec); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *KernelGatewayAppSettings) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CustomImages":
            if err := json.Unmarshal([]byte(v), &strct.CustomImages); err != nil {
                return err
             }
        case "DefaultResourceSpec":
            if err := json.Unmarshal([]byte(v), &strct.DefaultResourceSpec); err != nil {
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
    // "DomainId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DomainId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainId\": ")
	if tmp, err := json.Marshal(strct.DomainId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SpaceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SpaceArn\": ")
	if tmp, err := json.Marshal(strct.SpaceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SpaceName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SpaceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SpaceName\": ")
	if tmp, err := json.Marshal(strct.SpaceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SpaceSettings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SpaceSettings\": ")
	if tmp, err := json.Marshal(strct.SpaceSettings); err != nil {
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
    DomainIdReceived := false
    SpaceNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DomainId":
            if err := json.Unmarshal([]byte(v), &strct.DomainId); err != nil {
                return err
             }
            DomainIdReceived = true
        case "SpaceArn":
            if err := json.Unmarshal([]byte(v), &strct.SpaceArn); err != nil {
                return err
             }
        case "SpaceName":
            if err := json.Unmarshal([]byte(v), &strct.SpaceName); err != nil {
                return err
             }
            SpaceNameReceived = true
        case "SpaceSettings":
            if err := json.Unmarshal([]byte(v), &strct.SpaceSettings); err != nil {
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
    // check if DomainId (a required property) was received
    if !DomainIdReceived {
        return errors.New("\"DomainId\" is required but was not present")
    }
    // check if SpaceName (a required property) was received
    if !SpaceNameReceived {
        return errors.New("\"SpaceName\" is required but was not present")
    }
    return nil
}

func (strct *ResourceSpec) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "InstanceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceType\": ")
	if tmp, err := json.Marshal(strct.InstanceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SageMakerImageArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SageMakerImageArn\": ")
	if tmp, err := json.Marshal(strct.SageMakerImageArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SageMakerImageVersionArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SageMakerImageVersionArn\": ")
	if tmp, err := json.Marshal(strct.SageMakerImageVersionArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResourceSpec) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "InstanceType":
            if err := json.Unmarshal([]byte(v), &strct.InstanceType); err != nil {
                return err
             }
        case "SageMakerImageArn":
            if err := json.Unmarshal([]byte(v), &strct.SageMakerImageArn); err != nil {
                return err
             }
        case "SageMakerImageVersionArn":
            if err := json.Unmarshal([]byte(v), &strct.SageMakerImageVersionArn); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *SpaceSettings) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "JupyterServerAppSettings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"JupyterServerAppSettings\": ")
	if tmp, err := json.Marshal(strct.JupyterServerAppSettings); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KernelGatewayAppSettings" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KernelGatewayAppSettings\": ")
	if tmp, err := json.Marshal(strct.KernelGatewayAppSettings); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SpaceSettings) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "JupyterServerAppSettings":
            if err := json.Unmarshal([]byte(v), &strct.JupyterServerAppSettings); err != nil {
                return err
             }
        case "KernelGatewayAppSettings":
            if err := json.Unmarshal([]byte(v), &strct.KernelGatewayAppSettings); err != nil {
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
