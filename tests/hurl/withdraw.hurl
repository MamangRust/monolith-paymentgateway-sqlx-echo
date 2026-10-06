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

# Find All Withdraws
GET {{baseUrl}}/api/withdraw-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Withdraw
POST {{baseUrl}}/api/withdraw-command/create
Authorization: Bearer {{accessToken}}
{
    "card_number": "{{cardNumber}}",
    "withdraw_amount": 50000,
    "withdraw_time": "2026-04-09T10:00:00Z"
}
HTTP 200
[Captures]
withdrawId: jsonpath "$.data.id"

# Find Withdraw by ID
GET {{baseUrl}}/api/withdraw-query/{{withdrawId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Withdraws by Card Number
GET {{baseUrl}}/api/withdraw-query/card-number/{{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Withdraw
POST {{baseUrl}}/api/withdraw-command/update/{{withdrawId}}
Authorization: Bearer {{accessToken}}
{
    "card_number": "{{cardNumber}}",
    "withdraw_amount": 100000,
    "withdraw_time": "2026-04-09T11:00:00Z"
}
HTTP 200

# Trash Withdraw
POST {{baseUrl}}/api/withdraw-command/trashed/{{withdrawId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Withdraw
POST {{baseUrl}}/api/withdraw-command/restore/{{withdrawId}}
Authorization: Bearer {{accessToken}}
HTTP 200
