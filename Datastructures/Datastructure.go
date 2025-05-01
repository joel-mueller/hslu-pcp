package Datastructures

import "fmt"

type Datastructures interface {
	Empty() bool
	Size() int
}

func GetStats(datastructures Datastructures) string {
	if datastructures.Empty() {
		return "The datastructures is empty"
	}
	return fmt.Sprintf("The size is %d", datastructures.Size())
}

func GetStatsStack[T any](stack *Stack[T]) string {
	if stack.Empty() {
		return "The stack is empty"
	}
	return fmt.Sprintf("The size is %d", stack.Size())
}
