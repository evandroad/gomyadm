#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /api/connection/tables"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connection/tables"

echo ""