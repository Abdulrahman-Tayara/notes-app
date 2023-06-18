package http

import "net/http"

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
