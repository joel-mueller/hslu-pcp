package advent_of_code

import "fmt"

func Advent(start []int, steps int) int {
	var storage = make(map[int]int)
	for i := 0; i < 9; i++ {
		storage[i] = 0
	}
	for _, val := range start {
		storage[val] = storage[val] + 1
	}

	for i := 0; i < steps; i++ {
		var zeros = storage[0]
		for i := 1; i < 9; i++ {
			storage[i-1] = storage[i]
		}
		storage[6] = storage[6] + zeros
		storage[8] = zeros
	}

	fmt.Printf("%v \n", storage)

	size := 0
	for i := 0; i < 9; i++ {
		size += storage[i]
	}
	return size
}
