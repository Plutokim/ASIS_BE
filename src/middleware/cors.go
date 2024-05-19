package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders: []string{
			headers.ContentType,
			headers.ContentLength,
			headers.AcceptEncoding,
			headers.XCSRFToken,
			headers.Authorization,
			headers.Accept,
			headers.Origin,
			headers.CacheControl,
			headers.XRequestedWith,
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
	})
}