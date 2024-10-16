// Code generated by go-swagger; DO NOT EDIT.

package advisor_service

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

// ChangeAdvisorChecksReader is a Reader for the ChangeAdvisorChecks structure.
type ChangeAdvisorChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeAdvisorChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeAdvisorChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeAdvisorChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeAdvisorChecksOK creates a ChangeAdvisorChecksOK with default headers values
func NewChangeAdvisorChecksOK() *ChangeAdvisorChecksOK {
	return &ChangeAdvisorChecksOK{}
}

/*
ChangeAdvisorChecksOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeAdvisorChecksOK struct {
	Payload interface{}
}

// IsSuccess returns true when this change advisor checks Ok response has a 2xx status code
func (o *ChangeAdvisorChecksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change advisor checks Ok response has a 3xx status code
func (o *ChangeAdvisorChecksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change advisor checks Ok response has a 4xx status code
func (o *ChangeAdvisorChecksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change advisor checks Ok response has a 5xx status code
func (o *ChangeAdvisorChecksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change advisor checks Ok response a status code equal to that given
func (o *ChangeAdvisorChecksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change advisor checks Ok response
func (o *ChangeAdvisorChecksOK) Code() int {
	return 200
}

func (o *ChangeAdvisorChecksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/advisors/checks:batchChange][%d] changeAdvisorChecksOk %s", 200, payload)
}

func (o *ChangeAdvisorChecksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/advisors/checks:batchChange][%d] changeAdvisorChecksOk %s", 200, payload)
}

func (o *ChangeAdvisorChecksOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeAdvisorChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAdvisorChecksDefault creates a ChangeAdvisorChecksDefault with default headers values
func NewChangeAdvisorChecksDefault(code int) *ChangeAdvisorChecksDefault {
	return &ChangeAdvisorChecksDefault{
		_statusCode: code,
	}
}

/*
ChangeAdvisorChecksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeAdvisorChecksDefault struct {
	_statusCode int

	Payload *ChangeAdvisorChecksDefaultBody
}

// IsSuccess returns true when this change advisor checks default response has a 2xx status code
func (o *ChangeAdvisorChecksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this change advisor checks default response has a 3xx status code
func (o *ChangeAdvisorChecksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this change advisor checks default response has a 4xx status code
func (o *ChangeAdvisorChecksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this change advisor checks default response has a 5xx status code
func (o *ChangeAdvisorChecksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this change advisor checks default response a status code equal to that given
func (o *ChangeAdvisorChecksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the change advisor checks default response
func (o *ChangeAdvisorChecksDefault) Code() int {
	return o._statusCode
}

func (o *ChangeAdvisorChecksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/advisors/checks:batchChange][%d] ChangeAdvisorChecks default %s", o._statusCode, payload)
}

func (o *ChangeAdvisorChecksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/advisors/checks:batchChange][%d] ChangeAdvisorChecks default %s", o._statusCode, payload)
}

func (o *ChangeAdvisorChecksDefault) GetPayload() *ChangeAdvisorChecksDefaultBody {
	return o.Payload
}

func (o *ChangeAdvisorChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeAdvisorChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangeAdvisorChecksBody change advisor checks body
swagger:model ChangeAdvisorChecksBody
*/
type ChangeAdvisorChecksBody struct {
	// params
	Params []*ChangeAdvisorChecksParamsBodyParamsItems0 `json:"params"`
}

