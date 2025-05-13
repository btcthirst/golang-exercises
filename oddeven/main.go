package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := getUserInput()
	// Check if the number is even or odd
	if num == 0 {
		fmt.Println("The number is zero.")
	} else if num%2 == 0 {
		fmt.Printf("The number %d is even.\n", num)
	} else {
		fmt.Printf("The number %d is odd.\n", num)
	}
}

func getUserInput() int {
	for {
		// Get user input
		var input string
		println("Please enter the number:")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		num := stringToInt(input)
		if num == -1 {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		return num
	}
}

func stringToInt(input string) int {
	if input == "" {
		return -1
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}
