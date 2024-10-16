// Code generated by go-swagger; DO NOT EDIT.

package management_service

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

// AddAzureDatabaseReader is a Reader for the AddAzureDatabase structure.
type AddAzureDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAzureDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAzureDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddAzureDatabaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAzureDatabaseOK creates a AddAzureDatabaseOK with default headers values
func NewAddAzureDatabaseOK() *AddAzureDatabaseOK {
	return &AddAzureDatabaseOK{}
}

/*
AddAzureDatabaseOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddAzureDatabaseOK struct {
	Payload interface{}
}

// IsSuccess returns true when this add azure database Ok response has a 2xx status code
func (o *AddAzureDatabaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add azure database Ok response has a 3xx status code
func (o *AddAzureDatabaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add azure database Ok response has a 4xx status code
func (o *AddAzureDatabaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add azure database Ok response has a 5xx status code
func (o *AddAzureDatabaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add azure database Ok response a status code equal to that given
func (o *AddAzureDatabaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add azure database Ok response
func (o *AddAzureDatabaseOK) Code() int {
	return 200
}

func (o *AddAzureDatabaseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/management/services/azure][%d] addAzureDatabaseOk %s", 200, payload)
}

func (o *AddAzureDatabaseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/management/services/azure][%d] addAzureDatabaseOk %s", 200, payload)
}

func (o *AddAzureDatabaseOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AddAzureDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAzureDatabaseDefault creates a AddAzureDatabaseDefault with default headers values
func NewAddAzureDatabaseDefault(code int) *AddAzureDatabaseDefault {
	return &AddAzureDatabaseDefault{
		_statusCode: code,
	}
}

/*
AddAzureDatabaseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddAzureDatabaseDefault struct {
	_statusCode int

	Payload *AddAzureDatabaseDefaultBody
}

// IsSuccess returns true when this add azure database default response has a 2xx status code
func (o *AddAzureDatabaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add azure database default response has a 3xx status code
func (o *AddAzureDatabaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add azure database default response has a 4xx status code
func (o *AddAzureDatabaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add azure database default response has a 5xx status code
func (o *AddAzureDatabaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add azure database default response a status code equal to that given
func (o *AddAzureDatabaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add azure database default response
func (o *AddAzureDatabaseDefault) Code() int {
	return o._statusCode
}

func (o *AddAzureDatabaseDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/management/services/azure][%d] AddAzureDatabase default %s", o._statusCode, payload)
}

func (o *AddAzureDatabaseDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/management/services/azure][%d] AddAzureDatabase default %s", o._statusCode, payload)
}

func (o *AddAzureDatabaseDefault) GetPayload() *AddAzureDatabaseDefaultBody {
	return o.Payload
}

func (o *AddAzureDatabaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddAzureDatabaseDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddAzureDatabaseBody add azure database body
swagger:model AddAzureDatabaseBody
*/
type AddAzureDatabaseBody struct {
	// Azure database location.
	Region string `json:"region,omitempty"`

	// Azure database availability zone.
	Az string `json:"az,omitempty"`

	// Azure database instance ID.
	InstanceID string `json:"instance_id,omitempty"`

	// Represents a purchasable Stock Keeping Unit (SKU) under a product.
	// https://docs.microsoft.com/en-us/partner-center/develop/product-resources#sku.
	NodeModel string `json:"node_model,omitempty"`

	// Address used to connect to it.
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Unique across all Nodes user-defined name. Defaults to Azure Database instance ID.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all Services user-defined name. Defaults to Azure Database instance ID.
	ServiceName string `json:"service_name,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Azure client ID.
	AzureClientID string `json:"azure_client_id,omitempty"`

	// Azure client secret.
	AzureClientSecret string `json:"azure_client_secret,omitempty"`

	// Azure tanant ID.
	AzureTenantID string `json:"azure_tenant_id,omitempty"`

	// Azure subscription ID.
	AzureSubscriptionID string `json:"azure_subscription_id,omitempty"`

	// Azure resource group.
	AzureResourceGroup string `json:"azure_resource_group,omitempty"`

	// If true, adds azure_database_exporter.
	AzureDatabaseExporter bool `json:"azure_database_exporter,omitempty"`

	// If true, adds qan-mysql-perfschema-agent or qan-postgresql-pgstatements-agent.
	QAN bool `json:"qan,omitempty"`

	// Custom user-assigned labels for Node and Service.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// Tablestats group collectors will be disabled if there are more than that number of tables.
	// If zero, server's default value is used.
	// Use negative value to disable them.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

	// DiscoverAzureDatabaseType describes supported Azure Database instance engines.
	//
	//  - DISCOVER_AZURE_DATABASE_TYPE_MYSQL: MySQL type: microsoft.dbformysql or MariaDB type: microsoft.dbformariadb
	//  - DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL: PostgreSQL type: microsoft.dbformysql
	// Enum: ["DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED","DISCOVER_AZURE_DATABASE_TYPE_MYSQL","DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this add azure database body
func (o *AddAzureDatabaseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addAzureDatabaseBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED","DISCOVER_AZURE_DATABASE_TYPE_MYSQL","DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addAzureDatabaseBodyTypeTypePropEnum = append(addAzureDatabaseBodyTypeTypePropEnum, v)
	}
}

const (

	// AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEUNSPECIFIED captures enum value "DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED"
	AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEUNSPECIFIED string = "DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED"

	// AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEMYSQL captures enum value "DISCOVER_AZURE_DATABASE_TYPE_MYSQL"
	AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEMYSQL string = "DISCOVER_AZURE_DATABASE_TYPE_MYSQL"

	// AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEPOSTGRESQL captures enum value "DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"
	AddAzureDatabaseBodyTypeDISCOVERAZUREDATABASETYPEPOSTGRESQL string = "DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"
)

// prop value enum
func (o *AddAzureDatabaseBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addAzureDatabaseBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddAzureDatabaseBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add azure database body based on context it is used
func (o *AddAzureDatabaseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseBody) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddAzureDatabaseDefaultBody add azure database default body
swagger:model AddAzureDatabaseDefaultBody
*/
type AddAzureDatabaseDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddAzureDatabaseDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add azure database default body
func (o *AddAzureDatabaseDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAzureDatabaseDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add azure database default body based on the context it is used
func (o *AddAzureDatabaseDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAzureDatabaseDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddAzureDatabaseDefaultBodyDetailsItems0 add azure database default body details items0
swagger:model AddAzureDatabaseDefaultBodyDetailsItems0
*/
type AddAzureDatabaseDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// add azure database default body details items0
	AddAzureDatabaseDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *AddAzureDatabaseDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AddAzureDatabaseDefaultBodyDetailsItems0

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
		o.AddAzureDatabaseDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o AddAzureDatabaseDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.AddAzureDatabaseDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.AddAzureDatabaseDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this add azure database default body details items0
func (o *AddAzureDatabaseDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add azure database default body details items0 based on context it is used
func (o *AddAzureDatabaseDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
