# 📦 Package `main`

**Source Path:** `service/apigateway/cmd`

## 🚀 Functions

### `main`

main starts the API Gateway service.

It sets up the gRPC clients for other microservices and starts the HTTP server.
When an interrupt signal is received, it gracefully shuts down the service.

```go
func main()
```

