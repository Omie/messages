package db

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var dbDriver, dbName string
var db *sqlx.DB

// checks if local database exists, saves the references for use if it does
// param: driver: string: sql driver name
// param: dbname: string: database name
// returns error or nil
func InitDB(driver, dbname string) error {
	var err error

	if _, err := os.Stat(dbname); os.IsNotExist(err) {
		return errors.New("database does not exist:" + dbname)
	}

	db, err := sqlx.Connect(driver, dbname)
	if err != nil {
		return errors.New("Error opening connection to the db:" + err.Error())
	}
	defer db.Close()
	dbDriver = driver
	dbName = dbname

	return nil
}

// returns a new connection to the database
// returns sql.DB instance on success or returns error
func GetDB() (*sqlx.DB, error) {
	return sqlx.Open(dbDriver, dbName)
}
