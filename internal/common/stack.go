package common

import (
	"errors"
	"fmt"
	"strconv"
)

// Stack is an implementation of a stack
type Stack struct {
	Contents []interface{}
	Pointer  int
}

// Len gets the length of the stack
func (stack *Stack) Len() int {
	return (stack.Pointer + 1)
}

// Push pushes a value to the stack
func (stack *Stack) push(value interface{}) error {
	if stack.Len() == len(stack.Contents) {
		return errors.New("stack full")
	}
	stack.Pointer++
	stack.Contents[stack.Pointer] = value
	return nil
}

// Pop takes the value from the top of the stack
func (stack *Stack) pop() (interface{}, error) {
	if stack.Len() > 0 {
		retVal := stack.Contents[stack.Pointer]
		stack.Contents[stack.Pointer] = nil
		stack.Pointer--
		return retVal, nil
	}
	return 0, errors.New("stack empty")
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

	num, ok := numInterface.(float32)
	if !ok {
		return 0, errors.New("Tried to get float from stack, but top element not float")
	}
	return num, nil
}

// AddOperation is a wrapper for adding an arithmetic operation to the stack
func (stack *Stack) AddOperation(operation string) {
	stack.push(operation)
}

// GetOperation is a convenience function for safely getting an operation from the stack
func (stack *Stack) GetOperation() (string, error) {
	opInterface, err := stack.pop()
	if err != nil {
		return "", err
	}

	operation, ok := opInterface.(string)
	if !ok {
		return "", errors.New("Tried to get string from stack, but top element not string")
	}
	if !IsOperation(operation) {
		return "", fmt.Errorf("String at top of stack %s is not a valid operation", operation)
	}
	return operation, nil
}
