// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeGenericNodeReader is a Reader for the ChangeGenericNode structure.
type ChangeGenericNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeGenericNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeGenericNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewChangeGenericNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeGenericNodeOK creates a ChangeGenericNodeOK with default headers values
func NewChangeGenericNodeOK() *ChangeGenericNodeOK {
	return &ChangeGenericNodeOK{}
}

/*ChangeGenericNodeOK handles this case with default header values.

A successful response.
*/
type ChangeGenericNodeOK struct {
	Payload *ChangeGenericNodeOKBody
}

func (o *ChangeGenericNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/ChangeGeneric][%d] changeGenericNodeOk  %+v", 200, o.Payload)
}

func (o *ChangeGenericNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeGenericNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeGenericNodeDefault creates a ChangeGenericNodeDefault with default headers values
func NewChangeGenericNodeDefault(code int) *ChangeGenericNodeDefault {
	return &ChangeGenericNodeDefault{
		_statusCode: code,
	}
}

/*ChangeGenericNodeDefault handles this case with default header values.

An error response.
*/
type ChangeGenericNodeDefault struct {
	_statusCode int

	Payload *ChangeGenericNodeDefaultBody
}

// Code gets the status code for the change generic node default response
func (o *ChangeGenericNodeDefault) Code() int {
	return o._statusCode
}

func (o *ChangeGenericNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/ChangeGeneric][%d] ChangeGenericNode default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeGenericNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeGenericNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeGenericNodeBody change generic node body
swagger:model ChangeGenericNodeBody
*/
type ChangeGenericNodeBody struct {

	// Address FIXME.
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution (if any).
	Distro string `json:"distro,omitempty"`

	// Linux distribution version (if any).
	DistroVersion string `json:"distro_version,omitempty"`

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this change generic node body
func (o *ChangeGenericNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeGenericNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeGenericNodeBody) UnmarshalBinary(b []byte) error {
	var res ChangeGenericNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeGenericNodeDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeGenericNodeDefaultBody
*/
type ChangeGenericNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change generic node default body
func (o *ChangeGenericNodeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeGenericNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeGenericNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeGenericNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeGenericNodeOKBody change generic node OK body
swagger:model ChangeGenericNodeOKBody
*/
type ChangeGenericNodeOKBody struct {

	// generic
	Generic *ChangeGenericNodeOKBodyGeneric `json:"generic,omitempty"`
}

// Validate validates this change generic node OK body
func (o *ChangeGenericNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeGenericNodeOKBody) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeGenericNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeGenericNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeGenericNodeOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeGenericNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeGenericNodeOKBodyGeneric GenericNode represents a bare metal server or virtual machine.
swagger:model ChangeGenericNodeOKBodyGeneric
*/
type ChangeGenericNodeOKBodyGeneric struct {

	// Address FIXME.
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution (if any). Can be changed.
	Distro string `json:"distro,omitempty"`

	// Linux distribution version (if any). Can be changed.
	DistroVersion string `json:"distro_version,omitempty"`

	// Linux machine-id. Can't be changed. Must be unique across all Generic Nodes if specified.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier, can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name, can be changed.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this change generic node OK body generic
func (o *ChangeGenericNodeOKBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeGenericNodeOKBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeGenericNodeOKBodyGeneric) UnmarshalBinary(b []byte) error {
	var res ChangeGenericNodeOKBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
