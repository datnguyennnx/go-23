package helper

import (
	"errors"
	"strings"
)

func ValidCountryCode(countryCode string) error {
	countryCode = strings.ToUpper(countryCode)
	switch countryCode {
	case "VN", "US":
		return nil
	default:
		return errors.New("invalid country code")
	}
}
