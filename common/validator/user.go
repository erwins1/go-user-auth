package validator

import (
	"errors"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
)

func ValidateBodyUpdateUser(param generated.PutProfileJSONRequestBody) (err error) {
	if param.FullName == nil && param.PhoneNumber == nil {
		return errors.New("at least need full_name or phone_number")
	}
	if param.FullName != nil {
		err = validateFullName(*param.PhoneNumber)
		if err != nil {
			return err
		}
	}
	if param.PhoneNumber != nil {
		err = validatePhoneNumber(*param.PhoneNumber)
		if err != nil {
			return err
		}
	}
	return
}

func validatePhoneNumber(phoneNumber string) (err error) {
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return errors.New("phone number must be between 10 and 13 characters")
	}

	if !strings.HasPrefix(phoneNumber, "+62") {
		return errors.New("phone number must start with +62")
	}
	return
}

func validateFullName(fullName string) (err error) {
	if len(fullName) < 3 || len(fullName) > 60 {
		return errors.New("full name must be between 3 and 60 characters")
	}
	return
}
