package utilinpvalid

import (
	"fmt"
	"net/mail"
)

func checkValidString(str string) bool {
	for _, char := range str {
		if char == ' ' {
			fmt.Println("String has space, invalid!")
			return false
		}

		if (char < '0' || char > '9') &&
			(char < 'a' || char > 'z') &&
			(char < 'A' || char > 'Z') {
			fmt.Println("String has strange character, invalid!")
			return false
		}
	}

	return true
}

func CheckValidUsername(username string) bool {
	if len(username) < 7 || len(username) > 20 {
		fmt.Println("Username's length is invalid!")
		return false
	}

	if !checkValidString(username) {
		return false
	}

	return true
}

func CheckValidPassword(password string) bool {
	if len(password) < 7 || len(password) > 255 {
		fmt.Println("Password's length is invalid!")
		return false
	}

	if !checkValidString(password) {
		return false
	}

	return true
}

func CheckValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Email is invalid!")
	}
	return err == nil
}
