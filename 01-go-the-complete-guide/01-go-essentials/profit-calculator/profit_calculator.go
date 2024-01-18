package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		panic("ERROR")
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		panic("ERROR")
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		panic("ERROR")
	}

	ebt, profit, ration := CalculateFinancials(revenue, expenses, taxRate)
	writeCalcToFile(ebt, profit, ration)

	PrintValues(ebt, profit, ration)

}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ration := ebt / profit
	return ebt, profit, ration
}

func PrintValues(ebt, profit, ration float64) {
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f", ration)
}

func writeCalcToFile(ebt, profit, ration float64) {
	calcText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ration)
	os.WriteFile("calculate.txt", []byte(calcText), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
