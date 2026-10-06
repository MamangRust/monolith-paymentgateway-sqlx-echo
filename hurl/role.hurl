# Role Service API Tests
# Base URL: http://localhost:8080

### Create a new role
POST http://localhost:8080/api/role
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "name": "admin",
  "description": "Administrator role with full access",
  "permissions": ["read", "write", "delete", "manage_users"]
}

HTTP 200

### Get all roles
GET http://localhost:8080/api/role
Authorization: Bearer {{auth_token}}

HTTP 200

### Get role by ID
GET http://localhost:8080/api/role/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update role by ID
PUT http://localhost:8080/api/role/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "name": "super_admin",
  "description": "Super administrator with all privileges",
  "permissions": ["read", "write", "delete", "manage_users", "system_config"]
}

HTTP 200

### Delete role by ID
DELETE http://localhost:8080/api/role/123
Authorization: Bearer {{auth_token}}

HTTP 200