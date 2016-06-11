#!/bin/bash

export MESSAGES_DB_DRIVER=sqlite3
export MESSAGES_DB_NAME=dbtest.sqlite

go get github.com/mattes/migrate

go get

migrate -url sqlite3://test/dbtest.sqlite -path ./migrations up

go test test/messages_test.go
rm test/dbtest.sqlite

