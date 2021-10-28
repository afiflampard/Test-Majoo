package middleware

import (
	"fmt"
	"majoo/controllers"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Split(authHeader, " ")
		token, err := controllers.JWTAuthService().ValidateToken(tokenString[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if !ok && !token.Valid {
			fmt.Println("Masuk sini")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user_id", claims["id"])
	}
}
