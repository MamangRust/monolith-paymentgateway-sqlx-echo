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

# Find All Users
GET {{baseUrl}}/api/user-query
Authorization: Bearer {{accessToken}}
HTTP 200

# Create User
POST {{baseUrl}}/api/user-command/create
Authorization: Bearer {{accessToken}}
{
    "firstname": "Alice",
    "lastname": "Smith",
    "email": "{{testEmail}}",
    "password": "password123",
    "confirm_password": "password123"
}
HTTP 200
[Captures]
userId: jsonpath "$.data.id"

# Find User by ID
GET {{baseUrl}}/api/user-query/{{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update User
POST {{baseUrl}}/api/user-command/update/{{userId}}
Authorization: Bearer {{accessToken}}
{
    "firstname": "Alicia",
    "lastname": "Smith",
    "email": "{{testEmail}}",
    "password": "password123456",
    "confirm_password": "password123456"
}
HTTP 200

# Find Active Users
GET {{baseUrl}}/api/user-query/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Trash User
POST {{baseUrl}}/api/user-command/trashed/{{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Trashed Users
GET {{baseUrl}}/api/user-query/trashed
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore User
POST {{baseUrl}}/api/user-command/restore/{{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Permanent Delete User
DELETE {{baseUrl}}/api/user-command/permanent/{{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200
