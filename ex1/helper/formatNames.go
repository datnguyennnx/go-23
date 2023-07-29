package helper

import (
	"errors"
	"fmt"
	"strings"
)

func FormatName(context []string, countryCode string) (string, error) {
	if len(context) <= 2 {
		panic(fmt.Errorf("invalid parameter"))
	}

	countryCode = strings.ToUpper(countryCode)

	firstNames := context[0]
	lastNames := context[1]
	middleNames := ""
	switch countryCode {
	case "VN":
		{
			if context[2] != countryCode {
				middleNames = strings.Join(context[2:len(context)-1], " ")
				return fmt.Sprintf("%s %s %s", lastNames, middleNames, firstNames), nil
			}
		}
		return fmt.Sprintf("%s %s", lastNames, firstNames), nil
	case "US":
		{
			if context[2] != countryCode {
				middleNames = strings.Join(context[2:len(context)-1], " ")
				return fmt.Sprintf("%s %s %s", firstNames, middleNames, lastNames), nil
			}
			return fmt.Sprintf("%s %s", firstNames, lastNames), nil
		}
	default:
		return "", errors.New("invalid parameter")
	}
}
