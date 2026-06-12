#!/usr/bin/sh

. "$(dirname "$0")/../common.sh"

TABLE="$1"
COLUMN="$2"

if [ -z "$TABLE" ]; then
  echo "Usage: ./getall.sh <table-name>"
  exit 1
fi

if [ -z "$COLUMN" ]; then
  echo "Usage: ./getall.sh <table-name> <column-name>"
  exit 1
fi

curl -sS -X DELETE "$BASE_URL/api/tables/column/$TABLE/$COLUMN"

echo ""