package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func ShowUserHandler(c *gin.Context) {
	// 1. Проверка авторизации текущего пользователя (кто делает запрос)
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	// Опционально: здесь можно добавить проверку прав.
	// Например, разрешать просмотр списка пользователей только админам (role == 1)

	// 2. Пагинация под MUI DataGrid
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

	// 3. СОРТИРОВКА: адаптирована под колонки таблицы users
	sortField := c.DefaultQuery("sort_field", "id")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	// White-list: разрешаем сортировку по реальным полям пользователей
	allowedFields := map[string]bool{
		"id":         true,
		"login":      true,
		"email":      true,
		"first_name": true,
		"last_name":  true,
		"role":       true,
	}

	// Если пришло что-то другое, сбрасываем на ID
	if !allowedFields[sortField] {
		sortField = "id"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// 4. ФИЛЬТРЫ: адаптированы под поиск пользователей
	searchQuery := c.Query("search") // Поиск по подстроке (login/email/имя)
	roleFilter := c.Query("role")    // Фильтр по конкретной роли

	// 5. Запрос к сервису пользователей
	users, totalCount, err := service.ShowUsersService( // Данные того, кто запрашивает (для разграничения прав внутри сервиса)
		offset, limit,
		sortField, sortOrder,
		searchQuery, roleFilter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список пользователей: " + err.Error()})
		return
	}

	// 6. Ответ для MUI DataGrid
	c.JSON(http.StatusOK, gin.H{
		"rows":  users, // Массив структур models.User (пароли из json тегов структуры лучше исключить через `json:"-"`)
		"total": totalCount,
	})
}
