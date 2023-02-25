// Code generated by schema-generate. DO NOT EDIT.

package cluster

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// ClusterLogging The cluster control plane logging configuration for your cluster. 
type ClusterLogging struct {
  EnabledTypes []*LoggingTypeConfig `json:"EnabledTypes,omitempty"`
}

// ControlPlanePlacement Specify the placement group of the control plane machines for your cluster.
type ControlPlanePlacement struct {

  // Specify the placement group name of the control place machines for your cluster.
  GroupName string `json:"GroupName,omitempty"`
}

// EncryptionConfig The encryption configuration for the cluster
type EncryptionConfig struct {

  // The encryption provider for the cluster.
  Provider *Provider `json:"Provider,omitempty"`

  // Specifies the resources to be encrypted. The only supported value is "secrets".
  Resources []string `json:"Resources,omitempty"`
}

// KubernetesNetworkConfig The Kubernetes network configuration for the cluster.
type KubernetesNetworkConfig struct {

  // Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
  IpFamily string `json:"IpFamily,omitempty"`

  // The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
  ServiceIpv4Cidr string `json:"ServiceIpv4Cidr,omitempty"`

  // The CIDR block to assign Kubernetes service IP addresses from.
  ServiceIpv6Cidr string `json:"ServiceIpv6Cidr,omitempty"`
}

// Logging Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.
type Logging struct {

  // The cluster control plane logging configuration for your cluster. 
  ClusterLogging *ClusterLogging `json:"ClusterLogging,omitempty"`
}

// LoggingTypeConfig Enabled Logging Type
type LoggingTypeConfig struct {

  // name of the log type
  Type string `json:"Type,omitempty"`
}

// OutpostConfig An object representing the Outpost configuration to use for AWS EKS outpost cluster.
type OutpostConfig struct {

  // Specify the Instance type of the machines that should be used to create your cluster.
  ControlPlaneInstanceType string `json:"ControlPlaneInstanceType"`

  // Specify the placement group of the control plane machines for your cluster.
  ControlPlanePlacement *ControlPlanePlacement `json:"ControlPlanePlacement,omitempty"`

  // Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.
  OutpostArns []string `json:"OutpostArns"`
}

// Provider 
type Provider struct {

  // Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
  KeyArn string `json:"KeyArn,omitempty"`
}

// Resource An object representing an Amazon EKS cluster.
type Resource struct {

  // The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
  Arn string `json:"Arn,omitempty"`

  // The certificate-authority-data for your cluster.
  CertificateAuthorityData string `json:"CertificateAuthorityData,omitempty"`

  // The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
  ClusterSecurityGroupId string `json:"ClusterSecurityGroupId,omitempty"`
  EncryptionConfig []*EncryptionConfig `json:"EncryptionConfig,omitempty"`

  // Amazon Resource Name (ARN) or alias of the customer master key (CMK).
  EncryptionConfigKeyArn string `json:"EncryptionConfigKeyArn,omitempty"`

  // The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
  Endpoint string `json:"Endpoint,omitempty"`

  // The unique ID given to your cluster.
  Id string `json:"Id,omitempty"`
  KubernetesNetworkConfig *KubernetesNetworkConfig `json:"KubernetesNetworkConfig,omitempty"`
  Logging *Logging `json:"Logging,omitempty"`

  // The unique name to give to your cluster.
  Name string `json:"Name,omitempty"`

  // The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.
  OpenIdConnectIssuerUrl string `json:"OpenIdConnectIssuerUrl,omitempty"`
  OutpostConfig *OutpostConfig `json:"OutpostConfig,omitempty"`
  ResourcesVpcConfig *ResourcesVpcConfig `json:"ResourcesVpcConfig"`

  // The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
  RoleArn string `json:"RoleArn"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
  Version string `json:"Version,omitempty"`
}

// ResourcesVpcConfig An object representing the VPC configuration to use for an Amazon EKS cluster.
type ResourcesVpcConfig struct {

  // Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
  EndpointPrivateAccess bool `json:"EndpointPrivateAccess,omitempty"`

  // Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
  EndpointPublicAccess bool `json:"EndpointPublicAccess,omitempty"`

  // The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
  PublicAccessCidrs []string `json:"PublicAccessCidrs,omitempty"`

  // Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.
  SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`

  // Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
  SubnetIds []string `json:"SubnetIds"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value"`
}

