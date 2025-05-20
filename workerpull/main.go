package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Path string
}

type Result struct {
	Job       Job
	LineCount int
	WorkerID  int
	Processed time.Duration
	Err       error
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		start := time.Now()
		lineCount, err := countLines(job.Path)

		result := Result{
			Job:       job,
			LineCount: lineCount,
			WorkerID:  id,
			Processed: time.Since(start),
			Err:       err,
		}
		results <- result
	}
}

// Простий підрахунок рядків у файлі
func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func main() {
	const numWorkers = 4

	// Знаходимо всі .txt файли в поточній директорії
	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println("Помилка пошуку файлів:", err)
		return
	}

	jobs := make(chan Job, len(files))
	results := make(chan Result, len(files))

	var wg sync.WaitGroup

	// Запускаємо воркерів
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Додаємо задачі у чергу
	for i, file := range files {
		jobs <- Job{ID: i + 1, Path: file}
	}
	close(jobs)

	wg.Wait()
	close(results)

	// Виводимо результати
	for result := range results {
		if result.Err != nil {
			fmt.Printf("[Worker %d] ❌ Job %d (%s) error: %v\n",
				result.WorkerID, result.Job.ID, result.Job.Path, result.Err)
		} else {
			fmt.Printf("[Worker %d] ✅ Job %d (%s): %d lines (%v)\n",
				result.WorkerID, result.Job.ID, result.Job.Path, result.LineCount, result.Processed)
		}
	}
}
