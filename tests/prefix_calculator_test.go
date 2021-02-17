package test

import (
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

var PrefixTestData = []struct {
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
	{
		"Combine two operations",
		"+ 1 * 2 3",
		7,
	},
	{
		"Combine more operations",
		"- / 10 + 1 1 * 1 2",
		3,
	},
	// TODO: failing test cases
}
