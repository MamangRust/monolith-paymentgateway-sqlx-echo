set shell := ["bash", "-c"]

COMPOSE_FILE := "deployments/local/docker-compose.yml"
SERVICES := "apigateway migrate auth user role card merchant saldo topup transaction transfer withdraw email"
BASE_URL := env_var_or_default("BASE_URL", "http://localhost:5000")
K6_TOKEN := env_var_or_default("K6_TOKEN", "")
K6_ENDPOINTS := env_var_or_default("K6_ENDPOINTS", "/health")
K6_VUS := env_var_or_default("K6_VUS", "1")
K6_RAMP_UP := env_var_or_default("K6_RAMP_UP", "5s")
K6_HOLD := env_var_or_default("K6_HOLD", "20s")
K6_RAMP_DOWN := env_var_or_default("K6_RAMP_DOWN", "5s")
K6_P95_MS := env_var_or_default("K6_P95_MS", "1000")
ENGINE := `(command -v docker >/dev/null 2>&1 && echo docker) || echo podman`
DOCKER_COMPOSE := `(command -v docker >/dev/null 2>&1 && echo "docker compose") || echo "podman-compose"`
# Testcontainers Ryuk needs privileged socket access; rootless Podman can't run it.
RYUK_DISABLED := `(command -v docker >/dev/null 2>&1 && echo "") || echo "TESTCONTAINERS_RYUK_DISABLED=true"`
PROTO_DIR := "proto"
OUTDIR_PROTO := "pb"

# List all recipes
default:
    @just --list

# Run database migrations up
migrate:
    go run service/migrate/cmd/main.go -dir service/migrate/migrations up

# Run database migrations down
migrate-down:
    go run service/migrate/cmd/main.go -dir service/migrate/migrations down

# Generate protocol buffers
generate-proto:
    protoc \
        --proto_path={{PROTO_DIR}} \
        --go_out={{OUTDIR_PROTO}} --go_opt=paths=source_relative \
        --go-grpc_out={{OUTDIR_PROTO}} --go-grpc_opt=paths=source_relative \
        $(find {{PROTO_DIR}} -name "*.proto")

# Generate swagger documentation
generate-swagger:
    swag init -g service/apigateway/cmd/main.go -o service/apigateway/docs

# Run seeder
seeder:
    go run service/seeder/main.go

# Build images for all services (docker or podman)
build-image:
    @for service in {{SERVICES}}; do \
        echo "🔨 Building $service-service..."; \
        {{ENGINE}} build -t $service-service:1.1 -f service/$service/Dockerfile . || exit 1; \
    done
    @echo "✅ All services built successfully."

# Load images to minikube
image-load:
    @for service in {{SERVICES}}; do \
        echo "🚚 Loading $service-service..."; \
        minikube image load $service-service:1.1 || exit 1; \
    done
    @echo "✅ All services loaded successfully."

# Delete images from minikube
image-delete:
    @for service in {{SERVICES}}; do \
        echo "🗑️ Deleting $service-service image..."; \
        minikube image rm $service-service:1.1 || echo "⚠️ Failed to delete $service-service (maybe not found)"; \
    done
    @echo "✅ All requested images deleted (if they existed)."

# Show docker compose process status
ps:
    {{DOCKER_COMPOSE}} -f {{COMPOSE_FILE}} ps

# Start docker compose services
up:
    {{DOCKER_COMPOSE}} -f {{COMPOSE_FILE}} up -d

# Start optional observability profile and enable application exporters.
up-observability:
    OTEL_ENABLED=true PYROSCOPE_ENABLED=true {{DOCKER_COMPOSE}} -f {{COMPOSE_FILE}} --profile observability up -d

# Stop docker compose services
down:
    {{DOCKER_COMPOSE}} -f {{COMPOSE_FILE}} down

# Build images and start docker compose
build-up: build-image up

# Start minikube with docker driver
kube-start:
    minikube start --driver=docker

# Apply kubernetes manifests
kube-up:
    kubectl apply -f deployments/kubernetes/namespace.yaml
    kubectl apply -f deployments/kubernetes

# Delete kubernetes manifests
kube-down:
    kubectl delete -f deployments/kubernetes --ignore-not-found
    kubectl delete -f deployments/kubernetes/namespace.yaml --ignore-not-found

# Show kubernetes status
kube-status:
    @echo "🔍 Checking Pods in payment-gateway..."
    @kubectl get pods -n payment-gateway
    @echo -e "\n🔍 Checking Services in payment-gateway..."
    @kubectl get svc -n payment-gateway
    @echo -e "\n🔍 Checking PVCs in payment-gateway..."
    @kubectl get pvc -n payment-gateway
    @echo -e "\n🔍 Checking Jobs in payment-gateway..."
    @kubectl get jobs -n payment-gateway

# Tunnel minikube services
kube-tunnel:
    minikube tunnel

# Generate mocks for all services
generate-mocks:
    @echo "🔧 Generating mocks..."
    @for svc in auth card merchant saldo topup transaction transfer user role withdraw; do \
        if [ -d "service/$$svc" ]; then \
            (cd service/$$svc && go generate ./...); \
        fi \
    done
    @cd pkg && go generate ./...
    @echo "✅ Mocks generated."

# Run unit tests in tests/ (fast mock-based tests)
test-unit:
    @echo "🧪 Running unit tests (mock-based)..."
    @cd tests && go test -race -count=1 -short -coverprofile=coverage.out ./...

