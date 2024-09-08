package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/USD"

func convertCurrency(from string, to string, amount float64) (float64, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("error decoding response: %v", err)
	}

	rates := result["rates"].(map[string]interface{})
	fromRate, okFrom := rates[from]
	toRate, okTo := rates[to]
	if !okFrom || !okTo {
		return 0, fmt.Errorf("currency not found: %s or %s", from, to)
	}

	baseAmount := amount / fromRate.(float64)

	convertedAmount := baseAmount * toRate.(float64)
	return convertedAmount, nil
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "convert [from_currency] [to_currency] [amount]",
		Short: "Convert currency between two different currencies",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			fromCurrency := args[0]
			toCurrency := args[1]
			amount, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				fmt.Println("Invalid amount:", args[2])
				os.Exit(1)
			}

			converted, err := convertCurrency(fromCurrency, toCurrency, amount)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, converted, toCurrency)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
