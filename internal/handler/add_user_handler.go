package handler

import (
	"net/http"
	"strings"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddUserHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	validate := validator.New()

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
	userIdValue, existsID := c.Get("userID")
	userId, ok1 := userIdValue.(int)
	if !existsID || !ok1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	err := service.AddUserService(userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
