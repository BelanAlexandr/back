package models

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"firstName"`
	Second_Name  string `json:"lastName"`
	Middle_Name  string `json:"middleName"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone"`
	Role         int    `json:"role"`
}
