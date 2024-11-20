package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DepartamentID uint `gorm:"column:departament_id"`
	Firstname     string
	Lastname      string
	Patronymic    string
	Phone         string
	Username      string
	Password      string
	Role          string
}
