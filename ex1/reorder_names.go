package main

import (
	"fmt"
	"os"
)

func printResult(firstnName, lastName, middleName string) {
	if middleName == "" {
		fmt.Print(firstnName, " ", lastName)
	} else {
		fmt.Print(firstnName, " ", middleName, " ", lastName)
	}
}

func validContext(context []string, formatCode string) bool {
	if len(context) > 4 {
		fmt.Print("Invalid Input")
		return true
	} else if formatCode != "VN" && formatCode != "US" {
		fmt.Print("Invaild Country Code")
		return true
	} else if context[1] == formatCode {
		fmt.Print("Invalid Name ")
		return true
	}
	return false
}

func orderVietNam(context []string) (string, string, string) {
	firstname := context[1]
	lastname := context[0]
	var middleName string
	if context[2] != "VN" {
		firstname = context[1]
		middleName = context[2]
		lastname = context[0]
	}
	return firstname, lastname, middleName
}

func orderUS(context []string) (string, string, string) {
	firstname := context[0]
	lastname := context[1]
	var middleName string
	if context[2] != "US" {
		firstname = context[0]
		middleName = context[1]
		lastname = context[2]
	}
	return firstname, lastname, middleName
}

func main() {
	context := os.Args[1:]
	formatCode := context[len(context)-1]
	if validContext(context, formatCode) {
		os.Exit(0)
	} else if formatCode == "VN" {
		printResult(orderVietNam(context))
	} else {
		printResult(orderUS(context))
	}
}
