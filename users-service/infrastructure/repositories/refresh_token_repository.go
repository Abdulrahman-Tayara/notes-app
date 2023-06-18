package repositories

import (
	"errors"
	errors2 "github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) interfaces.IRefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r RefreshTokenRepository) Save(token *auth.RefreshToken) (err error) {
	err = r.db.Save(token).Error

	return
}

func (r RefreshTokenRepository) GetByToken(token string) (*auth.RefreshToken, error) {
	where := auth.RefreshToken{Token: token}

	var model auth.RefreshToken

	res := r.db.Where(where).First(&model)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors2.ErrEntityNotFound
		}
		return nil, res.Error
	}

	return &model, nil
}

func (r RefreshTokenRepository) Delete(token *auth.RefreshToken) (err error) {
	err = r.db.Delete(token).Error
	return
}
