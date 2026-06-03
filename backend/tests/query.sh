#!/usr/bin/sh

. "$(dirname "$0")/common.sh"

echo "=== SELECT ==="

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{"query":"SELECT * FROM usuarios"}' \
  "$BASE_URL/api/query"

echo ""
echo ""

echo "=== INSERT ==="

curl \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{"query":"INSERT INTO usuarios (nome) VALUES (\"Evandro\")"}' \
  "$BASE_URL/api/query"

echo ""