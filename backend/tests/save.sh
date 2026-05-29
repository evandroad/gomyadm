#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{
    "name": "super-api",
    "driver": "mysql",
    "host": "localhost",
    "port": 9906,
    "username": "evandro",
    "password": "evandro",
    "database": "2804618_contas"
  }' \
  "$BASE_URL/api/connection/save"

echo ""