// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/user.proto

package userv1

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

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserResponseMultiError, or nil if none found.
func (m *GetUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ProductTourCompleted

	// no validation rules for AlertingTourCompleted

	// no validation rules for SnoozedPmmVersion

	// no validation rules for SnoozedApiKeysMigration

	if len(errors) > 0 {
		return GetUserResponseMultiError(errors)
	}

	return nil
}

// GetUserResponseMultiError is an error wrapping multiple validation errors
// returned by GetUserResponse.ValidateAll() if the designated constraints
// aren't met.
type GetUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserResponseMultiError) AllErrors() []error { return m }

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.ProductTourCompleted != nil {
		// no validation rules for ProductTourCompleted
	}

	if m.AlertingTourCompleted != nil {
		// no validation rules for AlertingTourCompleted
	}

	if m.SnoozedPmmVersion != nil {
		// no validation rules for SnoozedPmmVersion
	}

	if m.SnoozedApiKeysMigration != nil {
		// no validation rules for SnoozedApiKeysMigration
	}

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on UpdateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserResponseMultiError, or nil if none found.
func (m *UpdateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ProductTourCompleted

	// no validation rules for AlertingTourCompleted

	// no validation rules for SnoozedPmmVersion

	// no validation rules for SnoozedApiKeysMigration

	if len(errors) > 0 {
		return UpdateUserResponseMultiError(errors)
	}

	return nil
}

// UpdateUserResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserResponseMultiError) AllErrors() []error { return m }

// UpdateUserResponseValidationError is the validation error returned by
// UpdateUserResponse.Validate if the designated constraints aren't met.
type UpdateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserResponseValidationError) ErrorName() string {
	return "UpdateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserResponseValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersRequestMultiError, or nil if none found.
func (m *ListUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUsersRequestMultiError(errors)
	}

	return nil
}

// ListUsersRequestMultiError is an error wrapping multiple validation errors
// returned by ListUsersRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersRequestMultiError) AllErrors() []error { return m }

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on ListUsersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponseMultiError, or nil if none found.
func (m *ListUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUsersResponseMultiError(errors)
	}

	return nil
}

// ListUsersResponseMultiError is an error wrapping multiple validation errors
// returned by ListUsersResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponseMultiError) AllErrors() []error { return m }

// ListUsersResponseValidationError is the validation error returned by
// ListUsersResponse.Validate if the designated constraints aren't met.
type ListUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponseValidationError) ErrorName() string {
	return "ListUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponseValidationError{}

// Validate checks the field values on ListUsersResponse_UserDetail with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse_UserDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse_UserDetail with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponse_UserDetailMultiError, or nil if none found.
func (m *ListUsersResponse_UserDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse_UserDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return ListUsersResponse_UserDetailMultiError(errors)
	}

	return nil
}

// ListUsersResponse_UserDetailMultiError is an error wrapping multiple
// validation errors returned by ListUsersResponse_UserDetail.ValidateAll() if
// the designated constraints aren't met.
type ListUsersResponse_UserDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponse_UserDetailMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponse_UserDetailMultiError) AllErrors() []error { return m }

// ListUsersResponse_UserDetailValidationError is the validation error returned
// by ListUsersResponse_UserDetail.Validate if the designated constraints
// aren't met.
type ListUsersResponse_UserDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponse_UserDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponse_UserDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponse_UserDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponse_UserDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponse_UserDetailValidationError) ErrorName() string {
	return "ListUsersResponse_UserDetailValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponse_UserDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse_UserDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponse_UserDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponse_UserDetailValidationError{}
