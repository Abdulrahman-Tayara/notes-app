package domain

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"time"
)

type Note struct {
	Id        core.ID   `gorm:"type:uuid;primary_key;" json:"id"`
	UserId    core.ID   `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNote(userId core.ID, title string, content string) (*Note, error) {
	if userId == "" {
		return nil, errors.New("invalid user id")
	}

	return &Note{
		UserId:    userId,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (n Note) GetId() core.ID {
	return n.Id
}

func (n *Note) Update(title, content string) error {
	n.Title = title
	n.Content = content
	return nil
}
