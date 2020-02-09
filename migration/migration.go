package migration

import (
	"github.com/BruzGus/edcomment/configuration"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

}
