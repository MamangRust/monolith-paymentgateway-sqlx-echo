# Card Service API Tests
# Base URL: http://localhost:8080

### Create a new card
POST http://localhost:8080/api/card
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "user_id": "user123",
  "card_number": "4111111111111111",
  "card_holder": "John Doe",
  "expiry_date": "12/25",
  "cvv": "123",
  "card_type": "credit"
}

HTTP 200

### Get all cards
GET http://localhost:8080/api/card
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card by ID
GET http://localhost:8080/api/card/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update card by ID
PUT http://localhost:8080/api/card/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "card_holder": "John Smith",
  "expiry_date": "12/26"
}

HTTP 200

### Delete card by ID
DELETE http://localhost:8080/api/card/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all card balances
GET http://localhost:8080/api/card/balance
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card balance by card ID
GET http://localhost:8080/api/card/balance/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card dashboard
GET http://localhost:8080/api/card/dashboard
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all card topups
GET http://localhost:8080/api/card/topup
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card topup by ID
GET http://localhost:8080/api/card/topup/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all card transactions
GET http://localhost:8080/api/card/transaction
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card transaction by ID
GET http://localhost:8080/api/card/transaction/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all card transfers
GET http://localhost:8080/api/card/transfer
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card transfer by ID
GET http://localhost:8080/api/card/transfer/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all card withdrawals
GET http://localhost:8080/api/card/withdraw
Authorization: Bearer {{auth_token}}

HTTP 200

### Get card withdrawal by ID
GET http://localhost:8080/api/card/withdraw/123
Authorization: Bearer {{auth_token}}

HTTP 200