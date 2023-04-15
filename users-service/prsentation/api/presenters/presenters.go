package presenters

import (
	"encoding/json"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/viewmodels"
	"net/http"
)

type Response struct {
	Code int
	Body map[string]any
	Err  error
}

func NewErrorResponse(err error, code int) Response {
	r := Response{
		Code: code,
		Body: map[string]any{},
		Err:  err,
	}
	r.Body["error"] = err.Error()
	return r
}

func NewSuccessResponse(data any) Response {
	r := Response{
		Code: http.StatusOK,
		Body: map[string]any{},
	}
	r.Body["data"] = data
	return r
}

// ---------

type BasePresenter[TResult any] struct {
	response Response
}

func (p *BasePresenter[TResult]) HandleError(err error) {
	p.response = NewErrorResponse(err, http.StatusBadRequest)
}
func (p *BasePresenter[TResult]) HandleResult(result TResult) {
	bytes, _ := json.Marshal(result)
	var res map[string]any
	_ = json.Unmarshal(bytes, &res)
	p.response = NewSuccessResponse(res)
}
func (p *BasePresenter[TResult]) Present() Response {
	return p.response
}

// ---------

type SignUpPresenter struct {
	BasePresenter[*entity.User]
}

func NewSingUpPresenter() *SignUpPresenter {
	return &SignUpPresenter{
		BasePresenter: BasePresenter[*entity.User]{},
	}
}
func (s *SignUpPresenter) HandleResult(result *entity.User) {
	s.response = NewSuccessResponse(viewmodels.UserToViewModel(result))
}

// ---------

type LoginPresenter struct {
	BasePresenter[*commands.LoginResult]
}

func NewLoginPresenter() *LoginPresenter {
	return &LoginPresenter{
		BasePresenter: BasePresenter[*commands.LoginResult]{},
	}
}

// ---------

type RefreshAccessTokenPresenter struct {
	BasePresenter[*commands.RefreshAccessTokenResult]
}

func NewRefreshAccessTokenPresenter() *RefreshAccessTokenPresenter {
	return &RefreshAccessTokenPresenter{
		BasePresenter: BasePresenter[*commands.RefreshAccessTokenResult]{},
	}
}

// ---------

type LogoutPresenter struct {
	BasePresenter[bool]
}

func NewLogoutPresenter() *LogoutPresenter {
	return &LogoutPresenter{
		BasePresenter: BasePresenter[bool]{},
	}
}
func (p *LogoutPresenter) HandleResult(res bool) {
	p.response = NewSuccessResponse(map[string]bool{"success": res})
}
