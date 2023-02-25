// Code generated by schema-generate. DO NOT EDIT.

package container

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// Container Describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.
type Container struct {

  // The launch command for the container.
  Command []string `json:"Command,omitempty"`

  // The name of the container.
  ContainerName string `json:"ContainerName,omitempty"`

  // The environment variables of the container.
  Environment []*EnvironmentVariable `json:"Environment,omitempty"`

  // The name of the image used for the container.
  Image string `json:"Image,omitempty"`

  // The open firewall ports of the container.
  Ports []*PortInfo `json:"Ports,omitempty"`
}

// ContainerServiceDeployment Describes a container deployment configuration of an Amazon Lightsail container service.
type ContainerServiceDeployment struct {

  // An object that describes the configuration for the containers of the deployment.
  Containers []*Container `json:"Containers,omitempty"`

  // An object that describes the endpoint of the deployment.
  PublicEndpoint *PublicEndpoint `json:"PublicEndpoint,omitempty"`
}

// EnvironmentVariable 
type EnvironmentVariable struct {
  Value string `json:"Value,omitempty"`
  Variable string `json:"Variable,omitempty"`
}

// HealthCheckConfig Describes the health check configuration of an Amazon Lightsail container service.
type HealthCheckConfig struct {

  // The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.
  HealthyThreshold int `json:"HealthyThreshold,omitempty"`

  // The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.
  IntervalSeconds int `json:"IntervalSeconds,omitempty"`

  // The path on the container on which to perform the health check. The default value is /.
  Path string `json:"Path,omitempty"`

  // The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).
  SuccessCodes string `json:"SuccessCodes,omitempty"`

  // The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.
  TimeoutSeconds int `json:"TimeoutSeconds,omitempty"`

  // The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.
  UnhealthyThreshold int `json:"UnhealthyThreshold,omitempty"`
}

// PortInfo 
type PortInfo struct {
  Port string `json:"Port,omitempty"`
  Protocol string `json:"Protocol,omitempty"`
}

// PublicDomainName The public domain name to use with the container service, such as example.com and www.example.com.
type PublicDomainName struct {
  CertificateName string `json:"CertificateName,omitempty"`

  // An object that describes the configuration for the containers of the deployment.
  DomainNames []string `json:"DomainNames,omitempty"`
}

// PublicEndpoint Describes the settings of a public endpoint for an Amazon Lightsail container service.
type PublicEndpoint struct {

  // The name of the container for the endpoint.
  ContainerName string `json:"ContainerName,omitempty"`

  // The port of the container to which traffic is forwarded to.
  ContainerPort int `json:"ContainerPort,omitempty"`

  // An object that describes the health check configuration of the container.
  HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitempty"`
}

