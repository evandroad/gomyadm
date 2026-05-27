#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connections"

echo ""