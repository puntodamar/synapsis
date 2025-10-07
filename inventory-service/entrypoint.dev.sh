#!/usr/bin/env sh
set -e

go mod tidy
go mod download

if [ -n "$DATABASE_URL" ]; then
  echo "[goose] dev migrations…"
  goose -dir ./migrations postgres "$DATABASE_URL" up || true
fi

echo "[order] starting service (http:${HTTP_ADDRESS})…"
echo "[dev] starting Air…"
exec air -c .air.toml
