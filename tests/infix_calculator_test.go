package test

import (
	"fmt"
	"testing"

	calculator "gitlab.com/lmoz25/kheiron-technical-test/internal/infix-calculator"
)

func TestInfixCalculator(t *testing.T) {
	for _, tc := range InfixTestData {
		testName := fmt.Sprintf("Infix Calculator: %s", tc.TestDescription)
		var calc calculator.InfixCalculator
		t.Run(testName, func(t *testing.T) {
			err := calc.ParseInput(tc.Sum)
			if err != nil {
				t.Error(err)
			}
			retVal, err := calc.Result()
			if err != nil {
				t.Error(err)
			}

			if retVal != tc.ExpectedResult {
				t.Errorf("Incorrect result %f from sum %s", retVal, tc.Sum)
			}
		})
	}
}

var InfixTestData = []struct {
	TestDescription string
	Sum             string
	ExpectedResult  float32
}{
	{
		"Add two numbers",
		"( 1 + 2 )",
		3,
	},
	{
		"Subtract two numbers",
		"( 1 - 2 )",
		-1,
	},
	{
		"Multiply two numbers",
		"( 3 *  4 )",
		12,
	},
	{
		"Divide two numbers",
		"( 3 / 4 )",
		0.75,
	},
	{
		"Combine two operations",
		"( 1 + ( 2 * 3 ) )",
		7,
	},
	{
		"Combine to operations again",
		"( ( 1 * 2 ) + 3 )",
		5,
	},
	{
		"Combine more operations",
		"( ( ( 1 + 1 ) / 10 ) - ( 1 * 2 ) )",
		-1.8,
	},
	// TODO: failing test cases
}
