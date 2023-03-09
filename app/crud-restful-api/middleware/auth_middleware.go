package middleware

import (
	"net/http"

	"github.com/yogiis/golang-api-automation/app/crud-restful-api/helpers"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/model"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("API-KEY") == "KODE" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		response := model.FormatterResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED"}

		helpers.WriteToResponseBody(writer, response)
	}
}
