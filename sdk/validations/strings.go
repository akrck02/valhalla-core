package validations

import (
	"errors"
	"net/mail"
	"strings"

	sdkerrors "github.com/akrck02/valhalla-core/sdk/errors"
)

func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New(sdkerrors.PasswordEmptyMessage)
	}

	if len(password) < 16 {
		return errors.New(sdkerrors.PasswordShortMessage)
	}

	if !strings.ContainsAny(password, "123456789") {
		return errors.New(sdkerrors.PasswordNoNumericMessage)
	}

	if !strings.ContainsAny(password, "*¡!¿?$%&/()@#~¬") {
		return errors.New(sdkerrors.PasswordNoSpecialCharacterMessage)
	}

	if password == strings.ToLower(password) {
		return errors.New(sdkerrors.PasswordNoUppercaseCharacterMessage)
	}

	if password == strings.ToUpper(password) {
		return errors.New(sdkerrors.PasswordNoLowercaseCharacterMessage)
	}

	return nil
}

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New(sdkerrors.EmailEmptyMessage)
	}

	_, err := mail.ParseAddress(email)
	if nil != err {
		return err
	}

	return nil
}
