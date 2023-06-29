package auth

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"time"
)

type UserClaimsPayload struct {
	UserId string
	Email  string
}

func NewUserClaimsPayload(m map[string]any) (res *UserClaimsPayload) {
	defer func() {
		if err := recover(); err != nil {
			res = nil
		}
	}()

	return &UserClaimsPayload{
		UserId: m["user_id"].(string),
		Email:  m["email"].(string),
	}
}

func (c UserClaimsPayload) AsPayload() map[string]any {
	return map[string]any{
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
