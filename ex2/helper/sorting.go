package helper

import (
	"fmt"
	"sort"
)

type MixDataTypes interface {
	string | int64 | float64
}

func SortSlice[T MixDataTypes](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func Print[T MixDataTypes](s []T) {
	for _, value := range s {
		fmt.Printf("%v ", value)
	}
}
