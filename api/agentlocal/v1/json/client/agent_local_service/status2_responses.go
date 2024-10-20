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

// Status2Reader is a Reader for the Status2 structure.
type Status2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Status2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatus2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatus2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatus2OK creates a Status2OK with default headers values
func NewStatus2OK() *Status2OK {
	return &Status2OK{}
}

/*
Status2OK describes a response with status code 200, with default header values.

A successful response.
*/
type Status2OK struct {
	Payload *Status2OKBody
}

// IsSuccess returns true when this status2 Ok response has a 2xx status code
func (o *Status2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status2 Ok response has a 3xx status code
func (o *Status2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status2 Ok response has a 4xx status code
func (o *Status2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status2 Ok response has a 5xx status code
func (o *Status2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this status2 Ok response a status code equal to that given
func (o *Status2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status2 Ok response
func (o *Status2OK) Code() int {
	return 200
}

func (o *Status2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /local/Status][%d] status2Ok %s", 200, payload)
}

func (o *Status2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /local/Status][%d] status2Ok %s", 200, payload)
}

func (o *Status2OK) GetPayload() *Status2OKBody {
	return o.Payload
}

func (o *Status2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(Status2OKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatus2Default creates a Status2Default with default headers values
func NewStatus2Default(code int) *Status2Default {
	return &Status2Default{
		_statusCode: code,
	}
}

/*
Status2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type Status2Default struct {
	_statusCode int

	Payload *Status2DefaultBody
}

// IsSuccess returns true when this status2 default response has a 2xx status code
func (o *Status2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this status2 default response has a 3xx status code
func (o *Status2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this status2 default response has a 4xx status code
func (o *Status2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this status2 default response has a 5xx status code
func (o *Status2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this status2 default response a status code equal to that given
func (o *Status2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the status2 default response
func (o *Status2Default) Code() int {
	return o._statusCode
}

func (o *Status2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /local/Status][%d] Status2 default %s", o._statusCode, payload)
}

func (o *Status2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /local/Status][%d] Status2 default %s", o._statusCode, payload)
}

func (o *Status2Default) GetPayload() *Status2DefaultBody {
	return o.Payload
}

func (o *Status2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(Status2DefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
Status2DefaultBody status2 default body
swagger:model Status2DefaultBody
*/
type Status2DefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*Status2DefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this status2 default body
func (o *Status2DefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2DefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this status2 default body based on the context it is used
func (o *Status2DefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2DefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *Status2DefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2DefaultBody) UnmarshalBinary(b []byte) error {
	var res Status2DefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
Status2DefaultBodyDetailsItems0 status2 default body details items0
swagger:model Status2DefaultBodyDetailsItems0
*/
type Status2DefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// status2 default body details items0
	Status2DefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *Status2DefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv Status2DefaultBodyDetailsItems0

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
		o.Status2DefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o Status2DefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.Status2DefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.Status2DefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this status2 default body details items0
func (o *Status2DefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status2 default body details items0 based on context it is used
func (o *Status2DefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2DefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2DefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res Status2DefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
Status2OKBody status2 OK body
swagger:model Status2OKBody
*/
type Status2OKBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// TODO: rename to node_id
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// node name
	NodeName string `json:"node_name,omitempty"`

	// agents info
	AgentsInfo []*Status2OKBodyAgentsInfoItems0 `json:"agents_info"`

	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `json:"config_filepath,omitempty"`

	// PMM Agent version.
	AgentVersion string `json:"agent_version,omitempty"`

	// Shows connection uptime in percentage between agent and server
	ConnectionUptime float32 `json:"connection_uptime,omitempty"`

	// server info
	ServerInfo *Status2OKBodyServerInfo `json:"server_info,omitempty"`
}

// Validate validates this status2 OK body
func (o *Status2OKBody) Validate(formats strfmt.Registry) error {
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

func (o *Status2OKBody) validateAgentsInfo(formats strfmt.Registry) error {
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
					return ve.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *Status2OKBody) validateServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.ServerInfo) { // not required
		return nil
	}

	if o.ServerInfo != nil {
		if err := o.ServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status2Ok" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status2Ok" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status2 OK body based on the context it is used
func (o *Status2OKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *Status2OKBody) contextValidateAgentsInfo(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.AgentsInfo); i++ {
		if o.AgentsInfo[i] != nil {

			if swag.IsZero(o.AgentsInfo[i]) { // not required
				return nil
			}

			if err := o.AgentsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *Status2OKBody) contextValidateServerInfo(ctx context.Context, formats strfmt.Registry) error {
	if o.ServerInfo != nil {

		if swag.IsZero(o.ServerInfo) { // not required
			return nil
		}

		if err := o.ServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status2Ok" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status2Ok" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBody) UnmarshalBinary(b []byte) error {
	var res Status2OKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
Status2OKBodyAgentsInfoItems0 AgentInfo contains information about Agent managed by pmm-agent.
swagger:model Status2OKBodyAgentsInfoItems0
*/
type Status2OKBodyAgentsInfoItems0 struct {
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

// Validate validates this status2 OK body agents info items0
func (o *Status2OKBodyAgentsInfoItems0) Validate(formats strfmt.Registry) error {
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

var status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_TYPE_UNSPECIFIED","AGENT_TYPE_PMM_AGENT","AGENT_TYPE_VM_AGENT","AGENT_TYPE_NODE_EXPORTER","AGENT_TYPE_MYSQLD_EXPORTER","AGENT_TYPE_MONGODB_EXPORTER","AGENT_TYPE_POSTGRES_EXPORTER","AGENT_TYPE_PROXYSQL_EXPORTER","AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT","AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT","AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT","AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT","AGENT_TYPE_EXTERNAL_EXPORTER","AGENT_TYPE_RDS_EXPORTER","AGENT_TYPE_AZURE_DATABASE_EXPORTER","AGENT_TYPE_NOMAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum = append(status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum, v)
	}
}

const (

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEUNSPECIFIED captures enum value "AGENT_TYPE_UNSPECIFIED"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEUNSPECIFIED string = "AGENT_TYPE_UNSPECIFIED"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPMMAGENT captures enum value "AGENT_TYPE_PMM_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPMMAGENT string = "AGENT_TYPE_PMM_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEVMAGENT captures enum value "AGENT_TYPE_VM_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEVMAGENT string = "AGENT_TYPE_VM_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPENODEEXPORTER captures enum value "AGENT_TYPE_NODE_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPENODEEXPORTER string = "AGENT_TYPE_NODE_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEMYSQLDEXPORTER captures enum value "AGENT_TYPE_MYSQLD_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEMYSQLDEXPORTER string = "AGENT_TYPE_MYSQLD_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEMONGODBEXPORTER captures enum value "AGENT_TYPE_MONGODB_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEMONGODBEXPORTER string = "AGENT_TYPE_MONGODB_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPOSTGRESEXPORTER captures enum value "AGENT_TYPE_POSTGRES_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPOSTGRESEXPORTER string = "AGENT_TYPE_POSTGRES_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPROXYSQLEXPORTER captures enum value "AGENT_TYPE_PROXYSQL_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEPROXYSQLEXPORTER string = "AGENT_TYPE_PROXYSQL_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLPERFSCHEMAAGENT captures enum value "AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLPERFSCHEMAAGENT string = "AGENT_TYPE_QAN_MYSQL_PERFSCHEMA_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLSLOWLOGAGENT captures enum value "AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMYSQLSLOWLOGAGENT string = "AGENT_TYPE_QAN_MYSQL_SLOWLOG_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMONGODBPROFILERAGENT captures enum value "AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANMONGODBPROFILERAGENT string = "AGENT_TYPE_QAN_MONGODB_PROFILER_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATEMENTSAGENT captures enum value "AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATEMENTSAGENT string = "AGENT_TYPE_QAN_POSTGRESQL_PGSTATEMENTS_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATMONITORAGENT captures enum value "AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEQANPOSTGRESQLPGSTATMONITORAGENT string = "AGENT_TYPE_QAN_POSTGRESQL_PGSTATMONITOR_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEEXTERNALEXPORTER captures enum value "AGENT_TYPE_EXTERNAL_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEEXTERNALEXPORTER string = "AGENT_TYPE_EXTERNAL_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPERDSEXPORTER captures enum value "AGENT_TYPE_RDS_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPERDSEXPORTER string = "AGENT_TYPE_RDS_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEAZUREDATABASEEXPORTER captures enum value "AGENT_TYPE_AZURE_DATABASE_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEAZUREDATABASEEXPORTER string = "AGENT_TYPE_AZURE_DATABASE_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPENOMAD captures enum value "AGENT_TYPE_NOMAD"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPENOMAD string = "AGENT_TYPE_NOMAD"
)

// prop value enum
func (o *Status2OKBodyAgentsInfoItems0) validateAgentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *Status2OKBodyAgentsInfoItems0) validateAgentType(formats strfmt.Registry) error {
	if swag.IsZero(o.AgentType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAgentTypeEnum("agent_type", "body", *o.AgentType); err != nil {
		return err
	}

	return nil
}

var status2OkBodyAgentsInfoItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_UNSPECIFIED","AGENT_STATUS_STARTING","AGENT_STATUS_INITIALIZATION_ERROR","AGENT_STATUS_RUNNING","AGENT_STATUS_WAITING","AGENT_STATUS_STOPPING","AGENT_STATUS_DONE","AGENT_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		status2OkBodyAgentsInfoItems0TypeStatusPropEnum = append(status2OkBodyAgentsInfoItems0TypeStatusPropEnum, v)
	}
}

const (

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSUNSPECIFIED captures enum value "AGENT_STATUS_UNSPECIFIED"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSUNSPECIFIED string = "AGENT_STATUS_UNSPECIFIED"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSSTARTING captures enum value "AGENT_STATUS_STARTING"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSSTARTING string = "AGENT_STATUS_STARTING"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSINITIALIZATIONERROR captures enum value "AGENT_STATUS_INITIALIZATION_ERROR"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSINITIALIZATIONERROR string = "AGENT_STATUS_INITIALIZATION_ERROR"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSRUNNING captures enum value "AGENT_STATUS_RUNNING"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSRUNNING string = "AGENT_STATUS_RUNNING"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSWAITING captures enum value "AGENT_STATUS_WAITING"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSWAITING string = "AGENT_STATUS_WAITING"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSSTOPPING captures enum value "AGENT_STATUS_STOPPING"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSSTOPPING string = "AGENT_STATUS_STOPPING"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSDONE captures enum value "AGENT_STATUS_DONE"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSDONE string = "AGENT_STATUS_DONE"

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSUNKNOWN captures enum value "AGENT_STATUS_UNKNOWN"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSUNKNOWN string = "AGENT_STATUS_UNKNOWN"
)

// prop value enum
func (o *Status2OKBodyAgentsInfoItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, status2OkBodyAgentsInfoItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *Status2OKBodyAgentsInfoItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this status2 OK body agents info items0 based on context it is used
func (o *Status2OKBodyAgentsInfoItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBodyAgentsInfoItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBodyAgentsInfoItems0) UnmarshalBinary(b []byte) error {
	var res Status2OKBodyAgentsInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
Status2OKBodyServerInfo ServerInfo contains information about the PMM Server.
swagger:model Status2OKBodyServerInfo
*/
type Status2OKBodyServerInfo struct {
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

// Validate validates this status2 OK body server info
func (o *Status2OKBodyServerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status2 OK body server info based on context it is used
func (o *Status2OKBodyServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBodyServerInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBodyServerInfo) UnmarshalBinary(b []byte) error {
	var res Status2OKBodyServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
