package handler

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func AddUserHandlerShow(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	tmpl, err := template.ParseFiles("internal/templates/add_user.html")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, nil)
}
func AddUserHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	err := service.AddUserService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
