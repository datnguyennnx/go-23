package helper

import "strconv"

func isInt64(data string) bool {
	_, err := strconv.ParseInt(data, 10, 64)

	if err == nil {
		return true
	} else {
		return false
	}
}

func isFloat64(data string) bool {
	_, err := strconv.ParseFloat(data, 64)

	if err == nil {
		return true
	} else {
		return false
	}
}
