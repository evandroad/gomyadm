#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing POST /api/connections/local/disconnect"

curl \
  --silent \
  --show-error \
  --request POST \
  "$BASE_URL/api/connections/local/disconnect"

echo ""