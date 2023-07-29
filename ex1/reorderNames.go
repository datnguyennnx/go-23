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

	err := helper.ValidCountryCode(countryCode)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	result, err := helper.FormatName(context, countryCode)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Println(result)

}
