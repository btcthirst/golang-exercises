package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseString("Hello, World!"), " runes")
	fmt.Println(reverseString2("Hello, World!"), " bytes")
	fmt.Println(reverseString3("Hello, World!"), " string.slice")
}

func reverseString(s string) string {
	// Convert the string to a slice of runes
	runes := []rune(s)
	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert the slice of runes back to a string
	return string(runes)
}

func reverseString2(s string) string {
	// Convert the string to a slice of bytes
	bytes := []byte(s)
	// Reverse the slice of bytes
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	// Convert the slice of bytes back to a string
	return string(bytes)
}
func reverseString3(s string) string {
	strs := strings.Split(s, "")
	// Reverse the slice of strings
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	// Convert the slice of strings back to a string
	return strings.Join(strs, "")
}
