package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Configuration ..., estructura para mapear el archivo de configuracion
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

//GetConfiguration ..., obtiene la configuracion del archivo de configuracion
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//GetConnection ..., realiza la coneccion a la base de datos
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset= utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
