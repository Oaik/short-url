// postgres practice

package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresDB struct {
	Db *sql.DB
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// TODO: remove init and make just new func
func New(config DbConfig) *PostgresDB {
	w := PostgresDB{}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)
	var err error
	w.Db, err = sql.Open("postgres", psqlconn)
	w.checkIfError(err)
	return &w
}

func (e *PostgresDB) checkIfError(err error) {
	if err != nil {
		log.Fatal("error", err)
	}
}
