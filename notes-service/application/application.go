package application

import "github.com/Abdulrahman-Tayara/notes-app/notes-service/application/commands"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateNote func() *commands.CreateNoteHandler
	UpdateNote func() *commands.UpdateNoteHandler
}
