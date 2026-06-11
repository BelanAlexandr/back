package routes

import (
	"github.com/BelanAlexandr/back/internal/handler"
	"github.com/BelanAlexandr/back/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	//Авторизация
	r.POST("/api/login", handler.LoginHandler)
	r.GET("/login", handler.LoginHandlerShow)
	auth := r.Group("/")
	auth.Use(middleware.AuthorisedCheck())
	//Добавление пользователя
	auth.POST("/api/adduser", handler.AddUserHandler)
	auth.GET("/adduser", handler.AddUserHandlerShow)
	//Главная
	auth.GET("/", handler.IndexHandlerShow)
	return r
}
