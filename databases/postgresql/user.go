package dbpostgres

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var userOrm *gorm.DB

// Only connect once to keep the connection being the same
func UserConnect() (err error) {

	once.Do(func() {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", host, user, password, dbname)

		userOrm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			return
		}
	})

	return err
}

func GetUserOrm() *gorm.DB {
	return userOrm
}
