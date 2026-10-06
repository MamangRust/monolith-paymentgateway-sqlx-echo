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

# Find All Transactions
GET {{baseUrl}}/api/transaction-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Transaction
# Uses merchant 2 (FashionHub, user 2): its settlement card (5425233430109903)
# is guaranteed distinct from the payer card, satisfying the atomic SQL rule $1 <> $3.
# merchant 1 belongs to user 1 whose primary card can equal the payer card -> 400.
POST {{baseUrl}}/api/transaction-command/create
Authorization: Bearer {{accessToken}}
X-Api-Key: API_KEY_002_IJKLMNOP789012
{
    "card_number": "{{cardNumber}}",
    "amount": 50000,
    "payment_method": "visa",
    "merchant_id": 2,
    "transaction_time": "2026-04-09T19:41:10Z"
}
HTTP 200
[Captures]
transactionId: jsonpath "$.data.id"

# Find Transaction by ID
GET {{baseUrl}}/api/transaction-query/{{transactionId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Transactions by Card Number
GET {{baseUrl}}/api/transaction-query/card-number/{{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Transaction
POST {{baseUrl}}/api/transaction-command/update/{{transactionId}}
Authorization: Bearer {{accessToken}}
X-Api-Key: API_KEY_002_IJKLMNOP789012
{
    "card_number": "{{cardNumber}}",
    "amount": 100000,
    "payment_method": "visa",
    "merchant_id": 2,
    "transaction_time": "2026-04-09T20:41:10Z"
}
HTTP 200

# Trash Transaction
POST {{baseUrl}}/api/transaction-command/trashed/{{transactionId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Transaction
POST {{baseUrl}}/api/transaction-command/restore/{{transactionId}}
Authorization: Bearer {{accessToken}}
HTTP 200
