#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

TABLE="$1"
ID=$(./tests/connections.sh | jq -r '.[0].id')

if [ -z "$ID" ] || [ -z "$TABLE" ]; then
  echo "Usage: ./describe-table.sh <connection-id> <table-name>"
  exit 1
fi

echo "Testing GET /api/connections/$ID/tables/$TABLE"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connections/$ID/tables/$TABLE"

echo ""