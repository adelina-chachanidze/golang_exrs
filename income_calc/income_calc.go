package main

import "fmt"

func main() {
	revenue := getInput("Revenue:")
	expenses := getInput("Expences:")
	taxRate := getInput("Tax rate:")

	beforeTax, afterTax, ratio := calculateFinance(revenue, expenses, taxRate)

	fmt.Printf("\nYour EBT: %.0f\n", beforeTax)
	fmt.Printf("Your income: %.0f\n", afterTax)
	fmt.Printf("Your ratio: %.1f\n", ratio)
}

func calculateFinance (revenue, expenses, taxRate float64) (float64, float64, float64) {
	beforeTax := revenue - expenses
	afterTax := beforeTax - (beforeTax * (taxRate / 100))
	ratio := beforeTax / afterTax
	return beforeTax, afterTax, ratio
}

func getInput (infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}