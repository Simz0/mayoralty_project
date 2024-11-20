package models

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`
	//не нашёл, как можно прям дату пихать :(
	Stop_at time.Time
}
