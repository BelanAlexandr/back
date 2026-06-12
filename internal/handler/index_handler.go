package handler

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func IndexHandlerShow(c *gin.Context) {
	_, existsRole := c.Get("userRole")
	_, existsID := c.Get("userID")
	if !existsRole || !existsID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, nil)
}

func IndexHandler(c *gin.Context) {

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

	lastIDStr := c.DefaultQuery("last_id", "2147483647")
	lastID, err := strconv.Atoi(lastIDStr)
	if err != nil || lastID <= 0 {
		lastID = 2147483647
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	const maxLimit = 100
	if limit > maxLimit {
		limit = maxLimit
	}

	statusFilter := c.Query("status")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	tables, err := service.IndexGetService(userId, userRole, lastID, limit, statusFilter, dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}
