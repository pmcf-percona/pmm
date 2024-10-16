// Code generated by go-swagger; DO NOT EDIT.

package platform_service

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
)

// DisconnectReader is a Reader for the Disconnect structure.
type DisconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisconnectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisconnectOK creates a DisconnectOK with default headers values
func NewDisconnectOK() *DisconnectOK {
	return &DisconnectOK{}
}

/*
DisconnectOK describes a response with status code 200, with default header values.

A successful response.
*/
type DisconnectOK struct {
	Payload interface{}
}

// IsSuccess returns true when this disconnect Ok response has a 2xx status code
func (o *DisconnectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disconnect Ok response has a 3xx status code
func (o *DisconnectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disconnect Ok response has a 4xx status code
func (o *DisconnectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disconnect Ok response has a 5xx status code
func (o *DisconnectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disconnect Ok response a status code equal to that given
func (o *DisconnectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disconnect Ok response
func (o *DisconnectOK) Code() int {
	return 200
}

func (o *DisconnectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/platform:disconnect][%d] disconnectOk %s", 200, payload)
}

func (o *DisconnectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/platform:disconnect][%d] disconnectOk %s", 200, payload)
}

func (o *DisconnectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DisconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisconnectDefault creates a DisconnectDefault with default headers values
func NewDisconnectDefault(code int) *DisconnectDefault {
	return &DisconnectDefault{
		_statusCode: code,
	}
}

/*
DisconnectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DisconnectDefault struct {
	_statusCode int

	Payload *DisconnectDefaultBody
}

// IsSuccess returns true when this disconnect default response has a 2xx status code
func (o *DisconnectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disconnect default response has a 3xx status code
func (o *DisconnectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disconnect default response has a 4xx status code
func (o *DisconnectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disconnect default response has a 5xx status code
func (o *DisconnectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disconnect default response a status code equal to that given
func (o *DisconnectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the disconnect default response
func (o *DisconnectDefault) Code() int {
	return o._statusCode
}

func (o *DisconnectDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/platform:disconnect][%d] Disconnect default %s", o._statusCode, payload)
}

func (o *DisconnectDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/platform:disconnect][%d] Disconnect default %s", o._statusCode, payload)
}

func (o *DisconnectDefault) GetPayload() *DisconnectDefaultBody {
	return o.Payload
}

func (o *DisconnectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(DisconnectDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DisconnectBody disconnect body
swagger:model DisconnectBody
*/
type DisconnectBody struct {
	// Forces the cleanup process for connected PMM instances regardless of the Portal API response
	Force bool `json:"force,omitempty"`
}

// Validate validates this disconnect body
func (o *DisconnectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disconnect body based on context it is used
func (o *DisconnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DisconnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisconnectBody) UnmarshalBinary(b []byte) error {
	var res DisconnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DisconnectDefaultBody disconnect default body
swagger:model DisconnectDefaultBody
*/
type DisconnectDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DisconnectDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this disconnect default body
func (o *DisconnectDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DisconnectDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Disconnect default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Disconnect default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this disconnect default body based on the context it is used
func (o *DisconnectDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DisconnectDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Disconnect default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Disconnect default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DisconnectDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisconnectDefaultBody) UnmarshalBinary(b []byte) error {
	var res DisconnectDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DisconnectDefaultBodyDetailsItems0 disconnect default body details items0
swagger:model DisconnectDefaultBodyDetailsItems0
*/
type DisconnectDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// disconnect default body details items0
	DisconnectDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *DisconnectDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DisconnectDefaultBodyDetailsItems0

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
		o.DisconnectDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o DisconnectDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.DisconnectDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.DisconnectDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this disconnect default body details items0
func (o *DisconnectDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disconnect default body details items0 based on context it is used
func (o *DisconnectDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DisconnectDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisconnectDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DisconnectDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
