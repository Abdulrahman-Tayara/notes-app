package http

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
}

func (c *Context) AppContext() *context.AppContext {
	return &context.AppContext{
		Context: c,
		UserId: func() core.ID {
			id, _ := core.ParseSafely(c.GetString("user_id"))
			return id
		}(),
		Email: c.GetString("email"),
	}
}

func (c *Context) Response(r Response) {
	if r.Err != nil {
		_ = c.Error(r.Err)
	}
	c.JSON(r.Code, r.Body)
}

func (c *Context) BindJsonOrReturnError(obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Response(NewErrorResponse(err, http.StatusBadRequest))
		return false
	}

	return true
}

func GinWrapper(handler func(ctx *Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		context := Context{Context: ctx}
		handler(&context)
	}
}
