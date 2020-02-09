package migration

import (
	"github.com/BruzGus/edcomment/configuration"
	"github.com/BruzGus/edcomment/models"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})

	db.Model(&models.Vote{}).AddUniqueIndex("commen_id_user_id_unique", "comment_id", "user_id")
}
