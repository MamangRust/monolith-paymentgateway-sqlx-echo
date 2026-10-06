# Topup Service API Tests
# Base URL: http://localhost:8080

### Create a new topup
POST http://localhost:8080/api/topup
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "user_id": "user123",
  "amount": 500.00,
  "currency": "USD",
  "payment_method": "credit_card",
  "reference": "TOPUP123456"
}

HTTP 200

### Get all topups
GET http://localhost:8080/api/topup
Authorization: Bearer {{auth_token}}

HTTP 200

### Get topup by ID
GET http://localhost:8080/api/topup/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update topup by ID
PUT http://localhost:8080/api/topup/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "amount": 600.00,
  "status": "completed"
}

HTTP 200

### Delete topup by ID
DELETE http://localhost:8080/api/topup/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all topup amounts
GET http://localhost:8080/api/topup/amount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get topup amount by ID
GET http://localhost:8080/api/topup/amount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all topup methods
GET http://localhost:8080/api/topup/method
Authorization: Bearer {{auth_token}}

HTTP 200

### Get topup method by ID
GET http://localhost:8080/api/topup/method/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all topup queries
GET http://localhost:8080/api/topup/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get topup query by ID
GET http://localhost:8080/api/topup/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all topup statuses
GET http://localhost:8080/api/topup/status
Authorization: Bearer {{auth_token}}

HTTP 200

### Get topup status by ID
GET http://localhost:8080/api/topup/status/123
Authorization: Bearer {{auth_token}}

HTTP 200