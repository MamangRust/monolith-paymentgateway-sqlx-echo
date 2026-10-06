# Kubernetes Resources Documentation

This documentation provides a comprehensive overview of the Kubernetes manifests in this directory. Each file defines a resource or set of resources required to deploy, configure, monitor, and run a monolith-based system.

## Table of Contents

- [Namespaces](#namespaces)
- [ConfigMaps & Secrets](#configmaps--secrets)
- [Core Services](#core-services)
  - [API Gateway](#api-gateway)
  - [Authentication](#authentication)
  - [User Management](#user-management)
  - [Role Management](#role-management)
  - [Card Management](#card-management)
  - [Merchant Management](#merchant-management)
  - [Saldo, TopUp, Withdraw, Transfer, Transaction](#saldo-topup-withdraw-transfer-transaction)
  - [Email Service](#email-service)
- [Database & Messaging](#database--messaging)
  - [PostgreSQL](#postgresql)
  - [Redis](#redis)
  - [Kafka & Zookeeper](#kafka--zookeeper)
- [Monitoring & Observability](#monitoring--observability)
  - [Prometheus](#prometheus)
  - [Grafana](#grafana)
  - [Node Exporter](#node-exporter)
  - [Kafka Exporter](#kafka-exporter)
  - [OpenTelemetry Collector](#opentelemetry-collector)
  - [Jaeger](#jaeger)
- [Networking & Load Balancing](#networking--load-balancing)
  - [Nginx](#nginx)
- [Jobs](#jobs)
  - [Migrations](#migrations)
- [Persistent Volumes](#persistent-volumes)
- [Environment Files](#environment-files)
- [How to Use](#how-to-use)

---

## Namespaces

- **namespace.yaml**: Declares the Kubernetes namespace for resource isolation.

## ConfigMaps & Secrets

- **configmaps.yaml**: General configuration for services.
- **nginx-configmap.yaml**: Nginx-specific configuration.
- **otel-collector-configmap.yaml**: OpenTelemetry Collector configuration.
- **promtheus-configmap.yaml**: Prometheus configuration.
- **secret.yaml**: Sensitive environment variables and credentials.

## Core Services

### API Gateway

- **apigateway-deployment.yaml**
- **apigateway-service.yaml**

Manages external traffic routing and acts as the main entry point to the system.

### Authentication

- **auth-deployment.yaml**
- **auth-service.yaml**

Handles authentication and authorization.

### User Management

- **user-deployment.yaml**
- **user-service.yaml**

Handles user-related operations.

### Role Management

- **role-deployment.yaml**
- **role-service.yaml**

Manages user roles and permissions.

### Card Management

- **card-deployment.yaml**
- **card-service.yaml**

Handles card-related operations.

### Merchant Management

- **merchant-deployment.yaml**
- **merchant-service.yaml**

Handles merchant-related operations.

### Saldo, TopUp, Withdraw, Transfer, Transaction

- **saldo-deployment.yaml**, **saldo-service.yaml**
- **topup-deployment.yaml**, **topup-service.yaml**
- **withdraw-deployment.yaml**, **withdraw-service.yaml**
- **transfer-deployment.yaml**, **transfer-service.yaml**
- **transaction-deployment.yaml**, **transaction-service.yaml**

All monolitic for financial operations.

### Email Service

- **email-deployment.yaml**
- **email-service.yaml**

Handles email notifications and communications.

## Database & Messaging

### PostgreSQL

- **postgres-deployment.yaml**
- **postgres-pvc.yaml**
- **postgres-service.yaml**

Primary database for persistent data storage.

### Redis

- **redis-deployment.yaml**
- **redis-pvc.yaml**
- **redis-service.yaml**

In-memory data store for caching and fast operations.

### Kafka & Zookeeper

- **kafka-deployment.yaml**
- **kafka-pvc.yaml**
- **kafka-service.yaml**
- **zookeeper-deployment.yaml**
- **zookeeper.-pvc.yaml**
- **zookeeper-service.yaml**

Kafka for messaging and event streaming. Zookeeper is used for Kafka coordination.

## Monitoring & Observability

### Prometheus

- **prometheus-deployment.yaml**
- **prometheus.yaml**
- **promtheus-configmap.yaml**

Monitoring and alerting toolkit.

### Grafana

- **grafana-deployment.yaml**
- **grafana-service.yaml**

Visualizes metrics and data from Prometheus.

### Node Exporter

- **node-exporter-deployment.yaml**
- **node-exporter-service.yaml**

Exports hardware and OS metrics.

### Kafka Exporter

- **kafka-exporter-deployment.yaml**
- **kafka-exporter-service.yaml**

Exports Kafka metrics for Prometheus.

### OpenTelemetry Collector

- **otel-collector-deployment.yaml**
- **otel-collector-service.yaml**
- **otel-collector-configmap.yaml**

Collects and forwards traces and metrics.

### Jaeger

- **jaeger-deployment.yaml**
- **jaeger-service.yaml**

Distributed tracing.

## Networking & Load Balancing

### Nginx

- **nginx-deployment.yaml**
- **nginx-service.yaml**
- **nginx-configmap.yaml**

Reverse proxy and load balancer.

## Jobs

### Migrations

- **migrate-job.yaml**

One-off job for running database migrations.

## Persistent Volumes

- **kafka-pvc.yaml**
- **postgres-pvc.yaml**
- **redis-pvc.yaml**
- **zookeeper.-pvc.yaml**

Defines persistent volume claims for stateful components.

---

## 🛠️ How to Use

You can deploy the entire stack locally using **Minikube** with the Docker driver.

### 📦 Prerequisites

Make sure you have the following installed:

- [Minikube](https://minikube.sigs.k8s.io/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- Docker (for local container runtime)

---

### ⚡ Quick Start (Recommended)

```sh
minikube start --driver=docker

make go-mod-tidy

make build-image

make image-load

kubectl apply -f namespace.yaml

kubectl apply -f .

minikube tunnel

kubectl get pods -n payment-gateway

kubectl get svc -n payment-gateway

kubectl get pvc -n payment-gateway
```
---

