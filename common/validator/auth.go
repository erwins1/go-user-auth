package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
)

// validate request for register
func ValidateRegisterJSONBody(param generated.PostRegisterJSONBody) (err error) {
	// validate required params
	if param.FullName == nil || param.Password == nil || param.PhoneNumber == nil {
		return errors.New("full_name, password, phone_number is required")
	}
	// === start phone number validation == //
	err = validatePhoneNumber(*param.PhoneNumber)
	if err != nil {
		return err
	}
	// == end phone number validation == //

	// === start full name validation == //
	err = validateFullName(*param.PhoneNumber)
	if err != nil {
		return err
	}
	// === end full name validation == //

	// === start password validation == //
	if len(*param.Password) < 6 || len(*param.Password) > 64 {
		return errors.New("password must be between 6 and 64 characters")
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialRegex := regexp.MustCompile(`[^A-Za-z0-9]`)

	if !uppercaseRegex.MatchString(*param.Password) {
		return errors.New("password must contain at least 1 uppercase letter")
	}

	if !lowercaseRegex.MatchString(*param.Password) {
		return errors.New("password must contain at least 1 lowercase letter")
	}

	if !digitRegex.MatchString(*param.Password) {
		return errors.New("password must contain at least 1 digit")
	}

	if !specialRegex.MatchString(*param.Password) {
		return errors.New("password must contain at least 1 special character")
	}
	// === end password validation == //
	return
}

func ValidateLoginJSONBody(param generated.PostLoginJSONRequestBody) (err error) {
	// validate input
	if param.PhoneNumber == nil || param.Password == nil {
		return errors.New("phone_number, password is reqired")
	}
	// === start phone number validation == //
	if len(*param.PhoneNumber) < 10 || len(*param.PhoneNumber) > 13 {
		return errors.New("phone number must be between 10 and 13 characters")
	}

	if !strings.HasPrefix(*param.PhoneNumber, "+62") {
		return errors.New("phone number must start with +62")
	}
	// == end phone number validation == //

	return
}
