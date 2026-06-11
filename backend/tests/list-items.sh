#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

TABLE="$1"

if [ -z "$TABLE" ]; then
  echo "Usage: ./describe-table.sh <table-name>"
  exit 1
fi

URI="api/tables/item/$TABLE"
echo "Testing GET /$URI"

curl \
  --silent \
  --show-error \
  "$BASE_URL/$URI"

echo ""