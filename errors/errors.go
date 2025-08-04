package errors

type VError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TODO() VError {
	return VError{
		Status:  500,
		Code:    000,
		Message: "Not yet implemented",
	}
}
