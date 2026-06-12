package handler

import (
	"net/http"
	"strconv"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var req models.Exp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации полей", "details": err.Error()})
		return
	}
	err = repository.CloseExpRepo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить данные: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
