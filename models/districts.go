package models

import "gorm.io/gorm"

type District struct {
	gorm.Model
	Name      string
	ManagerID uint `gorm:"column:manager_id"`
}
