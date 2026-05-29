#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /api/connection/list"

curl -sS "$BASE_URL/api/connection/list"

echo ""