package common

import (
	"errors"
	"strconv"
)

type item struct {
	value interface{}
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
func (stack *Stack) push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

// Pop takes the value from the top of the stack
func (stack *Stack) pop() (interface{}, error) {
	if stack.Len() > 0 {
		value := stack.top.value
		stack.top = stack.top.next
		stack.size--
		return value, nil
	}
	return 0, errors.New("Stack empty")
}

// AddNumberString is a convenience function for adding a string containing a number to the stack
func (stack *Stack) AddNumberString(toAdd string) error {
	number, err := strconv.ParseInt(toAdd, 10, 64)
	if err != nil {
		return err
	}
	stack.push(float32(number))
	return nil
}

// AddNumber is a wrapper for adding a number to the stack
func (stack *Stack) AddNumber(toAdd float32) {
	stack.push(toAdd)
}

// GetNumber is a convenience function for getting a number safely from the stack
func (stack *Stack) GetNumber() (float32, error) {
	numInterface, err := stack.pop()
	if err != nil {
		return 0, err
	}

	retval, ok := numInterface.(float32)
	if !ok {
		return 0, errors.New("Tried to get float from stack, but top element not float")
	}
	return retval, nil
}

// AddOperation is a wrapper for adding an arithmetic operation to the stack
func (stack *Stack) AddOperation(operation string) {
	stack.push(operation)
}
