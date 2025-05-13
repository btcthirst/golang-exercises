package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	userInput()
}

func userInput() {
	for {
		println("Enter a message to log (or 'exit' to quit):")
		bufioReader := bufio.NewReader(os.Stdin)
		input, err := bufioReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		// Remove the newline character from the input
		input = input[:len(input)-1]
		// Check if the user wants to exit
		if input == "exit" {
			logMessage("exit \n")
			break
		}
		logMessage(input)
	}
}

func logMessage(message string) {
	// Open the log file in append mode
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Write the message to the log file
	timestmp := time.Now().Format("2006-01-02 15:04:05")
	message = fmt.Sprintf("%s: %s", timestmp, message)
	if _, err := file.WriteString(message + "\n"); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}
