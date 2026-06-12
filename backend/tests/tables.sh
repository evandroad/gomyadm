#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

URI="api/tables"
echo "Testing GET /$URI"

curl -sS "$BASE_URL/$URI"

echo ""