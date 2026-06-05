#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{"database": "my_system"}' \
  "$BASE_URL/api/database/select"

echo ""