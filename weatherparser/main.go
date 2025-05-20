package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Структура для парсингу JSON-відповіді (тільки потрібні поля)
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	apiKey, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		log.Fatal("Не знайдено змінну середовища WEATHER_API_KEY")
	}

	city := "Kyiv"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Помилка HTTP-запиту:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP-помилка: статус %d", resp.StatusCode)
	}

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		log.Fatal("Помилка декодування JSON:", err)
	}

	fmt.Printf("Температура в %s: %.1f°C\n", city, weather.Main.Temp)
}
