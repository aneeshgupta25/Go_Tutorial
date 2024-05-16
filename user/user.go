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