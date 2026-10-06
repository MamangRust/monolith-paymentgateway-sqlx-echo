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

# Find All Topups
GET {{baseUrl}}/api/topup-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Topup
POST {{baseUrl}}/api/topup-command/create
Authorization: Bearer {{accessToken}}
{
    "card_number": "{{cardNumber}}",
    "topup_amount": 50000,
    "topup_method": "visa"
}
HTTP 200
[Captures]
topupId: jsonpath "$.data.id"

# Find Topup by ID
GET {{baseUrl}}/api/topup-query/{{topupId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Topups by Card Number
GET {{baseUrl}}/api/topup-query/card-number/{{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Topup
POST {{baseUrl}}/api/topup-command/update/{{topupId}}
Authorization: Bearer {{accessToken}}
{
    "card_number": "{{cardNumber}}",
    "topup_amount": 100000,
    "topup_method": "visa"
}
HTTP 200

# Trash Topup
POST {{baseUrl}}/api/topup-command/trashed/{{topupId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Topup
POST {{baseUrl}}/api/topup-command/restore/{{topupId}}
Authorization: Bearer {{accessToken}}
HTTP 200
