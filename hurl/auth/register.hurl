POST http://localhost:5000/api/auth/register
Content-Type: application/json

{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@hellodota.com",
  "password": "password123",
  "confirm_password": "password123"
}

HTTP 200