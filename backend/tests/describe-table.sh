#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

TABLE="$1"

if [ -z "$TABLE" ]; then
  echo "Usage: ./describe-table.sh <table-name>"
  exit 1
fi

echo "Testing GET /api/connection/tables/$TABLE"

curl \
  --silent \
  --show-error \
  "$BASE_URL/api/connection/tables/$TABLE"

echo ""