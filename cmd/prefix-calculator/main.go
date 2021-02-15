package calculator

import (
	"fmt"

	calculator "gitlab.com/lmoz25/kheiron-technical-test/internal/prefix-calculator"
)

func main() {
	var calc calculator.Calculator
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter sum in prefix notation: ")

	// var then variable name then variable type
	var sum string

	// Taking input from user
	fmt.Scanln(&sum)

	err := calc.ParseInput(sum)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	result, err := calc.Result()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(result)
}
