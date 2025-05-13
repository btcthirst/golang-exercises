package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := getUserInput()

	multiplicationTable(num)
}

func multiplicationTable(num int) {
	for i := 1; i <= 20; i++ {
		result := num * i
		fmt.Printf("%d x %d = %d\n", num, i, result)
	}
	fmt.Println("Multiplication table completed.")
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
