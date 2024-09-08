package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

const apiURL = "http://api.openweathermap.org/data/2.5/weather"
const apiKey = "832f1946e2dffbc7bc8c44ffc3d3de61"

func getWeather(city string) (string, error) {
	cityEncoded := url.QueryEscape(city)
	fullURL := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", apiURL, cityEncoded, apiKey)

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		return "", fmt.Errorf("failed to fetch data, status code: %d, error: %v", resp.StatusCode, result["message"])
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	main, ok := result["main"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected data format")
	}

	temp, ok := main["temp"].(float64)
	if !ok {
		return "", fmt.Errorf("temperature not found")
	}

	return fmt.Sprintf("%.2f°C", temp), nil
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "weather [city]",
		Short: "Get the current weather for a city",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			city := args[0]

			weather, err := getWeather(city)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Printf("The current temperature in %s is %s\n", city, weather)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
