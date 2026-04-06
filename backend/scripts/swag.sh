#!/bin/bash
 
export PATH=$PATH:$(go env GOPATH)/bin

swag fmt
swag init -g ./cmd/server/main.go
