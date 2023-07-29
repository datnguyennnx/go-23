package helper

import (
	"fmt"
	"testing"
)

func ErrorCase() error {
	return ValidCountryCode("")
}
func TestValidCountryCode(t *testing.T) {
	testCase := []struct {
		countryCode string
		expected    error
	}{
		{"VN", nil},
		{"Vn", nil},
		{"vN", nil},
		{"vn", nil},
		{"US", nil},
		{"Us", nil},
		{"uS", nil},
		{"us", nil},
	}
	for _, tc := range testCase {
		t.Run(fmt.Sprintf("Input: %s", tc.countryCode), func(t *testing.T) {
			result := ValidCountryCode(tc.countryCode)
			if result != tc.expected {
				t.Errorf("For input %s, expected %s, but got %s", tc.countryCode, tc.expected, result)
			}
		})
	}
	t.Run("Error", func(t *testing.T) {
		result := ErrorCase()
		expected := "invalid country code"
		if result.Error() != expected {
			t.Errorf("Error result = %s, and Expected = %s.", result, expected)
		}
	})
}
