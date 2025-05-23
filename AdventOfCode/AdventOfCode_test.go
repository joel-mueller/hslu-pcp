package AdventOfCode

import (
	"testing"
)

func TestExercise1(t *testing.T) {
	start := []int{3, 4, 3, 1, 2}
	number := Advent(start, 80)
	if number != 5934 {
		t.Errorf(`test failed %v`, number)
	}
}

func TestExercise2(t *testing.T) {
	start := []int{3, 4, 3, 1, 2}
	number := Advent(start, 256)
	if number != 26984457539 {
		t.Errorf(`test failed %v`, number)
	}
}
