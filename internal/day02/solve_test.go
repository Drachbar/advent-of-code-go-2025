package day02

import (
	"os"
	"testing"
)

// PART 1

func TestSolvePart1ID11To22ShouldReturn33(t *testing.T) {
	input := "11-22"
	want := 33
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1ID95To115ShouldReturn99(t *testing.T) {
	input := "95-115"
	want := 99
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1ID1188511880To1188511890ShouldReturn99(t *testing.T) {
	input := "1188511880-1188511890"
	want := 1188511885
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1ExampleShouldReturn1227775554(t *testing.T) {
	data, _ := os.ReadFile("example.txt")
	want := 1227775554
	got := SolvePart1(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

// PART 2

func TestSolvePart2ID11To22ShouldReturn33(t *testing.T) {
	input := "11-22"
	want := 33
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart2ID95To115ShouldReturn210(t *testing.T) {
	input := "95-115"
	want := 210
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestSolvePart2ID11To22And95To115ShouldReturn210(t *testing.T) {
	input := "11-22,95-115"
	want := 243
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestSolvePart2ID565653To565659ShouldReturn565565(t *testing.T) {
	input := "565653-565659"
	want := 565656
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestSolvePart2ExampleShouldReturn16793817782(t *testing.T) {
	data, _ := os.ReadFile("example.txt")
	want := 4174379265
	got := SolvePart2(string(data))
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}
