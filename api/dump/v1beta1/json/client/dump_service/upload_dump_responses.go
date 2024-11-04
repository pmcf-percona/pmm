// Code generated by go-swagger; DO NOT EDIT.

package dump_service

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

// UploadDumpReader is a Reader for the UploadDump structure.
type UploadDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadDumpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadDumpOK creates a UploadDumpOK with default headers values
func NewUploadDumpOK() *UploadDumpOK {
	return &UploadDumpOK{}
}

/*
UploadDumpOK describes a response with status code 200, with default header values.

A successful response.
*/
type UploadDumpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this upload dump Ok response has a 2xx status code
func (o *UploadDumpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload dump Ok response has a 3xx status code
func (o *UploadDumpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload dump Ok response has a 4xx status code
func (o *UploadDumpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload dump Ok response has a 5xx status code
func (o *UploadDumpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload dump Ok response a status code equal to that given
func (o *UploadDumpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload dump Ok response
func (o *UploadDumpOK) Code() int {
	return 200
}

func (o *UploadDumpOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:upload][%d] uploadDumpOk %s", 200, payload)
}

func (o *UploadDumpOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:upload][%d] uploadDumpOk %s", 200, payload)
}

func (o *UploadDumpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadDumpDefault creates a UploadDumpDefault with default headers values
func NewUploadDumpDefault(code int) *UploadDumpDefault {
	return &UploadDumpDefault{
		_statusCode: code,
	}
}

/*
UploadDumpDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UploadDumpDefault struct {
	_statusCode int

	Payload *UploadDumpDefaultBody
}

// IsSuccess returns true when this upload dump default response has a 2xx status code
func (o *UploadDumpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upload dump default response has a 3xx status code
func (o *UploadDumpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upload dump default response has a 4xx status code
func (o *UploadDumpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upload dump default response has a 5xx status code
func (o *UploadDumpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upload dump default response a status code equal to that given
func (o *UploadDumpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upload dump default response
func (o *UploadDumpDefault) Code() int {
	return o._statusCode
}

func (o *UploadDumpDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:upload][%d] UploadDump default %s", o._statusCode, payload)
}

func (o *UploadDumpDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:upload][%d] UploadDump default %s", o._statusCode, payload)
}

func (o *UploadDumpDefault) GetPayload() *UploadDumpDefaultBody {
	return o.Payload
}

func (o *UploadDumpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UploadDumpDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UploadDumpBody upload dump body
swagger:model UploadDumpBody
*/
type UploadDumpBody struct {
	// dump ids
	DumpIds []string `json:"dump_ids"`

	// sftp parameters
	SftpParameters *UploadDumpParamsBodySftpParameters `json:"sftp_parameters,omitempty"`
}

// Validate validates this upload dump body
func (o *UploadDumpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSftpParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpBody) validateSftpParameters(formats strfmt.Registry) error {
	if swag.IsZero(o.SftpParameters) { // not required
		return nil
	}

	if o.SftpParameters != nil {
		if err := o.SftpParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "sftp_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "sftp_parameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upload dump body based on the context it is used
func (o *UploadDumpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSftpParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpBody) contextValidateSftpParameters(ctx context.Context, formats strfmt.Registry) error {
	if o.SftpParameters != nil {

		if swag.IsZero(o.SftpParameters) { // not required
			return nil
		}

		if err := o.SftpParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "sftp_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "sftp_parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpBody) UnmarshalBinary(b []byte) error {
	var res UploadDumpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpDefaultBody upload dump default body
swagger:model UploadDumpDefaultBody
*/
type UploadDumpDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UploadDumpDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this upload dump default body
func (o *UploadDumpDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this upload dump default body based on the context it is used
func (o *UploadDumpDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpDefaultBody) UnmarshalBinary(b []byte) error {
	var res UploadDumpDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpDefaultBodyDetailsItems0 upload dump default body details items0
swagger:model UploadDumpDefaultBodyDetailsItems0
*/
type UploadDumpDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// upload dump default body details items0
	UploadDumpDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *UploadDumpDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv UploadDumpDefaultBodyDetailsItems0

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
		o.UploadDumpDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o UploadDumpDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.UploadDumpDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.UploadDumpDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this upload dump default body details items0
func (o *UploadDumpDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upload dump default body details items0 based on context it is used
func (o *UploadDumpDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UploadDumpDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpParamsBodySftpParameters upload dump params body sftp parameters
swagger:model UploadDumpParamsBodySftpParameters
*/
type UploadDumpParamsBodySftpParameters struct {
	// address
	Address string `json:"address,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// directory
	Directory string `json:"directory,omitempty"`
}

// Validate validates this upload dump params body sftp parameters
func (o *UploadDumpParamsBodySftpParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upload dump params body sftp parameters based on context it is used
func (o *UploadDumpParamsBodySftpParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpParamsBodySftpParameters) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpParamsBodySftpParameters) UnmarshalBinary(b []byte) error {
	var res UploadDumpParamsBodySftpParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
