package models

import "time"

type Budget struct {
	ID              string    `json:"id" gorm:"type:uuid;primary_key"`
	Category        string    `json:"category" gorm:"type:varchar(255);not null"`
	AllocatedAmount float64   `json:"allocated_amount" gorm:"type:numeric(10,2);not null"`
	SpentAmount     float64   `json:"spent_amount" gorm:"type:numeric(10,2);default:0.00"`
	StartDate       time.Time `json:"start_date" gorm:"not null"`
	EndDate         time.Time `json:"end_date" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
