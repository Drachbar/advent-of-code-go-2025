package day03

import (
	"os"
	"testing"
)

// PART 1

func TestSolvePart1Input12Eq12(t *testing.T) {
	input := "12"
	want := 12
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1Input123Eq23(t *testing.T) {
	input := "123"
	want := 23
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1Input321Eq32(t *testing.T) {
	input := "321"
	want := 32
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1Input321n12Eq44(t *testing.T) {
	input := "321\n12"
	want := 44
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolvePart1ExampleShouldReturn357(t *testing.T) {
	data, _ := os.ReadFile("example.txt")
	want := 357
	got := SolvePart1(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}
