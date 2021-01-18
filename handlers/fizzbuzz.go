package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/schema"
	"go.uber.org/zap"

	"github.com/aftouh/fizzbuzz/core"
	"github.com/aftouh/fizzbuzz/models"
)

var decoder = schema.NewDecoder()

func Fizzbuzz(w http.ResponseWriter, r *http.Request) {
	// Parse and check query parameters
	var qr models.FizzbuzzRequest
	if err := decoder.Decode(&qr, r.URL.Query()); err != nil {
		_ = render.Render(w, r, models.BadRequestErrRendrer("Failed to parse query parameters", err))
		zap.L().Error(fmt.Sprint("Failed to parse query parameters ", err), zap.Any("request", qr))
		return
	}

	res, err := core.Fizzbuzz(qr.Int1, qr.Int2, qr.Limit, qr.Str1, qr.Str2)
	if err != nil {
		_ = render.Render(w, r, models.BadRequestErrRendrer("Failed to calculate fizzbuzz result", err))
		zap.L().Error(fmt.Sprint("Failed to calculate fizzbuzz result ", err), zap.Any("request", qr))
		return
	}

	_ = render.Render(w, r, models.FizzbuzzResponseRendrer(res))
}
