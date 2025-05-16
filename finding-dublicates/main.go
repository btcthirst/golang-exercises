package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 4, 6}
	duplicates := findDuplicates(arr)
	fmt.Println("Duplicates:", duplicates)
}

func findDuplicates(arr []int) []int {
	duplicates := make(map[int]int)
	duplicatesList := []int{}

	for _, num := range arr {
		duplicates[num]++
		if duplicates[num] == 2 {
			duplicatesList = append(duplicatesList, num)
		}
	}

	return duplicatesList
}

/////////////// big int array ////////////
/*
const (
	numWorkers = 8 // кількість горутин
	dataSize   = 10_000_000 // розмір масиву
)

func generateLargeData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Intn(100_000) // обмежена кількість унікальних значень -> більше дублікатів
	}
	return data
}

// Горутіна, яка обробляє частину даних і надсилає локальну мапу в канал
func worker(data []int, out chan map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()

	localMap := make(map[int]int)
	for _, val := range data {
		localMap[val]++
	}
	out <- localMap
}

func mergeMaps(maps []map[int]int) map[int]int {
	result := make(map[int]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] += v
		}
	}
	return result
}

func findDuplicatesParallel(data []int) []int {
	chunkSize := len(data) / numWorkers
	out := make(chan map[int]int, numWorkers)
	var wg sync.WaitGroup

	// запуск горутин
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}
		wg.Add(1)
		go worker(data[start:end], out, &wg)
	}

	// збір результатів
	wg.Wait()
	close(out)

	var partialMaps []map[int]int
	for m := range out {
		partialMaps = append(partialMaps, m)
	}

	// злиття локальних мап
	merged := mergeMaps(partialMaps)

	// витяг дублікатів
	var duplicates []int
	for k, v := range merged {
		if v > 1 {
			duplicates = append(duplicates, k)
		}
	}

	return duplicates
}

func main() {
	data := generateLargeData(dataSize)

	start := time.Now()
	duplicates := findDuplicatesParallel(data)
	elapsed := time.Since(start)

	fmt.Printf("Знайдено %d дублікатів за %s\n", len(duplicates), elapsed)
}*/

/////////////// big string array ////////////
