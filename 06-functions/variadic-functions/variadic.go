package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3}
	sum := sumup(1, 2, 3)
	fmt.Println(sum)
	anotherSum := sumup(numbers...)
	fmt.Println(anotherSum)

}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
