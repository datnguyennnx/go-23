package cmd

import (
	"flag"
	"fmt"

	"github.com/datnguyennnx/go-23/ex2/helper"
)

func Execute() {

	intFlag := flag.Bool("int", false, "is Int")
	floatFlag := flag.Bool("float", false, "is Float")
	stringFlag := flag.Bool("string", false, "is String")
	mixFlag := flag.Bool("mix", false, "")

	flag.Parse()
	args := flag.Args()

	if *intFlag {
		intArr, err := helper.StringToInt(args)
		if err != nil {
			panic(fmt.Errorf("%s", err))
		}
		helper.SortSlice(intArr)
		helper.Print(intArr)

	} else if *floatFlag {
		floatArr, err := helper.StringToFloat(args)
		if err != nil {
			panic(fmt.Errorf("%s", err))
		}
		helper.SortSlice(floatArr)
		helper.Print(floatArr)

	} else if *stringFlag {
		helper.SortSlice(args)
		helper.SortSlice(args)
		helper.Print(args)

	} else if *mixFlag {
		array, err := helper.CovertMixDataTypes(args)
		numArr, stringArr, err := helper.MixDataType(array, err)
		if err != nil {
			panic(fmt.Errorf("%s", err))
		}
		helper.SortSlice(numArr)
		helper.SortSlice(stringArr)
		helper.Print(numArr)
		helper.Print(stringArr)

	}
}
