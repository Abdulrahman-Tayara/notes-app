package initializers

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
)

func ConnectToDB(config *configs.Config) (err error) {
	err = db.ConnectToDB(config.DbDSN)

	if err != nil {
		return
	}

	fmt.Println("Connected successfully to the database")

	return
}
