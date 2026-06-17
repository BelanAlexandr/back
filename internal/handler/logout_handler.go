package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	c.SetCookie("tokenn", " ", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
