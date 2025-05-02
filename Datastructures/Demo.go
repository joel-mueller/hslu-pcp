package Datastructures

import "fmt"

func Demo() {
	stack := Stack[int]{}
	stackList := StackList[int]{}
	stackList.Push(4)
	stack.Push(10)
	stack.Push(20)
	stack.Pop()
	stack.Pop()
	fmt.Println(GetStats(&stackList))
	fmt.Println(GetStats(&stack))
	fmt.Println(GetStatsStack(&stack))
	// fmt.Println(GetStatsStack(&stackList))
}
