package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	userYear := getUserInput()

	calculator(userYear)

	fmt.Println("\nThank you for using the age calculator!")
	fmt.Println("Have a great day!")
	fmt.Println("If you want to calculate again, please run the program again.")
	fmt.Println("Goodbye!")
}

func calculator(year int) {
	currentYear := time.Now().Year()
	age := currentYear - year
	fmt.Println()
	if age == 0 {
		fmt.Println("You are not born yet!")
		return
	}
	fmt.Printf("You are %d years old.\n", age)
}

func getUserInput() int {
	for {
		// Get user input
		var input string
		println("Please enter the year of your birth:")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		year := stringToInt(input)

		if !checkIsValidYear(year) {
			fmt.Println("Please enter a valid year between 1900 and 2025.")
			continue
		}
		return year
	}
}

func stringToInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return num
}

func checkIsValidYear(year int) bool {
	if year < 1900 || year > 2025 {
		return false
	}
	return true
}
