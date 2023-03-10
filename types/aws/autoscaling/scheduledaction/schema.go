// Code generated by schema-generate. DO NOT EDIT.

package scheduledaction

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Resource The AWS::AutoScaling::ScheduledAction resource specifies an Amazon EC2 Auto Scaling scheduled action so that the Auto Scaling group can change the number of instances available for your application in response to predictable load changes.
type Resource struct {

  // The name of the Auto Scaling group.
  AutoScalingGroupName string `json:"AutoScalingGroupName"`

  // The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
  DesiredCapacity int `json:"DesiredCapacity,omitempty"`

  // The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
  EndTime string `json:"EndTime,omitempty"`

  // The minimum size of the Auto Scaling group.
  MaxSize int `json:"MaxSize,omitempty"`

  // The minimum size of the Auto Scaling group.
  MinSize int `json:"MinSize,omitempty"`

  // The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.
  Recurrence string `json:"Recurrence,omitempty"`

  // Auto-generated unique identifier
  ScheduledActionName string `json:"ScheduledActionName,omitempty"`

  // The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
  StartTime string `json:"StartTime,omitempty"`

  // The time zone for the cron expression.
  TimeZone string `json:"TimeZone,omitempty"`
}

func (strct *Resource) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "AutoScalingGroupName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AutoScalingGroupName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AutoScalingGroupName\": ")
	if tmp, err := json.Marshal(strct.AutoScalingGroupName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DesiredCapacity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DesiredCapacity\": ")
	if tmp, err := json.Marshal(strct.DesiredCapacity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "EndTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"EndTime\": ")
	if tmp, err := json.Marshal(strct.EndTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MaxSize" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MaxSize\": ")
	if tmp, err := json.Marshal(strct.MaxSize); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "MinSize" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"MinSize\": ")
	if tmp, err := json.Marshal(strct.MinSize); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Recurrence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Recurrence\": ")
	if tmp, err := json.Marshal(strct.Recurrence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ScheduledActionName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ScheduledActionName\": ")
	if tmp, err := json.Marshal(strct.ScheduledActionName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StartTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StartTime\": ")
	if tmp, err := json.Marshal(strct.StartTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "TimeZone" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"TimeZone\": ")
	if tmp, err := json.Marshal(strct.TimeZone); err != nil {
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
    AutoScalingGroupNameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AutoScalingGroupName":
            if err := json.Unmarshal([]byte(v), &strct.AutoScalingGroupName); err != nil {
                return err
             }
            AutoScalingGroupNameReceived = true
        case "DesiredCapacity":
            if err := json.Unmarshal([]byte(v), &strct.DesiredCapacity); err != nil {
                return err
             }
        case "EndTime":
            if err := json.Unmarshal([]byte(v), &strct.EndTime); err != nil {
                return err
             }
        case "MaxSize":
            if err := json.Unmarshal([]byte(v), &strct.MaxSize); err != nil {
                return err
             }
        case "MinSize":
            if err := json.Unmarshal([]byte(v), &strct.MinSize); err != nil {
                return err
             }
        case "Recurrence":
            if err := json.Unmarshal([]byte(v), &strct.Recurrence); err != nil {
                return err
             }
        case "ScheduledActionName":
            if err := json.Unmarshal([]byte(v), &strct.ScheduledActionName); err != nil {
                return err
             }
        case "StartTime":
            if err := json.Unmarshal([]byte(v), &strct.StartTime); err != nil {
                return err
             }
        case "TimeZone":
            if err := json.Unmarshal([]byte(v), &strct.TimeZone); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AutoScalingGroupName (a required property) was received
    if !AutoScalingGroupNameReceived {
        return errors.New("\"AutoScalingGroupName\" is required but was not present")
    }
    return nil
}
