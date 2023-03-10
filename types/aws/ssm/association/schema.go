// Code generated by schema-generate. DO NOT EDIT.

package association

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// InstanceAssociationOutputLocation 
type InstanceAssociationOutputLocation struct {
  S3Location *S3OutputLocation `json:"S3Location,omitempty"`
}

// Parameters Parameter values that the SSM document uses at runtime.
type Parameters struct {
}

// Resource The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.
type Resource struct {
  ApplyOnlyAtCronInterval bool `json:"ApplyOnlyAtCronInterval,omitempty"`

  // Unique identifier of the association.
  AssociationId string `json:"AssociationId,omitempty"`

  // The name of the association.
  AssociationName string `json:"AssociationName,omitempty"`
  AutomationTargetParameterName string `json:"AutomationTargetParameterName,omitempty"`
  CalendarNames []string `json:"CalendarNames,omitempty"`
  ComplianceSeverity string `json:"ComplianceSeverity,omitempty"`

  // The version of the SSM document to associate with the target.
  DocumentVersion string `json:"DocumentVersion,omitempty"`

  // The ID of the instance that the SSM document is associated with.
  InstanceId string `json:"InstanceId,omitempty"`
  MaxConcurrency string `json:"MaxConcurrency,omitempty"`
  MaxErrors string `json:"MaxErrors,omitempty"`

  // The name of the SSM document.
  Name string `json:"Name"`
  OutputLocation *InstanceAssociationOutputLocation `json:"OutputLocation,omitempty"`

  // Parameter values that the SSM document uses at runtime.
  Parameters *Parameters `json:"Parameters,omitempty"`

  // A Cron or Rate expression that specifies when the association is applied to the target.
  ScheduleExpression string `json:"ScheduleExpression,omitempty"`
  ScheduleOffset int `json:"ScheduleOffset,omitempty"`
  SyncCompliance string `json:"SyncCompliance,omitempty"`

  // The targets that the SSM document sends commands to.
  Targets []*Target `json:"Targets,omitempty"`
  WaitForSuccessTimeoutSeconds int `json:"WaitForSuccessTimeoutSeconds,omitempty"`
}

// S3OutputLocation 
type S3OutputLocation struct {
  OutputS3BucketName string `json:"OutputS3BucketName,omitempty"`
  OutputS3KeyPrefix string `json:"OutputS3KeyPrefix,omitempty"`
  OutputS3Region string `json:"OutputS3Region,omitempty"`
}

// Target 
type Target struct {
  Key string `json:"Key"`
  Values []string `json:"Values"`
}

