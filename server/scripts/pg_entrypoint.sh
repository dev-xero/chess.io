#!/bin/bash
set -e

cp /certs/server.crt /var/lib/postgresql/server.crt
cp /certs/server.key /var/lib/postgresql/server.key
cp /certs/rootCA.crt /var/lib/postgresql/rootCA.crt

chmod 600 /var/lib/postgresql/server.* /var/lib/postgresql/rootCA.crt
chown postgres:postgres /var/lib/postgresql/server.* /var/lib/postgresql/rootCA.crt

exec docker-entrypoint.sh postgres
