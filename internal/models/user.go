package models

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"firstName" validate:"required"`
	Second_Name  string `json:"lastName" validate:"required"`
	Middle_Name  string `json:"middleName"`
	Login        string `json:"login" validate:"required"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone"`
	Role         int    `json:"role" validate:"required"`
}
