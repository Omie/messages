#!/bin/bash

export MESSAGES_SERVER_HOST=localhost
export MESSAGES_SERVER_PORT=8000

export MESSAGES_DB_DRIVER=sqlite3
export MESSAGES_DB_NAME=dbtest.sqlite

go get
go test test/messages_test.go

