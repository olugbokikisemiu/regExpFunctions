package scenario

import (
	"fmt"
	"regexp"
)

// IsStringNoSpace returns true if s contains only letters without a space
func IsStringNoSpace(s string) (bool, error) {
	isValid, err := regexp.MatchString(`^[A-z]+$`, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsStringSpace returns true if s contains only letters with a space
func IsStringSpace(s string) (bool, error) {
	isValid, err := regexp.MatchString(`^[A-z' ']+$`, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsStringLength returns true if s contains only letters and the length specified matches
func IsStringLength(s string, length int) (bool, error) {
	pattern := fmt.Sprintf(`^[A-z]{%d}$`, length)
	isValid, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsNumberOnly returns true if s contains only number
func IsNumber(s string) (bool, error) {
	isValid, err := regexp.MatchString(`^\d+$`, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsNumberLength returns true if s contains only number
func IsNumberLength(s string, length int) (bool, error) {
	pattern := fmt.Sprintf(`^\d{%d}$`, length)
	isValid, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsEmail returns true if s is a valid email
func IsEmail(s string) (bool, error) {
	isValid, err := regexp.MatchString(`@`, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}
