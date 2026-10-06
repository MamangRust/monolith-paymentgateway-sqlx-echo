POST http://localhost:5000/api/auth/login
Content-Type: application/json
{
  "email": "john.doe@hellodotaa.com",
  "password": "password123"
}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data.access_token" exists
jsonpath "$.data.refresh_token" exists
[Captures]
auth_token: jsonpath "$.data.access_token"
refresh_token: jsonpath "$.data.refresh_token"

GET http://localhost:5000/api/auth/me
Authorization: Bearer {{auth_token}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data.email" == "john.doe@hellodotaa.com"


POST http://localhost:5000/api/auth/refresh-token
Content-Type: application/json
Authorization: Bearer {{auth_token}}
{
  "refresh_token": "{{refresh_token}}"
}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data.access_token" exists
jsonpath "$.data.refresh_token" exists