package util

import (
	"strconv"
)

func ToIntSafely(input string) int {
	intS, err := strconv.Atoi(input)
	if err != nil {
		// fmt.Println("ToIntSafely error:", err, input)
		intS = 0
	}

	return intS
}
