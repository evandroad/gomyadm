#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

TABLE="$1"

if [ -z "$TABLE" ]; then
  echo "Usage: ./getall.sh <table-name>"
  exit 1
fi

URI="api/tables/item/$TABLE"
echo "Testing GET /$URI"

curl -sS "$BASE_URL/$URI"

echo ""