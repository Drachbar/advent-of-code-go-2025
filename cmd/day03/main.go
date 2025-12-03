package main

import (
	"fmt"
	"os"

	"main.go/internal/day03"
)

func main() {
	data, err := os.ReadFile("internal/day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(day03.SolvePart1(string(data)))
	fmt.Println(day03.SolvePart2(string(data)))

}
