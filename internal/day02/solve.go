package day02

import (
	"fmt"
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

func SolvePart2(data string) int {
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
			found := false
			for j := 2; length/j >= 1 && !found; j++ {
				result, err := splitInParts(valueToCheck, j)
				if err == nil && allEqual(result) {
					found = true
					value += valueToCheck
				}
			}
		}
	}

	return value
}

func splitInParts(n int, parts int) ([]int, error) {
	s := strconv.Itoa(n)
	if len(s)%parts != 0 {
		return nil, fmt.Errorf("går inte att dela jämnt")
	}

	size := len(s) / parts
	out := make([]int, parts)

	for i := 0; i < parts; i++ {
		part := s[i*size : (i+1)*size]
		v, _ := strconv.Atoi(part)
		out[i] = v
	}

	return out, nil
}

func allEqual[T comparable](a []T) bool {
	if len(a) == 0 {
		return false
	}

	first := a[0]
	for _, v := range a {
		if v != first {
			return false
		}
	}
	return true
}
