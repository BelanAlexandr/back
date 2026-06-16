package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func EditExpHandlerShow(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleDirector {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	row, err := repository.GetJournalRow(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, row)
}
func EditExpHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleDirector {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req models.Exp
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.EditExpService(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, err)
}
