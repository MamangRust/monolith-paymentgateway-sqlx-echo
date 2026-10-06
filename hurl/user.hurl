# User Service API Tests
# Base URL: http://localhost:8080

### Create a new user
POST http://localhost:8080/api/user
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "first_name": "Jane",
  "last_name": "Smith",
  "email": "jane.smith@example.com",
  "phone": "+1234567890",
  "address": "456 User Street, User City",
  "date_of_birth": "1990-01-15",
  "status": "active"
}

HTTP 200

### Get all users
GET http://localhost:8080/api/user
Authorization: Bearer {{auth_token}}

HTTP 200

### Get user by ID
GET http://localhost:8080/api/user/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update user by ID
PUT http://localhost:8080/api/user/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "first_name": "Jane Marie",
  "last_name": "Smith-Johnson",
  "phone": "+1234567891",
  "address": "789 Updated Street, New City"
}

HTTP 200

### Delete user by ID
DELETE http://localhost:8080/api/user/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all user commands
GET http://localhost:8080/api/user/command
Authorization: Bearer {{auth_token}}

HTTP 200

### Get user command by ID
GET http://localhost:8080/api/user/command/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Get all user queries
GET http://localhost:8080/api/user/query
Authorization: Bearer {{auth_token}}

HTTP 200

### Get user query by ID
GET http://localhost:8080/api/user/query/123
Authorization: Bearer {{auth_token}}

HTTP 200