package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Username     string    `json:"username" gorm:"type:varchar(255);unique;not null"`
	PasswordHash string    `json:"-" gorm:"column:password_hash;type:text;not null"` // Coluna correta para armazenar o hash da senha
	Email        string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
