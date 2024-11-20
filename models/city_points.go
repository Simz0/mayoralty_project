package models

import "gorm.io/gorm"

type City_points struct {
	gorm.Model
	Longitude  float64
	Latitude   float64
	DistrictID uint `gorm:"column:district_id"`
}
