# Login
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find all roles
GET {{baseUrl}}/api/role-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Find active roles
GET {{baseUrl}}/api/role-query/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Find trashed roles
GET {{baseUrl}}/api/role-query/trashed
Authorization: Bearer {{accessToken}}
HTTP 200

# Find role by user ID
GET {{baseUrl}}/api/role-query/user/{{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200
