#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

TABLE="$1"

if [ -z "$TABLE" ]; then
  echo "Usage: ./getall.sh <table-name>"
  exit 1
fi

curl -sS "$BASE_URL/api/tables/column/$TABLE"

echo ""