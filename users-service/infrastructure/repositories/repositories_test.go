package repositories

import (
	"errors"
	sharederrors "github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"log"
	"testing"
)

func init() {
	log.Println("init")

	config, err := initializers.LoadTestConfig("../../")

	if err != nil {
		log.Fatal(err)
	}

	err = initializers.ConnectToDB(&config)

	if err != nil {
		log.Fatal(err)
	}
}

func TestUserRepository_Save(t *testing.T) {
	repo := NewUserRepository(db.Instance())

	user, _ := entity.NewUser("", "abdulrahman@gmail.com", "a1234456")

	savedUser, err := repo.Save(user)

	if err != nil {
		t.Fatal(err)
	}

	retrievedUser, err := repo.GetById(savedUser.GetId())

	if err != nil {
		t.Fatal(err)
	}

	if savedUser.Id != retrievedUser.Id || savedUser.Email != retrievedUser.Email {
		t.Error("error in user saving")
	}
}

func TestUserRepository_DeleteById(t *testing.T) {
	repo := NewUserRepository(db.Instance())

	user, _ := entity.NewUser("", "fordelete@gmail.com", "a1234456")

	savedUser, err := repo.Save(user)

	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteById(savedUser.GetId())

	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.GetById(savedUser.GetId())

	if !errors.Is(err, sharederrors.ErrEntityNotFound) {
		t.Error("user wasn't deleted")
	}
}

func TestUserRepository_Count(t *testing.T) {
	repo := NewUserRepository(db.Instance())

	email := "fordelete@gmail.com"

	count := repo.Count(interfaces.UsersFilter{Email: email})

	if count > 0 {
		t.Errorf("count should be 0")
		return
	}

	user, _ := entity.NewUser("", email, "a1234456")

	savedUser, err := repo.Save(user)

	if err != nil {
		t.Fatal(err)
	}

	count = repo.Count(interfaces.UsersFilter{Email: email})

	if count == 0 {
		t.Errorf("count should be gte 1")
		return
	}

	_ = repo.DeleteById(savedUser.Id)
}
