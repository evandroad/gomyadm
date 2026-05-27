#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing POST /api/connections/local/connect"

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{
    "id": "local",
    "name": "Local MySQL",
    "driver": "mysql",
    "host": "127.0.0.1",
    "port": 9906,
    "username": "evandro",
    "password": "evandro",
    "database": "2804618_contas"
  }' \
  "$BASE_URL/api/connections/connect"

echo ""