# ArgoCD Applications

This directory intentionally contains no ArgoCD Application manifests.

The ArgoCD app-of-apps is a **single application**: `payment-gateway-production`
(defined in `../production/production.yaml`), which points to the flat
`deployments/kubernetes` directory of the Go monolith
(`MamangRust/monolith-payment-gateway-grpc`). That directory contains plain
manifests for every service and infrastructure component of the
`payment-gateway` namespace.

`deployments/kubernetes` is flat on purpose — no base/overlays or
`GHCR_OWNER` templating is required. Image tags are pinned directly in each
manifest (e.g. `transaction-service:1.1`); CI replaces them with registry
qualified images when needed. Adding a new service requires no new ArgoCD
Application — it is picked up automatically when the manifest is added to
`deployments/kubernetes`.

The root app (`../root-app.yaml`) bootstraps this production application and
lives in the `payment-gateway` ArgoCD AppProject.
