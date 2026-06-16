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

	// Пагинация под MUI DataGrid
	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0
	}

	pageSizeStr := c.DefaultQuery("limit", "25")
	limit, err := strconv.Atoi(pageSizeStr)
	if err != nil || limit <= 0 {
		limit = 25
	}
	const maxLimit = 100
	if limit > maxLimit {
		limit = maxLimit
	}

	offset := page * limit

	// СОРТИРОВКА: оставляем строго только 3 параметра, которые есть в таблице
	sortField := c.DefaultQuery("sort_field", "id")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	// White-list: разрешаем сортировку ТОЛЬКО по этим трем колонкам
	allowedFields := map[string]bool{
		"id":        true,
		"data_post": true,
		"is_closed": true,
	}

	// Если пришло что-то другое, сбрасываем на сортировку по умолчанию (по ID)
	if !allowedFields[sortField] {
		sortField = "id"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// Фильтры
	statusFilter := c.Query("status")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	// Запрос к сервису
	rows, totalCount, err := service.IndexGetService(userId, userRole, offset, limit, sortField, sortOrder, statusFilter, dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные: " + err.Error()})
		return
	}

	// Ответ для MUI DataGrid
	c.JSON(http.StatusOK, gin.H{
		"rows":  rows,
		"total": totalCount,
	})
}
