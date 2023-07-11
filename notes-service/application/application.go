package application

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateNote func() *commands.CreateNoteHandler
	UpdateNote func() *commands.UpdateNoteHandler
	DeleteNote func() *commands.DeleteNoteHandler
}

type Queries struct {
	GetNoteById  func() *queries.GetNoteHandler
	GetUserNotes func() *queries.GetUserNotesHandler
}
