#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste2"
    "column": {
      "name": "teste"
      "type": "INT"
      "length": 10
      "nullable": true
      "primary": true
      "unique": true
      "autoIncrement": true
      "defaultValue": "teste"
    }
  }' \
  "$BASE_URL/api/tables/column"

echo ""