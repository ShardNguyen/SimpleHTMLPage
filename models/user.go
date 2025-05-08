package models

// Perform all CRUD related to User in here?

import (
	// Removed import to avoid import cycle
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	"SimpleHTMLPage/responses"
	"SimpleHTMLPage/utilities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
}

func CreateOrUpdateUserTable() error {
	userOrm := dbpostgres.GetUserOrm()
	return userOrm.AutoMigrate(&User{})
}

func CreateUser(userRes *responses.UserResponse) error {
	salt := utilities.GenerateRandomSalt()
	hashedPassword := utilities.HashPassword(userRes.Password, salt)

	user := &User{
		Username: userRes.Username,
		Email:    userRes.Email,
		Password: hashedPassword,
		Salt:     salt,
	}

	userOrm := dbpostgres.GetUserOrm()
	result := userOrm.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUser(userRes *responses.UserResponse) (*User, error) {
	userOrm := dbpostgres.GetUserOrm()
	var user User
	result := userOrm.Where(&User{Username: userRes.Username}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
