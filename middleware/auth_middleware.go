package middleware

import (
	"goelster/helper"
	"goelster/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	//"12345678" == request.Header.Get("X-API-Key") &&
	if request.Header.Get("Authorization") != "" {
		jwt := request.Header.Get("Authorization")

		iss := helper.ValidateJWT(jwt, "secret-key")

		if iss {
			// ok
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	} else {
		if request.URL.Path == "/simta/login" {
			if request.Header.Get("X-Api-Key") == "simta" {
				// ok
				middleware.Handler.ServeHTTP(writer, request)
			} else {
				// error
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}

				helper.WriteToResponseBody(writer, webResponse)
			}
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	}

}
