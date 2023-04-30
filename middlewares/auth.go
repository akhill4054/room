package middlewares

import (
	"net/http"
	"strings"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/pkg/e"
	"github.com/akhill4054/room-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			code = e.ERROR_AUTH_TOKEN_MISSING
		} else {
			authHeader := strings.Split(authHeader, " ")
			tokenPrefix := authHeader[0]

			if tokenPrefix != "Bearer" {
				code = e.ERROR_AUTH_INVALID_TOKEN
			} else {
				authToken := authHeader[1]

				claims, err := util.ParseToken(authToken)

				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
					default:
						code = e.ERROR_AUTH_INVALID_TOKEN
					}
				} else {
					user, _ := models.GetUserWithClaims(
						claims.UID, claims.Username,
					)
					if user != nil {
						code = e.SUCCESS
						c.Set("user", user)
					} else {
						code = e.ERROR_AUTH_INVALID_TOKEN
					}
				}
			}
		}

		if code != e.SUCCESS {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"code": code,
			})
			c.Abort()
		}
		c.Next()
	}
}
