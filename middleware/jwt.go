package middleware

import (
	"fmt"
	auth "github.com/avinashb98/myfin/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BearerSchema = "Bearer "

func AuthorizeJWT(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "authentication header missing",
			})
			c.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if token.Valid {
			claims := token.Claims.(*auth.Payload)
			c.Set("handle", claims.Handle)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
