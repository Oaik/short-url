package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "yourpassword"
	dbname   = "postgres"
)

type PostgresDB struct {
	Db *sql.DB
}

func New() *PostgresDB {
	w := PostgresDB{}
	w.init()
	return &w
}

func (e *PostgresDB) init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	e.Db, err = sql.Open("postgres", psqlconn)
	e.checkIfError(err)
}

func (e *PostgresDB) CheckIfShortUrlNotExists(shortUrl string) bool {
	website := e.SelectWebsiteMatchShortUrl(shortUrl)
	if website == "" {
		return true
	}
	return false
}

func (e *PostgresDB) SelectWebsiteMatchShortUrl(shortUrl string) string {
	insertDynStmt := `SELECT "website" FROM "url" WHERE "shortUrl" = $1`
	rows, err := e.Db.Query(insertDynStmt, shortUrl)
	e.checkIfError(err)
	for rows.Next() {
		var website string
		err = rows.Scan(&website)
		e.checkIfError(err)
		return website
	}
	return ""
}

func (e *PostgresDB) InsertShortUrlAndWebsite(shortUrl, website string) {
	insertDynStmt := `insert into "url"("shortUrl", "website") values($1, $2)`
	_, err := e.Db.Exec(insertDynStmt, shortUrl, website)
	e.checkIfError(err)
}

func (e *PostgresDB) checkIfError(err error) {
	if err != nil {
		log.Fatal("error", err)
	}
}
