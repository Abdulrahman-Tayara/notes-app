package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
}

type AuthorizedContext struct {
	*Context

	UserId string
	Email  string
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

func GinAuthorizedWrapper(handler func(ctx *AuthorizedContext)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c := AuthorizedContext{
			Context: &Context{Context: ctx},
			UserId:  ctx.GetString("user_id"),
			Email:   ctx.GetString("email"),
		}

		handler(&c)
	}
}
