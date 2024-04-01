package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Args()
}

func Args() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must enter arguments!")
	} else {
		mealTotal, _ := strconv.ParseFloat(args[0], 32)
		tipAmount, _ := strconv.ParseFloat(args[1], 32)
		fmt.Printf("Your meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
	}
}

func calculateTotal(mealTotal float32, tipAmount float32) float32 {
	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))
	return totalPrice
}

func NameReader() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %v", text)
}
