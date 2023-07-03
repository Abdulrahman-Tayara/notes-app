package injection

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure/db"
)

func InitApplication() application.Application {
	return application.Application{
		Commands: application.Commands{
			CreateNote: func() *commands.CreateNoteHandler {
				return initCreateNoteHandler()
			},
			UpdateNote: func() *commands.UpdateNoteHandler {
				return initUpdateNoteHandler()
			},
		},
	}
}

func initCreateNoteHandler() *commands.CreateNoteHandler {
	return commands.NewCreateNoteHandler(initUnitOfWork())
}

func initUpdateNoteHandler() *commands.UpdateNoteHandler {
	return commands.NewUpdateNoteHandler(initUnitOfWork())
}

func initUnitOfWork() interfaces.IUnitOfWork {
	return infrastructure.NewAppUnitOfWork(db.Instance())
}
