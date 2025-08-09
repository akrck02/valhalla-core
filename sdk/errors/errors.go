// Package verrors provides models and functions for better error handling
package verrors

import (
	"net/http"
)

type VError struct {
	Code    VErrorCode `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}

type APIError struct {
	Status int `json:"status,omitempty"`
	VError
}

func TODO() VError {
	return VError{
		Code:    NotImplemented,
		Message: "Not yet implemented",
	}
}

func Unexpected(message string) *VError {
	return &VError{
		Code:    UnexpectedError,
		Message: message,
	}
}

func New(code VErrorCode, message string) *VError {
	return &VError{
		Code:    code,
		Message: message,
	}
}

func NewAPIError(verror *VError) *APIError {
	var status int

	if 0 <= verror.Code && verror.Code <= 999 {
		status = http.StatusInternalServerError
	} else if 1000 <= verror.Code && verror.Code <= 3999 {
		status = http.StatusBadRequest
	} else if 4000 <= verror.Code && verror.Code <= 4999 {
		status = http.StatusNotFound
	} else if 5000 <= verror.Code && verror.Code <= 5999 {
		status = http.StatusUnauthorized
	} else if 6000 <= verror.Code && verror.Code <= 6999 {
		status = http.StatusForbidden
	} else {
		status = http.StatusTeapot
	}

	return &APIError{
		Status: status,
		VError: *verror,
	}
}

type VErrorCode int

const (
	// 0 --> 999 | SYSTEM UNEXPECTED ERRORS
	UnexpectedError            VErrorCode = 0
	DatabaseError              VErrorCode = 1
	NotImplemented             VErrorCode = 2
	NothingChanged             VErrorCode = 3
	CannotGenerateAuthToken    VErrorCode = 4
	CannotCreateValidationCode VErrorCode = 5

	// 1000 -> 3999 | VALIDATION ERRORS
	InvalidRequest        VErrorCode = 1000
	InvalidID             VErrorCode = 1001
	InvalidToken          VErrorCode = 1002
	InvalidPassword       VErrorCode = 1003
	InvalidEmail          VErrorCode = 1004
	InvalidValidationCode VErrorCode = 1005

	// 1100 -> 1299 | USER RELATED VALIDATION ERRORS
	UserAlreadyExists    VErrorCode = 1100
	UserAlreadyValidated VErrorCode = 1101

	// 1300 -> 1499 | PROJECT RELATED VALIDATION ERRORS
	ProjectAlreadyExists VErrorCode = 1300

	// 4000 -> 4999 | LOOKUP ERRORS
	NotFound VErrorCode = 4000

	// 5000 -> 5999 | AUTHORITATION ERRORS
	NotAuthorized VErrorCode = 5000

	// 6000 -> 7999 | PERMISSION ERRORS
	AccessDenied          VErrorCode = 6000
	NotEnoughtPermissions VErrorCode = 6001
)

const (
	AccessDeniedMessage string = "access denied"

	DatabaseConnectionEmptyMessage string = "database connection cannot be empty"
	ServiceIDEmptyMessage          string = "service id cannot be empty"
	RegisteredDomainsEmptyMessage  string = "registered domains cannot be empty"
	SecretEmptyMessage             string = "secret cannot be empty"

	TokenEmptyMessage   string = "token cannot be empty"
	TokenInvalidMessage string = "invalid token"

	PasswordEmptyMessage                string = "password cannot be empty"
	PasswordShortMessage                string = "password is short"
	PasswordNoNumericMessage            string = "password must contain at least one numeric character"
	PasswordNoSpecialCharacterMessage   string = "password must contain at least one special character"
	PasswordNoLowercaseCharacterMessage string = "password must contain at least one uppercase character"
	PasswordNoUppercaseCharacterMessage string = "password must contain at least one lowercase character"

	EmailEmptyMessage string = "email cannot be empty"

	UserEmptyMessage         string = "user cannot be empty"
	UserNotFoundMessage      string = "user not found"
	UserIDNegativeMessage    string = "user id must be positive"
	UserCannotDeleteMessage  string = "cannot delete user"
	UserCannotUpdateMessage  string = "cannot update user"
	UserAlreadyExistsMessage string = "user already exists"

	DeviceEmptyMessage          string = "device cannot be empty"
	DeviceAddressEmptyMessage   string = "device address cannot be empty"
	DeviceUserAgentEmptyMessage string = "device user agent cannot be empty"
)
