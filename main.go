package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/omie/messages/api"
	_ "github.com/omie/messages/api/messages"
	"github.com/omie/messages/lib/db"
)

// check environment parameters and setup database
func setupDB() error {
	// get parameters and initialize the database
	dbDriver := os.Getenv("MESSAGES_DB_DRIVER")
	dbName := os.Getenv("MESSAGES_DB_NAME")
	if dbDriver == "" || dbName == "" {
		return errors.New("main: db driver or db name not set")
	}

	if err := db.InitDB(dbDriver, dbName); err != nil {
		return err
	}
	return nil
}

// check environment parameters and setup http server
func setupServer() (*http.Server, error) {
	// get parameters and initialize the server
	host := os.Getenv("MESSAGES_SERVER_HOST")
	port := os.Getenv("MESSAGES_SERVER_PORT")
	if host == "" || port == "" {
		return nil, errors.New("main: host or port not set")
	}
	// make sure port is a number
	intPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil, errors.New("port must be a number")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, intPort),
		Handler: api.Container,
	}
	return server, nil
}

/*
entry point of the program
	- initializes database
	- initializes api
	- initializes server
return fatal error on failure
*/
func main() {
	log.SetLevel(log.DebugLevel)

	if err := setupDB(); err != nil {
		log.Fatal(err)
	}

	server, err := setupServer()
	if err != nil {
		log.Fatal("error setting up server: ", err)
	}

	log.Println("--- listening on ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
