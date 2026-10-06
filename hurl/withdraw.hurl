# Withdraw Service API Tests
# Base URL: http://localhost:8080

### Create a new withdrawal
POST http://localhost:8080/api/withdraw
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "user_id": "user123",
  "amount": 300.00,
  "currency": "USD",
  "withdrawal_method": "bank_transfer",
  "bank_account": "1234567890",
  "bank_name": "Test Bank",
  "reference": "WDR789012",
  "description": "Withdrawal to bank account"
}

HTTP 200

### Get all withdrawals
GET http://localhost:8080/api/withdraw
Authorization: Bearer {{auth_token}}

HTTP 200

### Get withdrawal by ID
GET http://localhost:8080/api/withdraw/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update withdrawal by ID
PUT http://localhost:8080/api/withdraw/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "status": "completed",
  "description": "Updated withdrawal description"
}

HTTP 200

### Delete withdrawal by ID
DELETE http://localhost:8080/api/withdraw/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all withdrawal amounts
GET http://localhost:8080/api/withdraw/amount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get withdrawal amount by ID
GET http://localhost:8080/api/withdraw/amount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all withdrawal queries
GET http://localhost:8080/api/withdraw/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get withdrawal query by ID
GET http://localhost:8080/api/withdraw/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all withdrawal statuses
GET http://localhost:8080/api/withdraw/status
Authorization: Bearer {{auth_token}}

HTTP 200

### Get withdrawal status by ID
GET http://localhost:8080/api/withdraw/status/123
Authorization: Bearer {{auth_token}}

HTTP 200