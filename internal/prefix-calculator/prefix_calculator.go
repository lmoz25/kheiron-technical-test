package pcalculator

import (
	"fmt"
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/internal/common"
)

type PrefixCalculator struct {
	stack common.Stack
}

func (calculator *PrefixCalculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	for i := len(sum) - 1; i >= 0; i-- {
		section := sum[i]
		if common.IsOperation(section) {
			result, err := calculator.operate(section)
			if err != nil {
				return err
			}
			calculator.stack.AddNumber(result)
		} else {
			calculator.stack.AddNumberString(section)
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
