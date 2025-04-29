package main

import (
	"fmt"
	"hslu-pcp/advent_of_code"
)

func main() {
	start := []int{3, 4, 3, 1, 2}
	fmt.Println(advent_of_code.Advent(start, 80))
	fmt.Println(advent_of_code.Advent(start, 256))
}
