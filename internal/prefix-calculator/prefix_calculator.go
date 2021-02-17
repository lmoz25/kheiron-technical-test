package pcalculator

import (
	"fmt"
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/internal/common"
)

// PrefixCalculator is the struct representing a calculator that operates on prefix expressions
type PrefixCalculator struct {
	stack common.Stack
}

// ParseInput is the function through which an infix calculator parses infix expressions
func (calculator *PrefixCalculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	// Allocate more than enough space on the stack ahead of time.
	// The highest ratio of numbers to operators is 2:1, i.e for the expression `+ a b` or similar,
	// since this is the sum that never involves calculations involving numbers put on the stack after previous
	// calculations. So allocating more than 2/3 of the length of the string is a decent upper bound for size.
	stackLen := int(2.0/3.0*float32(len(sum))) + 1
	calculator.stack.Contents = make([]interface{}, stackLen)
	for i := len(sum) - 1; i >= 0; i-- {
		section := sum[i]
		if common.IsOperation(section) {
			result, err := calculator.operate(section)
			if err != nil {
				return err
			}
			calculator.stack.AddNumber(result)
		} else {
			err := calculator.stack.AddNumberString(section)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Result outputs the result of the calculator's calculation
func (calculator *PrefixCalculator) Result() (float32, error) {
	result, err := calculator.stack.GetNumber()
	if err != nil {
		return 0, err
	}

	return result, nil
}

// ClearData wipes the stack for the calculator
func (calculator *PrefixCalculator) ClearData() {
	calculator.stack = common.Stack{}
}

func (calculator *PrefixCalculator) operate(operation string) (float32, error) {
	stackLength := calculator.stack.Len()
	// Need at least 2 numbers in the stack to operate on them
	if stackLength < 2 {
		return 0, fmt.Errorf("Cannot perform operation with only %d numbers", stackLength)
	}
	num1, err := calculator.stack.GetNumber()
	if err != nil {
		return 0, err
	}

	num2, err := calculator.stack.GetNumber()
	if err != nil {
		return 0, err
	}

	return common.PerformOperation(num1, num2, operation)
}
