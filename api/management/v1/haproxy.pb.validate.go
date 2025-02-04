// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: management/v1/haproxy.proto

package managementv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AddHAProxyServiceParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddHAProxyServiceParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddHAProxyServiceParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddHAProxyServiceParamsMultiError, or nil if none found.
func (m *AddHAProxyServiceParams) ValidateAll() error {
	return m.validate(true)
}

func (m *AddHAProxyServiceParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	// no validation rules for NodeName

	if all {
		switch v := interface{}(m.GetAddNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddHAProxyServiceParamsValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddHAProxyServiceParamsValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddHAProxyServiceParamsValidationError{
				field:  "AddNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := AddHAProxyServiceParamsValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Scheme

	// no validation rules for MetricsPath

	if val := m.GetListenPort(); val <= 0 || val >= 65536 {
		err := AddHAProxyServiceParamsValidationError{
			field:  "ListenPort",
			reason: "value must be inside range (0, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for CustomLabels

	// no validation rules for MetricsMode

	// no validation rules for SkipConnectionCheck

	if len(errors) > 0 {
		return AddHAProxyServiceParamsMultiError(errors)
	}

	return nil
}

// AddHAProxyServiceParamsMultiError is an error wrapping multiple validation
// errors returned by AddHAProxyServiceParams.ValidateAll() if the designated
// constraints aren't met.
type AddHAProxyServiceParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddHAProxyServiceParamsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddHAProxyServiceParamsMultiError) AllErrors() []error { return m }

// AddHAProxyServiceParamsValidationError is the validation error returned by
// AddHAProxyServiceParams.Validate if the designated constraints aren't met.
type AddHAProxyServiceParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddHAProxyServiceParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddHAProxyServiceParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddHAProxyServiceParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddHAProxyServiceParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddHAProxyServiceParamsValidationError) ErrorName() string {
	return "AddHAProxyServiceParamsValidationError"
}

// Error satisfies the builtin error interface
func (e AddHAProxyServiceParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddHAProxyServiceParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddHAProxyServiceParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddHAProxyServiceParamsValidationError{}

// Validate checks the field values on HAProxyServiceResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HAProxyServiceResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HAProxyServiceResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HAProxyServiceResultMultiError, or nil if none found.
func (m *HAProxyServiceResult) ValidateAll() error {
	return m.validate(true)
}

func (m *HAProxyServiceResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HAProxyServiceResultValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HAProxyServiceResultValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HAProxyServiceResultValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExternalExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HAProxyServiceResultValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HAProxyServiceResultValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExternalExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HAProxyServiceResultValidationError{
				field:  "ExternalExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HAProxyServiceResultMultiError(errors)
	}

	return nil
}

// HAProxyServiceResultMultiError is an error wrapping multiple validation
// errors returned by HAProxyServiceResult.ValidateAll() if the designated
// constraints aren't met.
type HAProxyServiceResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HAProxyServiceResultMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HAProxyServiceResultMultiError) AllErrors() []error { return m }

// HAProxyServiceResultValidationError is the validation error returned by
// HAProxyServiceResult.Validate if the designated constraints aren't met.
type HAProxyServiceResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HAProxyServiceResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HAProxyServiceResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HAProxyServiceResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HAProxyServiceResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HAProxyServiceResultValidationError) ErrorName() string {
	return "HAProxyServiceResultValidationError"
}

// Error satisfies the builtin error interface
func (e HAProxyServiceResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHAProxyServiceResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HAProxyServiceResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HAProxyServiceResultValidationError{}
