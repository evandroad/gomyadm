#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{"database": "2804618_contas"}' \
  "$BASE_URL/api/connection/database/select"

echo ""