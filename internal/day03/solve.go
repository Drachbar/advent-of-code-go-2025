package day03

import (
	"strconv"
	"strings"
)

func SolvePart1(data string) int {
	result := 0

	for _, line := range strings.Split(data, "\n") {
		chars := []rune(line)
		firstValAsInt, _ := strconv.Atoi(string(chars[0]))
		lastValAsInt, _ := strconv.Atoi(string(chars[len(chars)-1]))
		firstValIndex := 0

		for i := 0; i < len(chars)-1; i++ {
			val, _ := strconv.Atoi(string(chars[i]))
			if val > firstValAsInt {
				firstValAsInt = val
				firstValIndex = i
			}
		}

		for i := firstValIndex + 1; i < len(chars); i++ {
			val, _ := strconv.Atoi(string(chars[i]))
			if val > lastValAsInt {
				lastValAsInt = val
			}
		}
		result += firstValAsInt*10 + lastValAsInt
	}

	return result
}

func SolvePart2(data string) int {
	return 0
}
