# 📡 Observability Configuration

This directory contains the **observability stack configuration** for the microservices platform. It includes definitions for monitoring, logging, tracing, and alerting components using tools like **Prometheus**, **Alertmanager**, **Grafana**, **Loki**, **Promtail**, and **OpenTelemetry Collector**.


## Project Structure

```
observability/
├── alertmanager.yml # Alertmanag er routing and notification configuration
├── loki-config.yaml # Loki logging backend configuration 
├── otel-collector.yaml  # OpenTelemetry Collector pipelines and exporters
├── prometheus.yaml  # Prometheus scrape configs and alert rule loading
├── promtail-config.yaml # Promtail log shipper configuration for Loki
├── grafana/         # Grafana provisioning + enterprise dashboards
│   ├── provisioning/
│   │   ├── datasources/datasource.yml  # Prometheus, Loki, Jaeger (fixed UIDs)
│   │   └── dashboards/dashboards.yaml    # file provider -> /var/lib/grafana/dashboards
│   └── dashboards/
│       ├── payment-gateway-overview.json    # SRE landing: health, RPS, errors, P50/95/99, cache, logs
│       ├── business-metrics.json            # Financial domain KPIs per service
│       ├── infrastructure.json              # Host, Kafka, PostgreSQL, OTel Collector
│       └── go-services-runtime.json         # Heap, goroutines, GC goal, allocations per service
├── README.md
└── rules  # Directory for service-specific alert rules
    ├── apigateway-alerts.yaml
    ├── auth-service-alerts.yaml
    ├── card-service-alerts.yaml
    ├── email-service-alerts.yaml
    ├── golang-runtime-alerts.yaml
    ├── kafka-exporter-alerts.yaml
    ├── merchant-service-alerts.yaml
    ├── node-exporter-alerts.yaml
    ├── otel-collector-alerts.yaml
    ├── role-service-alerts.yaml
    ├── saldo-service-alerts.yaml
    ├── topup-service-alerts.yaml
    ├── transaction-service-alerts.yaml
    ├── transfer-service-alerts.yaml
    ├── user-service-alerts.yaml
    └── withdrawal-service-alerts.yaml
```

## Grafana Dashboards

Start the stack with the observability profile, then open **http://localhost:3000** (admin/admin).

The four dashboards are provisioned automatically from `grafana/dashboards/` into the **Payment Gateway** folder:

| Dashboard | UID | Purpose |
| :--- | :--- | :--- |
| **Payment Gateway Overview** | `payment-gateway-overview` | SRE landing — services up, RPS, error rate, P50/P95/P99 latency, cache hit ratio, goroutines, Loki logs |
| **Payment Gateway Business Metrics** | `payment-gateway-business` | Financial domain volume (topup/transfer/withdraw/transaction/merchant/saldo), success vs error, top operations table |
| **Payment Gateway Infrastructure** | `payment-gateway-infra` | Host (CPU/mem/disk/network), Kafka (lag/partitions/throughput), PostgreSQL (conns/locks/cache-hit/size), OTel Collector self-metrics |
| **Payment Gateway Go Runtime** | `payment-gateway-go-runtime` | Per-service heap, goroutines, allocation rate, GC goal/used, processor limit |

### Data sources

- **Prometheus** (`uid: prometheus`) — app metrics via OTel collector (`golang_app_*`), node-exporter, kafka-exporter, postgres-exporter, otel-collector self-metrics
- **Loki** (`uid: loki`) — structured logs with `service_name` label (OTLP via collector); derived `TraceID` field links to Jaeger
- **Jaeger** (`uid: jaeger`) — traces (UI mirror at http://localhost:16687)

### Editing dashboards

Dashboards are file-based (provider polls every 15s). Edit the JSON files under `grafana/dashboards/` and they reload automatically; changes made in the Grafana UI are exported back to disk via `allowUiUpdates`.

### Metric naming notes

- Application metrics are exported by the OTel collector Prometheus exporter with the `golang_app_` namespace: `golang_app_requests_total`, `golang_app_request_duration_seconds`, `golang_app_errors_total`, `golang_app_cache_*`, and OTel runtime metrics `golang_app_go_*` (heap, goroutines, allocations, GC goal).
- Collector self-metrics (`otelcol_*`) are exposed on `otel-collector:8888` and scraped by the `otel-collector` Prometheus job.
- `go_*` / `process_*` series exist only for the exporter jobs (node/kafka/postgres), not for app services.