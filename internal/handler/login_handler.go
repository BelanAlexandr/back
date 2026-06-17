package handler

import (
	"net/http"
	"strings"

	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	token, role, err := service.LoginService(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"role": role})
}
