package day03

import (
	"testing"
)

// PART 1

func TestSolvePart1(t *testing.T) {
	input := ""
	want := 0
	got := SolvePart1(input)
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}
