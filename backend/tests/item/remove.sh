#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

URI="api/tables/item"
echo "Testing GET /$URI"

curl \
  -X DELETE \
  -H "Content-Type: application/json" \
  -d '{
    "table": "teste",
    "key": { "id": 9 }
  }' \
  "$BASE_URL/$URI"

echo ""