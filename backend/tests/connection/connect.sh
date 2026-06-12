#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "driver": "mysql",
    "host": "127.0.0.1",
    "port": 3306,
    "username": "admin",
    "password": "senha123"
  }' \
  "$BASE_URL/api/connection/connect"

echo ""