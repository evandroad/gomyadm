#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "Testing GET /health"

curl -sS "$BASE_URL/health"

echo ""