#!/bin/bash

set -e

mkdir -p certs 
cd ./certs

openssl genrsa -out rootCA.key 2048
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 365 -out rootCA.crt -subj "/CN=MyLocalCA"

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out server.crt -days 365 -sha256

openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=postgres"
openssl x509 -req -in client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out client.crt -days 365 -sha256

echo "Certificate files have been generated and saved to $(pwd)"
