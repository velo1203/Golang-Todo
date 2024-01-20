package models

type User struct {
	ID       string `json:"id" db:"id" gorm:"column:id;size:255;primary_key;"`
	Username string `json:"username" db:"username" gorm:"column:username;size:50;uniqueIndex;"`
	Password string `json:"-" db:"password" gorm:"column:password;size:255;not null;"`
}
