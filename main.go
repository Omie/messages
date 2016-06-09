package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/omie/messages/api"
	_ "github.com/omie/messages/api/messages"
	"github.com/omie/messages/lib/db"
)

/*
entry point of the program
	- initializes database
	- initializes api
	- initializes server
return fatal error on failure
*/
func main() {
	log.SetLevel(log.DebugLevel)

	// get parameters and initialize the database
	dbDriver := os.Getenv("MESSAGES_DB_DRIVER")
	dbName := os.Getenv("MESSAGES_DB_NAME")
	if dbDriver == "" || dbName == "" {
		log.Error("main: db driver or db name not set")
		return
	}

	if err := db.InitDB(dbDriver, dbName); err != nil {
		log.Error("Error initializing database", err.Error())
		return
	}

	// get parameters and initialize the server
	host := os.Getenv("MESSAGES_SERVER_HOST")
	port := os.Getenv("MESSAGES_SERVER_PORT")
	if host == "" || port == "" {
		log.Error("main: host or port not set")
		return
	}
	// make sure port is a number
	intPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		log.Fatal("port must be a number")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, intPort),
		Handler: api.Container,
	}
	log.Println("--- listening on ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
