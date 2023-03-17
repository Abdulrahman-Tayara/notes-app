package presenters

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/viewmodels"
	"net/http"
)

type Response struct {
	Code int
	Body map[string]any
}

func NewErrorResponse(err error, code int) Response {
	r := Response{
		Code: code,
		Body: map[string]any{},
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

type SignUpPresenter struct {
	response Response
}

func NewSingUpPresenter() *SignUpPresenter {
	return &SignUpPresenter{}
}
func (s *SignUpPresenter) HandleError(err error) {
	s.response = NewErrorResponse(err, http.StatusBadRequest)
}
func (s *SignUpPresenter) HandleResult(result *entity.User) {
	s.response = NewSuccessResponse(viewmodels.UserToViewModel(result))
}
func (s *SignUpPresenter) Present() Response {
	return s.response
}