func (strct *ClusterLogging) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EnabledTypes" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EnabledTypes\": ")
	if tmp, err := json.Marshal(strct.EnabledTypes); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ClusterLogging) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EnabledTypes":
            if err := json.Unmarshal([]byte(v), &strct.EnabledTypes); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *ControlPlanePlacement) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "GroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GroupName\": ")
	if tmp, err := json.Marshal(strct.GroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ControlPlanePlacement) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "GroupName":
            if err := json.Unmarshal([]byte(v), &strct.GroupName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *EncryptionConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Provider" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Provider\": ")
	if tmp, err := json.Marshal(strct.Provider); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Resources" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Resources\": ")
	if tmp, err := json.Marshal(strct.Resources); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EncryptionConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Provider":
            if err := json.Unmarshal([]byte(v), &strct.Provider); err != nil {
                return err
             }
        case "Resources":
            if err := json.Unmarshal([]byte(v), &strct.Resources); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *KubernetesNetworkConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "IpFamily" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IpFamily\": ")
	if tmp, err := json.Marshal(strct.IpFamily); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServiceIpv4Cidr" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceIpv4Cidr\": ")
	if tmp, err := json.Marshal(strct.ServiceIpv4Cidr); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ServiceIpv6Cidr" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ServiceIpv6Cidr\": ")
	if tmp, err := json.Marshal(strct.ServiceIpv6Cidr); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *KubernetesNetworkConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "IpFamily":
            if err := json.Unmarshal([]byte(v), &strct.IpFamily); err != nil {
                return err
             }
        case "ServiceIpv4Cidr":
            if err := json.Unmarshal([]byte(v), &strct.ServiceIpv4Cidr); err != nil {
                return err
             }
        case "ServiceIpv6Cidr":
            if err := json.Unmarshal([]byte(v), &strct.ServiceIpv6Cidr); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Logging) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ClusterLogging" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClusterLogging\": ")
	if tmp, err := json.Marshal(strct.ClusterLogging); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Logging) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ClusterLogging":
            if err := json.Unmarshal([]byte(v), &strct.ClusterLogging); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *LoggingTypeConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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

func (strct *LoggingTypeConfig) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *OutpostConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ControlPlaneInstanceType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "ControlPlaneInstanceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ControlPlaneInstanceType\": ")
	if tmp, err := json.Marshal(strct.ControlPlaneInstanceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ControlPlanePlacement" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ControlPlanePlacement\": ")
	if tmp, err := json.Marshal(strct.ControlPlanePlacement); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "OutpostArns" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "OutpostArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutpostArns\": ")
	if tmp, err := json.Marshal(strct.OutpostArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *OutpostConfig) UnmarshalJSON(b []byte) error {
    ControlPlaneInstanceTypeReceived := false
    OutpostArnsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ControlPlaneInstanceType":
            if err := json.Unmarshal([]byte(v), &strct.ControlPlaneInstanceType); err != nil {
                return err
             }
            ControlPlaneInstanceTypeReceived = true
        case "ControlPlanePlacement":
            if err := json.Unmarshal([]byte(v), &strct.ControlPlanePlacement); err != nil {
                return err
             }
        case "OutpostArns":
            if err := json.Unmarshal([]byte(v), &strct.OutpostArns); err != nil {
                return err
             }
            OutpostArnsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if ControlPlaneInstanceType (a required property) was received
    if !ControlPlaneInstanceTypeReceived {
        return errors.New("\"ControlPlaneInstanceType\" is required but was not present")
    }
    // check if OutpostArns (a required property) was received
    if !OutpostArnsReceived {
        return errors.New("\"OutpostArns\" is required but was not present")
    }
    return nil
}

