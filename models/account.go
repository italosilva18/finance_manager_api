package models

import "time"

type Account struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Balance   float64   `json:"balance" gorm:"type:numeric(10,2);default:0.00"`
	UserID    string    `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
