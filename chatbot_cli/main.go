package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	chatbot()
	fmt.Println("Thank you for using the Chatbot CLI. Goodbye!")

}

func chatbot() {
	fmt.Println("Welcome to the Chatbot CLI!")
	fmt.Println("Type 'exit' to quit the program.")

	for {
		input := getUserInput()
		if input == "exit" {
			break
		}
		response := processInput(input)
		fmt.Println("Chatbot:", response)
	}
}

func processInput(input string) string {

	if len(input) == 1 {
		return simpleAnswer(input)
	}

	switch input {
	case "hello":
		return "Hello! How can I assist you today?"
	case "how are you?":
		return "I'm just a program, but thanks for asking! How can I help you?"
	case "what's your name?":
		return "I'm a simple chatbot created to assist you. You can call me Chatbot!"
	case "tell me a joke":
		return "Why did the scarecrow win an award? Because he was outstanding in his field!"
	case "what's the weather like?":
		return "I don't have real-time data, but you can check a weather website for the latest updates."
	case "what's your favorite color?":
		return "I don't have preferences, but I think blue is a nice color!"
	case "what's your favorite food?":
		return "I don't eat, but I hear pizza is quite popular!"
	case "tell me a riddle":
		return "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I? (Answer: An echo)"
	case "what's the meaning of life?":
		return "The meaning of life is a philosophical question that has been debated for centuries. Some say it's 42, others say it's about finding happiness and purpose."
	case "who created you?":
		return "I was created by a programmer who wanted to build a simple chatbot for fun!"
	default:
		return "I'm sorry, I don't understand that. Can you ask something else?"
	}
}

func simpleAnswer(input string) string {
	switch input {
	case "", " ":
		return "Why so quiet? Speak up!"
	case "0":
		return "1"
	case "1":
		return "2"
	case "2":
		return "3"
	case "3":
		return "4"
	case "4":
		return "5"
	case "5":
		return "6"
	case "6":
		return "7"
	case "7":
		return "8"
	case "8":
		return "9"
	case "9":
		return "10"
	case "a":
		return "b"
	case "b":
		return "c"
	case "c":
		return "d"
	default:
		return "I don't have a specific response for that single character."
	}
}

func getUserInput() string {
	for {
		fmt.Print("Enter your message: ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = input[:len(input)-1] // Remove the newline characterq
		return input
	}
}
