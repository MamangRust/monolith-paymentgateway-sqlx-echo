# Payment Gateway API Gateway

## Overview

The **API Gateway** serves as the main entry point for all client requests in the Payment Gateway distributed system. By centralizing access, it simplifies how clients interact with various backend microservices such as `AuthService`, `UserService`, `MerchantService`, `TransactionService`, and others (e.g., Topup, Transfer, Withdraw).

This gateway is typically implemented using a high-performance HTTP framework (such as [Echo](https://echo.labstack.com/)), and it communicates with backend services using gRPC. This architecture enables scalable, secure, and maintainable separation of concerns across the entire system.

---

## ✨ Key Features

- **Single Entry Point**
  All external requests from clients (web/mobile) go through the API Gateway, which acts as a reverse proxy.

- **Request Routing**
  The gateway forwards HTTP requests to the appropriate backend microservice using gRPC.

- **Security**
  Handles authentication and authorization, including JWT token validation, to enforce access control at the edge.

- **Data Transformation**
  Adjusts the request and response formats as needed between clients and backend microservices (e.g., JSON ↔ Protobuf).

- **Centralized Logging and Monitoring**
  All traffic passes through the gateway, making it a suitable place for logging, metrics, and tracing integration.

---

## 🔄 Request Flow Through the API Gateway

1. **Client Sends a Request**
   The client (web or mobile app) sends an HTTP request to the API Gateway endpoint.

2. **Validation & Security**
   The API Gateway validates the request, checking authentication (such as JWTs) and authorization policies.

3. **Routing to Monolith**
   After validation, the gateway identifies the correct microservice (e.g., `AuthService`, `UserService`, `MerchantService`, `TransactionService`, etc.) and forwards the request via gRPC.

4. **Monolith Processing**
   The microservice processes the business logic and returns a response to the API Gateway.

5. **Response to Client**
   The API Gateway may transform the response format and sends it back to the client.

---

## 🛠️ Monoliths Behind the API Gateway

- **Auth Service**
  Handles user authentication, JWT issuance, and verification.

- **User Service**
  Manages user profiles and related operations.

- **Merchant Service**
  Handles merchant accounts and merchant-specific logic.

- **Transaction Service**
  Manages financial transactions, including payments, refunds, etc.

- **Topup, Transfer, Withdraw Services**
  Specialized services for account top-ups, fund transfers, and withdrawals.

---
