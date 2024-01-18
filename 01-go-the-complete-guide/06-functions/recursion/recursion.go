package main

import "fmt"

type transformFn func(int) int

func main() {
	factorialResult := factorial(6)
	fmt.Println(factorialResult)

}
func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}
