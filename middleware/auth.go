package middleware

import (
	"net/http"
	"strings"

	"github.com/CkCreative/rest/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtVerify function
func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {

		var header = c.Request.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			c.JSON(403, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Writer.Header().Add("user", header)
		c.Next()
	}
}
