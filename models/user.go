package models

import (
	"SimpleHTMLPage/consts"
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	"SimpleHTMLPage/requests"
	"SimpleHTMLPage/utilities"
	utilpass "SimpleHTMLPage/utilities/password"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
}

// End check section //

func CreateOrUpdateUserTable() error {
	userOrm := dbpostgres.GetUserOrm()
	return userOrm.AutoMigrate(&User{})
}

func CreateUser(userReq *requests.UserSignUpRequest) error {
	userOrm := dbpostgres.GetUserOrm()

	// Check if username exists
	userCheck, err := GetUser(userReq.Username)

	// Since getting user returns error when record is not found
	// Void the error if the record is not found
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Check if user existed
	if userCheck != nil {
		return consts.ErrUsernameExisted
	}

	// Generate a random salt and a hashed password for user
	salt := utilities.GenerateRandomSalt()
	hashedPassword := utilpass.HashPassword(userReq.RawPassword, salt)

	user := &User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: hashedPassword,
		Salt:     salt,
	}

	// Create a user and store the info into the user's database
	result := userOrm.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUser(username string) (*User, error) {
	userOrm := dbpostgres.GetUserOrm()

	// Find the record based on username
	var user User
	result := userOrm.Where(&User{Username: username}).First(&user)

	// Return error if user is not found
	// Or if there are other errors
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
