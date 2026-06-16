package handler

import (
	"fmt"
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

	// 1. ПАГИНАЦИЯ под MUI DataGrid (переходим с last_id на page/pageSize)
	// MUI присылает страницу, начиная с 0 (page=0 — первая страница)
	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0
	}

	pageSizeStr := c.DefaultQuery("limit", "25") // Переименовано или оставлено limit под pageSizeOptions={[25, 50, 100]}
	limit, err := strconv.Atoi(pageSizeStr)
	if err != nil || limit <= 0 {
		limit = 25
	}
	const maxLimit = 100
	if limit > maxLimit {
		limit = maxLimit
	}

	// Вычисляем OFFSET для SQL-запроса: (номер страницы * размер страницы)
	offset := page * limit

	// 2. СОРТИРОВКА (MUI DataGrid присылает их при клике на заголовки колонок)
	sortField := c.DefaultQuery("sort_field", "id")   // по умолчанию сортируем по id
	sortOrder := c.DefaultQuery("sort_order", "desc") // по умолчанию свежие сверху

	// Валидация полей сортировки во избежание SQL-инъекций
	allowedFields := map[string]bool{"id": true, "data_post": true, "is_closed": true}
	if !allowedFields[sortField] {
		sortField = "id"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// 3. ФИЛЬТРЫ
	statusFilter := c.Query("status")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	// 4. ЗАПРОС К СЕРВИСУ
	// Внимание: ваш service.IndexGetService должен быть обновлен!
	// Теперь вместо lastID передается offset, limit, а также sortField и sortOrder.
	// И теперь сервис должен возвращать ДВА значения: срез структур (rows) и общее число записей int (totalCount)
	rows, totalCount, err := service.IndexGetService(userId, userRole, offset, limit, sortField, sortOrder, statusFilter, dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные: " + err.Error()})
		return
	}

	fmt.Printf("Fetched %d rows out of %d total\n", len(rows), totalCount)

	// 5. ОТВЕТ ДЛЯ СЕРВЕРНОЙ ТАБЛИЦЫ
	// Возвращаем объект, который без проблем распарсит фронтенд для DataGrid
	c.JSON(http.StatusOK, gin.H{
		"rows":  rows,       // Массив записей для текущей страницы
		"total": totalCount, // Общее количество подходящих под фильтр строк в БД
	})
}
