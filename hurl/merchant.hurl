# Merchant Service API Tests
# Base URL: http://localhost:8080

### Create a new merchant
POST http://localhost:8080/api/merchant
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "name": "Test Store",
  "description": "A test store for payments",
  "email": "merchant@teststore.com",
  "phone": "+1234567890",
  "address": "123 Test Street, Test City",
  "category": "retail"
}

HTTP 200

### Get all merchants
GET http://localhost:8080/api/merchant
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant by ID
GET http://localhost:8080/api/merchant/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update merchant by ID
PUT http://localhost:8080/api/merchant/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "name": "Updated Store Name",
  "description": "Updated description",
  "phone": "+1234567891"
}

HTTP 200

### Delete merchant by ID
DELETE http://localhost:8080/api/merchant/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all merchant amounts
GET http://localhost:8080/api/merchant/amount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant amount by ID
GET http://localhost:8080/api/merchant/amount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all merchant methods
GET http://localhost:8080/api/merchant/method
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant method by ID
GET http://localhost:8080/api/merchant/method/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all merchant queries
GET http://localhost:8080/api/merchant/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant query by ID
GET http://localhost:8080/api/merchant/query/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all merchant total amounts
GET http://localhost:8080/api/merchant/totalamount
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant total amount by ID
GET http://localhost:8080/api/merchant/totalamount/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all merchant transactions
GET http://localhost:8080/api/merchant/transaction
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant transaction by ID
GET http://localhost:8080/api/merchant/transaction/123
Authorization: Bearer {{auth_token}}

HTTP 200