package Datastructures

import "fmt"

type Datastructures[T any] interface {
	Empty() bool
	Size() int
}

func GetStats(datastructures Datastructures[any]) string {
	if datastructures.Empty() {
		return "The datastructures is empty"
	}
	return fmt.Sprintf("The size is %d", datastructures.Size())
}
