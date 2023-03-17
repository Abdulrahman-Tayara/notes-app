package injection

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	services2 "github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/services"
	"gorm.io/gorm"
)

func InitSignUpCommand() *commands.SingUpHandler {
	return commands.NewSingUpHandler(InitUnitOfWork(), InitHashService())
}

func InitUnitOfWork() interfaces.IUnitOfWork {
	return infrastructure.NewAppUnitOfWork(InitPostgresDBInstance())
}

func InitHashService() services.IHashService {
	return services2.NewMD5HashService()
}

func InitPostgresDBInstance() *gorm.DB {
	return db.DBInstance()
}
