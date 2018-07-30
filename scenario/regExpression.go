package scenario

import (
	"fmt"
	"regexp"
)

var (
	matchStringWithNoSpace = regexp.MustCompile(`^[A-z]+$`)
	matchStringWithSpace   = regexp.MustCompile(`^[A-z' ']+$`)
	matchNumbers           = regexp.MustCompile(`^\d+$`)
	matchEmail             = regexp.MustCompile(`@`)
)

// IsStringNoSpace returns true if s contains only letters without a space
func IsStringNoSpace(s string) bool {
	return matchStringWithNoSpace.MatchString(s)
}

// IsStringSpace returns true if s contains only letters with a space
func IsStringSpace(s string) bool {
	return matchStringWithNoSpace.MatchString(s)
}

// IsAlphaNumericWithLength returns true if s contains only letters and the length specified matches
func IsAlphaNumericWithLength(s string, length int) (bool, error) {
	pattern := fmt.Sprintf(`^[A-z]{%d}$`, length)
	isValid, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsNumeric returns true if s contains only number
func IsNumeric(s string) bool {
	return matchNumbers.MatchString(s)
}

// IsNumericWithLength returns true if s contains only number
func IsNumericWithLength(s string, length int) (bool, error) {
	pattern := fmt.Sprintf(`^\d{%d}$`, length)
	isValid, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, fmt.Errorf("An error occured: %s ", err)
	}
	return isValid, nil
}

// IsEmail returns true if s is a valid email
func IsEmail(s string) bool {
	return matchEmail.MatchString(s)
}
