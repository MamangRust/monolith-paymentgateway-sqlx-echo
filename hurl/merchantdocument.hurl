# Merchant Document Service API Tests
# Base URL: http://localhost:8080

### Create a new merchant document
POST http://localhost:8080/api/merchantdocument
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "merchant_id": "123",
  "document_type": "business_license",
  "document_name": "Business License.pdf",
  "document_url": "https://example.com/documents/business_license.pdf",
  "status": "pending"
}

HTTP 200

### Get all merchant documents
GET http://localhost:8080/api/merchantdocument
Authorization: Bearer {{auth_token}}

HTTP 200

### Get merchant document by ID
GET http://localhost:8080/api/merchantdocument/123
Authorization: Bearer {{auth_token}}

HTTP 200

### Update merchant document by ID
PUT http://localhost:8080/api/merchantdocument/123
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "document_name": "Updated Business License.pdf",
  "status": "approved"
}

HTTP 200

### Delete merchant document by ID
DELETE http://localhost:8080/api/merchantdocument/123
Authorization: Bearer {{auth_token}}

HTTP 200