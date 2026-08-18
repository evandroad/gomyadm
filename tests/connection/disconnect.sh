#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

URI="api/connection/disconnect"
echo "Testing POST /$URI"

curl -sS -X POST "$BASE_URL/$URI"

echo ""