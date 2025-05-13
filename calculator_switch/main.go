package main

import (
	"fmt"
	"strconv"
)

func main() {
	calculator()
	fmt.Println("Thank you for using the calculator!")
	fmt.Println("Goodbye!")
}

func calculator() {
	fmt.Println("Welcome to the calculator!")
	fmt.Println("You can perform basic operations: +, -, *, /")
	fmt.Println("To exit, type 'exit'.")

	for {
		calculation()
		fmt.Println("Do you want to perform another calculation? (yes/no)")
		var userInput string
		fmt.Scanln(&userInput)
		if userInput == "no" || userInput == "exit" {
			break
		}
	}
}

func getUserNumber() string {
	for {
		fmt.Println("Please enter a number:")
		var userInput string
		fmt.Scanln(&userInput)
		if userInput == "" {
			fmt.Println("You must enter a number.")
			continue
		}
		num := strToInt(userInput)
		if num == -1 {
			fmt.Println("Invalid number. Please try again.")
			continue
		}
		return userInput
	}
}

func getUserOperation() string {
	for {
		fmt.Println("Please enter an operation (+, -, *, /):")
		var userInput string
		fmt.Scanln(&userInput)
		if userInput == "" {
			fmt.Println("You must enter an operation.")
			continue
		}
		if userInput != "+" && userInput != "-" && userInput != "*" && userInput != "/" {
			fmt.Println("Invalid operation. Please try again.")
			continue
		}
		return userInput
	}
}

func calculation() {
	num1 := getUserNumber()
	num2 := getUserNumber()
	operation := getUserOperation()

	switch operation {
	case "+":
		fmt.Println("Result:", strToInt(num1)+strToInt(num2))
	case "-":
		fmt.Println("Result:", strToInt(num1)-strToInt(num2))
	case "*":
		fmt.Println("Result:", strToInt(num1)*strToInt(num2))
	case "/":
		if strToInt(num2) == 0 {
			fmt.Println("Cannot divide by zero.")
			return
		}
		fmt.Println("Result:", strToInt(num1)/strToInt(num2))
	default:
		fmt.Println("Invalid operation.")
	}
}

func strToInt(str string) int {
	// Convert string to int
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return -1
	}
	return num
}
