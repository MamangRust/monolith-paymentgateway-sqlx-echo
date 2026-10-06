#!/usr/bin/env bash
# Stop all natively-run Go services started by scripts/run-local.sh.
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT="$(pwd)"

PID_DIR="${PID_DIR:-$ROOT/tmp/local-pids}"
LOG_DIR="${LOG_DIR:-$ROOT/tmp/local-logs}"

if [[ -d "$PID_DIR" ]]; then
  for pidfile in "$PID_DIR"/*.pid; do
    [[ -e "$pidfile" ]] || continue
    pid="$(cat "$pidfile" 2>/dev/null || true)"
    if [[ -n "$pid" ]] && kill -0 "$pid" 2>/dev/null; then
      echo "⏹  Stopping pid $pid ($(basename "$pidfile"))"
      kill "$pid" 2>/dev/null || true
    fi
    rm -f "$pidfile"
  done
else
  echo "No pid directory ($PID_DIR); nothing to stop."
fi

echo "✅ All local services stopped. (logs kept in $LOG_DIR)"
