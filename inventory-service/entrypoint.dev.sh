#!/usr/bin/env sh
set -e

MIG_DIR="${MIG_DIR:-./db/migrations}"
SEED_DIR="${SEED_DIR:-./db/seeds}"


if [ -n "${DATABASE_URL:-}" ]; then
  if find "$MIG_DIR" -type f -name '*.go' -print -quit | grep -q .; then

    # migrate from go file
    [ -d ./cmd/migrate ] && go build -o ./bin_migrate ./cmd/migrate
    DB_DSN="$DATABASE_URL" ./bin_migrate || true
  else
    # migrate from sql
    goose -dir "$MIG_DIR" postgres "$DATABASE_URL" up || true
  fi

  if [ "${RUN_SEEDS:-false}" = "true" ] && [ -d "$SEED_DIR" ]; then
    if find "$SEED_DIR" -type f -name '*.go' -print -quit | grep -q .; then
      [ -d ./cmd/seed ] && go build -o ./bin/seed ./cmd/seed
      DB_DSN="$DATABASE_URL" ./bin/seed || true
    else
      goose -dir "$SEED_DIR" postgres "$DATABASE_URL" up || true
    fi
  fi
else
  echo "[warn] DATABASE_URL not set; skipping migrations/seeds."
fi

exec air -c .air.toml
