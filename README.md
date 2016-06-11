Messages
------


What is it
-----------

A simple API to store and retrive messages

[![circleci status](https://circleci.com/gh/Omie/messages.svg?&style=shield&circletoken=b69d9c0f2264fcff2222d6b431ab38a6e2a825a9)](https://circleci.com/gh/Omie/messages)

Running the program
--------------------

Edit `run.sh` to set environment variables accordingly and do `./run.sh`

In order to run using docker, make sure you have GOPATH set properly.

make sure docker and docker-compose is installed on your system and do `./docker_run.sh`


Testing the program
--------------------

Edit `test_run.sh` to set environment variables accordingly and do `./test_run.sh`

Building docker image
---------------------

    `$ docker build -t yourname/imagename .`

make sure to update image name in docker-compose.yml


API
----

documentation is also at http://docs.messages19.apiary.io/

- endpoint `/messages`

    - POST:
    param `text` : string : message to store

- endpoint `/messages/<id>`

    - GET:
    retrive stored message

