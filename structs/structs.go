package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {

	// appUser.firstName = getUserData("Please enter your first name: ")
	// appUser.lastName = getUserData("Please enter your last name: ")
	// appUser.birthdate = getUserData("Please enter your first birthdate (MM/DD/YYYY): ")
	UserfirstName := getUserData("Please enter your first name: ")
	UserlastName := getUserData("Please enter your last name: ")
	Userbirthdate := getUserData("Please enter your first birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser = newUser(UserfirstName, UserlastName, Userbirthdate)
	// appUser = user{
	// 	firstName: UserfirstName,
	// 	lastName:  UserlastName,
	// 	birthdate: Userbirthdate,
	// 	createdAt: time.Now(),
	// }

	//.. do something awsome with that gathered data!
	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
