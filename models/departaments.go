package models

import "gorm.io/gorm"

type Departament struct {
	gorm.Model
	Name string
}
