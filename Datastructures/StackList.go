package Datastructures

type Element[T any] struct {
	next  *Element[T]
	value T
}

type StackList[T any] struct {
	head *Element[T]
}

func (stack *StackList[T]) Push(value T) {
	newElem := &Element[T]{
		value: value,
		next:  stack.head,
	}
	stack.head = newElem
}

func (stack *StackList[T]) Pop() (value T) {
	if stack.head == nil {
		panic("stack is empty")
	}
	value = stack.head.value
	stack.head = stack.head.next
	return value
}

func (stack *StackList[T]) Peek() (value T) {
	if stack.head == nil {
		panic("stack is empty")
	}
	return stack.head.value
}

func (stack *StackList[T]) Empty() bool {
	return stack.head == nil
}

func (stack *StackList[T]) Size() int {
	size := 0
	head := stack.head
	for head != nil {
		head = head.next
		size++
	}
	return size
}
