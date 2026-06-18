package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	noti, err := repository.GetNotifications(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка загрузки уведомлений"})
		return
	}
	c.JSON(http.StatusOK, noti)
}
func SetMarkNotification(c *gin.Context) {
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = repository.MarkNotification(id, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка отметить уведомление как прочитанное"})
		return
	}
	c.JSON(http.StatusOK, err)
}
func MarkAllNotification(c *gin.Context) {
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	err := repository.MarkAllNotifications(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка отметить уведомление как прочитанное"})
		return
	}
	c.JSON(http.StatusOK, err)
}
