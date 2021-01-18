package models

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type FizzbuzzRequest struct {
	Int1  int    `schema:"int1,required"`
	Int2  int    `schema:"int2,required"`
	Limit int    `schema:"limit,required"`
	Str1  string `schema:"str1,required"`
	Str2  string `schema:"str2,required"`
}

type FizzbuzzResponse struct {
	Result   []string `json:"result"`
	Status   string   `json:"status"`
	HTTPCode int      `json:"-"`
}

func (e *FizzbuzzResponse) Render(w http.ResponseWriter, r *http.Request) error {
	zap.L().Debug("Send response", zap.Any("response", e))
	render.Status(r, e.HTTPCode)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}

func FizzbuzzResponseRendrer(res []string) *FizzbuzzResponse {
	return &FizzbuzzResponse{
		Result:   res,
		Status:   "Success",
		HTTPCode: http.StatusOK,
	}
}
