#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"database": "my_system"}' \
  "$BASE_URL/api/database/select"

echo ""