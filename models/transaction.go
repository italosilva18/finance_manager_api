package models

import "time"

type Transaction struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Amount    float64   `json:"amount" gorm:"type:numeric(10,2);not null"`
	Date      time.Time `json:"date" gorm:"not null"`
	Category  string    `json:"category" gorm:"type:varchar(255)"`
	IsExpense bool      `json:"is_expense" gorm:"type:boolean;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