// Resource Resource Type definition for AWS::Lightsail::Container
type Resource struct {
  ContainerArn string `json:"ContainerArn,omitempty"`

  // Describes a container deployment configuration of an Amazon Lightsail container service.
  ContainerServiceDeployment *ContainerServiceDeployment `json:"ContainerServiceDeployment,omitempty"`

  // A Boolean value to indicate whether the container service is disabled.
  IsDisabled bool `json:"IsDisabled,omitempty"`

  // The power specification for the container service.
  Power string `json:"Power"`

  // The public domain names to use with the container service, such as example.com and www.example.com.
  PublicDomainNames []*PublicDomainName `json:"PublicDomainNames,omitempty"`

  // The scale specification for the container service.
  Scale int `json:"Scale"`

  // The name for the container service.
  ServiceName string `json:"ServiceName"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // The publicly accessible URL of the container service.
  Url string `json:"Url,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value,omitempty"`
}

func (strct *Container) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Command" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Command\": ")
	if tmp, err := json.Marshal(strct.Command); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ContainerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContainerName\": ")
	if tmp, err := json.Marshal(strct.ContainerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Environment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Environment\": ")
	if tmp, err := json.Marshal(strct.Environment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Image" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Image\": ")
	if tmp, err := json.Marshal(strct.Image); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Ports" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ports\": ")
	if tmp, err := json.Marshal(strct.Ports); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Container) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Command":
            if err := json.Unmarshal([]byte(v), &strct.Command); err != nil {
                return err
             }
        case "ContainerName":
            if err := json.Unmarshal([]byte(v), &strct.ContainerName); err != nil {
                return err
             }
        case "Environment":
            if err := json.Unmarshal([]byte(v), &strct.Environment); err != nil {
                return err
             }
        case "Image":
            if err := json.Unmarshal([]byte(v), &strct.Image); err != nil {
                return err
             }
        case "Ports":
            if err := json.Unmarshal([]byte(v), &strct.Ports); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ContainerServiceDeployment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Containers" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Containers\": ")
	if tmp, err := json.Marshal(strct.Containers); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicEndpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicEndpoint\": ")
	if tmp, err := json.Marshal(strct.PublicEndpoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ContainerServiceDeployment) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Containers":
            if err := json.Unmarshal([]byte(v), &strct.Containers); err != nil {
                return err
             }
        case "PublicEndpoint":
            if err := json.Unmarshal([]byte(v), &strct.PublicEndpoint); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *EnvironmentVariable) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "Variable" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Variable\": ")
	if tmp, err := json.Marshal(strct.Variable); err != nil {
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
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
        case "Variable":
            if err := json.Unmarshal([]byte(v), &strct.Variable); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *HealthCheckConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "HealthyThreshold" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HealthyThreshold\": ")
	if tmp, err := json.Marshal(strct.HealthyThreshold); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IntervalSeconds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IntervalSeconds\": ")
	if tmp, err := json.Marshal(strct.IntervalSeconds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Path" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Path\": ")
	if tmp, err := json.Marshal(strct.Path); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SuccessCodes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SuccessCodes\": ")
	if tmp, err := json.Marshal(strct.SuccessCodes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TimeoutSeconds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TimeoutSeconds\": ")
	if tmp, err := json.Marshal(strct.TimeoutSeconds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UnhealthyThreshold" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UnhealthyThreshold\": ")
	if tmp, err := json.Marshal(strct.UnhealthyThreshold); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *HealthCheckConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "HealthyThreshold":
            if err := json.Unmarshal([]byte(v), &strct.HealthyThreshold); err != nil {
                return err
             }
        case "IntervalSeconds":
            if err := json.Unmarshal([]byte(v), &strct.IntervalSeconds); err != nil {
                return err
             }
        case "Path":
            if err := json.Unmarshal([]byte(v), &strct.Path); err != nil {
                return err
             }
        case "SuccessCodes":
            if err := json.Unmarshal([]byte(v), &strct.SuccessCodes); err != nil {
                return err
             }
        case "TimeoutSeconds":
            if err := json.Unmarshal([]byte(v), &strct.TimeoutSeconds); err != nil {
                return err
             }
        case "UnhealthyThreshold":
            if err := json.Unmarshal([]byte(v), &strct.UnhealthyThreshold); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *PortInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Port" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Port\": ")
	if tmp, err := json.Marshal(strct.Port); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PortInfo) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Port":
            if err := json.Unmarshal([]byte(v), &strct.Port); err != nil {
                return err
             }
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *PublicDomainName) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CertificateName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CertificateName\": ")
	if tmp, err := json.Marshal(strct.CertificateName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DomainNames" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DomainNames\": ")
	if tmp, err := json.Marshal(strct.DomainNames); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PublicDomainName) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CertificateName":
            if err := json.Unmarshal([]byte(v), &strct.CertificateName); err != nil {
                return err
             }
        case "DomainNames":
            if err := json.Unmarshal([]byte(v), &strct.DomainNames); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *PublicEndpoint) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ContainerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContainerName\": ")
	if tmp, err := json.Marshal(strct.ContainerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ContainerPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContainerPort\": ")
	if tmp, err := json.Marshal(strct.ContainerPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "HealthCheckConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"HealthCheckConfig\": ")
	if tmp, err := json.Marshal(strct.HealthCheckConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PublicEndpoint) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContainerName":
            if err := json.Unmarshal([]byte(v), &strct.ContainerName); err != nil {
                return err
             }
        case "ContainerPort":
            if err := json.Unmarshal([]byte(v), &strct.ContainerPort); err != nil {
                return err
             }
        case "HealthCheckConfig":
            if err := json.Unmarshal([]byte(v), &strct.HealthCheckConfig); err != nil {
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
    // Marshal the "ContainerArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContainerArn\": ")
	if tmp, err := json.Marshal(strct.ContainerArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ContainerServiceDeployment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ContainerServiceDeployment\": ")
	if tmp, err := json.Marshal(strct.ContainerServiceDeployment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IsDisabled" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IsDisabled\": ")
	if tmp, err := json.Marshal(strct.IsDisabled); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Power" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Power" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Power\": ")
	if tmp, err := json.Marshal(strct.Power); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicDomainNames" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicDomainNames\": ")
	if tmp, err := json.Marshal(strct.PublicDomainNames); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Scale" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Scale" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Scale\": ")
	if tmp, err := json.Marshal(strct.Scale); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ServiceName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ServiceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceName\": ")
	if tmp, err := json.Marshal(strct.ServiceName); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resource) UnmarshalJSON(b []byte) error {
    PowerReceived := false
    ScaleReceived := false
    ServiceNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ContainerArn":
            if err := json.Unmarshal([]byte(v), &strct.ContainerArn); err != nil {
                return err
             }
        case "ContainerServiceDeployment":
            if err := json.Unmarshal([]byte(v), &strct.ContainerServiceDeployment); err != nil {
                return err
             }
        case "IsDisabled":
            if err := json.Unmarshal([]byte(v), &strct.IsDisabled); err != nil {
                return err
             }
        case "Power":
            if err := json.Unmarshal([]byte(v), &strct.Power); err != nil {
                return err
             }
            PowerReceived = true
        case "PublicDomainNames":
            if err := json.Unmarshal([]byte(v), &strct.PublicDomainNames); err != nil {
                return err
             }
        case "Scale":
            if err := json.Unmarshal([]byte(v), &strct.Scale); err != nil {
                return err
             }
            ScaleReceived = true
        case "ServiceName":
            if err := json.Unmarshal([]byte(v), &strct.ServiceName); err != nil {
                return err
             }
            ServiceNameReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "Url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Power (a required property) was received
    if !PowerReceived {
        return errors.New("\"Power\" is required but was not present")
    }
    // check if Scale (a required property) was received
    if !ScaleReceived {
        return errors.New("\"Scale\" is required but was not present")
    }
    // check if ServiceName (a required property) was received
    if !ServiceNameReceived {
        return errors.New("\"ServiceName\" is required but was not present")
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