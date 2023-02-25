// Code generated by schema-generate. DO NOT EDIT.

package analysis

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AnalysisError <p>A metadata error structure for an analysis.</p>
type AnalysisError struct {

  // <p>The message associated with the analysis error.</p>
  Message string `json:"Message,omitempty"`
  Type string `json:"Type,omitempty"`
}

// AnalysisSourceEntity <p>The source entity of an analysis.</p>
type AnalysisSourceEntity struct {
  SourceTemplate *AnalysisSourceTemplate `json:"SourceTemplate,omitempty"`
}

// AnalysisSourceTemplate <p>The source template of an analysis.</p>
type AnalysisSourceTemplate struct {

  // <p>The Amazon Resource Name (ARN) of the source template of an analysis.</p>
  Arn string `json:"Arn"`

  // <p>The dataset references of the source template of an analysis.</p>
  DataSetReferences []*DataSetReference `json:"DataSetReferences"`
}

// DataSetReference <p>Dataset reference.</p>
type DataSetReference struct {

  // <p>Dataset Amazon Resource Name (ARN).</p>
  DataSetArn string `json:"DataSetArn"`

  // <p>Dataset placeholder.</p>
  DataSetPlaceholder string `json:"DataSetPlaceholder"`
}

// DateTimeParameter <p>A date-time parameter.</p>
type DateTimeParameter struct {

  // <p>A display name for the date-time parameter.</p>
  Name string `json:"Name"`

  // <p>The values for the date-time parameter.</p>
  Values []string `json:"Values"`
}

// DecimalParameter <p>A decimal parameter.</p>
type DecimalParameter struct {

  // <p>A display name for the decimal parameter.</p>
  Name string `json:"Name"`

  // <p>The values for the decimal parameter.</p>
  Values []float64 `json:"Values"`
}

// IntegerParameter <p>An integer parameter.</p>
type IntegerParameter struct {

  // <p>The name of the integer parameter.</p>
  Name string `json:"Name"`

  // <p>The values for the integer parameter.</p>
  Values []float64 `json:"Values"`
}

// Parameters <p>A list of QuickSight parameters and the list's override values.</p>
type Parameters struct {

  // <p>Date-time parameters.</p>
  DateTimeParameters []*DateTimeParameter `json:"DateTimeParameters,omitempty"`

  // <p>Decimal parameters.</p>
  DecimalParameters []*DecimalParameter `json:"DecimalParameters,omitempty"`

  // <p>Integer parameters.</p>
  IntegerParameters []*IntegerParameter `json:"IntegerParameters,omitempty"`

  // <p>String parameters.</p>
  StringParameters []*StringParameter `json:"StringParameters,omitempty"`
}

