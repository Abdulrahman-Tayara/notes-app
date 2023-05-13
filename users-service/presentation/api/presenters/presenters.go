package presenters

import (
	"encoding/json"
	"github.com/Abdulrahman-Tayara/notes-app/shared/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/presentation/viewmodels"
	nethttp "net/http"
)

type BasePresenter[TResult any] struct {
	response http.Response
}

func (p *BasePresenter[TResult]) HandleError(err error) {
	p.response = http.NewErrorResponse(err, nethttp.StatusBadRequest)
}
func (p *BasePresenter[TResult]) HandleResult(result TResult) {
	bytes, _ := json.Marshal(result)
	var res map[string]any
	_ = json.Unmarshal(bytes, &res)
	p.response = http.NewSuccessResponse(res)
}
func (p *BasePresenter[TResult]) Present() http.Response {
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
	s.response = http.NewSuccessResponse(viewmodels.UserToViewModel(result))
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
	p.response = http.NewSuccessResponse(map[string]bool{"success": res})
}
