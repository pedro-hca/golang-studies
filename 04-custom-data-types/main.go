package main

import (
	"fmt"
	"pluralsight/organization"
)

func main() {
	// var p organization.Identifiable = organization.Person{}
	p := organization.NewPerson("Mairo", "Vergara")
	err := p.SetTwitterHandler("@mairo_vergara")
	if err != nil {
		fmt.Printf("An error occurred while setting twitter handler: %s\n", err.Error())
	}
	fmt.Println(p.ID())
	fmt.Println(p.GetTwitterHandler())
	fmt.Println(p.TwitterHandler.RedirectUrl())

}
