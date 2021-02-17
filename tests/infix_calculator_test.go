package test

import (
	"errors"
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
			if tc.ExpectedError == nil && err != nil {
				fmt.Println("Unexpected error:")
				t.Error(err)
			} else if err != nil {
				// The failing test cases have "passed"
				return
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
	ExpectedError   error
}{
	{
		"Add two numbers",
		"( 1 + 2 )",
		3,
		nil,
	},
	{
		"Subtract two numbers",
		"( 1 - 2 )",
		-1,
		nil,
	},
	{
		"Multiply two numbers",
		"( 3 *  4 )",
		12,
		nil,
	},
	{
		"Divide two numbers",
		"( 3 / 4 )",
		0.75,
		nil,
	},
	{
		"Combine two operations",
		"( 1 + ( 2 * 3 ) )",
		7,
		nil,
	},
	{
		"Combine to operations again",
		"( ( 1 * 2 ) + 3 )",
		5,
		nil,
	},
	{
		"Combine more operations",
		"( ( ( 1 + 1 ) / 10 ) - ( 1 * 2 ) )",
		-1.8,
		nil,
	},
	{
		"Invalid sum",
		"+ 4 3",
		0,
		errors.New("Sum not in infix notation"),
	},
}
