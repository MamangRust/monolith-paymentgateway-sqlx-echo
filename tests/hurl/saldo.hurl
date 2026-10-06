# Login
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"
refreshToken: jsonpath "$.data.refresh_token"

# Find All Saldos
GET {{baseUrl}}/api/saldo-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Saldo
POST {{baseUrl}}/api/saldo-command/create
Authorization: Bearer {{accessToken}}
{
    "card_number": "{{cardNumber}}",
    "total_balance": 1000000
}
HTTP 200
[Captures]
saldoId: jsonpath "$.data.id"

# Find Saldo by ID
GET {{baseUrl}}/api/saldo-query/{{saldoId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Saldo by Card Number
GET {{baseUrl}}/api/saldo-query/card_number/{{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Saldo
POST {{baseUrl}}/api/saldo-command/update/{{saldoId}}
Authorization: Bearer {{accessToken}}
{
    "saldo_id": {{saldoId}},
    "card_number": "{{cardNumber}}",
    "total_balance": 2000000
}
HTTP 200
