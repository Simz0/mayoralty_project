package models

import "time"

// Sessions описывает сессию пользователя
type Sessions struct {
	// ID сессии
	ID uint `json:"id" example:"1"`
	// Время создания записи
	CreatedAt time.Time `json:"created_at" example:"2023-11-22T12:00:00Z"`
	// Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-22T12:30:00Z"`
	// Время мягкого удаления записи
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-11-22T13:00:00Z"`

	// ID пользователя, связанного с сессией
	UserID uint `json:"user_id" example:"5"`
	// Время окончания сессии
	StopAt time.Time `json:"stop_at" example:"2023-11-22T18:00:00Z"`
}
