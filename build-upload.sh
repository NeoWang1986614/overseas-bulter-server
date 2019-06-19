#!/bin/bash

GOOS=linux GOARCH=amd64 go build main.go

scp -r ./apiclient_cert.pem ./main root@129.28.57.139:/root/opt
