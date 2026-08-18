#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste",
    "column": {
      "name": "teste",
      "type": "INT",
      "length": 10,
      "nullable": true,
      "primary": false,
      "unique": true,
      "autoIncrement": false,
      "defaultValue": ""
    }
  }' \
  "$BASE_URL/api/tables/column"

echo ""