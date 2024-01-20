package models

import "time"

type Token struct {
	ID           string    `json:"id" db:"id" gorm:"column:id;size:255;primary_key;"`
	UserID       string    `json:"user_id" db:"user_id" gorm:"column:user_id;size:255;index;"`
	Token        string    `json:"token" gorm:"column:token;size:255;index;"`
	RefreshToken string    `json:"refresh_token" gorm:"column:token;size:255;index;"`
	Status       string    `json:"status" gorm:"column:status;size:10;index;"`
	ExpireAt     time.Time `json:"expire_at" gorm:"column:expire_at;index;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;index;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;index;"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	TokenBody TokenBody `json:"tokenBody"`
}

type TokenBody struct {
	ExpireAt  time.Time `json:"tokenExpiredDate"`
	TokenIdx  int       `json:"tokenIdx"`
	TokenType int       `json:"tokenType"`
}
