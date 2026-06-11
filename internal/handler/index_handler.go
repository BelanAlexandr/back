package handler

import (
	"net/http"
	"text/template"

	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func IndexHandlerShow(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	userIdValue, existsID := c.Get("userID")
	if !existsRole || !existsID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userId, ok1 := userIdValue.(int)
	userRole, ok2 := userRoleValue.(int)
	if !ok1 || !ok2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID или роли пользователя"})
		return
	}
	tmpl, err := template.ParseFiles("internal/templates/add_user.html")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tables := service.IndexGetService(userId, userRole)
	c.JSON(http.StatusOK, tables)
	tmpl.Execute(c.Writer, nil)
}
