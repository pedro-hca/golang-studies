package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	// 1)
	hobbiesArr := [3]string{"Gym", "Videogames", "Biking"}
	fmt.Println(hobbiesArr)
	// 2)
	fmt.Println(hobbiesArr[0])
	fmt.Println(hobbiesArr[1:3])
	// 3)
	hobbiesSlice := hobbiesArr[:2]
	fmt.Println(hobbiesSlice, cap(hobbiesSlice))
	// 4)
	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println(hobbiesSlice)
	// 5)
	goaslSlice := []string{"Learn Go", "Understand concurrency"}
	fmt.Println(goaslSlice)
	// 6)
	goaslSlice[1] = "Learn concurrency"
	fmt.Println(goaslSlice)
	goaslSlice = append(goaslSlice, "Learn how to create a good API in Go")
	fmt.Println(goaslSlice)
	// 7)
	product := []Product{
		{
			title: "monitor",
			id:    1,
			price: 1500,
		},
		{title: "keyboard",
			id:    2,
			price: 500,
		},
	}
	product = append(product, Product{
		title: "mouse",
		id:    3,
		price: 200,
	})
	fmt.Println(product)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
