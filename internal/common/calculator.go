package common

// Contains common functions for use by calculators

import (
	"fmt"
)

// Operations is the list of characters that correspond to mathematical operations
// Holding a list of strings in this way is idiomatic in Golang, as it allows them to be
// hashed for fast lookups
var Operations = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
}

// IsOperation checks that the input string is a valid arithmetic operator
func IsOperation(candidate string) bool {
	_, isOperation := Operations[candidate]
	return isOperation
}

// PerformOperation performs the given operation on the given numbers
func PerformOperation(num1, num2 float32, operation string) (float32, error) {
	switch operation {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Operation %s not supported", operation)
	}
}
