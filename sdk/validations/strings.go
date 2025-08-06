package validations

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidatePassword(password string) error {

	if "" == strings.TrimSpace(password) {
		return errors.New("Password cannot be empty.")
	}

	if len(password) < 16 {
		return errors.New("Password is short.")
	}

	if !strings.ContainsAny(password, "123456789") {
		return errors.New("Password must contain at least one numeric character.")
	}

	if !strings.ContainsAny(password, "*¡!¿?$%&/()@#~¬") {
		return errors.New("Password must contain at least one special character.")
	}

	if password == strings.ToLower(password) {
		return errors.New("Password must contain at least one uppercase character.")
	}

	if password == strings.ToUpper(password) {
		return errors.New("Password must contain at least one lowercase character.")
	}

	return nil
}

func ValidateEmail(email string) error {

	if "" == strings.TrimSpace(email) {
		return errors.New("Email cannot be empty.")
	}

	_, err := mail.ParseAddress(email)
	if nil != err {
		return err
	}

	return nil
}
