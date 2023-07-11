package injection

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/queries"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure/db"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure/repositories"
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
			DeleteNote: func() *commands.DeleteNoteHandler {
				return initDeleteNoteHandler()
			},
		},
		Queries: application.Queries{
			GetUserNotes: func() *queries.GetUserNotesHandler {
				return initGetUserNotesHandler()
			},
			GetNoteById: func() *queries.GetNoteHandler {
				return initGetNoteHandler()
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

func initDeleteNoteHandler() *commands.DeleteNoteHandler {
	return commands.NewDeleteNoteHandler(initNoteReadRepo(), initNoteWriteRepo())
}

func initGetNoteHandler() *queries.GetNoteHandler {
	return queries.NewGetNoteHandler(initNoteReadRepo())
}

func initGetUserNotesHandler() *queries.GetUserNotesHandler {
	return queries.NewGetUserNotesHandler(initNoteReadRepo())
}

func initUnitOfWork() interfaces.IUnitOfWork {
	return infrastructure.NewAppUnitOfWork(db.Instance())
}

func initNoteReadRepo() interfaces.INoteReadRepository {
	return repositories.NewNoteRepository(db.Instance())
}

func initNoteWriteRepo() interfaces.INoteWriteRepository {
	return repositories.NewNoteRepository(db.Instance())
}
