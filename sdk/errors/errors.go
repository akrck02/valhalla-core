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
		Code:    NotImplementedErrorCode,
		Message: "Not yet implemented",
	}
}

func Unexpected(message string) *VError {
	return &VError{
		Code:    UnexpectedErrorCode,
		Message: message,
	}
}

func DatabaseError(message string) *VError {
	return &VError{
		Code:    InvalidRequestErrorCode,
		Message: message,
	}
}

func NotFound(message string) *VError {
	return &VError{
		Code:    NotFoundErrorCode,
		Message: message,
	}
}

func AccessDenied(message string) *VError {
	return &VError{
		Code:    AccessDeniedErrorCode,
		Message: message,
	}
}

func Unauthorized(message string) *VError {
	return &VError{
		Code:    UnauthorizedErrorCode,
		Message: message,
	}
}

func InvalidRequest(message string) *VError {
	return &VError{
		Code:    InvalidRequestErrorCode,
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
	UnexpectedErrorCode                 VErrorCode = 0
	DatabaseErrorCode                   VErrorCode = 1
	NotImplementedErrorCode             VErrorCode = 2
	NothingChangedErrorCode             VErrorCode = 3
	CannotGenerateAuthTokenErrorCode    VErrorCode = 4
	CannotCreateValidationCodeErrorCode VErrorCode = 5

	// 1000 -> 3999 | VALIDATION ERRORS
	InvalidRequestErrorCode VErrorCode = 1000

	// 1100 -> 1299 | USER RELATED VALIDATION ERRORS
	UserAlreadyExistsErrorCode    VErrorCode = 1100
	UserAlreadyValidatedErrorCode VErrorCode = 1101

	// 1300 -> 1499 | PROJECT RELATED VALIDATION ERRORS
	ProjectAlreadyExistsErrorCode VErrorCode = 1300

	// 4000 -> 4999 | LOOKUP ERRORS
	NotFoundErrorCode VErrorCode = 4000

	// 5000 -> 5999 | AUTHORITATION ERRORS
	UnauthorizedErrorCode VErrorCode = 5000

	// 6000 -> 7999 | PERMISSION ERRORS
	AccessDeniedErrorCode          VErrorCode = 6000
	NotEnoughtPermissionsErrorCode VErrorCode = 6001
)

const (
	AccessDeniedMessage string = "access denied"

	CannotConnectToDatabaseMessage string = "cannot connect to database"
	DatabaseConnectionEmptyMessage string = "database connection cannot be empty"
	ServiceIDEmptyMessage          string = "service id cannot be empty"
	RegisteredDomainsEmptyMessage  string = "registered domains cannot be empty"
	SecretEmptyMessage             string = "secret cannot be empty"

	TokenEmptyMessage   string = "token cannot be empty"
	TokenInvalidMessage string = "invalid token"

	FileTooLargeMessage string = "%s is too long; the maximum size is %dMB"

	PasswordEmptyMessage                string = "password cannot be empty"
	PasswordShortMessage                string = "password is short"
	PasswordNoNumericMessage            string = "password must contain at least one numeric character"
	PasswordNoSpecialCharacterMessage   string = "password must contain at least one special character"
	PasswordNoLowercaseCharacterMessage string = "password must contain at least one uppercase character"
	PasswordNoUppercaseCharacterMessage string = "password must contain at least one lowercase character"

	EmailEmptyMessage string = "email cannot be empty"

	UserEmptyMessage               string = "user cannot be empty"
	UserNotFoundMessage            string = "user not found"
	UserIDNegativeMessage          string = "user id must be positive"
	UserCannotDeleteMessage        string = "cannot delete user"
	UserCannotUpdateMessage        string = "cannot update user"
	UserAlreadyExistsMessage       string = "user already exists"
	UserProfilePictureEmptyMessage string = "no profile picture provided"

	DeviceEmptyMessage          string = "device cannot be empty"
	DeviceAddressEmptyMessage   string = "device address cannot be empty"
	DeviceUserAgentEmptyMessage string = "device user agent cannot be empty"
)
