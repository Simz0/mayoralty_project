package models

import (
	"time"
)

// WorkPoints описывает модель точек работы
// @Description Модель для работы с точками
type Work_points struct {
	// ID записи
	ID uint `json:"id" example:"1"`
	// Время создания записи
	CreatedAt time.Time `json:"created_at" example:"2023-11-22T12:00:00Z"`
	// Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-22T12:30:00Z"`
	// Время мягкого удаления записи
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-11-22T13:00:00Z"`
	// Долгота точки
	Longitude float64 `json:"longitude" example:"37.7749"`
	// Широта точки
	Latitude float64 `json:"latitude" example:"55.7512"`
	// Идентификатор сессии
	SessionID uint `json:"session_id" gorm:"column:session_id" example:"42"`
}
