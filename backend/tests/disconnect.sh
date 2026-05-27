#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

ID="$1"

if [ -z "$ID" ]; then
  ID=$(./tests/connections.sh | jq -r '.[0].id')
fi

if [ -z "$ID" ]; then
  echo "Usage: ./disconnect.sh <connection-id>"
  exit 1
fi

echo "Testing POST /api/connections/$ID/disconnect"

curl \
  --silent \
  --show-error \
  --request POST \
  "$BASE_URL/api/connections/$ID/disconnect"

echo ""