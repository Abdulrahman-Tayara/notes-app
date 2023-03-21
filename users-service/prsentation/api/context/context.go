package context

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
}

func (c *Context) Response(r presenters.Response) {
	if r.Err != nil {
		_ = c.Error(r.Err)
	}
	c.JSON(r.Code, r.Body)
}

func (c *Context) BindJsonOrReturnError(obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Response(presenters.NewErrorResponse(err, http.StatusBadRequest))
		return false
	}

	return true
}

func GinWrapper(handler func(ctx *Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c := Context{Context: ctx}
		handler(&c)
	}
}
