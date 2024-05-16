package main

import (	
	"fmt"	
	"example.com/mygo/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	// creating user instance using struct literal notation
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)	
	if(err != nil) {
		fmt.Println(err)
		return
	}

	adminUser, err := user.NewAdmin("aneesh", "gupta")
	adminUser.OutputUserDetails() //adminUser.User.OutputUserDetails() will also work -> for OutputUserDetails, User struct must be exposed since this method is associated with User
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
