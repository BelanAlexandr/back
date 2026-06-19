package handler

import (
	"net/http"

	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func GetRegionsHandler(c *gin.Context) {
	_, existsRole := c.Get("userRole")
	if !existsRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	regions, err := service.GetRegionsService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	vid, err := service.GetVidService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"regions": regions,
		"vids":    vid,
	})
}
