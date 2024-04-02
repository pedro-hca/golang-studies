package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Scan()
}

func Args() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnetotal <Total Meal Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 20 10>")
	} else {

		if len(args) != 2 {
			fmt.Println("You must enter 2 arguments! ty")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Your meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
		}
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

func Flag() {
	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "IA64":
		fmt.Println("Remember IA64")
	}

	fmt.Println("Process Complete")
}

func Scan() {
	var name string
	fmt.Println("Whats is your name")
	inputs, _ := fmt.Scanf("%q", &name)

	switch inputs {
	case 0:
		fmt.Printf("You must enter a name! \n")
	case 1:
		fmt.Printf("Hello %s! Nice to meet you \n", name)
	}
}
