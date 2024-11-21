package models

import "time"

// Departament описывает департамент
type Departament struct {
	// ID департамента
	ID uint `json:"id" example:"1"`
	// Время создания записи
	CreatedAt time.Time `json:"created_at" example:"2023-11-22T12:00:00Z"`
	// Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-22T12:30:00Z"`
	// Время мягкого удаления записи
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-11-22T13:00:00Z"`

	// Название департамента
	Name string `json:"name" example:"IT"`
}
