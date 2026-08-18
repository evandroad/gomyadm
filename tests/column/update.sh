#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste",
    "oldName": "teste",
    "column": {
      "name": "teste",
      "type": "VARCHAR",
      "length": 100,
      "nullable": true,
      "primary": false,
      "unique": true,
      "autoIncrement": false,
      "defaultValue": "teste"
    }
  }' \
  "$BASE_URL/api/tables/column"

echo ""