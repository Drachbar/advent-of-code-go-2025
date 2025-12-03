package day01

import (
	"os"
	"testing"
)

// PART 1

func TestSolvePart1ExampleShouldBe3(t *testing.T) {
	data, _ := os.ReadFile("example.txt")
	want := 3
	got := SolvePart1(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestPart1ShouldReturn1WhenDialHitZeroOnce(t *testing.T) {
	input := "R5\nL55"
	want := 1

	got := SolvePart1(input)

	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}

}

func TestPart1ShouldReturn1WhenDialHitOneHundredOnce(t *testing.T) {
	input := "R25\nR25"
	want := 1

	got := SolvePart1(input)

	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestSolve1WithCorrectAnswers(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	want := 1177
	got := SolvePart1(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

// PART 2

func TestSolvePart2ExampleShouldBe6(t *testing.T) {
	data, _ := os.ReadFile("example.txt")
	want := 6
	got := SolvePart2(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}

func TestPart2ShouldReturn1WhenDialHit0Once(t *testing.T) {
	input := "R50"
	want := 1
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestPart2ShouldReturn1WhenDialClicksOnce(t *testing.T) {
	input := "R55\nR5"
	want := 1
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestPart2ShouldReturn3WhenDialClicksThreeTimes(t *testing.T) {
	input := "R250"
	want := 3
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestPart2ShouldReturn2WhenDialClicksTwiceWhileNegative(t *testing.T) {
	input := "L100\nL50"
	want := 2
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := "L50\nL300"
	want := 4
	got := SolvePart2(input)
	if got != want {
		t.Fatalf("SolvePart2() = %d, want %d", got, want)
	}
}

func TestSolve2WithCorrectAnswers(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	want := 6768
	got := SolvePart2(string(data))
	if got != want {
		t.Fatalf("SolvePart1() = %d, want %d", got, want)
	}
}
