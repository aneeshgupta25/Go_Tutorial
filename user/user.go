package user

import (
	"fmt"
	"errors"		
	"time"
)

type User struct {
	firstName, lastName, birthdate string
	createdAt time.Time	
}

type Admin struct {
	email, password string
	// user User // user won't allow to access this property in other package
	User // User allows to access this property in other package
}

func NewAdmin(email, password string) (*Admin, error) {
	if(email == "" || password == "") {
		return nil, errors.New("Can't be empty")
	}
	return &Admin{
		email: email,
		password: password,
		User: User {
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "25/01/2002",
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if(firstName == "" || lastName == "" || birthDate == "") {
		return nil, errors.New("Can't be empty!")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthDate,
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}