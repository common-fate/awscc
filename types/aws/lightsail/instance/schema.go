// Code generated by schema-generate. DO NOT EDIT.

package instance

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// AddOn A addon associate with a resource.
type AddOn struct {

  // The add-on type
  AddOnType string `json:"AddOnType"`
  AutoSnapshotAddOnRequest *AutoSnapshotAddOn `json:"AutoSnapshotAddOnRequest,omitempty"`

  // Status of the Addon
  Status string `json:"Status,omitempty"`
}

// AutoSnapshotAddOn An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
type AutoSnapshotAddOn struct {

  // The daily time when an automatic snapshot will be created.
  SnapshotTimeOfDay string `json:"SnapshotTimeOfDay,omitempty"`
}

// Disk Disk associated with the Instance.
type Disk struct {

  // Instance attached to the disk.
  AttachedTo string `json:"AttachedTo,omitempty"`

  // Attachment state of the disk.
  AttachmentState string `json:"AttachmentState,omitempty"`

  // The names to use for your new Lightsail disk.
  DiskName string `json:"DiskName"`

  // IOPS of disk.
  IOPS int `json:"IOPS,omitempty"`

  // Is the Attached disk is the system disk of the Instance.
  IsSystemDisk bool `json:"IsSystemDisk,omitempty"`

  // Path of the disk attached to the instance.
  Path string `json:"Path"`

  // Size of the disk attached to the Instance.
  SizeInGb string `json:"SizeInGb,omitempty"`
}

// Hardware Hardware of the Instance.
type Hardware struct {

  // CPU count of the Instance.
  CpuCount int `json:"CpuCount,omitempty"`

  // Disks attached to the Instance.
  Disks []*Disk `json:"Disks,omitempty"`

  // RAM Size of the Instance.
  RamSizeInGb int `json:"RamSizeInGb,omitempty"`
}

// Location Location of a resource.
type Location struct {

  // The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
  AvailabilityZone string `json:"AvailabilityZone,omitempty"`

  // The Region Name in which to create your instance.
  RegionName string `json:"RegionName,omitempty"`
}

// MonthlyTransfer Monthly Transfer of the Instance.
type MonthlyTransfer struct {

  // GbPerMonthAllocated of the Instance.
  GbPerMonthAllocated string `json:"GbPerMonthAllocated,omitempty"`
}

// Networking Networking of the Instance.
type Networking struct {
  MonthlyTransfer *MonthlyTransfer `json:"MonthlyTransfer,omitempty"`

  // Ports to the Instance.
  Ports []*Port `json:"Ports"`
}

// Port Port of the Instance.
type Port struct {

  // Access Direction for Protocol of the Instance(inbound/outbound).
  AccessDirection string `json:"AccessDirection,omitempty"`

  // Access From Protocol of the Instance.
  AccessFrom string `json:"AccessFrom,omitempty"`

  // Access Type Protocol of the Instance.
  AccessType string `json:"AccessType,omitempty"`
  CidrListAliases []string `json:"CidrListAliases,omitempty"`
  Cidrs []string `json:"Cidrs,omitempty"`

  // CommonName for Protocol of the Instance.
  CommonName string `json:"CommonName,omitempty"`

  // From Port of the Instance.
  FromPort int `json:"FromPort,omitempty"`
  Ipv6Cidrs []string `json:"Ipv6Cidrs,omitempty"`

  // Port Protocol of the Instance.
  Protocol string `json:"Protocol,omitempty"`

  // To Port of the Instance.
  ToPort int `json:"ToPort,omitempty"`
}

// Resource Resource Type definition for AWS::Lightsail::Instance
type Resource struct {

  // An array of objects representing the add-ons to enable for the new instance.
  AddOns []*AddOn `json:"AddOns,omitempty"`

  // The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
  AvailabilityZone string `json:"AvailabilityZone,omitempty"`

  // The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).
  BlueprintId string `json:"BlueprintId"`

  // The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).
  BundleId string `json:"BundleId"`
  Hardware *Hardware `json:"Hardware,omitempty"`
  InstanceArn string `json:"InstanceArn,omitempty"`

  // The names to use for your new Lightsail instance.
  InstanceName string `json:"InstanceName"`

  // Is the IP Address of the Instance is the static IP
  IsStaticIp bool `json:"IsStaticIp,omitempty"`

  // The name of your key pair.
  KeyPairName string `json:"KeyPairName,omitempty"`
  Location *Location `json:"Location,omitempty"`
  Networking *Networking `json:"Networking,omitempty"`

  // Private IP Address of the Instance
  PrivateIpAddress string `json:"PrivateIpAddress,omitempty"`

  // Public IP Address of the Instance
  PublicIpAddress string `json:"PublicIpAddress,omitempty"`

  // Resource type of Lightsail instance.
  ResourceType string `json:"ResourceType,omitempty"`

  // SSH Key Name of the  Lightsail instance.
  SshKeyName string `json:"SshKeyName,omitempty"`
  State *State `json:"State,omitempty"`

  // Support code to help identify any issues
  SupportCode string `json:"SupportCode,omitempty"`

  // An array of key-value pairs to apply to this resource.
  Tags []*Tag `json:"Tags,omitempty"`

  // A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update.
  UserData string `json:"UserData,omitempty"`

  // Username of the  Lightsail instance.
  UserName string `json:"UserName,omitempty"`
}

