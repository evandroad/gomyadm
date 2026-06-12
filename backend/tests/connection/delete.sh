#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

curl -sS -X DELETE "$BASE_URL/api/connection/d8m2va6sd8766cdpbnm0"

echo ""