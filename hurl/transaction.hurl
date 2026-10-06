# Transaction Service API Tests
# Base URL: http://localhost:8080

### Create a new transaction
POST http://localhost:8080/api/transaction
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "user_id": "user123",
  "merchant_id": "merchant456",
  "amount": 100.50,
  "currency": "USD",
  "payment_method": "credit_card",
  "reference": "TXN789012",
  "description": "Payment for online purchase"
}

HTTP 200

### Get all transactions
GET http://localhost:8080/api/transaction
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transaction by ID
GET http://localhost:8080/api/transaction/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update transaction by ID
PUT http://localhost:8080/api/transaction/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "status": "completed",
  "description": "Updated transaction description"
}

HTTP 200

### Delete transaction by ID
DELETE http://localhost:8080/api/transaction/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transaction amounts
GET http://localhost:8080/api/transaction/amount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transaction amount by ID
GET http://localhost:8080/api/transaction/amount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transaction methods
GET http://localhost:8080/api/transaction/method
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transaction method by ID
GET http://localhost:8080/api/transaction/method/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transaction queries
GET http://localhost:8080/api/transaction/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transaction query by ID
GET http://localhost:8080/api/transaction/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transaction statuses
GET http://localhost:8080/api/transaction/status
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transaction status by ID
GET http://localhost:8080/api/transaction/status/123
Authorization: Bearer {{auth_token}}

HTTP 200