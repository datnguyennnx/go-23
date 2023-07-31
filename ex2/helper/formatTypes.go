package helper

import (
	"errors"
	"reflect"
	"strconv"
)

func StringToInt(arrString []string) ([]int64, error) {
	result := make([]int64, len(arrString))
	var err error

	for index, value := range arrString {
		result[index], err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, errors.New("invalid datatype")
		}
	}
	return result, nil
}

func StringToFloat(arrString []string) ([]float64, error) {
	result := make([]float64, len(arrString))
	var err error

	for index, value := range arrString {
		result[index], err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("invalid datatype")
		}
	}
	return result, nil
}

func MixDataType(array []interface{}, err error) ([]float64, []string, error) {
	if err != nil {
		return nil, nil, errors.New("invalid datatype")
	}
	var numberArray []float64
	var stringArray []string

	arrayLen := len(array)

	for i := 0; i < arrayLen; i++ {
		switch dataType := reflect.TypeOf(array[i]).String(); dataType {
		case "int64":
			numberArray = append(numberArray, float64(array[i].(int64)))
		case "float64":
			numberArray = append(numberArray, array[i].(float64))
		case "string":
			stringArray = append(stringArray, array[i].(string))
		default:
			return nil, nil, errors.New("this value is not expected: ")
		}
	}
	return numberArray, stringArray, nil
}

func CovertMixDataTypes(array []string) ([]interface{}, error) {
	var result []interface{}
	arrayLen := len(array)

	for i := 0; i < arrayLen; i++ {
		temp := array[i]
		if isInt64(temp) {
			intElement, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				return nil, errors.New("error during conversion")
			}
			result = append(result, intElement)
		} else if isFloat64(temp) {
			floatElement, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				return nil, errors.New("error during conversion")
			}
			result = append(result, floatElement)
		} else {
			result = append(result, temp)
		}
	}
	return result, nil
}
