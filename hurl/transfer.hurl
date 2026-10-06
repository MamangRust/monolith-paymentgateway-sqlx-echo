# Transfer Service API Tests
# Base URL: http://localhost:8080

### Create a new transfer
POST http://localhost:8080/api/transfer
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "from_user_id": "user123",
  "to_user_id": "user456",
  "amount": 250.00,
  "currency": "USD",
  "reference": "TRF789012",
  "description": "Transfer to friend"
}

HTTP 200

### Get all transfers
GET http://localhost:8080/api/transfer
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transfer by ID
GET http://localhost:8080/api/transfer/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update transfer by ID
PUT http://localhost:8080/api/transfer/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "status": "completed",
  "description": "Updated transfer description"
}

HTTP 200

### Delete transfer by ID
DELETE http://localhost:8080/api/transfer/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transfer amounts
GET http://localhost:8080/api/transfer/amount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transfer amount by ID
GET http://localhost:8080/api/transfer/amount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transfer queries
GET http://localhost:8080/api/transfer/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transfer query by ID
GET http://localhost:8080/api/transfer/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all transfer statuses
GET http://localhost:8080/api/transfer/status
Authorization: Bearer {{auth_token}}

HTTP 200

### Get transfer status by ID
GET http://localhost:8080/api/transfer/status/123
Authorization: Bearer {{auth_token}}

HTTP 200