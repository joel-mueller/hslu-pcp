package AdventOfCode

import "fmt"

func Demo() {
	steps := 30
	start := []int{3, 4, 3, 1, 2}
	fmt.Printf("The slice is %v big\n", Advent(start, steps))
}
