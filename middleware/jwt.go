package middleware

import (
	"fmt"
	auth "github.com/avinashb98/myfin/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := authService.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(*auth.Payload)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
