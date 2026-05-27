#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

ID="$1"

if [ -z "$ID" ]; then
  ID=$(./tests/connections.sh | jq -r '.[0].id')
fi

if [ -z "$ID" ]; then
  echo "Usage: ./list-tables.sh <connection-id>"
  exit 1
fi

echo "Testing GET /api/connections/$ID/tables"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connections/$ID/tables"

echo ""