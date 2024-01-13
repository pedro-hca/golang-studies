package main

import "fmt"

func main() {
	prices := []float64{1, 2, 3, 4}
	fmt.Println(prices[:1])
	prices[0] = 5
	fmt.Println(prices)
	prices = append(prices, 5, 6, 7, 8)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{1, 2, 3, 4}
// 	fmt.Println(prices)
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	fmt.Println("featuredPrices", featuredPrices)
// 	featuredPrices[0] = 5
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println("featuredPrices", featuredPrices)
// 	fmt.Println("highlightedPrices", highlightedPrices)
// 	fmt.Println("prices", prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
