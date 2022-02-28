package insertInstance

import (
	"fmt"
	"log"
	"short-url/postgres"
)

type SqlInsert struct {
	currentFields map[string]string
}

func New() *SqlInsert {
	w := SqlInsert{
		currentFields: map[string]string{},
	}
	return &w
}

func (e *SqlInsert) AddPair(key, value string) {
	if e.checkIfMapContainKey(key) {
		log.Fatal("Adding value in already existed key", key, value)
		return
	}
	e.currentFields[key] = value
}

func (e *SqlInsert) checkIfMapContainKey(key string) bool {
	_, found := e.currentFields[key]
	return found
}

func (e *SqlInsert) Execute(dbConnection *postgres.PostgresDB, tableName string) {
	// insertDynStmt := `INSERT INTO "url"("shortUrl", "website") values($1, $2)`
	mapKeys, mapValues := "", ""
	prv := false
	for key, value := range e.currentFields {
		if prv {
			mapKeys = mapKeys + ", "
			mapValues = mapValues + "', "
		}
		mapKeys = mapKeys + key
		mapValues = mapValues + "'" + value
		prv = true
	}

	insertDynStmt := "INSERT INTO " + tableName + " (" + mapKeys + ") values(" + mapValues + "')"
	fmt.Println(insertDynStmt)
	_, err := dbConnection.Db.Exec(insertDynStmt)
	e.checkIfError(err)
}

func (e *SqlInsert) checkIfError(err error) {
	if err != nil {
		log.Fatal("error", err)
	}
}
