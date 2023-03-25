package injection

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/types"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/repositories"
	services2 "github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/services"
	"gorm.io/gorm"
	"time"
)

func InitSignUpCommand() *commands.SingUpHandler {
	return commands.NewSingUpHandler(InitUnitOfWork(), InitHashService())
}

func InitLoginCommand() *commands.LoginHandler {
	return commands.NewLoginHandler(
		types.AuthOptions{
			AccessTokenAge:  time.Minute * time.Duration(configs.AppConfig.JWTAccessTokenAgeMinutes),
			RefreshTokenAge: time.Minute * time.Duration(configs.AppConfig.JWTRefreshTokenAgeMinutes),
		},
		InitUserReadRepository(),
		InitRefreshTokenRepository(),
		InitTokenService(),
		InitHashService(),
	)
}

func InitUserReadRepository() interfaces.IUserReadRepository {
	return repositories.NewUserRepository(InitPostgresDBInstance())
}

func InitRefreshTokenRepository() interfaces.IRefreshTokenRepository {
	return repositories.NewRefreshTokenRepository(InitPostgresDBInstance())
}

func InitTokenService() services.ITokenService {
	return services2.NewJWTService(services2.Config{
		SigningKey: configs.AppConfig.JWTKey,
		Issuer:     configs.AppConfig.JWTIssuer,
	})
}

func InitUnitOfWork() interfaces.IUnitOfWork {
	return infrastructure.NewAppUnitOfWork(InitPostgresDBInstance())
}

func InitHashService() services.IHashService {
	return services2.NewMD5HashService()
}

func InitPostgresDBInstance() *gorm.DB {
	return db.Instance()
}
