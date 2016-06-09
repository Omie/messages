package db

import (
	"database/sql"
	"errors"
	"os"

	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)

var dbDriver, dbName string

// checks if local database exists, creates database if it doesnt
// param: driver: string: sql driver name
// param: dbname: string: database name
// returns error or nil
func InitDB(driver, dbname string) error {
	var err error
	createDb := false

	if _, err := os.Stat(dbname); os.IsNotExist(err) {
		log.Info("InitDB: database does not exist ", dbname)
		createDb = true
	}

	db, err := sql.Open(driver, dbname)
	if err != nil {
		return errors.New("Error opening connection to the db")
	}
	defer db.Close()

	if createDb {
		if err = syncDB(db); err != nil {
			return errors.New("Error creating database")
		}
	}
	dbDriver = driver
	dbName = dbname

	return nil
}

// run DDL queries to create the database
// param: db: *sql.DB: sql db instance
// returns error or nil
func syncDB(db *sql.DB) error {
	log.Debug("--- syncDB")
	//create messages table
	createMessagesTable := `CREATE TABLE messages (
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                uuid char(32) UNIQUE NOT NULL,
                text text  NOT NULL,
                created_on TIMESTAMP default CURRENT_TIMESTAMP
            );`
	_, err := db.Exec(createMessagesTable, nil)
	return err
}

// returns a new connection to the database
// returns sql.DB instance on success or returns error
func GetDB() (*sql.DB, error) {
	return sql.Open(dbDriver, dbName)
}
