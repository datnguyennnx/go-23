package helper

import "fmt"

func ValidCountryCode(countryCode string) bool {
	switch countryCode {
	case "VN", "US":
		return true
	default:
		fmt.Printf("Invalid Country Code")
		return false
	}
}
