package routes

import (
	"net/http"

	"github.com/BelanAlexandr/back/internal/handler"
	"github.com/BelanAlexandr/back/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Адрес вашего фронта
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	//Авторизация
	r.POST("/api/login", handler.LoginHandler)
	r.GET("/login", handler.LoginHandlerShow)

	auth := r.Group("/")

	auth.Use(middleware.AuthorisedCheck())
	auth.GET("/api/auth/me", func(c *gin.Context) {
		userRoleValue, existsRole := c.Get("userRole")
		if !existsRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"role": userRoleValue})
	})
	//Добавление пользователя
	auth.POST("/api/users", handler.AddUserHandler)
	auth.GET("/api/users", handler.ShowUserHandler)
	//Главная
	auth.GET("/api/expertiza/list", handler.IndexHandler)
	auth.GET("/api/", handler.IndexHandler)

	//Добавление экспертизы
	auth.GET("/api/regions", handler.GetRegionsHandler)
	auth.POST("/api/expertiza/save", handler.AddExpHandler)
	//Закрытие экспертизы
	auth.GET("/closeexp", handler.CloseExpHandlerShow)
	auth.PUT("/api/expertize/:id/complete", handler.CloseExpHandler)
	//Редактирование
	auth.GET("/api/expertiza/:id", handler.EditExpHandlerShow)
	auth.PUT("/api/expertiza/update/:id", handler.EditExpHandler)
	return r
}
