# Login
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find all merchant documents
GET {{baseUrl}}/api/merchant-document-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Find active merchant documents
GET {{baseUrl}}/api/merchant-document-query/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Find trashed merchant documents
GET {{baseUrl}}/api/merchant-document-query/trashed
Authorization: Bearer {{accessToken}}
HTTP 200