// State Current State of the Instance.
type State struct {

  // Status code of the Instance.
  Code int `json:"Code,omitempty"`

  // Status code of the Instance.
  Name string `json:"Name,omitempty"`
}

// Tag A key-value pair to associate with a resource.
type Tag struct {

  // The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Key string `json:"Key"`

  // The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
  Value string `json:"Value,omitempty"`
}

func (strct *AddOn) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AddOnType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AddOnType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AddOnType\": ")
	if tmp, err := json.Marshal(strct.AddOnType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AutoSnapshotAddOnRequest" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoSnapshotAddOnRequest\": ")
	if tmp, err := json.Marshal(strct.AutoSnapshotAddOnRequest); err != nil {
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

func (strct *AddOn) UnmarshalJSON(b []byte) error {
    AddOnTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AddOnType":
            if err := json.Unmarshal([]byte(v), &strct.AddOnType); err != nil {
                return err
             }
            AddOnTypeReceived = true
        case "AutoSnapshotAddOnRequest":
            if err := json.Unmarshal([]byte(v), &strct.AutoSnapshotAddOnRequest); err != nil {
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
    // check if AddOnType (a required property) was received
    if !AddOnTypeReceived {
        return errors.New("\"AddOnType\" is required but was not present")
    }
    return nil
}

func (strct *AutoSnapshotAddOn) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "SnapshotTimeOfDay" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SnapshotTimeOfDay\": ")
	if tmp, err := json.Marshal(strct.SnapshotTimeOfDay); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AutoSnapshotAddOn) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "SnapshotTimeOfDay":
            if err := json.Unmarshal([]byte(v), &strct.SnapshotTimeOfDay); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Disk) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AttachedTo" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachedTo\": ")
	if tmp, err := json.Marshal(strct.AttachedTo); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AttachmentState" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AttachmentState\": ")
	if tmp, err := json.Marshal(strct.AttachmentState); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DiskName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DiskName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DiskName\": ")
	if tmp, err := json.Marshal(strct.DiskName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IOPS" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IOPS\": ")
	if tmp, err := json.Marshal(strct.IOPS); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IsSystemDisk" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IsSystemDisk\": ")
	if tmp, err := json.Marshal(strct.IsSystemDisk); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Path" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "SizeInGb" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SizeInGb\": ")
	if tmp, err := json.Marshal(strct.SizeInGb); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Disk) UnmarshalJSON(b []byte) error {
    DiskNameReceived := false
    PathReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AttachedTo":
            if err := json.Unmarshal([]byte(v), &strct.AttachedTo); err != nil {
                return err
             }
        case "AttachmentState":
            if err := json.Unmarshal([]byte(v), &strct.AttachmentState); err != nil {
                return err
             }
        case "DiskName":
            if err := json.Unmarshal([]byte(v), &strct.DiskName); err != nil {
                return err
             }
            DiskNameReceived = true
        case "IOPS":
            if err := json.Unmarshal([]byte(v), &strct.IOPS); err != nil {
                return err
             }
        case "IsSystemDisk":
            if err := json.Unmarshal([]byte(v), &strct.IsSystemDisk); err != nil {
                return err
             }
        case "Path":
            if err := json.Unmarshal([]byte(v), &strct.Path); err != nil {
                return err
             }
            PathReceived = true
        case "SizeInGb":
            if err := json.Unmarshal([]byte(v), &strct.SizeInGb); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DiskName (a required property) was received
    if !DiskNameReceived {
        return errors.New("\"DiskName\" is required but was not present")
    }
    // check if Path (a required property) was received
    if !PathReceived {
        return errors.New("\"Path\" is required but was not present")
    }
    return nil
}

