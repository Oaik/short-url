package sqlInstance

import (
	"log"
	"short-url/postgres"
)

type SqlInstance struct {
	*postgres.PostgresDB
	currentFields map[string]string
}

func New(postgresDB *postgres.PostgresDB) *SqlInstance {
	w := SqlInstance{postgresDB, map[string]string{}}
	return &w
}

func (e *SqlInstance) CheckIfShortUrlNotExists(shortUrl string) bool {
	website := e.SelectWebsiteMatchShortUrl(shortUrl)
	if website == "" {
		return true
	}
	return false
}

func (e *SqlInstance) SelectWebsiteMatchShortUrl(shortUrl string) string {
	insertDynStmt := `SELECT website FROM url WHERE shorturl = $1`
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

// pass Map or struct as parameter
// add addPair function to insert into db
// func (e *SqlInstance) InsertShortUrlAndWebsite(shortUrl, website string) {
// 	insertDynStmt := `
// 	INSERT INTO url (shorturl, website)
// 	VALUES ($1, $2)`
// 	_, err := e.Db.Exec(insertDynStmt, shortUrl, website)
// 	e.checkIfError(err)
// }

func (e *SqlInstance) checkIfError(err error) {
	if err != nil {
		log.Fatal("error", err)
	}
}
