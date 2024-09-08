package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "calculator",
		Short: "A simple CLI calculator",
		Long:  `This is a command-line calculator written in Go using the Cobra library.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				fmt.Println("Please provide a valid expression (e.g., 2 * 5)")
				return
			}
			num1, err1 := strconv.ParseFloat(args[0], 64)
			operator := args[1]
			num2, err2 := strconv.ParseFloat(args[2], 64)

			if err1 != nil || err2 != nil {
				fmt.Println("Please provide valid numbers")
				return
			}

			switch operator {
			case "+":
				fmt.Printf("Result: %f\n", num1+num2)
			case "-":
				fmt.Printf("Result: %f\n", num1-num2)
			case "*":
				fmt.Printf("Result: %f\n", num1*num2)
			case "/":
				if num2 == 0 {
					fmt.Println("Cannot divide by zero")
				} else {
					fmt.Printf("Result: %f\n", num1/num2)
				}
			case "^":
				fmt.Printf("Result: %f\n", math.Pow(num1, num2))
			default:
				fmt.Println("Unsupported operator! Use one of +, -, *, /, ^")
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
