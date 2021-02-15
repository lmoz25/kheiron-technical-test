package test

import (
	"fmt"
	"testing"

	calculator "gitlab.com/lmoz25/kheiron-technical-test/m/internal/prefix-calculator"
)

func TestPrefixCalculator(t *testing.T) {
	for _, tc := range TestData {
		testName := fmt.Sprintf("Add Question: %s", tc.TestDescription)
		var calc calculator.Calculator
		t.Run(testName, func(t *testing.T) {
			retVal, err := calc.ParseInput(tc.Sum)
			if err != nil {
				t.Error(err)
			}
			if retVal != tc.ExpectedResult {
				t.Errorf("Incorrent result %f from sum %s", retVal, tc.Sum)
			}
		})
	}
}

var TestData = []struct {
	TestDescription string
	Sum             string
	ExpectedResult  float32
}{
	{
		"Add two numbers",
		"+ 3 4",
		7,
	},
	{
		"Subtract two numbers",
		"- 3 4",
		-1,
	},
	{
		"Multiply two numbers",
		"* 3 4",
		12,
	},
	{
		"Divide two numbers",
		"/ 3 4",
		0.75,
	},
}
