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

# Find All Cards
GET {{baseUrl}}/api/card-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Card
POST {{baseUrl}}/api/card-command/create
Authorization: Bearer {{accessToken}}
{
    "user_id": {{userId}},
    "card_type": "debit",
    "expire_date": "2030-12-31T23:59:59Z",
    "cvv": "123",
    "card_provider": "visa"
}
HTTP 200
[Captures]
cardId: jsonpath "$.data.id"
cardNumber: jsonpath "$.data.card_number"

# Find Card by ID
GET {{baseUrl}}/api/card-query/{{cardId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Card by Card Number
GET {{baseUrl}}/api/card-query/card_number/{{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Card
POST {{baseUrl}}/api/card-command/update/{{cardId}}
Authorization: Bearer {{accessToken}}
{
    "card_id": {{cardId}},
    "user_id": {{userId}},
    "card_type": "credit",
    "expire_date": "2031-12-31T23:59:59Z",
    "cvv": "321",
    "card_provider": "mastercard"
}
HTTP 200

# Trash Card
POST {{baseUrl}}/api/card-command/trashed/{{cardId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Card
POST {{baseUrl}}/api/card-command/restore/{{cardId}}
Authorization: Bearer {{accessToken}}
HTTP 200
