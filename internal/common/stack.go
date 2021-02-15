package common

import "errors"

type item struct {
	value float32
	next  *item
}

// Stack is an implementation of a stack
type Stack struct {
	top  *item
	size int
}

// Len gets the length of the stack
func (stack *Stack) Len() int {
	return stack.size
}

// Push pushes a value to the stack
func (stack *Stack) Push(value float32) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

// Pop takes the value from the top of the stack
func (stack *Stack) Pop() (float32, error) {
	if stack.Len() > 0 {
		value := stack.top.value
		stack.top = stack.top.next
		stack.size--
		return value, nil
	}
	return 0, errors.New("Stack empty")
}
