package dbpostgres

import (
	"SimpleHTMLPage/config"
	"SimpleHTMLPage/consts"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var userOrm *gorm.DB

// Only connect once to keep the connection being the same
func UserConnect() (err error) {

	once.Do(func() {
		dsn := config.GetConfig().GetPostgresqlDSN()

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

func CloseUserConnection() error {
	if userOrm != nil {
		sqlDB, err := userOrm.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return consts.ErrOrmNotExist
}
