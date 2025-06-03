package utilvalidate

import (
	"SimpleHTMLPage/consts"
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

func CheckValidUsername(username string) error {
	if len(username) < 7 || len(username) > 20 {
		fmt.Println("Username's length is invalid!")
		return consts.ErrUsernameInvalid
	}

	if !checkValidString(username) {
		return consts.ErrUsernameInvalid
	}

	return nil
}

func CheckValidPassword(password string) error {
	if len(password) < 7 || len(password) > 255 {
		fmt.Println("Password's length is invalid!")
		return consts.ErrPasswordInvalid
	}

	if !checkValidString(password) {
		return consts.ErrPasswordInvalid
	}

	return nil
}

func CheckValidEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		fmt.Println("Email is invalid!")
	}

	return err
}
