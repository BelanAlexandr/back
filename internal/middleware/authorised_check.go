package middleware

import (
	"net/http"

	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/BelanAlexandr/back/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthorisedCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusSeeOther, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(cookie)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
			c.Abort()
			return
		}
		idFloat, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный формат ID в токене"})
			c.Abort()
			return
		}
		role, session, err := repository.AuthorisedCheck(int(idFloat))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			c.Abort()
			return
		}
		tokenSessionID, ok := claims["session_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный формат сессии в токене"})
			c.Abort()
			return
		}
		if tokenSessionID != session {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Сессия устарела. Выполнен вход с другого устройства."})
			c.Abort()
			return
		}
		c.Set("userID", int(idFloat))
		c.Set("userRole", role)

		c.Next()
	}
}
