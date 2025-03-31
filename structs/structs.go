package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// creating user using utility function
	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return

	}

	// creating Admin using utility function
	appAdmin := user.NewAdmin("", "")
	appAdmin.OutputUserDetails()

	appUser.OutputUserDetails()
	// or if you want to call the outputUserDetails function that hasn't been bound to struct
	// outputUserDetails(&appUser)
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	// creating a variable using custom type str
	var name str = "Farmaan"
	// since variable "name" is of custom type str, we can use the function log on it as it has been bound to this
	// struct
	name.log()
}

// outputUserDetails function declaration without binding to struct

// func outputUserDetails(user *user) {
// we dont use *u.lastName here because Go allows us to use this shortcut for structs
// the correct syntax without the shortcut is (*u).lastName
// 	fmt.Println(user.firstName)
// 	fmt.Println(user.lastName)
// 	fmt.Println(user.birthdate)
// 	fmt.Println(user.createdAt)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// declaring custom types on pre defined types can be used to add methods to that type
type str string

// since str is a custom type build over string, we can bind methods to it which we can then use
// we using variables of this type
func (s str) log() {
	fmt.Println(s)
}
