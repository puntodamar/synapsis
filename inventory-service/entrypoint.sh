#!/usr/bin/env sh
set -e

if [ -z "$DATABASE_URL" ]; then
  echo "DATABASE_URL is not set" >&2
  exit 1
fi

echo "[goose] perfoming migration"
goose -dir /app/db/migrations postgres "$DATABASE_URL" up

echo "[inventory] starting service (http:${HTTP_ADDRESS}, grpc:${GRPC_ADDRESS})..."
exec /usr/local/bin/inventory
