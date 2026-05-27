#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /api/connections/local/tables"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connections/local/tables"

echo ""