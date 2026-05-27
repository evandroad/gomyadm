#!/usr/bin/bash

source "$(dirname "$0")/common.sh"

echo "Testing health endpoint..."

curl \
  --silent \
  --show-error \
  --location \
  "$BASE_URL/health"

echo ""