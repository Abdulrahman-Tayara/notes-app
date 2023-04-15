package types

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"time"
)

type UserClaimsPayload struct {
	UserId string
	Email  string
}

func (c UserClaimsPayload) AsPayload() services.Payload {
	return services.Payload{
		"user_id": c.UserId,
		"email":   c.Email,
	}
}

type RefreshToken struct {
	Token     string  `gorm:"primaryKey;autoIncrement:false"`
	UserId    core.ID `gorm:"primaryKey;autoIncrement:false"`
	ExpiresIn time.Time
}

func NewRefreshToken(token string, userId core.ID, expiresIn time.Time) *RefreshToken {
	return &RefreshToken{
		Token:     token,
		UserId:    userId,
		ExpiresIn: expiresIn,
	}
}

type AuthOptions struct {
	AccessTokenAge  time.Duration
	RefreshTokenAge time.Duration
}
