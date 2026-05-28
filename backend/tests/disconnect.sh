#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing POST /api/connection/disconnect"

curl \
  --silent \
  --show-error \
  --request POST \
  "$BASE_URL/api/connection/disconnect"

echo ""