# Run unit tests in pkg/
test-pkg:
    @echo "🧪 Running unit tests in pkg/..."
    @cd pkg && go test -race -count=1 -coverprofile=coverage.out ./...

# Run integration tests in tests/ (testcontainers; auto-detects Docker/Podman)
test-integration:
    @echo "🧪 Running integration tests (testcontainers; Docker/Podman auto-detected)..."
    @cd tests && APP_ENV=development {{RYUK_DISABLED}} go test -count=1 -v \
        -run "Test.*Suite" \
        ./auth/... \
        ./card/... \
        ./merchant/... \
        ./saldo/... \
        ./topup/... \
        ./transaction/... \
        ./transfer/... \
        ./user/... \
        ./role/... \
        ./withdraw/...

# Run all tests (unit + integration)
test-all: test-unit test-pkg test-integration

# Run a safe, read-only P2 performance baseline.
# Example: just p2-read-baseline K6_TOKEN=... K6_ENDPOINTS=/health
p2-read-baseline:
    BASE_URL={{BASE_URL}} K6_TOKEN={{K6_TOKEN}} K6_ENDPOINTS={{K6_ENDPOINTS}} K6_VUS={{K6_VUS}} \
        K6_RAMP_UP={{K6_RAMP_UP}} K6_HOLD={{K6_HOLD}} K6_RAMP_DOWN={{K6_RAMP_DOWN}} \
        K6_P95_MS={{K6_P95_MS}} k6 run k6/p2_read_baseline.js

# Run configurable CRUD lifecycle k6 suite. Financial deletes stay disabled by default.
k6-crud-lifecycle:
    BASE_URL={{BASE_URL}} K6_TOKEN={{K6_TOKEN}} \
        K6_API_KEY={{env_var_or_default("K6_API_KEY", "")}} \
        K6_USER_ID={{env_var_or_default("K6_USER_ID", "1")}} \
        K6_MERCHANT_ID={{env_var_or_default("K6_MERCHANT_ID", "0")}} \
        K6_CARD_NUMBER={{env_var_or_default("K6_CARD_NUMBER", "")}} \
        K6_RECEIVER_CARD_NUMBER={{env_var_or_default("K6_RECEIVER_CARD_NUMBER", "")}} \
        K6_SALDO_ID={{env_var_or_default("K6_SALDO_ID", "0")}} \
        K6_DOMAINS={{env_var_or_default("K6_DOMAINS", "user,role,merchant,card,merchant-document")}} \
        K6_ALLOW_FINANCIAL_DELETE={{env_var_or_default("K6_ALLOW_FINANCIAL_DELETE", "false")}} \
        K6_VUS={{K6_VUS}} K6_RAMP_UP={{K6_RAMP_UP}} K6_HOLD={{K6_HOLD}} K6_RAMP_DOWN={{K6_RAMP_DOWN}} \
        k6 run k6/crud_lifecycle.js

# Run financial idempotency replay/concurrency for one selected domain.
k6-financial-concurrency:
    BASE_URL={{BASE_URL}} K6_TOKEN={{K6_TOKEN}} \
        K6_API_KEY={{env_var_or_default("K6_API_KEY", "")}} \
        K6_CARD_NUMBER={{env_var_or_default("K6_CARD_NUMBER", "")}} \
        K6_RECEIVER_CARD_NUMBER={{env_var_or_default("K6_RECEIVER_CARD_NUMBER", "")}} \
        K6_MERCHANT_ID={{env_var_or_default("K6_MERCHANT_ID", "0")}} \
        K6_FINANCIAL_DOMAIN={{env_var_or_default("K6_FINANCIAL_DOMAIN", "topup")}} \
        K6_VUS={{K6_VUS}} K6_RAMP_UP={{K6_RAMP_UP}} K6_HOLD={{K6_HOLD}} K6_RAMP_DOWN={{K6_RAMP_DOWN}} \
        k6 run k6/financial/financial_concurrency.js

# Run auth service integration tests
test-auth:
    @echo "🧪 Running auth integration tests..."
    @cd tests && APP_ENV=development {{RYUK_DISABLED}} go test ./auth/... -v

# Run just the unit tests with race detection (fast)
test-ci:
    @echo "🧪 CI test suite (mock-based only)..."
    @cd tests && go test -race -count=1 -short -coverprofile=coverage.out ./...
    @cd pkg && go test -race -count=1 -coverprofile=coverage.out ./...
    @echo "✅ All CI tests passed."

# Build all Go service binaries (from api-gateway to withdraw)
build:
    @mkdir -p bin
    @for mod in service/*/go.mod; do \
        dir=$(dirname $mod); \
        service=$(basename $dir); \
        echo "🔨 Building $service..."; \
        (cd $dir && go build -o ../../bin/$service ./cmd/main.go) || exit 1; \
    done
    @echo "✅ All services built successfully in bin/ folder."

# Run go mod tidy for all services
tidy-all:
    @echo "🧹 Tidying all service modules..."
    @for service in {{SERVICES}}; do \
        if [ -d "service/$service" ]; then \
            echo "📦 Tidying $service..."; \
            (cd service/$service && go mod tidy) || echo "⚠️ Failed to tidy $service"; \
        fi \
    done
    @echo "✅ All services tidied."
