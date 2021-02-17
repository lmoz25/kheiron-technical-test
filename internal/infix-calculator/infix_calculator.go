package icalculator

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/internal/common"
)

type InfixCalculator struct {
	numberStack    common.Stack
	operationStack common.Stack
}

func (calculator *InfixCalculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	if sum[0] != "(" {
		return errors.New("Sum not in infix notation")
	}
	for _, section := range sum {
		if common.IsOperation(section) {
			calculator.operationStack.AddOperation(section)
		} else if section == "(" {
			// Don't care about open brackets
			continue
		} else if section == ")" {
			result, err := calculator.operate()
			if err != nil {
				return err
			}
			calculator.numberStack.AddNumber(result)
		} else {
			err := calculator.numberStack.AddNumberString(section)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Result outputs the result of the calculator's calculation
func (calculator *InfixCalculator) Result() (float32, error) {
	result, err := calculator.numberStack.GetNumber()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (calculator *InfixCalculator) ClearData() {
	calculator.numberStack = common.Stack{}
	calculator.operationStack = common.Stack{}
}

func (calculator *InfixCalculator) operate() (float32, error) {
	stackLength := calculator.numberStack.Len()
	// Need at least 2 numbers in the stack to operate on them
	if stackLength < 2 {
		return 0, fmt.Errorf("Cannot perform operation with only %d numbers", stackLength)
	}

	stackLength = calculator.operationStack.Len()
	// Need at least one operation on the stack to operate
	if stackLength < 1 {
		return 0, errors.New("Cannot perform operation with no operations on the stack")
	}

	// For division and subtraction, last on the stack will be second in the operation, and next last will be first
	num2, err := calculator.numberStack.GetNumber()
	if err != nil {
		return 0, err
	}

	num1, err := calculator.numberStack.GetNumber()
	if err != nil {
		return 0, err
	}

	operation, err := calculator.operationStack.GetOperation()

	return common.PerformOperation(num1, num2, operation)
}
