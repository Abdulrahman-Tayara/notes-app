package repositories

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/types"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	"reflect"
	"testing"
	"time"
)

func TestRefreshTokenRepository_Save(t *testing.T) {
	token := types.NewRefreshToken("some-string", "user-id", time.Now().Add(time.Hour))

	repo := NewRefreshTokenRepository(db.Instance())

	err := repo.Save(token)

	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Delete(token)

	retrievedToken, err := repo.GetByToken("some-string")

	if err != nil {
		t.Error(err)
		return
	}

	retrievedToken.ExpiresIn = token.ExpiresIn

	if !reflect.DeepEqual(*token, *retrievedToken) {
		t.Errorf("expected: %v, actual: %v", *token, *retrievedToken)
	}
}
