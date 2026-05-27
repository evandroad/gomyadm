#!/usr/bin/sh

source "$(dirname "$0")/common.sh"

echo "Testing GET /health"

curl \
  --silent \
  --show-error \
  --location \
  "$BASE_URL/health"

echo ""