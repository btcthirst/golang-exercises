package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task

const filename = "tasks.json"

func main() {
	loadTasks()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTODO CLI — доступні команди:")
		fmt.Println("1. list         - Показати всі завдання")
		fmt.Println("2. add [назва]  - Додати нове завдання")
		fmt.Println("3. done [id]    - Позначити завдання як виконане")
		fmt.Println("4. exit         - Вийти")

		fmt.Print("\nВведи команду: ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.SplitN(input, " ", 2)

		switch args[0] {
		case "list":
			printTasks()
		case "add":
			if len(args) < 2 {
				fmt.Println("⚠️  Введи назву завдання: add Назва завдання")
				continue
			}
			addTask(args[1])
		case "done":
			if len(args) < 2 {
				fmt.Println("⚠️  Введи ID завдання: done 2")
				continue
			}
			markDone(args[1])
		case "exit":
			saveTasks()
			fmt.Println("👋 До зустрічі!")
			return
		default:
			fmt.Println("⛔ Невідома команда")
		}
	}
}

func addTask(title string) {
	id := len(tasks) + 1
	task := Task{ID: id, Title: title, Completed: false}
	tasks = append(tasks, task)
	fmt.Println("✅ Завдання додано!")
	saveTasks()
}

func printTasks() {
	if len(tasks) == 0 {
		fmt.Println("📭 Немає завдань")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
	}
}

func markDone(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("⚠️  Невірний ID")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("🎉 Завдання виконане!")
			saveTasks()
			return
		}
	}
	fmt.Println("❓ Завдання з таким ID не знайдено")
}

func saveTasks() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("⛔ Помилка збереження:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("⛔ Помилка кодування JSON:", err)
	}
}

func loadTasks() {
	file, err := os.Open(filename)
	if err != nil {
		// Немає збереженого файлу — починаємо з нуля
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("⚠️  Помилка завантаження задач:", err)
	}
}

// made by chatgpt
