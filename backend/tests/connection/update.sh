#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{
    "id": "d8m2va6sd8766cdpbnm0",
    "name": "super-api",
    "driver": "mysql",
    "host": "localhost",
    "port": 9906,
    "username": "evandro",
    "password": "evandro",
    "database": "2804618_contas"
  }' \
  "$BASE_URL/api/connection"

echo ""