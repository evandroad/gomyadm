#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl -sS "$BASE_URL/api/connection"

echo ""