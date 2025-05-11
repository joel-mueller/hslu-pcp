package Datastructures

import "fmt"

// PCP-Übung Woche 1: Einstieg - C & Java revisited, 2. Aufgabe: ADT Stack in C (Array-Implementierung)

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
	// fmt.Println(GetStatsStack(&stackList)) // error, StackList ist nicht vom Typ Stack
}
