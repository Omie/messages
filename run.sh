#!/bin/bash

export MESSAGES_SERVER_HOST=localhost
export MESSAGES_SERVER_PORT=8000

export MESSAGES_DB_DRIVER=sqlite3
export MESSAGES_DB_NAME=db.sqlite

go get github.com/mattes/migrate

go get
go build

migrate -url sqlite3://db.sqlite -path ./migrations up

./messages

