# Auth Service

The **AuthService** is a microservice specializing in user authentication and account management in a distributed payment gateway system. It orchestrates user registration, login, password reset, identity verification, and token management, exposing these functionalities via gRPC and integrating with the broader system using Kafka and observability tooling.

---

## ✨ Features

- **User Registration & Account Creation**
- **Login & JWT Token Issuance**
- **Password Reset (Forgot Password, OTP, Verification)**
- **Identity (Me) and Token Refresh**
- **Kafka Event Publishing for Email Notifications**
- **Prometheus Metrics for Monitoring**
- **OpenTelemetry Distributed Tracing**
- **Structured Logging with Zap**

---

## 🏗️ Service Architecture

The AuthService is modularized into multiple interfaces for maintainability:

- **RegistrationService:** User registration and account creation.
- **LoginService:** User authentication and JWT token issuance.
- **PasswordResetService:** Password reset workflow (OTP sending/verification, password update).
- **IdentifyService:** Token refresh, and retrieving information about the current user session.

**Integrations:**

- **Kafka:** Publishes email-related events for the email-service.
- **Prometheus:** Provides metrics for observability.
- **OpenTelemetry:** Distributed tracing for debugging and performance.
- **Zap:** Structured logging for operational and error logs.

---

## 📚 Available RPC Methods

| Method            | Description                                                  |
|-------------------|-------------------------------------------------------------|
| `RegisterUser`    | Registers a new user into the system                        |
| `LoginUser`       | Authenticates a user and issues a JWT token                 |
| `RefreshToken`    | Refreshes access token                                      |
| `GetMe`           | Returns info about the currently authenticated user         |
| `ForgotPassword`  | Sends OTP (via Kafka/email) for password reset              |
| `VerifyCode`      | Verifies the OTP code sent to the user's email              |
| `ResetPassword`   | Resets password after successful code verification          |


---

## 📦 Kafka Integration

The AuthService publishes events to Kafka topics, which are then consumed by the `email-service` to notify users.

| Kafka Topic                                | Event Triggered                   | Purpose                                              | Consumer Group         |
|--------------------------------------------|-----------------------------------|------------------------------------------------------|------------------------|
| `email-service-topic-auth-register`        | User registration                 | Send welcome/confirmation email                      | `email-service-group`  |
| `email-service-topic-auth-forgot-password` | Forgot password initiated         | Send OTP to user's email                             | `email-service-group`  |
| `email-service-topic-auth-verify-code-success` | Successful OTP verification    | Notify user of successful verification               | `email-service-group`  |

---

## 📊 Observability

### Prometheus Metrics

- Request counters per RPC method and status
- Request duration histograms per RPC method

**Example metrics:**

- `login_service_requests_total`
- `login_service_request_duration_seconds`
- `password_reset_service_requests_total`
- `password_reset_service_request_duration_seconds`
- `register_service_requests_total`
- `register_service_request_duration_seconds`

### Distributed Tracing

- Integrated with OpenTelemetry
- Each service interface has its own tracer:
    - `login-service`
    - `register-service`
    - `password-reset-service`
    - `identity-service`

---

## 📂 Logging

- All actions and errors are logged using Zap for structured, high-performance logging.
- Includes trace/span IDs for end-to-end debugging.

---

## 🛡️ Security

- All authentication and token-related endpoints are secured.
- JWT tokens are validated for all protected RPC methods.

---

## 📝 Example Usage

- **Register:**
  `RegisterUser` → triggers Kafka event to `email-service-topic-auth-register`
- **Forgot Password:**
  `ForgotPassword` → triggers Kafka event to `email-service-topic-auth-forgot-password`
- **Verify Code:**
  `VerifyCode` → triggers Kafka event to `email-service-topic-auth-verify-code-success`

---

## 🚀 Deployment

- Deploy as a standalone service (container/pod) in the `payment-gateway` namespace.
- Requires access to Kafka, Prometheus, and OpenTelemetry collector endpoints.

---

> _The AuthService is the secure gateway for all user authentication and account management operations, designed for reliability, scalability, and seamless integration with distributed systems._