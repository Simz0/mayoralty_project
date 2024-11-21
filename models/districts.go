package models

import "time"

// District описывает район города
type District struct {
	// ID района
	ID uint `json:"id" example:"1"`
	// Время создания записи
	CreatedAt time.Time `json:"created_at" example:"2023-11-22T12:00:00Z"`
	// Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-22T12:30:00Z"`
	// Время мягкого удаления записи
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-11-22T13:00:00Z"`

	// Название района
	Name string `json:"name" example:"Центральный"`
	// ID менеджера, закреплённого за районом
	ManagerID uint `json:"manager_id" example:"7"`
}
