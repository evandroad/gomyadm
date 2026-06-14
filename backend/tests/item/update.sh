#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste",
    "key": { "id": 10 },
    "values": {
      "name": "teste",
      "age": 1020
    }
  }' \
  "$BASE_URL/api/tables/item"

echo ""