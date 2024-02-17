package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// fields can be added according to the json data required and the api specs
type WeatherData struct {
	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Humidity    float64 `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

// user variables
var (
	APIKey      string
	APIEndpoint = "http://api.openweathermap.org/data/2.5/weather"
	interval    = 5 * time.Second
	cities      = []string{"Kolkata", "New York", "London", "Jaipur"}
)

func main() {
	godotenv.Load()
	APIKey = os.Getenv("API_KEY")

	fmt.Println(APIKey)
	dataChannel := make(chan *WeatherData)
	errorChannel := make(chan error)

	go startPoller(dataChannel, errorChannel, interval)

	handleData(dataChannel, errorChannel)
}

func startPoller(dataChannel chan *WeatherData, errorChannel chan error, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, city := range cities {
			go fetchWeatherData(city, dataChannel, errorChannel)
		}
	}
}

func handleData(dataChannel chan *WeatherData, errorChannel chan error) {
	for {
		select {
		case weatherData := <-dataChannel:
			fmt.Printf("Temperature in %s: %.1f°C\n", weatherData.Name, weatherData.Main.Temperature)
			fmt.Printf("Feels Like %.1f°C\n", weatherData.Main.FeelsLike)
			fmt.Printf("Humidity in %s: %.1f%%\n", weatherData.Name, weatherData.Main.Humidity)
		case err := <-errorChannel:
			fmt.Printf("Error fetching weather data: %v\n", err)
		}
	}
}

func fetchWeatherData(city string, dataChannel chan *WeatherData, errorChannel chan error) {
	//different parameters can be passed according to the API specs
	uri := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", APIEndpoint, city, APIKey)
	resp, err := http.Get(uri)
	if err != nil {
		errorChannel <- err
		return
	}
	defer resp.Body.Close()

	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		errorChannel <- err
		return
	}

	dataChannel <- &weatherData
}
