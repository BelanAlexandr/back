package routes

import (
	"github.com/BelanAlexandr/back/internal/handler"
	"github.com/BelanAlexandr/back/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)
	auth := r.Group("/")
	auth.Use(middleware.AuthorisedCheck())
	return r
}
