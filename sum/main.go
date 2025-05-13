package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := getUserInput()

	sum := calculateSum(nums)
	fmt.Printf("The sum of the numbers is: %d\n", sum)
	fmt.Println("Sum calculation completed.")
}

func calculateSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getUserInput() []int {
	for {
		// Get user input
		var input string
		println("Please enter the numbers separated by point:")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		nums := stringToIntSlice(input)
		if len(nums) == 0 {
			fmt.Println("Invalid input. Please enter valid numbers.")
			continue
		}
		return nums
	}
}

func stringToIntSlice(input string) []int {
	if input == "" {
		return []int{}
	}
	var nums []int

	for _, str := range strings.Split(input, ".") {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("%s is not a number\n", str)
		}
		nums = append(nums, num)
	}
	return nums
}
