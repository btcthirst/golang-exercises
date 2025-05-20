package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/unidoc/unioffice/document"
	"github.com/xuri/excelize/v2"
)

type DataRow struct {
	Source string
	Text   string
}

func readExcelFile(path string, out chan<- []DataRow, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println("Помилка Excel:", err)
		return
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		fmt.Println("Помилка зчитування рядків:", err)
		return
	}

	var result []DataRow
	for _, row := range rows {
		if len(row) > 0 && strings.TrimSpace(row[0]) != "" {
			result = append(result, DataRow{Source: path, Text: row[0]})
		}
	}
	out <- result
}

func readWordFile(path string, out chan<- []DataRow, wg *sync.WaitGroup) {
	defer wg.Done()
	doc, err := document.Open(path)
	if err != nil {
		fmt.Println("Помилка Word:", err)
		return
	}

	var result []DataRow
	for _, para := range doc.Paragraphs() {
		var textParts []string
		for _, run := range para.Runs() {
			textParts = append(textParts, run.Text())
		}
		text := strings.Join(textParts, "")
		if strings.TrimSpace(text) != "" {
			result = append(result, DataRow{Source: path, Text: text})
		}
	}
	out <- result
}

func writeResultToExcel(data []DataRow, path string) error {
	f := excelize.NewFile()
	sheet := "Result"
	f.NewSheet(sheet)

	f.SetCellValue(sheet, "A1", "Source")
	f.SetCellValue(sheet, "B1", "Text")

	for i, row := range data {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), row.Source)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), row.Text)
	}
	return f.SaveAs(path)
}

func main() {
	files := []string{"data1.xlsx", "data2.docx", "data3.xlsx"}
	var wg sync.WaitGroup
	out := make(chan []DataRow, len(files))

	for _, file := range files {
		wg.Add(1)
		if strings.HasSuffix(file, ".xlsx") {
			go readExcelFile(file, out, &wg)
		} else if strings.HasSuffix(file, ".docx") {
			go readWordFile(file, out, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var allData []DataRow
	for part := range out {
		allData = append(allData, part...)
	}

	// Фільтрація: тільки ті, хто містять "Go"
	filtered := []DataRow{}
	for _, row := range allData {
		if strings.Contains(strings.ToLower(row.Text), "go") {
			filtered = append(filtered, row)
		}
	}

	// Сортування за текстом
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Text < filtered[j].Text
	})

	err := writeResultToExcel(filtered, "result.xlsx")
	if err != nil {
		fmt.Println("Помилка запису:", err)
	} else {
		fmt.Println("✅ Результат збережено у result.xlsx")
	}
}
