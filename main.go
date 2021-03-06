package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"short-url/httpServer"
	"short-url/postgres"
	"short-url/sqlInstance"
	"short-url/workerGenerator"

	_ "github.com/lib/pq"
)

type Config struct {
	NumOfRandomLetters int
	NumOfRandomDigits  int
	NumOfWorkers       int
	DbConfig           DbConfig
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func ReadConfigFile() *Config {
	jsonFile, err := os.Open("config.json")
	exitOnError(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Config
	json.Unmarshal(byteValue, &result)
	return &result
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := ReadConfigFile()
	databaseConnection := postgres.New(postgres.DbConfig(config.DbConfig))
	defer databaseConnection.Db.Close()
	sqlInstance := sqlInstance.New(databaseConnection)

	workerGen := workerGenerator.New(sqlInstance, config.NumOfWorkers, config.NumOfRandomLetters, config.NumOfRandomDigits)
	defer workerGen.Close()

	server := httpServer.New(workerGen)
	server.StartServer()
}
