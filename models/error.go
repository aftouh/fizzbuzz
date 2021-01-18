package models

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Message  string `json:"message"`
	Status   string `json:"status"`
	HTTPCode int    `json:"-"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPCode)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}

func BadRequestErrRendrer(message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Message:  fmt.Sprintf("%s: %v", message, err),
		Status:   "Bad request",
		HTTPCode: http.StatusBadRequest,
	}
}
