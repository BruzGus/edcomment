package models

import "github.com/jinzhu/gorm"

// Vote ..., modelo de tipo estructura para el manejo de los votos, permite control solo Vote
// una unica vez por cada comentario.
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
