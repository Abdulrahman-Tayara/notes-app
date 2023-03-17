package initializers

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDB(config *Config) (err error) {
	err = db.ConnectToDB(config.DbDSN)

	if err != nil {
		return
	}

	fmt.Println("Connected successfully to the database")

	return
}
