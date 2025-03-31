package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// (u *user) is a reciever to the function outputUserDetails. This binds this function to the user struct
// making it a method of this struct.
// we can also pass (u user) to this method. Doing so will create a copy of our original struct which is
// okay if we need to print data but when we need to change the data of the struct, a pointer is required
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
}

// utility function to create user
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthdate arer required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "firstName",
			lastName:  "lastName",
			birthdate: "birthdate",
			createdAt: time.Now(),
		},
	}
}
