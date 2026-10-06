#!/usr/bin/env bash
# Bring up ONLY the infrastructure services from deployments/local/docker-compose.yml
# (postgres, pgbouncer, per-service Redis, kafka) plus the observability profile.
# The Go services themselves are intentionally NOT started here — run them
# natively with scripts/run-local.sh.
set -euo pipefail
cd "$(dirname "$0")/.."

COMPOSE_FILE="deployments/local/docker-compose.yml"

INFRA_SERVICES=(
  postgres
  pgbouncer
  redis-apigateway
  redis-auth
  redis-user
  redis-role
  redis-card
  redis-merchant
  redis-saldo
  redis-topup
  redis-transaction
  redis-transfer
  redis-withdraw
  kafka
)

OBSERVABILITY_SERVICES=(
  postgres-exporter
  pyroscope
  node-exporter
  alertmanager
  prometheus
  grafana
  otel-collector
  jaeger
  kafka-exporter
  loki
)

echo "🚀 Starting infra services (pg, pgbouncer, redis, kafka)..."
docker compose -f "$COMPOSE_FILE" up -d "${INFRA_SERVICES[@]}"

if [[ "${WITH_OBSERVABILITY:-0}" == "1" ]]; then
  echo "🚀 Starting observability profile services..."
  docker compose -f "$COMPOSE_FILE" --profile observability up -d "${OBSERVABILITY_SERVICES[@]}"
fi

echo "✅ Infra services started."
echo
echo "Next: run migrations + seed + Go services natively with:"
echo "  scripts/run-local.sh"
