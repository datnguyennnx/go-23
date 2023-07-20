package main

import (
	"fmt"
	"os"
	"strings"
)

func isValidContext(context []string, formatCode string) bool {
	if formatCode != "VN" && formatCode != "US" {
		fmt.Print("Invaild Country Code")
		return true
	} else if context[1] == formatCode {
		fmt.Print("Invalid Input ")
		return true
	}
	return false
}

func orderNameFollowCountryCode(context []string, countryCode string) {
	firstName := context[0]
	lastName := context[1]
	switch countryCode {
	case "VN":
		{
			if context[2] != countryCode {
				firstName = context[1]
				lastName = context[0]
				middleName := strings.Join(context[2:len(context)-1], " ")

				fmt.Printf("%s %s %s", lastName, middleName, firstName)
				break
			}
			fmt.Printf("%s %s", lastName, firstName)
		}
	case "US":
		{
			if context[2] != countryCode {
				middleName := strings.Join(context[2:len(context)-1], " ")
				fmt.Printf("%s %s %s", firstName, middleName, lastName)
				break
			}
			fmt.Printf("%s %s", firstName, lastName)
		}
	}
}

func main() {
	context := os.Args[1:]
	formatCode := context[len(context)-1]
	if isValidContext(context, formatCode) {
		os.Exit(0)
	} else {
		orderNameFollowCountryCode(context, formatCode)
	}
}
