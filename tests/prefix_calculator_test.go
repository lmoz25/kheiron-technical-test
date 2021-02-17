package test

import (
	"errors"
	"fmt"
	"testing"

	calculator "gitlab.com/lmoz25/kheiron-technical-test/internal/prefix-calculator"
)

func TestPrefixCalculator(t *testing.T) {
	for _, tc := range PrefixTestData {
		testName := fmt.Sprintf("Prefix Calculator: %s", tc.TestDescription)
		var calc calculator.PrefixCalculator
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

var PrefixTestData = []struct {
	TestDescription string
	Sum             string
	ExpectedResult  float32
	ExpectedError   error
}{
	{
		"Add two numbers",
		"+ 3 4",
		7,
		nil,
	},
	{
		"Subtract two numbers",
		"- 3 4",
		-1,
		nil,
	},
	{
		"Multiply two numbers",
		"* 3 4",
		12,
		nil,
	},
	{
		"Divide two numbers",
		"/ 3 4",
		0.75,
		nil,
	},
	{
		"Combine two operations",
		"+ 1 * 2 3",
		7,
		nil,
	},
	{
		"Combine more operations",
		"- / 10 + 1 1 * 1 2",
		3,
		nil,
	},
	{
		"Invalid sum",
		"( 4 + 3 )",
		0,
		errors.New("strconv.ParseInt: parsing \")\": invalid syntax"),
	},
}
