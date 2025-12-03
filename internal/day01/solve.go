package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(data string) int {
	var dial = 50
	var numberOfZeroes = 0

	for _, line := range strings.Split(data, "\n") {
		direction := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Gick inte att konvertera raden:", line, err)
			panic(err)
		}

		if strings.EqualFold(direction, "r") {
			dial += amount
		} else {
			dial -= amount
		}

		if dial%100 == 0 {
			dial = 0
		}

		if dial == 0 {
			numberOfZeroes++
		}
	}

	return numberOfZeroes
}

func SolvePart2(data string) int {
	var dial = 50
	var numberOfZeroClicks = 0

	for _, line := range strings.Split(data, "\n") {
		direction := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Gick inte att konvertera raden:", line, err)
			panic(err)
		}

		for amount >= 100 {
			amount -= 100
			numberOfZeroClicks++
		}

		if strings.EqualFold(direction, "r") {
			if dial < 0 {
				dial += amount
				if dial >= 0 {
					numberOfZeroClicks++
				}
			} else {
				dial += amount
				if dial >= 100 {
					numberOfZeroClicks++
					dial -= 100
				}
			}
		} else {
			if dial > 0 {
				dial -= amount
				if dial <= 0 {
					numberOfZeroClicks++
				}
			} else {
				dial += 100
				dial -= amount
				if dial <= 0 {
					numberOfZeroClicks++
				}
			}
		}

	}

	return numberOfZeroClicks
}
