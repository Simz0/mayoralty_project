package models

import "time"

// User описывает пользователя системы
type User struct {
	// ID пользователя
	ID uint `json:"id" example:"1"`
	// Время создания записи
	CreatedAt time.Time `json:"created_at" example:"2023-11-22T12:00:00Z"`
	// Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-22T12:30:00Z"`
	// Время мягкого удаления записи
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-11-22T13:00:00Z"`

	// ID департамента, к которому принадлежит пользователь
	DepartamentID uint `json:"departament_id" example:"42"`
	// Имя пользователя
	Firstname string `json:"firstname" example:"Иван"`
	// Фамилия пользователя
	Lastname string `json:"lastname" example:"Иванов"`
	// Отчество пользователя
	Patronymic string `json:"patronymic" example:"Иванович"`
	// Телефонный номер
	Phone string `json:"phone" example:"+79998887766"`
	// Логин пользователя
	Username string `json:"username" example:"ivan.ivanov"`
	// Пароль пользователя
	Password string `json:"password" example:"password123"`
	// Роль пользователя
	Role string `json:"role" example:"admin"`
}
