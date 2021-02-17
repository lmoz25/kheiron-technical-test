package main

import (
	"bufio"
	"fmt"
	"os"

	calculator "gitlab.com/lmoz25/kheiron-technical-test/internal/infix-calculator"
)

func main() {
	var calc calculator.InfixCalculator
	for {
		// Println function is used to
		// display output in the next line
		fmt.Println("Enter sum in infix notation: ")

		in := bufio.NewReader(os.Stdin)
		sum, err := in.ReadString('\n')

		fmt.Printf("Sum: %s\n", sum)

		err = calc.ParseInput(sum)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		result, err := calc.Result()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Printf("Result: %.1f\n", result)
		fmt.Println("*************************************")
		calc.ClearData()
	}
}
