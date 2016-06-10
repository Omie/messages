Messages
------


What is it
-----------

A simple API to store and retrive messages


Running the program
--------------------

Edit `run.sh` to set environment variables accordingly and do `./run.sh`


Testing the program
--------------------

Edit `test_run.sh` to set environment variables accordingly and do `./test_run.sh`


API
----

documentation is at https://docs.messages19.apiary.io/

- endpoint `/messages`

    - POST:
    param `text` : string : message to store

- endpoint `/messages/<id>`

    - GET:
    retrive stored message

