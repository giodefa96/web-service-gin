package middleware

import (
	"net/http"
	"strings"

	"example/web-service-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware di autenticazione JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ottieni il token dall'header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mancante"})
			c.Abort()
			return
		}

		// Verifica che il token sia nel formato "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato token non valido"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Verifica il token
		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token non valido", "details": err.Error()})
			c.Abort()
			return
		}

		// Estrai i claim dal token (opzionale)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"]) // Salva il nome utente nel contesto di Gin
		}

		c.Next()
	}
}
