package errors

type VError struct {
	Code    VErrorCode `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}

type ApiError struct {
	Status int `json:"status,omitempty"`
	Error  VError
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

type VErrorCode int

const (
	// 0 --> 999 | GENERAL ERRORS
	UnexpectedError            VErrorCode = 0
	AccessDenied               VErrorCode = 1
	NotImplemented             VErrorCode = 2
	InvalidRequest             VErrorCode = 3
	DatabaseError              VErrorCode = 4
	InvalidId                  VErrorCode = 5
	NotFound                   VErrorCode = 6
	NotEnoughtPermissions      VErrorCode = 7
	InvalidToken               VErrorCode = 8
	InvalidPassword            VErrorCode = 9
	InvalidEmail               VErrorCode = 10
	CannotCreateValidationCode VErrorCode = 11
	InvalidValidationCode      VErrorCode = 12
	NothingChanged             VErrorCode = 13

	// 1000 --> 1999 | USER ERRORS
	UserAlreadyExists    VErrorCode = 1000
	UserAlreadyValidated VErrorCode = 1001
	UserNotAuthorized    VErrorCode = 1002

	// 2000 --> 2999 | PROJECT ERRORS
	ProjectAlreadyExists VErrorCode = 2000

	// 3000 --> 3999 | TEAM ERRORS
	TeamAlreadyExists   VErrorCode = 3000
	UserIsAlreadyMember VErrorCode = 3001
	UserIsNotMember     VErrorCode = 3002

	// 4000 --> 4999 | DEVICE ERRORS
	DeviceNotFound          VErrorCode = 4000
	CannotGenerateAuthToken VErrorCode = 4001
)
