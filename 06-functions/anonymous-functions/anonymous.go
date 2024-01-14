package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	//anonymous function
	doubledArray := transformNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(doubledArray)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	numbersDoubled := []int{}
	for _, val := range *numbers {
		numbersDoubled = append(numbersDoubled, transform(val))
	}
	return numbersDoubled
}