// Validate validates this change advisor checks body
func (o *ChangeAdvisorChecksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAdvisorChecksBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change advisor checks body based on the context it is used
func (o *ChangeAdvisorChecksBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAdvisorChecksBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Params); i++ {
		if o.Params[i] != nil {

			if swag.IsZero(o.Params[i]) { // not required
				return nil
			}

			if err := o.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAdvisorChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAdvisorChecksBody) UnmarshalBinary(b []byte) error {
	var res ChangeAdvisorChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeAdvisorChecksDefaultBody change advisor checks default body
swagger:model ChangeAdvisorChecksDefaultBody
*/
type ChangeAdvisorChecksDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeAdvisorChecksDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change advisor checks default body
func (o *ChangeAdvisorChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAdvisorChecksDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeAdvisorChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeAdvisorChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change advisor checks default body based on the context it is used
func (o *ChangeAdvisorChecksDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAdvisorChecksDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeAdvisorChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeAdvisorChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAdvisorChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAdvisorChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeAdvisorChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeAdvisorChecksDefaultBodyDetailsItems0 change advisor checks default body details items0
swagger:model ChangeAdvisorChecksDefaultBodyDetailsItems0
*/
type ChangeAdvisorChecksDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// change advisor checks default body details items0
	ChangeAdvisorChecksDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ChangeAdvisorChecksDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ChangeAdvisorChecksDefaultBodyDetailsItems0

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
		o.ChangeAdvisorChecksDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ChangeAdvisorChecksDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.ChangeAdvisorChecksDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ChangeAdvisorChecksDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this change advisor checks default body details items0
func (o *ChangeAdvisorChecksDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change advisor checks default body details items0 based on context it is used
func (o *ChangeAdvisorChecksDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAdvisorChecksDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAdvisorChecksDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeAdvisorChecksDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeAdvisorChecksParamsBodyParamsItems0 ChangeAdvisorCheckParams specifies a single check parameters.
swagger:model ChangeAdvisorChecksParamsBodyParamsItems0
*/
type ChangeAdvisorChecksParamsBodyParamsItems0 struct {
	// The name of the check to change.
	Name string `json:"name,omitempty"`

	// enable
	Enable *bool `json:"enable,omitempty"`

	// AdvisorCheckInterval represents possible execution interval values for checks.
	// Enum: ["ADVISOR_CHECK_INTERVAL_UNSPECIFIED","ADVISOR_CHECK_INTERVAL_STANDARD","ADVISOR_CHECK_INTERVAL_FREQUENT","ADVISOR_CHECK_INTERVAL_RARE"]
	Interval *string `json:"interval,omitempty"`
}

// Validate validates this change advisor checks params body params items0
func (o *ChangeAdvisorChecksParamsBodyParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeAdvisorChecksParamsBodyParamsItems0TypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADVISOR_CHECK_INTERVAL_UNSPECIFIED","ADVISOR_CHECK_INTERVAL_STANDARD","ADVISOR_CHECK_INTERVAL_FREQUENT","ADVISOR_CHECK_INTERVAL_RARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeAdvisorChecksParamsBodyParamsItems0TypeIntervalPropEnum = append(changeAdvisorChecksParamsBodyParamsItems0TypeIntervalPropEnum, v)
	}
}

const (

	// ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALUNSPECIFIED captures enum value "ADVISOR_CHECK_INTERVAL_UNSPECIFIED"
	ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALUNSPECIFIED string = "ADVISOR_CHECK_INTERVAL_UNSPECIFIED"

	// ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALSTANDARD captures enum value "ADVISOR_CHECK_INTERVAL_STANDARD"
	ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALSTANDARD string = "ADVISOR_CHECK_INTERVAL_STANDARD"

	// ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALFREQUENT captures enum value "ADVISOR_CHECK_INTERVAL_FREQUENT"
	ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALFREQUENT string = "ADVISOR_CHECK_INTERVAL_FREQUENT"

	// ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALRARE captures enum value "ADVISOR_CHECK_INTERVAL_RARE"
	ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALRARE string = "ADVISOR_CHECK_INTERVAL_RARE"
)

// prop value enum
func (o *ChangeAdvisorChecksParamsBodyParamsItems0) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeAdvisorChecksParamsBodyParamsItems0TypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeAdvisorChecksParamsBodyParamsItems0) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(o.Interval) { // not required
		return nil
	}

	// value enum
	if err := o.validateIntervalEnum("interval", "body", *o.Interval); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change advisor checks params body params items0 based on context it is used
func (o *ChangeAdvisorChecksParamsBodyParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAdvisorChecksParamsBodyParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAdvisorChecksParamsBodyParamsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeAdvisorChecksParamsBodyParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
