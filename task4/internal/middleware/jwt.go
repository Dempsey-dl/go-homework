package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuth(jwtscret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerstr := ctx.GetHeader("Authorization")
		if headerstr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Header is empty"})
			ctx.Abort()
			return
		}

		//解析验证
		token, err := jwt.Parse(headerstr, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtscret), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("UserID", claims["userID"])
		ctx.Next()
	}
}
