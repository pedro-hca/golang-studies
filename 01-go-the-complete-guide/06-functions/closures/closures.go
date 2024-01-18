package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubleFunc := createTransformer(2)
	tripleFunc := createTransformer(3)
	doubledArray := transformNumbers(&numbers, doubleFunc)
	fmt.Println(doubledArray)
	doubledArray = transformNumbers(&numbers, tripleFunc)
	fmt.Println(doubledArray)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	numbersDoubled := []int{}
	for _, val := range *numbers {
		numbersDoubled = append(numbersDoubled, transform(val))
	}
	return numbersDoubled
}

// factory function and closure
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
