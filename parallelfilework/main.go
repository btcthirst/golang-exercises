package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readFile(filename string, out chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Помилка відкриття %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Помилка зчитування %s: %v\n", filename, err)
		return
	}

	out <- lines
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	var wg sync.WaitGroup
	results := make(chan []string, len(files))

	for _, f := range files {
		wg.Add(1)
		go readFile(f, results, &wg)
	}

	// Закриємо канал, коли всі горутини завершать роботу
	go func() {
		wg.Wait()
		close(results)
	}()

	// Збираємо всі рядки у єдиний масив
	var allLines []string
	for lines := range results {
		allLines = append(allLines, lines...)
	}

	// Вивід результату
	fmt.Println("Усі рядки з усіх файлів:")
	for _, line := range allLines {
		fmt.Println(line)
	}
}
