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

	userCheck, err := GetUser(userReq.Username)

	// Check if there's any error
	// Void the error if it is record not found
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Check if user existed
	if userCheck != nil {
		return consts.ErrUsernameExisted
	}

	salt := utilities.GenerateRandomSalt()
	hashedPassword := utilpass.HashPassword(userReq.RawPassword, salt)

	user := &User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: hashedPassword,
		Salt:     salt,
	}

	result := userOrm.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUser(username string) (*User, error) {
	userOrm := dbpostgres.GetUserOrm()
	var user User
	result := userOrm.Where(&User{Username: username}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
