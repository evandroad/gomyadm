#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste",
    "values": {
      "name": "teste",
      "age": 10,
      "id": null
    }
  }' \
  "$BASE_URL/api/tables/item"

echo ""