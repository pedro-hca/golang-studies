package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(&age)
	// fmt.Println((adultYears))
	fmt.Println((age))
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
