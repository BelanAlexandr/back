package handler

import (
	"net/http"
	"text/template"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func AddExpHandlerShow(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleEmployee {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, nil)
}
func AddExpHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleEmployee {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	err := service.AddExpService(userId)
	c.JSON(http.StatusOK, err)
}
