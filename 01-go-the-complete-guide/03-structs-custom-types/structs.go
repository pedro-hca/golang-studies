package main

import (
	"fmt"

	"github.com/pedro-hca/go-studies/03-structs-custom-types/user"
)

func main() {

	// appUser.firstName = getUserData("Please enter your first name: ")
	// appUser.lastName = getUserData("Please enter your last name: ")
	// appUser.birthdate = getUserData("Please enter your first birthdate (MM/DD/YYYY): ")
	UserfirstName := user.GetUserData("Please enter your first name: ")
	UserlastName := user.GetUserData("Please enter your last name: ")
	Userbirthdate := user.GetUserData("Please enter your first birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(UserfirstName, UserlastName, Userbirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@exemple.com", "1234")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}
