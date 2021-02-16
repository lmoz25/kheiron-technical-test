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

func IsOperation(candidate string) bool {
	_, isOperation := Operations[candidate]
	return isOperation
}

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
