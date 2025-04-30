package Stack

type Stack[T any] struct {
	stack []T
}

func (stack *Stack[T]) Push(value T) {
	stack.stack = append(stack.stack, value)
}

func (stack *Stack[T]) Pop() T {
	if len(stack.stack) == 0 {
		panic("stack is empty")
	}
	top := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return top
}

func (stack *Stack[T]) Peek() T {
	if len(stack.stack) == 0 {
		panic("stack is empty")
	}
	return stack.stack[len(stack.stack)-1]
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.stack) == 0
}

func (stack *Stack[T]) Len() int {
	return len(stack.stack)
}
