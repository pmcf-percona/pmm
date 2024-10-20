// Code generated by go-swagger; DO NOT EDIT.

package agent_local_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusReader is a Reader for the Status structure.
type StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusOK creates a StatusOK with default headers values
func NewStatusOK() *StatusOK {
	return &StatusOK{}
}

/*
StatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type StatusOK struct {
	Payload *StatusOKBody
}

// IsSuccess returns true when this status Ok response has a 2xx status code
func (o *StatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status Ok response has a 3xx status code
func (o *StatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status Ok response has a 4xx status code
func (o *StatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status Ok response has a 5xx status code
func (o *StatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status Ok response a status code equal to that given
func (o *StatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status Ok response
func (o *StatusOK) Code() int {
	return 200
}

func (o *StatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /local/Status][%d] statusOk %s", 200, payload)
}

func (o *StatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /local/Status][%d] statusOk %s", 200, payload)
}

func (o *StatusOK) GetPayload() *StatusOKBody {
	return o.Payload
}

func (o *StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusDefault creates a StatusDefault with default headers values
func NewStatusDefault(code int) *StatusDefault {
	return &StatusDefault{
		_statusCode: code,
	}
}

/*
StatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StatusDefault struct {
	_statusCode int

	Payload *StatusDefaultBody
}

// IsSuccess returns true when this status default response has a 2xx status code
func (o *StatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this status default response has a 3xx status code
func (o *StatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this status default response has a 4xx status code
func (o *StatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this status default response has a 5xx status code
func (o *StatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this status default response a status code equal to that given
func (o *StatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the status default response
func (o *StatusDefault) Code() int {
	return o._statusCode
}

func (o *StatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /local/Status][%d] Status default %s", o._statusCode, payload)
}

func (o *StatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /local/Status][%d] Status default %s", o._statusCode, payload)
}

func (o *StatusDefault) GetPayload() *StatusDefaultBody {
	return o.Payload
}

func (o *StatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StatusBody status body
swagger:model StatusBody
*/
type StatusBody struct {
	// Returns network info (latency and clock_drift) if true.
	GetNetworkInfo bool `json:"get_network_info,omitempty"`
}

