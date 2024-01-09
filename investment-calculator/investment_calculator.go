package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)
	if investmentAmount <= 0 {
		fmt.Println("ERROR")
		fmt.Println("-------")
		panic("Cant continue, sorry")
	}
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	if expectedReturnRate <= 0 {
		fmt.Println("ERROR")
		fmt.Println("-------")
		panic("Cant continue, sorry")
	}
	fmt.Print("Years: ")
	fmt.Scan(&years)
	if years <= 0 {
		fmt.Println("ERROR")
		fmt.Println("-------")
		panic("Cant continue, sorry")
	}

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	writeCalcToFile(futureValue, futureRealValue)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// fmt.Printf("expected return value: %.2f\n real return value: %.2f", futureValue, futureRealValue)
	formattedFV := fmt.Sprintf("expected return value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("real return value: %.2f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

}
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func writeCalcToFile(futureValue, futureRealValue float64) {
	calcText := fmt.Sprint(futureValue, " - ", futureRealValue)
	os.WriteFile("calculate.txt", []byte(calcText), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
