package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/m/internal/common"
)

type Calculator struct {
	stack common.Stack
}

// Operations is the list of characters that correspond to mathematical operations
// Holding a list of strings in this way is idiomatic in Golang, as it allows them to be
// hashed for fast lookups
var Operations = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
}

func (calculator *Calculator) isOperation(candidate string) bool {
	_, isOperation := Operations[candidate]
	return isOperation
}

func (calculator *Calculator) addNumberToStack(toAdd string) error {
	number, err := strconv.ParseInt(toAdd, 10, 64)
	if err != nil {
		return err
	}
	calculator.stack.Push(float32(number))
	return nil
}

func (calculator *Calculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	for i := len(sum) - 1; i >= 0; i-- {
		character := sum[i]
		if calculator.isOperation(character) {
			result, err := calculator.operate(character)
			if err != nil {
				return err
			}
			calculator.stack.Push(result)
		} else {
			calculator.addNumberToStack(character)
		}
	}
	return nil
}

// Result outputs the result of the calculator's calculation
func (calculator *Calculator) Result() (float32, error) {
	result, err := calculator.stack.Pop()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (calculator *Calculator) operate(operation string) (float32, error) {
	stackLength := calculator.stack.Len()
	if stackLength < 2 {
		return 0, fmt.Errorf("Cannot perform operation with only %d numbers", stackLength)
	}
	num1, err := calculator.stack.Pop()
	if err != nil {
		return 0, err
	}
	num2, err := calculator.stack.Pop()
	if err != nil {
		return 0, err
	}

	return calculator.performOperation(num1, num2, operation)
}

func (calculator *Calculator) performOperation(num1, num2 float32, operation string) (float32, error) {
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
		return 0, fmt.Errorf("Operation %r not supported", operation)
	}
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
