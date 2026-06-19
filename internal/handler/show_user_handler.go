package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func ShowUserHandler(c *gin.Context) {

	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

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

	sortField := c.DefaultQuery("sort_field", "id")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	allowedFields := map[string]bool{
		"id":         true,
		"login":      true,
		"email":      true,
		"first_name": true,
		"last_name":  true,
		"role":       true,
	}

	if !allowedFields[sortField] {
		sortField = "id"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	searchQuery := c.Query("search")
	roleFilter := c.Query("role")

	users, totalCount, err := service.ShowUsersService(
		offset, limit,
		sortField, sortOrder,
		searchQuery, roleFilter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список пользователей: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rows":  users,
		"total": totalCount,
	})
}
