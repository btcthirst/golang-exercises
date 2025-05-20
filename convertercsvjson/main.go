package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Відкриваємо CSV-файл
	csvFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalf("Помилка відкриття файлу: %v", err)
	}
	defer csvFile.Close()

	// Читаємо CSV
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Помилка читання CSV: %v", err)
	}

	if len(records) < 1 {
		log.Fatalf("CSV-файл порожній або некоректний")
	}

	// Перший рядок — заголовки
	headers := records[0]
	var jsonData []map[string]string

	// Обробка кожного рядка
	for _, row := range records[1:] {
		if len(row) != len(headers) {
			log.Printf("Пропускаємо рядок з невідповідною кількістю колонок: %v", row)
			continue
		}
		entry := make(map[string]string)
		for i, value := range row {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}

	// Конвертація в JSON
	jsonOutput, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("Помилка конвертації в JSON: %v", err)
	}

	// Запис у файл
	err = os.WriteFile("output.json", jsonOutput, 0644)
	if err != nil {
		log.Fatalf("Помилка запису JSON у файл: %v", err)
	}

	fmt.Println("Конвертацію завершено успішно. Дані записано у output.json")
}
