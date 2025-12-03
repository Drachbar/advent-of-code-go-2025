package main

import (
	"fmt"
	"os"

	"main.go/internal/day02"
)

func main() {
	data, err := os.ReadFile("internal/day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(day02.SolvePart1(string(data)))
	//fmt.Println(day01.SolvePart2(string(data)))

}