// Resource Definition of the AWS::QuickSight::Analysis Resource Type.
type Resource struct {
  AnalysisId string `json:"AnalysisId"`

  // <p>The Amazon Resource Name (ARN) of the analysis.</p>
  Arn string `json:"Arn,omitempty"`
  AwsAccountId string `json:"AwsAccountId"`

  // <p>The time that the analysis was created.</p>
  CreatedTime string `json:"CreatedTime,omitempty"`

  // <p>The ARNs of the datasets of the analysis.</p>
  DataSetArns []string `json:"DataSetArns,omitempty"`

  // <p>Errors associated with the analysis.</p>
  Errors []*AnalysisError `json:"Errors,omitempty"`

  // <p>The time that the analysis was last updated.</p>
  LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`

  // <p>The descriptive name of the analysis.</p>
  Name string `json:"Name,omitempty"`
  Parameters *Parameters `json:"Parameters,omitempty"`

  // <p>A structure that describes the principals and the resource-level permissions on an
  //             analysis. You can use the <code>Permissions</code> structure to grant permissions by
  //             providing a list of AWS Identity and Access Management (IAM) action information for each
  //             principal listed by Amazon Resource Name (ARN). </p>
  // 
  //         <p>To specify no permissions, omit <code>Permissions</code>.</p>
  Permissions []*ResourcePermission `json:"Permissions,omitempty"`

  // <p>A list of the associated sheets with the unique identifier and name of each sheet.</p>
  Sheets []*Sheet `json:"Sheets,omitempty"`
  SourceEntity *AnalysisSourceEntity `json:"SourceEntity"`
  Status string `json:"Status,omitempty"`

  // <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
  //             analysis.</p>
  Tags []*Tag `json:"Tags,omitempty"`

  // <p>The ARN of the theme of the analysis.</p>
  ThemeArn string `json:"ThemeArn,omitempty"`
}

// ResourcePermission <p>Permission for the resource.</p>
type ResourcePermission struct {

  // <p>The IAM action to grant or revoke permissions on.</p>
  Actions []string `json:"Actions"`

  // <p>The Amazon Resource Name (ARN) of the principal. This can be one of the
  //             following:</p>
  //         <ul>
  //             <li>
  //                 <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
  //             </li>
  //             <li>
  //                 <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
  //             </li>
  //             <li>
  //                 <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight
  //                     ARN. Use this option only to share resources (templates) across AWS accounts.
  //                     (This is less common.) </p>
  //             </li>
  //          </ul>
  Principal string `json:"Principal"`
}

// Sheet <p>A <i>sheet</i>, which is an object that contains a set of visuals that
//             are viewed together on one page in the Amazon QuickSight console. Every analysis and dashboard
//             contains at least one sheet. Each sheet contains at least one visualization widget, for
//             example a chart, pivot table, or narrative insight. Sheets can be associated with other
//             components, such as controls, filters, and so on.</p>
type Sheet struct {

  // <p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight
  //             console.</p>
  Name string `json:"Name,omitempty"`

  // <p>The unique identifier associated with a sheet.</p>
  SheetId string `json:"SheetId,omitempty"`
}

// StringParameter <p>A string parameter.</p>
type StringParameter struct {

  // <p>A display name for a string parameter.</p>
  Name string `json:"Name"`

  // <p>The values of a string parameter.</p>
  Values []string `json:"Values"`
}

// Tag <p>The key or keys of the key-value pairs for the resource tag or tags assigned to the
//             resource.</p>
type Tag struct {

  // <p>Tag key.</p>
  Key string `json:"Key"`

  // <p>Tag value.</p>
  Value string `json:"Value"`
}

func (strct *AnalysisError) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "Message" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

func (strct *AnalysisError) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Message":
            if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
                return err
             }
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

func (strct *AnalysisSourceEntity) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "SourceTemplate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceTemplate\": ")
	if tmp, err := json.Marshal(strct.SourceTemplate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AnalysisSourceEntity) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "SourceTemplate":
            if err := json.Unmarshal([]byte(v), &strct.SourceTemplate); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *AnalysisSourceTemplate) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Arn" field is required
    // only required object types supported for marshal checking (for now)
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
    // "DataSetReferences" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DataSetReferences" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataSetReferences\": ")
	if tmp, err := json.Marshal(strct.DataSetReferences); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AnalysisSourceTemplate) UnmarshalJSON(b []byte) error {
    ArnReceived := false
    DataSetReferencesReceived := false
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
            ArnReceived = true
        case "DataSetReferences":
            if err := json.Unmarshal([]byte(v), &strct.DataSetReferences); err != nil {
                return err
             }
            DataSetReferencesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Arn (a required property) was received
    if !ArnReceived {
        return errors.New("\"Arn\" is required but was not present")
    }
    // check if DataSetReferences (a required property) was received
    if !DataSetReferencesReceived {
        return errors.New("\"DataSetReferences\" is required but was not present")
    }
    return nil
}

func (strct *DataSetReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "DataSetArn" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DataSetArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataSetArn\": ")
	if tmp, err := json.Marshal(strct.DataSetArn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "DataSetPlaceholder" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "DataSetPlaceholder" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataSetPlaceholder\": ")
	if tmp, err := json.Marshal(strct.DataSetPlaceholder); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DataSetReference) UnmarshalJSON(b []byte) error {
    DataSetArnReceived := false
    DataSetPlaceholderReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "DataSetArn":
            if err := json.Unmarshal([]byte(v), &strct.DataSetArn); err != nil {
                return err
             }
            DataSetArnReceived = true
        case "DataSetPlaceholder":
            if err := json.Unmarshal([]byte(v), &strct.DataSetPlaceholder); err != nil {
                return err
             }
            DataSetPlaceholderReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if DataSetArn (a required property) was received
    if !DataSetArnReceived {
        return errors.New("\"DataSetArn\" is required but was not present")
    }
    // check if DataSetPlaceholder (a required property) was received
    if !DataSetPlaceholderReceived {
        return errors.New("\"DataSetPlaceholder\" is required but was not present")
    }
    return nil
}

func (strct *DateTimeParameter) MarshalJSON() ([]byte, error) {
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

func (strct *DateTimeParameter) UnmarshalJSON(b []byte) error {
    NameReceived := false
    ValuesReceived := false
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
        case "Values":
            if err := json.Unmarshal([]byte(v), &strct.Values); err != nil {
                return err
             }
            ValuesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Values (a required property) was received
    if !ValuesReceived {
        return errors.New("\"Values\" is required but was not present")
    }
    return nil
}

func (strct *DecimalParameter) MarshalJSON() ([]byte, error) {
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

func (strct *DecimalParameter) UnmarshalJSON(b []byte) error {
    NameReceived := false
    ValuesReceived := false
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
        case "Values":
            if err := json.Unmarshal([]byte(v), &strct.Values); err != nil {
                return err
             }
            ValuesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Values (a required property) was received
    if !ValuesReceived {
        return errors.New("\"Values\" is required but was not present")
    }
    return nil
}

func (strct *IntegerParameter) MarshalJSON() ([]byte, error) {
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

func (strct *IntegerParameter) UnmarshalJSON(b []byte) error {
    NameReceived := false
    ValuesReceived := false
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
        case "Values":
            if err := json.Unmarshal([]byte(v), &strct.Values); err != nil {
                return err
             }
            ValuesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Values (a required property) was received
    if !ValuesReceived {
        return errors.New("\"Values\" is required but was not present")
    }
    return nil
}

func (strct *Parameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "DateTimeParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DateTimeParameters\": ")
	if tmp, err := json.Marshal(strct.DateTimeParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DecimalParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DecimalParameters\": ")
	if tmp, err := json.Marshal(strct.DecimalParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "IntegerParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IntegerParameters\": ")
	if tmp, err := json.Marshal(strct.IntegerParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "StringParameters" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"StringParameters\": ")
	if tmp, err := json.Marshal(strct.StringParameters); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

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
    for k, v := range jsonMap {
        switch k {
        case "DateTimeParameters":
            if err := json.Unmarshal([]byte(v), &strct.DateTimeParameters); err != nil {
                return err
             }
        case "DecimalParameters":
            if err := json.Unmarshal([]byte(v), &strct.DecimalParameters); err != nil {
                return err
             }
        case "IntegerParameters":
            if err := json.Unmarshal([]byte(v), &strct.IntegerParameters); err != nil {
                return err
             }
        case "StringParameters":
            if err := json.Unmarshal([]byte(v), &strct.StringParameters); err != nil {
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
    // "AnalysisId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AnalysisId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AnalysisId\": ")
	if tmp, err := json.Marshal(strct.AnalysisId); err != nil {
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
    // "AwsAccountId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "AwsAccountId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"AwsAccountId\": ")
	if tmp, err := json.Marshal(strct.AwsAccountId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "CreatedTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CreatedTime\": ")
	if tmp, err := json.Marshal(strct.CreatedTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "DataSetArns" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"DataSetArns\": ")
	if tmp, err := json.Marshal(strct.DataSetArns); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Errors" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Errors\": ")
	if tmp, err := json.Marshal(strct.Errors); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "LastUpdatedTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"LastUpdatedTime\": ")
	if tmp, err := json.Marshal(strct.LastUpdatedTime); err != nil {
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
    // Marshal the "Permissions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Permissions\": ")
	if tmp, err := json.Marshal(strct.Permissions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "Sheets" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Sheets\": ")
	if tmp, err := json.Marshal(strct.Sheets); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "SourceEntity" field is required
    if strct.SourceEntity == nil {
        return nil, errors.New("SourceEntity is a required field")
    }
    // Marshal the "SourceEntity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SourceEntity\": ")
	if tmp, err := json.Marshal(strct.SourceEntity); err != nil {
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
    // Marshal the "ThemeArn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ThemeArn\": ")
	if tmp, err := json.Marshal(strct.ThemeArn); err != nil {
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
    AnalysisIdReceived := false
    AwsAccountIdReceived := false
    SourceEntityReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "AnalysisId":
            if err := json.Unmarshal([]byte(v), &strct.AnalysisId); err != nil {
                return err
             }
            AnalysisIdReceived = true
        case "Arn":
            if err := json.Unmarshal([]byte(v), &strct.Arn); err != nil {
                return err
             }
        case "AwsAccountId":
            if err := json.Unmarshal([]byte(v), &strct.AwsAccountId); err != nil {
                return err
             }
            AwsAccountIdReceived = true
        case "CreatedTime":
            if err := json.Unmarshal([]byte(v), &strct.CreatedTime); err != nil {
                return err
             }
        case "DataSetArns":
            if err := json.Unmarshal([]byte(v), &strct.DataSetArns); err != nil {
                return err
             }
        case "Errors":
            if err := json.Unmarshal([]byte(v), &strct.Errors); err != nil {
                return err
             }
        case "LastUpdatedTime":
            if err := json.Unmarshal([]byte(v), &strct.LastUpdatedTime); err != nil {
                return err
             }
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "Parameters":
            if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
                return err
             }
        case "Permissions":
            if err := json.Unmarshal([]byte(v), &strct.Permissions); err != nil {
                return err
             }
        case "Sheets":
            if err := json.Unmarshal([]byte(v), &strct.Sheets); err != nil {
                return err
             }
        case "SourceEntity":
            if err := json.Unmarshal([]byte(v), &strct.SourceEntity); err != nil {
                return err
             }
            SourceEntityReceived = true
        case "Status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "Tags":
            if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
                return err
             }
        case "ThemeArn":
            if err := json.Unmarshal([]byte(v), &strct.ThemeArn); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if AnalysisId (a required property) was received
    if !AnalysisIdReceived {
        return errors.New("\"AnalysisId\" is required but was not present")
    }
    // check if AwsAccountId (a required property) was received
    if !AwsAccountIdReceived {
        return errors.New("\"AwsAccountId\" is required but was not present")
    }
    // check if SourceEntity (a required property) was received
    if !SourceEntityReceived {
        return errors.New("\"SourceEntity\" is required but was not present")
    }
    return nil
}

func (strct *ResourcePermission) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Actions" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Actions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Actions\": ")
	if tmp, err := json.Marshal(strct.Actions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Principal" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Principal" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Principal\": ")
	if tmp, err := json.Marshal(strct.Principal); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResourcePermission) UnmarshalJSON(b []byte) error {
    ActionsReceived := false
    PrincipalReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Actions":
            if err := json.Unmarshal([]byte(v), &strct.Actions); err != nil {
                return err
             }
            ActionsReceived = true
        case "Principal":
            if err := json.Unmarshal([]byte(v), &strct.Principal); err != nil {
                return err
             }
            PrincipalReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Actions (a required property) was received
    if !ActionsReceived {
        return errors.New("\"Actions\" is required but was not present")
    }
    // check if Principal (a required property) was received
    if !PrincipalReceived {
        return errors.New("\"Principal\" is required but was not present")
    }
    return nil
}

func (strct *Sheet) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "SheetId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"SheetId\": ")
	if tmp, err := json.Marshal(strct.SheetId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Sheet) UnmarshalJSON(b []byte) error {
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
        case "SheetId":
            if err := json.Unmarshal([]byte(v), &strct.SheetId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}

func (strct *StringParameter) MarshalJSON() ([]byte, error) {
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

func (strct *StringParameter) UnmarshalJSON(b []byte) error {
    NameReceived := false
    ValuesReceived := false
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
        case "Values":
            if err := json.Unmarshal([]byte(v), &strct.Values); err != nil {
                return err
             }
            ValuesReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if Values (a required property) was received
    if !ValuesReceived {
        return errors.New("\"Values\" is required but was not present")
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