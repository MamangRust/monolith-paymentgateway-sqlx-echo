# Saldo Service API Tests
# Base URL: http://localhost:8080

### Create a new saldo record
POST http://localhost:8080/api/saldo
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "user_id": "user123",
  "amount": 1000.00,
  "currency": "USD",
  "type": "credit"
}

HTTP 200

### Get all saldo records
GET http://localhost:8080/api/saldo
Authorization: Bearer {{auth_token}}

HTTP 200

### Get saldo by ID
GET http://localhost:8080/api/saldo/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update saldo by ID
PUT http://localhost:8080/api/saldo/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "amount": 1500.00,
  "currency": "USD"
}

HTTP 200

### Delete saldo by ID
DELETE http://localhost:8080/api/saldo/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all saldo balances
GET http://localhost:8080/api/saldo/balance
Authorization: Bearer {{auth_token}}

HTTP 200

### Get saldo balance by ID
GET http://localhost:8080/api/saldo/balance/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all saldo queries
GET http://localhost:8080/api/saldo/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get saldo query by ID
GET http://localhost:8080/api/saldo/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all saldo totals
GET http://localhost:8080/api/saldo/total
Authorization: Bearer {{auth_token}}

HTTP 200

### Get saldo total by ID
GET http://localhost:8080/api/saldo/total/123
Authorization: Bearer {{auth_token}}

HTTP 200