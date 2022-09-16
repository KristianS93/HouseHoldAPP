package validation

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"unicode"
)

// CheckEmail validates whether the input email
// is a valid email or not, also does an MX lookup on domain.
//
// Returns: a bool representing the validity of the provided email
// true meaning valid and false meaning invalid.
func CheckEmail(e string) []string {
	// splitting in prefix and domain
	//! should refactor magic strings, also possibly invalidChar helpers
	xs := strings.Split(e, "@")
	if len(xs) != 2 {
		// a valid email only contains one instance of @
		return []string{"An email must only contain one instance of @."}
	}

	// checking prefix format, may contains a-z, A-Z, 0-9 and .-_ (however not consecutive)
	// anything other than this, or consecutive cases of .-_ is considered invalid
	var invalidChars, errMessages []string
	for i, v := range xs[0] {
		switch {
		case unicode.IsUpper(v) || unicode.IsLower(v) || unicode.IsNumber(v):
			continue
		default:
			if strings.ContainsAny(string(v), ".-_") {
				// checking not out of bounds and if consecutive case of .-_
				if (i+1 <= len(xs[0])) && (strings.ContainsAny(string(xs[0][i+1]), ".-_")) {
					errMessages = append(errMessages, "Invalid case of repeated punctuation (.-_).")
				}
				// valid case of punctuation => skip to next rune
				continue
			}
			// only reached if invalid rune
			invalidChars = append(invalidChars, string(v))
		}
	}

	// looking up domain, returns a list of valid domains under lookup
	// err is returned when no host under specified domain exists
	// very very giga unlikely, but technically error could be
	// caused by the DNS message not being unmarshalled correctly
	_, err := net.LookupMX(xs[1])
	if err != nil {
		if errors.Is(err, errors.New("no such host")) {
			errMessages = append(errMessages, "Email host is invalid.")
		} else {
			return []string{"External DNS error."}
		}
	}

	if len(invalidChars) != 0 {
		errMessages = append(errMessages, emailErrInvalidChar(invalidChars))
	}

	return errMessages
	//todo: testing for the function is not done currently
	// should technically do an email verification at this point
	// both to double check, but also to verify that the correct
	// email is associated with a given user, as they might have
	// misspelled or had a typo when entering their email
}

func emailErrInvalidChar(xs []string) string {
	return fmt.Sprintf("Email contains invalid characters: %s", strings.Join(xs, ", "))
}

type errNrPassword uint8

const (
	passwordErrLength = "Password is not an appropriate length."
	passwordErrUpper  = "Password does not contain an upper case letter."
	passwordErrLower  = "Password does not contain a lower case letter."
	passwordErrNumber = "Password does not contain a number."
	passwordErrSymbol = "Password does not contain a symbol."
)

const (
	upperErrNr errNrPassword = iota
	lowerErrNr
	numberErrNr
	symbolErrNr
)

// CheckPassword validates if the input password
// upholds the criteria specified in the function.
// By default the password must contain:
//
// - 1 uppercase letter
//
// - 1 lowercase letter
//
// - 1 number
//
// - 1 special character (! @ # $ % ^ & *)
//
// And have a length between 8 and 32, inclusive.
//
// Additionally, the function validates that no invalid
// characters are present in the input password.
//
// Returns: a string array containing all the errors
// found in the input password, it is empty for valid passwords.
func CheckPassword(p string) []string {
	requirements := [4]bool{}
	var invalidChars []string

	if !(len(p) >= 8) || !(len(p) <= 32) {
		// failsafe to not check millions of characters
		return []string{passwordErrLength}
	}

	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			requirements[upperErrNr] = true
		case unicode.IsLower(char):
			requirements[lowerErrNr] = true
		case unicode.IsNumber(char):
			requirements[numberErrNr] = true
		case strings.ContainsAny(string(char), "!@#$%^&*"):
			requirements[symbolErrNr] = true
		default:
			invalidChars = append(invalidChars, string(char))
		}
	}

	var errMessages []string
	for i, v := range requirements {
		// checking if all requirements are true => must flip sign,
		// so, negating value, so true -> false and vice versa
		if !v {
			errMessages = append(errMessages, passwordErrString(errNrPassword(i)))
		}
	}
	if len(invalidChars) != 0 {
		errMessages = append(errMessages, passwordErrInvalidChar(invalidChars))
	}

	if len(errMessages) != 0 {
		return errMessages
	} else {
		return nil
	}
}

// passwordErrInvalidChar is a helper function for CheckPassword
// used to generate a string representing an error message
// containing all invalid characters in the provided password.
func passwordErrInvalidChar(xs []string) string {
	return fmt.Sprintf("Password contains invalid characters: %s", strings.Join(xs, ", "))
}

// passwordErrString is a wrapper for a switch used in CheckPassword.
//
// Returns: a string value representing an error.
func passwordErrString(index errNrPassword) string {
	switch index {
	case upperErrNr:
		return passwordErrUpper
	case lowerErrNr:
		return passwordErrLower
	case numberErrNr:
		return passwordErrNumber
	case symbolErrNr:
		return passwordErrSymbol
	}
	return ""
}
