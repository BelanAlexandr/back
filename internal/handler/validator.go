package handler

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex  = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	digitsRegex = regexp.MustCompile(`\D`)

	nameRegex = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁәӘғҒқҚңҢөӨұҰүҮһҺіІ\s-]+$`)
)

func NewValidator() *validator.Validate {
	v := validator.New()

	v.RegisterValidation("custom_email", func(fl validator.FieldLevel) bool {
		email := strings.ToLower(strings.TrimSpace(fl.Field().String()))
		return emailRegex.MatchString(email)
	})

	v.RegisterValidation("custom_phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()

		cleanPhone := digitsRegex.ReplaceAllString(phone, "")

		if strings.HasPrefix(cleanPhone, "8") && len(cleanPhone) == 11 {
			cleanPhone = "7" + cleanPhone[1:]
		}

		return len(cleanPhone) >= 10 && len(cleanPhone) <= 15
	})

	v.RegisterValidation("custom_name", func(fl validator.FieldLevel) bool {
		name := strings.TrimSpace(fl.Field().String())

		return nameRegex.MatchString(name)
	})

	return v
}
