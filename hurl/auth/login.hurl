POST http://localhost:5000/api/auth/login
Content-Type: application/json
{
  "email": "john.doe@hellodota.com",
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