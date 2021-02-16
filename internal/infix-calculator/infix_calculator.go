package calculator

import (
	"strings"

	"gitlab.com/lmoz25/kheiron-technical-test/internal/common"
)

type InfixCalculator struct {
	numberStack    common.Stack
	operationStack common.Stack
}

// TODO: consider invalid input?
func (calculator *InfixCalculator) ParseInput(input string) error {
	sum := strings.Fields(input)
	for i := len(sum) - 1; i >= 0; i-- {
		section := sum[i]
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
			calculator.numberStack.AddNumberString(section)
		}
	}
	return nil
}

func (calculator *InfixCalculator) operate() (float32, error) {

}
