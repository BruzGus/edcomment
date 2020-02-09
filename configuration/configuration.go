package configuration

import (
	"encoding/json"
	"log"
	"os"

	
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
