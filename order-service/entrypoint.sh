#!/usr/bin/env sh
set -e
if [ -z "$DATABASE_URL" ]; then
  echo "DATABASE_URL is not set" >&2
  exit 1
fi
echo "[order][goose] performing migration…"
goose -dir /app/db/migrations postgres "$DATABASE_URL" up
echo "[order] starting service (http:${HTTP_ADDRESS})…"
exec /usr/local/bin/order
