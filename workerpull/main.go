package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Структура задачі
type Job struct {
	ID       int
	Workload int
}

// Структура результату
type Result struct {
	Job       Job
	Output    int
	WorkerID  int
	Processed time.Duration
}

// Воркер
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		start := time.Now()

		// Симуляція роботи
		time.Sleep(time.Millisecond * time.Duration(job.Workload))

		result := Result{
			Job:       job,
			Output:    job.Workload * 2, // умовна обробка
			WorkerID:  id,
			Processed: time.Since(start),
		}
		results <- result
	}
}

// Головна функція
func main() {
	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Запускаємо воркерів
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Створюємо задачі
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:       j,
			Workload: rand.Intn(100) + 10, // випадкове навантаження
		}
	}
	close(jobs) // Завершення подачі задач

	wg.Wait() // Очікуємо завершення воркерів
	close(results)

	// Виводимо результати
	for result := range results {
		fmt.Printf("Worker %d processed Job %d (load %d) in %v → Output: %d\n",
			result.WorkerID, result.Job.ID, result.Job.Workload, result.Processed, result.Output)
	}
}
