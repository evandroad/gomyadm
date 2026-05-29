#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /api/connection"

curl -sS "$BASE_URL/api/connection"

echo ""