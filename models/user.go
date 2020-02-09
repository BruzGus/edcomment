package models

import "github.com/jinzhu/gorm"

// User ..., Modelo usuario para el formatear la DATA
type User struct {
	gorm.Model
	Userame     string    `json:"username" gorm:"not null;unique"`
	Email       string    `json:"email" gorm:"notnull;unique"`
	Fullname    string    `json:"fullname" gorm:"not null"`
	Pass        string    `json:"pass,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPass string    `json:"confirmPass,omitempty" gorm:"-"`
	Picture     string    `json:"picture"`
	Comments    []Comment `json:"comments,omitempty"`
}
