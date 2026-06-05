#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /api/tables"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/tables"

echo ""