func (strct *InstanceAssociationOutputLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "S3Location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"S3Location\": ")
	if tmp, err := json.Marshal(strct.S3Location); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *InstanceAssociationOutputLocation) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "S3Location":
            if err := json.Unmarshal([]byte(v), &strct.S3Location); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Parameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Parameters) UnmarshalJSON(b []byte) error {
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

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ApplyOnlyAtCronInterval" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ApplyOnlyAtCronInterval\": ")
	if tmp, err := json.Marshal(strct.ApplyOnlyAtCronInterval); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AssociationId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssociationId\": ")
	if tmp, err := json.Marshal(strct.AssociationId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AssociationName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AssociationName\": ")
	if tmp, err := json.Marshal(strct.AssociationName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "AutomationTargetParameterName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutomationTargetParameterName\": ")
	if tmp, err := json.Marshal(strct.AutomationTargetParameterName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CalendarNames" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CalendarNames\": ")
	if tmp, err := json.Marshal(strct.CalendarNames); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ComplianceSeverity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ComplianceSeverity\": ")
	if tmp, err := json.Marshal(strct.ComplianceSeverity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DocumentVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DocumentVersion\": ")
	if tmp, err := json.Marshal(strct.DocumentVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "InstanceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InstanceId\": ")
	if tmp, err := json.Marshal(strct.InstanceId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaxConcurrency" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxConcurrency\": ")
	if tmp, err := json.Marshal(strct.MaxConcurrency); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaxErrors" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxErrors\": ")
	if tmp, err := json.Marshal(strct.MaxErrors); err != nil {
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
    // Marshal the "OutputLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputLocation\": ")
	if tmp, err := json.Marshal(strct.OutputLocation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Parameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ScheduleExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScheduleExpression\": ")
	if tmp, err := json.Marshal(strct.ScheduleExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ScheduleOffset" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScheduleOffset\": ")
	if tmp, err := json.Marshal(strct.ScheduleOffset); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "SyncCompliance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SyncCompliance\": ")
	if tmp, err := json.Marshal(strct.SyncCompliance); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Targets" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Targets\": ")
	if tmp, err := json.Marshal(strct.Targets); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "WaitForSuccessTimeoutSeconds" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"WaitForSuccessTimeoutSeconds\": ")
	if tmp, err := json.Marshal(strct.WaitForSuccessTimeoutSeconds); err != nil {
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
        case "ApplyOnlyAtCronInterval":
            if err := json.Unmarshal([]byte(v), &strct.ApplyOnlyAtCronInterval); err != nil {
                return err
             }
        case "AssociationId":
            if err := json.Unmarshal([]byte(v), &strct.AssociationId); err != nil {
                return err
             }
        case "AssociationName":
            if err := json.Unmarshal([]byte(v), &strct.AssociationName); err != nil {
                return err
             }
        case "AutomationTargetParameterName":
            if err := json.Unmarshal([]byte(v), &strct.AutomationTargetParameterName); err != nil {
                return err
             }
        case "CalendarNames":
            if err := json.Unmarshal([]byte(v), &strct.CalendarNames); err != nil {
                return err
             }
        case "ComplianceSeverity":
            if err := json.Unmarshal([]byte(v), &strct.ComplianceSeverity); err != nil {
                return err
             }
        case "DocumentVersion":
            if err := json.Unmarshal([]byte(v), &strct.DocumentVersion); err != nil {
                return err
             }
        case "InstanceId":
            if err := json.Unmarshal([]byte(v), &strct.InstanceId); err != nil {
                return err
             }
        case "MaxConcurrency":
            if err := json.Unmarshal([]byte(v), &strct.MaxConcurrency); err != nil {
                return err
             }
        case "MaxErrors":
            if err := json.Unmarshal([]byte(v), &strct.MaxErrors); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "OutputLocation":
            if err := json.Unmarshal([]byte(v), &strct.OutputLocation); err != nil {
                return err
             }
        case "Parameters":
            if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
                return err
             }
        case "ScheduleExpression":
            if err := json.Unmarshal([]byte(v), &strct.ScheduleExpression); err != nil {
                return err
             }
        case "ScheduleOffset":
            if err := json.Unmarshal([]byte(v), &strct.ScheduleOffset); err != nil {
                return err
             }
        case "SyncCompliance":
            if err := json.Unmarshal([]byte(v), &strct.SyncCompliance); err != nil {
                return err
             }
        case "Targets":
            if err := json.Unmarshal([]byte(v), &strct.Targets); err != nil {
                return err
             }
        case "WaitForSuccessTimeoutSeconds":
            if err := json.Unmarshal([]byte(v), &strct.WaitForSuccessTimeoutSeconds); err != nil {
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

func (strct *S3OutputLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "OutputS3BucketName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputS3BucketName\": ")
	if tmp, err := json.Marshal(strct.OutputS3BucketName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OutputS3KeyPrefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputS3KeyPrefix\": ")
	if tmp, err := json.Marshal(strct.OutputS3KeyPrefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "OutputS3Region" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"OutputS3Region\": ")
	if tmp, err := json.Marshal(strct.OutputS3Region); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *S3OutputLocation) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "OutputS3BucketName":
            if err := json.Unmarshal([]byte(v), &strct.OutputS3BucketName); err != nil {
                return err
             }
        case "OutputS3KeyPrefix":
            if err := json.Unmarshal([]byte(v), &strct.OutputS3KeyPrefix); err != nil {
                return err
             }
        case "OutputS3Region":
            if err := json.Unmarshal([]byte(v), &strct.OutputS3Region); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *Target) MarshalJSON() ([]byte, error) {
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
    // "Values" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Values" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Values\": ")
	if tmp, err := json.Marshal(strct.Values); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Target) UnmarshalJSON(b []byte) error {
    KeyReceived := false
    ValuesReceived := false
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
        case "Values":
            if err := json.Unmarshal([]byte(v), &strct.Values); err != nil {
                return err
             }
            ValuesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Key (a required property) was received
    if !KeyReceived {
        return errors.New("\"Key\" is required but was not present")
    }
    // check if Values (a required property) was received
    if !ValuesReceived {
        return errors.New("\"Values\" is required but was not present")
    }
    return nil
}
