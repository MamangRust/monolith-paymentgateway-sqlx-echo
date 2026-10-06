#!/usr/bin/env bash
# Run the full Go stack natively (host processes) against the docker-compose
# infrastructure (postgres/pgbouncer, per-service Redis, Kafka).
#
# Prereqs: infra must be up (scripts/up-infra.sh). This script then:
#   1. loads deployments/local/local.env (localhost hosts, host-mapped ports)
#   2. runs migrations
#   3. seeds the DB (seeder/seeder.sql) if users are missing
#   4. builds and starts every Go service in the background
#   5. waits for the API gateway on :5000
#
# Usage:
#   scripts/run-local.sh            # start everything
#   scripts/stop-local.sh           # stop everything
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT="$(pwd)"

LOCAL_ENV="deployments/local/local.env"
MIGRATIONS_DIR="pkg/database/migrations"
LOG_DIR="${LOG_DIR:-$ROOT/tmp/local-logs}"
PID_DIR="${PID_DIR:-$ROOT/tmp/local-pids}"
BIN_DIR="${BIN_DIR:-$ROOT/bin}"
mkdir -p "$LOG_DIR" "$PID_DIR" "$BIN_DIR"

# gRPC services (ports hardcoded in cmd/main.go) + email worker + apigateway.
GRPC_SERVICES=(auth role user card merchant saldo topup transaction transfer withdraw)
EXTRA_SERVICES=(email apigateway)

echo "📦 Loading env from $LOCAL_ENV"
set -a
# shellcheck disable=SC1090
source "$LOCAL_ENV"
set +a

echo "🐘 Waiting for Postgres/PgBouncer on localhost:6432 ..."
DB_READY=0
for i in $(seq 1 60); do
  if (exec 3<>/dev/tcp/localhost/6432) 2>/dev/null; then
    exec 3>&- 3<&-
    DB_READY=1
    break
  fi
  sleep 2
done
if [[ "$DB_READY" != "1" ]]; then
  echo "❌ Postgres/PgBouncer not reachable. Run scripts/up-infra.sh first."
  exit 1
fi

echo "🚀 Running migrations (dir: $MIGRATIONS_DIR)"
go run ./service/migrate/cmd/main.go -dir ./"$MIGRATIONS_DIR" up

# Seed only when users are missing (seeder.sql is not idempotent).
USER_COUNT=$(docker exec postgres psql -U DRAGON -d PAYMENT_GATEWAY -tAc "SELECT count(*) FROM users;" 2>/dev/null | tr -d ' ' || echo 0)
if [[ "${USER_COUNT:-0}" == "0" ]]; then
  echo "🌱 Seeding database (seeder/seeder.sql)..."
  docker exec -i postgres psql -U DRAGON -d PAYMENT_GATEWAY < seeder/seeder.sql
else
  echo "🌱 Users already seeded ($USER_COUNT rows), skipping seed."
fi

start_service() {
  local name="$1"
  local log="$LOG_DIR/$name.log"
  local pid="$PID_DIR/$name.pid"
  echo "▶️  Starting $name (log: $log)"
  (cd "$ROOT/service/$name" && nohup go run ./cmd/main.go >"$log" 2>&1 & echo $! > "$pid")
}

# Stop any previous run first (best-effort).
scripts/stop-local.sh >/dev/null 2>&1 || true

for svc in "${GRPC_SERVICES[@]}"; do
  start_service "$svc"
done
for svc in "${EXTRA_SERVICES[@]}"; do
  start_service "$svc"
done

echo
echo "⏳ Waiting for API Gateway on http://localhost:5000 ..."
READY=0
for i in $(seq 1 90); do
  if curl -fsS -o /dev/null http://localhost:5000/health 2>/dev/null; then
    READY=1
    break
  fi
  sleep 2
done

echo
if [[ "$READY" == "1" ]]; then
  echo "✅ API Gateway is up: http://localhost:5000"
  echo "   PIDs:  $(cat "$PID_DIR"/*.pid | tr '\n' ' ')"
  echo "   Logs:  $LOG_DIR"
  echo
  echo "   Run the E2E hurl suite with:"
  echo "     bash hurl/run_tests.sh"
else
  echo "❌ API Gateway did not become ready within the timeout."
  echo "   Inspect logs:"
  for svc in "${GRPC_SERVICES[@]}" "${EXTRA_SERVICES[@]}"; do
    echo "     tail -50 $LOG_DIR/$svc.log"
  done
  exit 1
fi
