package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
)

func EditUserHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req models.User
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	validate := NewValidator()

	if err := validate.Struct(req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации полей", "details": err.Error()})
		return
	}
	if req.Email != "" {
		req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	}
	if req.Phone_Number != "" {
		req.Phone_Number = digitsRegex.ReplaceAllString(req.Phone_Number, "")
		if strings.HasPrefix(req.Phone_Number, "8") && len(req.Phone_Number) == 11 {
			req.Phone_Number = "7" + req.Phone_Number[1:]
		}
	}
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	err = service.EditUserService(userId, req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
}
