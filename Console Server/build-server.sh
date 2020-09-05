#!/bin/bash
sudo service mysql start
go get github.com/gorilla/securecookie
go build -o Server Server.go Crypto.go FileWork.go Session.go