package calculator

import (
	"fmt"
	"unicode"

	"gitlab.com/lmoz25/kheiron-technical-test/m/internal/common"
)

type Calculator struct {
	stack common.Stack
}

// Operations is the list of characters that correspond to mathematical operations
// Holding a list of strings in this way is idiomatic in Golang, as it allows them to be
// hashed for fast lookups
var Operations = map[rune]struct{}{
	'+': {},
	'-': {},
	'*': {},
	'/': {},
}

func (calculator *Calculator) isOperation(character rune) bool {
	_, isOperation := Operations[character]
	return isOperation
}

func (calculator *Calculator) addNumberToStack(toAdd rune) {
	number := int(toAdd - '0')
	calculator.stack.Push(float32(number))
}

func (calculator *Calculator) ParseInput(input string) (float32, error) {
	var finalResult float32
	sum := Reverse(input)
	for _, character := range sum {
		if calculator.isOperation(character) {
			result, err := calculator.operate(character)
			if err != nil {
				return 0, err
			}
			finalResult += result
		} else if unicode.IsSpace(character) {
			continue
		} else {
			calculator.addNumberToStack(character)
		}
	}
	return finalResult, nil
}

func (calculator *Calculator) operate(operation rune) (float32, error) {
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

func (calculator *Calculator) performOperation(num1, num2 float32, operation rune) (float32, error) {
	switch operation {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
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
