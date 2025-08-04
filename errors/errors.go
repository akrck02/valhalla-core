package errors

import (
	"errors"
)

func TODO() error {
	return errors.New("Not yet implemented")
}
