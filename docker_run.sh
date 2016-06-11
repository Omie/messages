#!/bin/bash

mkdir data

go get github.com/mattes/migrate

echo "running migrations.."
migrate -url sqlite3://data/db.sqlite -path ./migrations up

echo "fetching latest image"
docker pull omie/messages

echo "starting up"
docker-compose up -d

