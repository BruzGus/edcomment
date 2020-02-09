package migration

import (
	"github.com/BruzGus/edcomment/configuration"
	"github.com/BruzGus/edcomment/models"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(models.User{})

}
