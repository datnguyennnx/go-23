package helper

import (
	"strings"
)

func FormatName(context []string, countryCode string) (firstNames string, middleNames string, lastNames string) {
	firstNames = context[0]
	lastNames = context[1]
	middleNames = ""
	switch countryCode {
	case "VN":
		{
			if context[2] != countryCode {
				middleNames = strings.Join(context[2:len(context)-1], " ")
			}
		}
		return firstNames, middleNames, lastNames
	case "US":
		{
			if context[2] != countryCode {
				middleNames = strings.Join(context[2:len(context)-1], " ")
			}
			return firstNames, middleNames, lastNames
		}

	default:
		return firstNames, middleNames, lastNames
	}
}
