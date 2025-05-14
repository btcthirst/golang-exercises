package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Todo represents a single todo item
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// TodoList represents a list of todo items
type TodoList struct {
	Todos []Todo `json:"todos"`
}

func newTodoList() *TodoList {
	return &TodoList{
		Todos: []Todo{},
	}
}
func (t *TodoList) addTodo(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

var todoList *TodoList

func main() {
	todoList = newTodoList()
	initDB()

	todoCli()
	fmt.Println("Exiting Todo CLI")
}

func initDB() {
	// Initialize the todo list if it doesn't exist
	if _, err := os.Stat("todo_list.json"); os.IsNotExist(err) {
		file, err := os.Create("todo_list.json")
		if err != nil {
			fmt.Println("Error creating todo list file:", err)
			return
		}
		jsonData, err := json.MarshalIndent(todoList, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling todo list:", err)
			return
		}
		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing todo list to file:", err)
			return
		}
		defer file.Close()
	}
}

func todoCli() {
	var err error
	todoList, err = readTodoList()
	if err != nil {
		fmt.Println("Error reading todo list:", err)
		return
	}

	for {
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			todo := getTodo()
			if todo != nil {
				todoList.Todos = append(todoList.Todos, *todo)
				fmt.Printf("Added Todo: %s\n", todo.Description)
				if err := writeTodoList(); err != nil {
					fmt.Println("Error writing todo list:", err)
				} else {
					fmt.Println("Todo list saved successfully.")
				}
			}
		case 2:
			fmt.Println("Todo List:", len(todoList.Todos))
			for _, todo := range todoList.Todos {
				fmt.Printf("%d: %s (Done: %t)\n", todo.ID, todo.Description, todo.Done)
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func getTodo() *Todo {
	for {
		var todo Todo
		if desc := getUserInput("input todo please"); desc != "" {
			todo.Description = desc
			todo.ID = len(todoList.Todos) + 1
			todo.Done = false
			return &todo
		} else {
			fmt.Println("No todo input, try again...")
		}
	}
}

func getUserInput(mes string) string {
	fmt.Println(mes)
	reader := bufio.NewReader(os.Stdin)
	// Read user input
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return input
}

func readTodoList() (*TodoList, error) {
	file, err := os.Open("todo_list.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func writeTodoList() error {
	file, err := os.OpenFile("todo_list.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(todoList)
	if err != nil {
		return err
	}

	return nil
}
