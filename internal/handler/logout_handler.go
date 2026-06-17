package handler

import (
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", " ", -1, "/", "", false, true)
}