func (strct *Hardware) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "CpuCount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CpuCount\": ")
	if tmp, err := json.Marshal(strct.CpuCount); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Disks" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Disks\": ")
	if tmp, err := json.Marshal(strct.Disks); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RamSizeInGb" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RamSizeInGb\": ")
	if tmp, err := json.Marshal(strct.RamSizeInGb); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Hardware) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CpuCount":
            if err := json.Unmarshal([]byte(v), &strct.CpuCount); err != nil {
                return err
             }
        case "Disks":
            if err := json.Unmarshal([]byte(v), &strct.Disks); err != nil {
                return err
             }
        case "RamSizeInGb":
            if err := json.Unmarshal([]byte(v), &strct.RamSizeInGb); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Location) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AvailabilityZone" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AvailabilityZone\": ")
	if tmp, err := json.Marshal(strct.AvailabilityZone); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "RegionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegionName\": ")
	if tmp, err := json.Marshal(strct.RegionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Location) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AvailabilityZone":
            if err := json.Unmarshal([]byte(v), &strct.AvailabilityZone); err != nil {
                return err
             }
        case "RegionName":
            if err := json.Unmarshal([]byte(v), &strct.RegionName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *MonthlyTransfer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "GbPerMonthAllocated" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"GbPerMonthAllocated\": ")
	if tmp, err := json.Marshal(strct.GbPerMonthAllocated); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MonthlyTransfer) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "GbPerMonthAllocated":
            if err := json.Unmarshal([]byte(v), &strct.GbPerMonthAllocated); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Networking) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "MonthlyTransfer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MonthlyTransfer\": ")
	if tmp, err := json.Marshal(strct.MonthlyTransfer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Ports" field is required
    // only required object types supported for marshal checking (for now)
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

func (strct *Networking) UnmarshalJSON(b []byte) error {
    PortsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "MonthlyTransfer":
            if err := json.Unmarshal([]byte(v), &strct.MonthlyTransfer); err != nil {
                return err
             }
        case "Ports":
            if err := json.Unmarshal([]byte(v), &strct.Ports); err != nil {
                return err
             }
            PortsReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Ports (a required property) was received
    if !PortsReceived {
        return errors.New("\"Ports\" is required but was not present")
    }
    return nil
}

func (strct *Port) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "AccessDirection" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessDirection\": ")
	if tmp, err := json.Marshal(strct.AccessDirection); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AccessFrom" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessFrom\": ")
	if tmp, err := json.Marshal(strct.AccessFrom); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AccessType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AccessType\": ")
	if tmp, err := json.Marshal(strct.AccessType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CidrListAliases" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CidrListAliases\": ")
	if tmp, err := json.Marshal(strct.CidrListAliases); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Cidrs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Cidrs\": ")
	if tmp, err := json.Marshal(strct.Cidrs); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CommonName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CommonName\": ")
	if tmp, err := json.Marshal(strct.CommonName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "FromPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"FromPort\": ")
	if tmp, err := json.Marshal(strct.FromPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Ipv6Cidrs" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Ipv6Cidrs\": ")
	if tmp, err := json.Marshal(strct.Ipv6Cidrs); err != nil {
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
    // Marshal the "ToPort" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ToPort\": ")
	if tmp, err := json.Marshal(strct.ToPort); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Port) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AccessDirection":
            if err := json.Unmarshal([]byte(v), &strct.AccessDirection); err != nil {
                return err
             }
        case "AccessFrom":
            if err := json.Unmarshal([]byte(v), &strct.AccessFrom); err != nil {
                return err
             }
        case "AccessType":
            if err := json.Unmarshal([]byte(v), &strct.AccessType); err != nil {
                return err
             }
        case "CidrListAliases":
            if err := json.Unmarshal([]byte(v), &strct.CidrListAliases); err != nil {
                return err
             }
        case "Cidrs":
            if err := json.Unmarshal([]byte(v), &strct.Cidrs); err != nil {
                return err
             }
        case "CommonName":
            if err := json.Unmarshal([]byte(v), &strct.CommonName); err != nil {
                return err
             }
        case "FromPort":
            if err := json.Unmarshal([]byte(v), &strct.FromPort); err != nil {
                return err
             }
        case "Ipv6Cidrs":
            if err := json.Unmarshal([]byte(v), &strct.Ipv6Cidrs); err != nil {
                return err
             }
        case "Protocol":
            if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
                return err
             }
        case "ToPort":
            if err := json.Unmarshal([]byte(v), &strct.ToPort); err != nil {
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
    // Marshal the "AddOns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AddOns\": ")
	if tmp, err := json.Marshal(strct.AddOns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AvailabilityZone" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AvailabilityZone\": ")
	if tmp, err := json.Marshal(strct.AvailabilityZone); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "BlueprintId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BlueprintId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BlueprintId\": ")
	if tmp, err := json.Marshal(strct.BlueprintId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "BundleId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BundleId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BundleId\": ")
	if tmp, err := json.Marshal(strct.BundleId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Hardware" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Hardware\": ")
	if tmp, err := json.Marshal(strct.Hardware); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InstanceArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceArn\": ")
	if tmp, err := json.Marshal(strct.InstanceArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "InstanceName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "InstanceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceName\": ")
	if tmp, err := json.Marshal(strct.InstanceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IsStaticIp" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IsStaticIp\": ")
	if tmp, err := json.Marshal(strct.IsStaticIp); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "KeyPairName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"KeyPairName\": ")
	if tmp, err := json.Marshal(strct.KeyPairName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Networking" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Networking\": ")
	if tmp, err := json.Marshal(strct.Networking); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PrivateIpAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PrivateIpAddress\": ")
	if tmp, err := json.Marshal(strct.PrivateIpAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "PublicIpAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"PublicIpAddress\": ")
	if tmp, err := json.Marshal(strct.PublicIpAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ResourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ResourceType\": ")
	if tmp, err := json.Marshal(strct.ResourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SshKeyName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SshKeyName\": ")
	if tmp, err := json.Marshal(strct.SshKeyName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "State" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"State\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SupportCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SupportCode\": ")
	if tmp, err := json.Marshal(strct.SupportCode); err != nil {
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
    // Marshal the "UserData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UserData\": ")
	if tmp, err := json.Marshal(strct.UserData); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "UserName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"UserName\": ")
	if tmp, err := json.Marshal(strct.UserName); err != nil {
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
    BlueprintIdReceived := false
    BundleIdReceived := false
    InstanceNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AddOns":
            if err := json.Unmarshal([]byte(v), &strct.AddOns); err != nil {
                return err
             }
        case "AvailabilityZone":
            if err := json.Unmarshal([]byte(v), &strct.AvailabilityZone); err != nil {
                return err
             }
        case "BlueprintId":
            if err := json.Unmarshal([]byte(v), &strct.BlueprintId); err != nil {
                return err
             }
            BlueprintIdReceived = true
        case "BundleId":
            if err := json.Unmarshal([]byte(v), &strct.BundleId); err != nil {
                return err
             }
            BundleIdReceived = true
        case "Hardware":
            if err := json.Unmarshal([]byte(v), &strct.Hardware); err != nil {
                return err
             }
        case "InstanceArn":
            if err := json.Unmarshal([]byte(v), &strct.InstanceArn); err != nil {
                return err
             }
        case "InstanceName":
            if err := json.Unmarshal([]byte(v), &strct.InstanceName); err != nil {
                return err
             }
            InstanceNameReceived = true
        case "IsStaticIp":
            if err := json.Unmarshal([]byte(v), &strct.IsStaticIp); err != nil {
                return err
             }
        case "KeyPairName":
            if err := json.Unmarshal([]byte(v), &strct.KeyPairName); err != nil {
                return err
             }
        case "Location":
            if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
                return err
             }
        case "Networking":
            if err := json.Unmarshal([]byte(v), &strct.Networking); err != nil {
                return err
             }
        case "PrivateIpAddress":
            if err := json.Unmarshal([]byte(v), &strct.PrivateIpAddress); err != nil {
                return err
             }
        case "PublicIpAddress":
            if err := json.Unmarshal([]byte(v), &strct.PublicIpAddress); err != nil {
                return err
             }
        case "ResourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
        case "SshKeyName":
            if err := json.Unmarshal([]byte(v), &strct.SshKeyName); err != nil {
                return err
             }
        case "State":
            if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
                return err
             }
        case "SupportCode":
            if err := json.Unmarshal([]byte(v), &strct.SupportCode); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "UserData":
            if err := json.Unmarshal([]byte(v), &strct.UserData); err != nil {
                return err
             }
        case "UserName":
            if err := json.Unmarshal([]byte(v), &strct.UserName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if BlueprintId (a required property) was received
    if !BlueprintIdReceived {
        return errors.New("\"BlueprintId\" is required but was not present")
    }
    // check if BundleId (a required property) was received
    if !BundleIdReceived {
        return errors.New("\"BundleId\" is required but was not present")
    }
    // check if InstanceName (a required property) was received
    if !InstanceNameReceived {
        return errors.New("\"InstanceName\" is required but was not present")
    }
    return nil
}

func (strct *State) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Code" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Code\": ")
	if tmp, err := json.Marshal(strct.Code); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *State) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
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
