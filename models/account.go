package models

import "time"

type Account struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`             // Identificador único da conta
	UserID    string    `json:"user_id" gorm:"type:uuid;not null"`           // Chave estrangeira para a tabela users
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`      // Nome da conta
	Balance   float64   `json:"balance" gorm:"type:decimal(10,2);not null"`  // Saldo da conta
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"` // Data de criação
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"` // Data de última atualização
}
