package main

import "fmt"

func main() {
	slc := []string{"apple", "banana", "apple", "orange", "banana", "kiwi"}
	fmt.Println("Original slice:", slc)
	fmt.Println(getUnique(slc))
}

func getUnique(s []string) []string {
	// Create a map to store unique strings
	uniqueMap := make(map[string]bool)
	// Create a slice to store the unique strings
	var uniqueStrings []string

	// Iterate over the input slice
	for _, str := range s {
		// If the string is not in the map, add it to the map and the unique slice
		if !uniqueMap[str] {
			uniqueMap[str] = true
			uniqueStrings = append(uniqueStrings, str)
		}
	}

	return uniqueStrings
}
