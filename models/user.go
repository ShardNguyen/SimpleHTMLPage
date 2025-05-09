package models

// Perform all CRUD related to User in here?

import (
	// Removed import to avoid import cycle
	"SimpleHTMLPage/consts"
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	"SimpleHTMLPage/requests"
	"SimpleHTMLPage/utilities"
	utilauth "SimpleHTMLPage/utilities/auth"

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

func CreateUser(userReq *requests.UserSignUpRequest) error {
	userOrm := dbpostgres.GetUserOrm()

	// Check if username existed
	userCheck, err := GetUser(userReq.ConvertToUserLoginRequest())
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if userCheck.ID > 0 {
		return consts.ErrUsernameExisted
	}

	salt := utilities.GenerateRandomSalt()
	hashedPassword := utilauth.HashPassword(userReq.RawPassword, salt)

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

func GetUser(userReq *requests.UserLoginRequest) (*User, error) {
	userOrm := dbpostgres.GetUserOrm()
	var user User
	result := userOrm.Where(&User{Username: userReq.Username}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
