package day02

import (
	"strconv"
	"strings"
)

func SolvePart1(data string) int {
	value := 0
	ranges := strings.Split(data, ",")
	for _, rangeToCheck := range ranges {
		temp := strings.Split(rangeToCheck, "-")
		startVal, _ := strconv.Atoi(temp[0])
		endVal, _ := strconv.Atoi(temp[1])
		for i := range endVal - startVal + 1 {
			valueToCheck := startVal + i
			stringValueToCheck := strconv.Itoa(valueToCheck)
			length := len(stringValueToCheck)
			// Bara jämnt antal tecken
			if length%2 == 0 {
				mid := len(stringValueToCheck) / 2
				left, _ := strconv.Atoi(stringValueToCheck[:mid])
				right, _ := strconv.Atoi(stringValueToCheck[mid:])
				if left == right {
					value += valueToCheck
				}
			}
		}
	}

	return value
}

func SolvePart2(data string) {
}
