package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/gin-gonic/gin"
)

func CloseExpHandlerShow(c *gin.Context) {

}
func CloseExpHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue == models.RoleDirector {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	// userIdValue, existsId := c.Get("userID")
	// userId, ok1 := userIdValue.(int)
	// if !existsId || !ok1 {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
	// 	return
	// }
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	table, err := repository.GetJournalRow(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, table)
}