// Validate validates this status body
func (o *StatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status body based on context it is used
func (o *StatusBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusBody) UnmarshalBinary(b []byte) error {
	var res StatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StatusDefaultBody status default body
swagger:model StatusDefaultBody
*/
type StatusDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StatusDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this status default body
func (o *StatusDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this status default body based on the context it is used
func (o *StatusDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res StatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StatusDefaultBodyDetailsItems0 status default body details items0
swagger:model StatusDefaultBodyDetailsItems0
*/
type StatusDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// status default body details items0
	StatusDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *StatusDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StatusDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.StatusDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o StatusDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.StatusDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.StatusDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this status default body details items0
func (o *StatusDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status default body details items0 based on context it is used
func (o *StatusDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StatusDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StatusOKBody status OK body
swagger:model StatusOKBody
*/
type StatusOKBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// TODO: rename to node_id
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// node name
	NodeName string `json:"node_name,omitempty"`

	// agents info
	AgentsInfo []*StatusOKBodyAgentsInfoItems0 `json:"agents_info"`

	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `json:"config_filepath,omitempty"`

	// PMM Agent version.
	AgentVersion string `json:"agent_version,omitempty"`

	// Shows connection uptime in percentage between agent and server
	ConnectionUptime float32 `json:"connection_uptime,omitempty"`

	// server info
	ServerInfo *StatusOKBodyServerInfo `json:"server_info,omitempty"`
}

// Validate validates this status OK body
func (o *StatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusOKBody) validateAgentsInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.AgentsInfo) { // not required
		return nil
	}

	for i := 0; i < len(o.AgentsInfo); i++ {
		if swag.IsZero(o.AgentsInfo[i]) { // not required
			continue
		}

		if o.AgentsInfo[i] != nil {
			if err := o.AgentsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statusOk" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statusOk" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *StatusOKBody) validateServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.ServerInfo) { // not required
		return nil
	}

	if o.ServerInfo != nil {
		if err := o.ServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOk" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOk" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status OK body based on the context it is used
func (o *StatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAgentsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusOKBody) contextValidateAgentsInfo(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.AgentsInfo); i++ {
		if o.AgentsInfo[i] != nil {

			if swag.IsZero(o.AgentsInfo[i]) { // not required
				return nil
			}

			if err := o.AgentsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statusOk" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statusOk" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *StatusOKBody) contextValidateServerInfo(ctx context.Context, formats strfmt.Registry) error {
	if o.ServerInfo != nil {

		if swag.IsZero(o.ServerInfo) { // not required
			return nil
		}

		if err := o.ServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOk" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOk" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBody) UnmarshalBinary(b []byte) error {
	var res StatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StatusOKBodyAgentsInfoItems0 AgentInfo contains information about Agent managed by pmm-agent.
swagger:model StatusOKBodyAgentsInfoItems0
*/
type StatusOKBodyAgentsInfoItems0 struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// AgentType describes supported Agent types.
	// Enum: ["AGENT_TYPE_UNSPECIFIED","AGENT_TYPE_PMM_AGENT","AGENT_TYPE_VM_AGENT","AGENT_TYPE_NODE_EXPORTER","AGENT_TYPE_MYSQLD_EXPORTER","AGENT_TYPE_MONGODB_EXPORTER","AGENT_TYPE_POSTGRES_EXPORTER","AGENT_TYPE_PROXYSQL_EXPORTER","AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT","AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT","AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT","AGENT_TYPE_EXTERNAL_EXPORTER","AGENT_TYPE_RDS_EXPORTER","AGENT_TYPE_AZURE_DATABASE_EXPORTER","AGENT_TYPE_NOMAD"]
	AgentType *string `json:"agent_type,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - AGENT_STATUS_STARTING: Agent is starting.
	//  - AGENT_STATUS_INITIALIZATION_ERROR: Agent encountered error when starting.
	//  - AGENT_STATUS_RUNNING: Agent is running.
	//  - AGENT_STATUS_WAITING: Agent encountered error and will be restarted automatically soon.
	//  - AGENT_STATUS_STOPPING: Agent is stopping.
	//  - AGENT_STATUS_DONE: Agent finished.
	//  - AGENT_STATUS_UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: ["AGENT_STATUS_UNSPECIFIED","AGENT_STATUS_STARTING","AGENT_STATUS_INITIALIZATION_ERROR","AGENT_STATUS_RUNNING","AGENT_STATUS_WAITING","AGENT_STATUS_STOPPING","AGENT_STATUS_DONE","AGENT_STATUS_UNKNOWN"]
	Status *string `json:"status,omitempty"`

	// The current listen port of this Agent (exporter or vmagent).
	// Zero for other Agent types, or if unknown or not yet supported.
	ListenPort int64 `json:"listen_port,omitempty"`

	// process exec path
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this status OK body agents info items0
func (o *StatusOKBodyAgentsInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statusOkBodyAgentsInfoItems0TypeAgentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_TYPE_UNSPECIFIED","AGENT_TYPE_PMM_AGENT","AGENT_TYPE_VM_AGENT","AGENT_TYPE_NODE_EXPORTER","AGENT_TYPE_MYSQLD_EXPORTER","AGENT_TYPE_MONGODB_EXPORTER","AGENT_TYPE_POSTGRES_EXPORTER","AGENT_TYPE_PROXYSQL_EXPORTER","AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT","AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT","AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT","AGENT_TYPE_EXTERNAL_EXPORTER","AGENT_TYPE_RDS_EXPORTER","AGENT_TYPE_AZURE_DATABASE_EXPORTER","AGENT_TYPE_NOMAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusOkBodyAgentsInfoItems0TypeAgentTypePropEnum = append(statusOkBodyAgentsInfoItems0TypeAgentTypePropEnum, v)
	}
}

const (

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEUNSPECIFIED captures enum value "AGENT_TYPE_UNSPECIFIED"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEUNSPECIFIED string = "AGENT_TYPE_UNSPECIFIED"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPMMAGENT captures enum value "AGENT_TYPE_PMM_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPMMAGENT string = "AGENT_TYPE_PMM_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEVMAGENT captures enum value "AGENT_TYPE_VM_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEVMAGENT string = "AGENT_TYPE_VM_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPENODEEXPORTER captures enum value "AGENT_TYPE_NODE_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPENODEEXPORTER string = "AGENT_TYPE_NODE_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEMYSQLDEXPORTER captures enum value "AGENT_TYPE_MYSQLD_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEMYSQLDEXPORTER string = "AGENT_TYPE_MYSQLD_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEMONGODBEXPORTER captures enum value "AGENT_TYPE_MONGODB_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEMONGODBEXPORTER string = "AGENT_TYPE_MONGODB_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPOSTGRESEXPORTER captures enum value "AGENT_TYPE_POSTGRES_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPOSTGRESEXPORTER string = "AGENT_TYPE_POSTGRES_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPROXYSQLEXPORTER captures enum value "AGENT_TYPE_PROXYSQL_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEPROXYSQLEXPORTER string = "AGENT_TYPE_PROXYSQL_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLPERFSCHEMAAGENT captures enum value "AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLPERFSCHEMAAGENT string = "AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLSLOWLOGAGENT captures enum value "AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLSLOWLOGAGENT string = "AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMONGODBPROFILERAGENT captures enum value "AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMONGODBPROFILERAGENT string = "AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATEMENTSAGENT captures enum value "AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATEMENTSAGENT string = "AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATMONITORAGENT captures enum value "AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATMONITORAGENT string = "AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEEXTERNALEXPORTER captures enum value "AGENT_TYPE_EXTERNAL_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEEXTERNALEXPORTER string = "AGENT_TYPE_EXTERNAL_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPERDSEXPORTER captures enum value "AGENT_TYPE_RDS_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPERDSEXPORTER string = "AGENT_TYPE_RDS_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEAZUREDATABASEEXPORTER captures enum value "AGENT_TYPE_AZURE_DATABASE_EXPORTER"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPEAZUREDATABASEEXPORTER string = "AGENT_TYPE_AZURE_DATABASE_EXPORTER"

	// StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPENOMAD captures enum value "AGENT_TYPE_NOMAD"
	StatusOKBodyAgentsInfoItems0AgentTypeAGENTTYPENOMAD string = "AGENT_TYPE_NOMAD"
)

// prop value enum
func (o *StatusOKBodyAgentsInfoItems0) validateAgentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusOkBodyAgentsInfoItems0TypeAgentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *StatusOKBodyAgentsInfoItems0) validateAgentType(formats strfmt.Registry) error {
	if swag.IsZero(o.AgentType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAgentTypeEnum("agent_type", "body", *o.AgentType); err != nil {
		return err
	}

	return nil
}

var statusOkBodyAgentsInfoItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_UNSPECIFIED","AGENT_STATUS_STARTING","AGENT_STATUS_INITIALIZATION_ERROR","AGENT_STATUS_RUNNING","AGENT_STATUS_WAITING","AGENT_STATUS_STOPPING","AGENT_STATUS_DONE","AGENT_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusOkBodyAgentsInfoItems0TypeStatusPropEnum = append(statusOkBodyAgentsInfoItems0TypeStatusPropEnum, v)
	}
}

const (

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSUNSPECIFIED captures enum value "AGENT_STATUS_UNSPECIFIED"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSUNSPECIFIED string = "AGENT_STATUS_UNSPECIFIED"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSSTARTING captures enum value "AGENT_STATUS_STARTING"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSSTARTING string = "AGENT_STATUS_STARTING"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSINITIALIZATIONERROR captures enum value "AGENT_STATUS_INITIALIZATION_ERROR"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSINITIALIZATIONERROR string = "AGENT_STATUS_INITIALIZATION_ERROR"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSRUNNING captures enum value "AGENT_STATUS_RUNNING"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSRUNNING string = "AGENT_STATUS_RUNNING"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSWAITING captures enum value "AGENT_STATUS_WAITING"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSWAITING string = "AGENT_STATUS_WAITING"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSSTOPPING captures enum value "AGENT_STATUS_STOPPING"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSSTOPPING string = "AGENT_STATUS_STOPPING"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSDONE captures enum value "AGENT_STATUS_DONE"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSDONE string = "AGENT_STATUS_DONE"

	// StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSUNKNOWN captures enum value "AGENT_STATUS_UNKNOWN"
	StatusOKBodyAgentsInfoItems0StatusAGENTSTATUSUNKNOWN string = "AGENT_STATUS_UNKNOWN"
)

// prop value enum
func (o *StatusOKBodyAgentsInfoItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusOkBodyAgentsInfoItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *StatusOKBodyAgentsInfoItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this status OK body agents info items0 based on context it is used
func (o *StatusOKBodyAgentsInfoItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyAgentsInfoItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyAgentsInfoItems0) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyAgentsInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StatusOKBodyServerInfo ServerInfo contains information about the PMM Server.
swagger:model StatusOKBodyServerInfo
*/
type StatusOKBodyServerInfo struct {
	// PMM Server URL in a form https://HOST:PORT/.
	URL string `json:"url,omitempty"`

	// PMM Server's TLS certificate validation should be skipped if true.
	InsecureTLS bool `json:"insecure_tls,omitempty"`

	// True if pmm-agent is currently connected to the server.
	Connected bool `json:"connected,omitempty"`

	// PMM Server version (if agent is connected).
	Version string `json:"version,omitempty"`

	// Ping time from pmm-agent to pmm-managed (if agent is connected).
	Latency string `json:"latency,omitempty"`

	// Clock drift from PMM Server (if agent is connected).
	ClockDrift string `json:"clock_drift,omitempty"`
}

// Validate validates this status OK body server info
func (o *StatusOKBodyServerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status OK body server info based on context it is used
func (o *StatusOKBodyServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyServerInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyServerInfo) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
