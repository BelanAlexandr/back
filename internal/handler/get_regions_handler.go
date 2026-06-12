package handler

import (
	"net/http"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func GetRegionsHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleDirector {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	regions, err := service.GetRegionsService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regions)
}
