package icalculator

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/internal/common"
)

// InfixCalculator is the struct representing a calculator that operates on infix expressions
type InfixCalculator struct {
	numberStack    common.Stack
	operationStack common.Stack
}

// ParseInput is the function through which an infix calculator parses infix expressions
func (calculator *InfixCalculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	// Allocate more than enough space on both stacks ahead of time.
	// For sums more complicated (i.e with more operations) than ( a + b ), when we add new symbols (brackets/numbers
	// operators), at most 2/5 of the added symbols will be numbers and at most a quarter will be an operator.
	// So we can safely allocate 2/5 of length of the input string as stack size to numbers, and 1/4 to operators.
	numberStackLen := int(2.0/5.0*float32(len(sum))) + 1
	operatorStackLen := int(1.0/4.0*float32(len(sum))) + 1
	calculator.operationStack.Contents = make([]interface{}, operatorStackLen)
	calculator.numberStack.Contents = make([]interface{}, numberStackLen)
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

// ClearData wipes the stacks for the calculator
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
