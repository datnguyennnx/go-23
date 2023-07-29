package helper

import (
	"errors"
	"fmt"
	"testing"
)

func PanicCase() (string, error) {
	context := []string{"", ""}
	countryCode := ""
	return FormatName(context, countryCode)
}
func TestFormatName(t *testing.T) {
	testCase := []struct {
		context     []string
		countryCode string
		expected    string
		expectedErr error
	}{
		{
			[]string{"Emily", "Rose", "Waston", "us"},
			"us",
			"Emily Waston Rose",
			nil,
		},
		{
			[]string{"Emily", "Rose", "Waston", "Us"},
			"Us",
			"Emily Waston Rose",
			nil,
		},
		{
			[]string{"Emily", "Rose", "Waston", "US"},
			"US",
			"Emily Waston Rose",
			nil,
		},
		{
			[]string{"Loan", "Cong", "Tang", "Ton", "Nu", "Ta", "Thi", "Bach", "VN"},
			"VN",
			"Cong Tang Ton Nu Ta Thi Bach Loan",
			nil,
		},
		{
			[]string{"Loan", "Cong", "Tang", "Ton", "Nu", "Ta", "Thi", "Bach", "vn"},
			"vn",
			"Cong Tang Ton Nu Ta Thi Bach Loan",
			nil,
		},
		{
			[]string{"Loan", "Cong", "Tang", "Ton", "Nu", "Ta", "Thi", "Bach", "Vn"},
			"Vn",
			"Cong Tang Ton Nu Ta Thi Bach Loan",
			nil,
		},
		{
			[]string{"Emily", "Rose", "US"},
			"US",
			"Emily Rose",
			nil,
		},
		{
			[]string{"Dat", "Nguyen", "VN"},
			"VN",
			"Nguyen Dat",
			nil,
		},
		{
			[]string{"Dat", "Nguyen", ""},
			"",
			"",
			errors.New("invalid parameter"),
		},
		{
			[]string{"Ronaldo", "Lima", "Santos", ""},
			"",
			"",
			errors.New("invalid parameter"),
		},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("Input: %s", tc.context), func(t *testing.T) {
			result, _ := FormatName(tc.context, tc.countryCode)
			if result != tc.expected {
				t.Errorf("For input %s, expected %s, but got %s", tc.context, tc.expected, result)
			}
		})
	}

	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()
		PanicCase()
	})
}
