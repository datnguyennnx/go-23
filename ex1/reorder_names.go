package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/datnguyennnx/go-23/ex1/helper"
)

func main() {
	context := os.Args[1:]
	countryCode := strings.ToUpper(context[len(context)-1])
	if len(context) <= 2 {
		fmt.Print("Invalid Input")
		os.Exit(0)
	}
	if helper.ValidCountryCode(countryCode) {
		firstNames, middleNames, lastNames := helper.FormatName(context, countryCode)
		switch countryCode {
		case "VN":
			{
				if middleNames == "" {
					fmt.Printf("%s %s", lastNames, firstNames)
				} else {
					fmt.Printf("%s %s %s", lastNames, middleNames, firstNames)
				}
			}
		case "US":
			{
				if middleNames == "" {
					fmt.Printf("%s %s", firstNames, lastNames)
				} else {
					fmt.Printf("%s %s %s", firstNames, middleNames, lastNames)
				}
			}
		}
	} else {
		os.Exit(0)
	}
}