func (strct *Provider) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "KeyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyArn\": ")
	if tmp, err := json.Marshal(strct.KeyArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Provider) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "KeyArn":
            if err := json.Unmarshal([]byte(v), &strct.KeyArn); err != nil {
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
    // Marshal the "CertificateAuthorityData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CertificateAuthorityData\": ")
	if tmp, err := json.Marshal(strct.CertificateAuthorityData); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ClusterSecurityGroupId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ClusterSecurityGroupId\": ")
	if tmp, err := json.Marshal(strct.ClusterSecurityGroupId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EncryptionConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionConfig\": ")
	if tmp, err := json.Marshal(strct.EncryptionConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EncryptionConfigKeyArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EncryptionConfigKeyArn\": ")
	if tmp, err := json.Marshal(strct.EncryptionConfigKeyArn); err != nil {
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
    // Marshal the "KubernetesNetworkConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KubernetesNetworkConfig\": ")
	if tmp, err := json.Marshal(strct.KubernetesNetworkConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Logging" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Logging\": ")
	if tmp, err := json.Marshal(strct.Logging); err != nil {
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
    // Marshal the "OpenIdConnectIssuerUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OpenIdConnectIssuerUrl\": ")
	if tmp, err := json.Marshal(strct.OpenIdConnectIssuerUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OutpostConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutpostConfig\": ")
	if tmp, err := json.Marshal(strct.OutpostConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourcesVpcConfig" field is required
    if strct.ResourcesVpcConfig == nil {
        return nil, errors.New("ResourcesVpcConfig is a required field")
    }
    // Marshal the "ResourcesVpcConfig" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourcesVpcConfig\": ")
	if tmp, err := json.Marshal(strct.ResourcesVpcConfig); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RoleArn" field is required
    // only required object types supported for marshal checking (for now)
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
    ResourcesVpcConfigReceived := false
    RoleArnReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "CertificateAuthorityData":
            if err := json.Unmarshal([]byte(v), &strct.CertificateAuthorityData); err != nil {
                return err
             }
        case "ClusterSecurityGroupId":
            if err := json.Unmarshal([]byte(v), &strct.ClusterSecurityGroupId); err != nil {
                return err
             }
        case "EncryptionConfig":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionConfig); err != nil {
                return err
             }
        case "EncryptionConfigKeyArn":
            if err := json.Unmarshal([]byte(v), &strct.EncryptionConfigKeyArn); err != nil {
                return err
             }
        case "Endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "KubernetesNetworkConfig":
            if err := json.Unmarshal([]byte(v), &strct.KubernetesNetworkConfig); err != nil {
                return err
             }
        case "Logging":
            if err := json.Unmarshal([]byte(v), &strct.Logging); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "OpenIdConnectIssuerUrl":
            if err := json.Unmarshal([]byte(v), &strct.OpenIdConnectIssuerUrl); err != nil {
                return err
             }
        case "OutpostConfig":
            if err := json.Unmarshal([]byte(v), &strct.OutpostConfig); err != nil {
                return err
             }
        case "ResourcesVpcConfig":
            if err := json.Unmarshal([]byte(v), &strct.ResourcesVpcConfig); err != nil {
                return err
             }
            ResourcesVpcConfigReceived = true
        case "RoleArn":
            if err := json.Unmarshal([]byte(v), &strct.RoleArn); err != nil {
                return err
             }
            RoleArnReceived = true
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
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
    // check if ResourcesVpcConfig (a required property) was received
    if !ResourcesVpcConfigReceived {
        return errors.New("\"ResourcesVpcConfig\" is required but was not present")
    }
    // check if RoleArn (a required property) was received
    if !RoleArnReceived {
        return errors.New("\"RoleArn\" is required but was not present")
    }
    return nil
}

func (strct *ResourcesVpcConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "EndpointPrivateAccess" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EndpointPrivateAccess\": ")
	if tmp, err := json.Marshal(strct.EndpointPrivateAccess); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EndpointPublicAccess" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EndpointPublicAccess\": ")
	if tmp, err := json.Marshal(strct.EndpointPublicAccess); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicAccessCidrs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicAccessCidrs\": ")
	if tmp, err := json.Marshal(strct.PublicAccessCidrs); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SecurityGroupIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SecurityGroupIds\": ")
	if tmp, err := json.Marshal(strct.SecurityGroupIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SubnetIds" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "SubnetIds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SubnetIds\": ")
	if tmp, err := json.Marshal(strct.SubnetIds); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResourcesVpcConfig) UnmarshalJSON(b []byte) error {
    SubnetIdsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "EndpointPrivateAccess":
            if err := json.Unmarshal([]byte(v), &strct.EndpointPrivateAccess); err != nil {
                return err
             }
        case "EndpointPublicAccess":
            if err := json.Unmarshal([]byte(v), &strct.EndpointPublicAccess); err != nil {
                return err
             }
        case "PublicAccessCidrs":
            if err := json.Unmarshal([]byte(v), &strct.PublicAccessCidrs); err != nil {
                return err
             }
        case "SecurityGroupIds":
            if err := json.Unmarshal([]byte(v), &strct.SecurityGroupIds); err != nil {
                return err
             }
        case "SubnetIds":
            if err := json.Unmarshal([]byte(v), &strct.SubnetIds); err != nil {
                return err
             }
            SubnetIdsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if SubnetIds (a required property) was received
    if !SubnetIdsReceived {
        return errors.New("\"SubnetIds\" is required but was not present")
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
