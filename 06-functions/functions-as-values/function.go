package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
	doubledArray := transformNumbers(&numbers, double)
	fmt.Println(doubledArray)
	doubledArray = transformNumbers(&numbers, triple)
	fmt.Println(doubledArray)
	doubledArray = transformNumbers(&numbers, getTransformerFunc(&numbers))
	fmt.Println(doubledArray)
	doubledArray = transformNumbers(&numbers, getTransformerFunc(&doubledArray))
	fmt.Println(doubledArray)

}

func doubleNumbersFor(numbers []int) []int {
	numbersDoubled := []int{}
	for i := 0; i < len(numbers); i++ {
		numbersDoubled = append(numbersDoubled, numbers[i]*2)
	}
	return numbersDoubled
}

func doubleNumbers(numbers *[]int) []int {
	numbersDoubled := []int{}
	for _, val := range *numbers {
		numbersDoubled = append(numbersDoubled, val*2)
	}
	return numbersDoubled
}
func transformNumbers(numbers *[]int, transform transformFn) []int {
	numbersDoubled := []int{}
	for _, val := range *numbers {
		numbersDoubled = append(numbersDoubled, transform(val))
	}
	return numbersDoubled
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}

func getTransformerFunc(numbers *[]int) transformFn {
	if (*numbers)[0] != 1 {
		return triple
	}

	return double
}
