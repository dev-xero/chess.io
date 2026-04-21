#!/bin/bash

set -e

echo "ssl = on" >> /var/lib/postgresql/data/postgresql.conf
echo "ssl_cert_file = '/var/lib/postgresql/server.crt'" >> /var/lib/postgresql/data/postgresql.conf
echo "ssl_key_file = '/var/lib/postgresql/server.key'" >> /var/lib/postgresql/data/postgresql.conf
echo "ssl_ca_file = '/var/lib/postgresql/rootCA.crt'" >> /var/lib/postgresql/data/postgresql.conf

echo "hostssl all all all cert clientcert=verify-full" > /var/lib/postgresql/data/pg_hba.conf
