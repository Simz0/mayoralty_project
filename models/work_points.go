package models

import "gorm.io/gorm"

type Work_points struct {
	gorm.Model
	Longitude float64
	Latitude  float64
	SessionID uint `gorm:"column:session_id"`
}
