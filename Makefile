COMPOSE_FILE=deployments/local/docker-compose.yml
SERVICES := apigateway migrate auth user role card merchant saldo topup transaction transfer withdraw email
ENGINE := $(shell command -v docker >/dev/null 2>&1 && echo docker || echo podman)
DOCKER_COMPOSE := $(shell command -v docker >/dev/null 2>&1 && echo "docker compose" || echo "podman-compose")
PROTO_DIR=proto
OUTDIR_PROTO=pb

migrate:
	go run service/migrate/main.go up

migrate-down:
	go run service/migrate/main.go down


generate-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUTDIR_PROTO) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTDIR_PROTO) --go-grpc_opt=paths=source_relative \
		$$(find $(PROTO_DIR) -name "*.proto")


generate-sql:
	GOTOOLCHAIN=local go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0 generate


generate-swagger:
	swag init -g service/apigateway/cmd/main.go -o service/apigateway/docs

seeder:
	go run service/seeder/main.go


build-image:
	@for service in $(SERVICES); do \
		echo "🔨 Building $$service-service..."; \
		$(ENGINE) build -t $$service-service:1.1 -f service/$$service/Dockerfile . || exit 1; \
	done
	@echo "✅ All services built successfully."

image-load:
	@for service in $(SERVICES); do \
		echo "🚚 Loading $$service-service..."; \
		minikube image load $$service-service:1.1 || exit 1; \
	done
	@echo "✅ All services loaded successfully."


image-delete:
	@for service in $(SERVICES); do \
		echo "🗑️ Deleting $$service-service image..."; \
		minikube image rm $$service-service:1.1 || echo "⚠️ Failed to delete $$service-service (maybe not found)"; \
	done
	@echo "✅ All requested images deleted (if they existed)."


ps:
	${DOCKER_COMPOSE} -f $(COMPOSE_FILE) ps

up:
	${DOCKER_COMPOSE} -f $(COMPOSE_FILE) up -d

# Start optional observability profile and enable application exporters.
up-observability:
	OTEL_ENABLED=true PYROSCOPE_ENABLED=true ${DOCKER_COMPOSE} -f $(COMPOSE_FILE) --profile observability up -d

down:
	${DOCKER_COMPOSE} -f $(COMPOSE_FILE) down

build-up:
	make build-image && make up

kube-start:
	minikube start --driver=docker

kube-up:
	kubectl apply -f deployments/kubernetes/namespace.yaml
	kubectl apply -f deployments/kubernetes

kube-down:
	kubectl delete -f deployments/kubernetes --ignore-not-found
	kubectl delete -f deployments/kubernetes/namespace.yaml --ignore-not-found

kube-status:
	@echo "🔍 Checking Pods in payment-gateway..."
	@kubectl get pods -n payment-gateway

	@echo "\n🔍 Checking Services in payment-gateway..."
	@kubectl get svc -n payment-gateway

	@echo "\n🔍 Checking PVCs in payment-gateway..."
	@kubectl get pvc -n payment-gateway

	@echo "\n🔍 Checking Jobs in payment-gateway..."
	@kubectl get jobs -n payment-gateway

kube-tunnel:
	minikube tunnel


test-auth:
	@cd tests && APP_ENV=development TESTCONTAINERS_RYUK_DISABLED=$$(command -v podman >/dev/null 2>&1 && echo true || echo false) go test ./auth/... -v

p2-read-baseline:
	BASE_URL=$${BASE_URL:-http://localhost:5000} \
	K6_TOKEN=$${K6_TOKEN:-} \
	K6_ENDPOINTS=$${K6_ENDPOINTS:-/health} \
	K6_VUS=$${K6_VUS:-1} \
	K6_RAMP_UP=$${K6_RAMP_UP:-5s} \
	K6_HOLD=$${K6_HOLD:-20s} \
	K6_RAMP_DOWN=$${K6_RAMP_DOWN:-5s} \
	K6_P95_MS=$${K6_P95_MS:-1000} \
	k6 run k6/p2_read_baseline.js

# Run configurable CRUD lifecycle k6 suite. Financial records are read/update-only unless explicitly enabled.
k6-crud-lifecycle:
	BASE_URL=$${BASE_URL:-http://localhost:5000} \
	K6_TOKEN=$${K6_TOKEN:-} \
	K6_API_KEY=$${K6_API_KEY:-} \
	K6_USER_ID=$${K6_USER_ID:-1} \
	K6_MERCHANT_ID=$${K6_MERCHANT_ID:-0} \
	K6_CARD_NUMBER=$${K6_CARD_NUMBER:-} \
	K6_RECEIVER_CARD_NUMBER=$${K6_RECEIVER_CARD_NUMBER:-} \
	K6_SALDO_ID=$${K6_SALDO_ID:-0} \
	K6_DOMAINS=$${K6_DOMAINS:-user,role,merchant,card,merchant-document} \
	K6_ALLOW_FINANCIAL_DELETE=$${K6_ALLOW_FINANCIAL_DELETE:-false} \
	K6_VUS=$${K6_VUS:-1} K6_RAMP_UP=$${K6_RAMP_UP:-5s} K6_HOLD=$${K6_HOLD:-20s} K6_RAMP_DOWN=$${K6_RAMP_DOWN:-5s} \
	k6 run k6/crud_lifecycle.js

# Run financial idempotency replay/concurrency for one selected domain.
k6-financial-concurrency:
	BASE_URL=$${BASE_URL:-http://localhost:5000} \
	K6_TOKEN=$${K6_TOKEN:-} \
	K6_API_KEY=$${K6_API_KEY:-} \
	K6_CARD_NUMBER=$${K6_CARD_NUMBER:-} \
	K6_RECEIVER_CARD_NUMBER=$${K6_RECEIVER_CARD_NUMBER:-} \
	K6_MERCHANT_ID=$${K6_MERCHANT_ID:-0} \
	K6_FINANCIAL_DOMAIN=$${K6_FINANCIAL_DOMAIN:-topup} \
	K6_VUS=$${K6_VUS:-1} K6_RAMP_UP=$${K6_RAMP_UP:-5s} K6_HOLD=$${K6_HOLD:-20s} K6_RAMP_DOWN=$${K6_RAMP_DOWN:-5s} \
	k6 run k6/financial/financial_concurrency.js
