package main

func main() {
	slc := make([]int, 100)
	for i := 0; i < 100; i++ {
		slc[i] = i
	}
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 0; w < 10; w++ {
		go worker(w, jobs, results)
	}
	for _, job := range slc {
		jobs <- job
	}
	close(jobs)
	for a := 0; a < len(slc); a++ {
		<-results
	}
	close(results)
	// Print results
	for i := 0; i < len(slc); i++ {
		println(<-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2

		// Print worker ID
		println("Worker", id, "finished processing job", j)
	}
}
