package context

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
)

type AppContext struct {
	context.Context

	UserId core.ID
	Email  string
}
