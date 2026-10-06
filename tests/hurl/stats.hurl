# =============================================================================
# Stats Endpoints E2E — covers all /api/*-stats-* routes (104 endpoints)
# Query params follow the seeded dataset (year=2026, month=4, card 4532123456789012,
# merchant_id=2, api_key API_KEY_002_IJKLMNOP789012) used by the stats probe.
# =============================================================================

# Login (capture tokens for authenticated requests)
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# -----------------------------------------------------------------------------
# Saldo stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/saldo-stats-balance/monthly-balances?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/saldo-stats-balance/yearly-balances?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/saldo-stats-total-balance/monthly-total-balance?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/saldo-stats-total-balance/yearly-total-balance?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Card stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/card-stats-balance/monthly-balance?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-balance/yearly-balance?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-balance/monthly-balance-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-balance/yearly-balance-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-topup/monthly-topup-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-topup/yearly-topup-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-topup/monthly-topup-amount-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-topup/yearly-topup-amount-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transaction/monthly-transaction-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transaction/yearly-transaction-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transaction/monthly-transaction-amount-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transaction/yearly-transaction-amount-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-withdraw/monthly-withdraw-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-withdraw/yearly-withdraw-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-withdraw/monthly-withdraw-amount-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-withdraw/yearly-withdraw-amount-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/monthly-transfer-sender-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/yearly-transfer-sender-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/monthly-transfer-receiver-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/yearly-transfer-receiver-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/monthly-transfer-sender-amount-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/yearly-transfer-sender-amount-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/monthly-transfer-receiver-amount-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/card-stats-transfer/yearly-transfer-receiver-amount-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Topup stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/topup-stats-amount/monthly-amounts?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-amount/yearly-amounts?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-amount/monthly-amounts-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-amount/yearly-amounts-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/monthly-success?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/yearly-success?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/monthly-failed?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/yearly-failed?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/monthly-success-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/yearly-success-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/monthly-failed-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-status/yearly-failed-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-method/monthly-methods?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-method/yearly-methods?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-method/monthly-methods-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/topup-stats-method/yearly-methods-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Transaction stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/transaction-stats-amount/monthly-amounts-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-amount/yearly-amounts-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-amount/monthly-amounts?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-amount/yearly-amounts?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/monthly-success?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/yearly-success?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/monthly-failed?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/yearly-failed?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/monthly-success-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/yearly-success-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/monthly-failed-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-status/yearly-failed-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-method/monthly-methods?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-method/yearly-methods?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-method/monthly-methods-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transaction-stats-method/yearly-methods-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Transfer stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/transfer-stats-amount/monthly-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-amount/yearly-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-amount/monthly-by-sender?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-amount/yearly-by-sender?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-amount/monthly-by-receiver?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-amount/yearly-by-receiver?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/monthly-success?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/yearly-success?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/monthly-failed?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/yearly-failed?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/monthly-success-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/yearly-success-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/monthly-failed-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/transfer-stats-status/yearly-failed-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Withdraw stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/withdraw-stats-amount/monthly-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-amount/yearly-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-amount/monthly-amount-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-amount/yearly-amount-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/monthly-success?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/yearly-success?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/monthly-failed?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/yearly-failed?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/monthly-success-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/yearly-success-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/monthly-failed-by-card?year=2026&month=4&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/withdraw-stats-status/yearly-failed-by-card?year=2026&card_number={{cardNumber}}
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Merchant stats
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/merchant-stats-amount/monthly-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-amount/yearly-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-amount/monthly-amount-by-merchant?year=2026&month=4&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-amount/yearly-amount-by-merchant?year=2026&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-amount/monthly-amount-by-apikey?year=2026&month=4&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-amount/yearly-amount-by-apikey?year=2026&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/monthly-payment-methods?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/yearly-payment-methods?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/monthly-payment-methods-by-merchant?year=2026&month=4&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/yearly-payment-methods-by-merchant?year=2026&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/monthly-payment-methods-by-apikey?year=2026&month=4&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-method/yearly-payment-methods-by-apikey?year=2026&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/monthly-total-amount?year=2026&month=4
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/yearly-total-amount?year=2026
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/monthly-totalamount-by-merchant?year=2026&month=4&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/yearly-totalamount-by-merchant?year=2026&merchant_id=2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/monthly-totalamount-by-apikey?year=2026&month=4&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-stats-totalamount/yearly-totalamount-by-apikey?year=2026&api_key=API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

# -----------------------------------------------------------------------------
# Merchant transactions (reporting)
# -----------------------------------------------------------------------------
GET {{baseUrl}}/api/merchant-transactions/transactions
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-transactions/transactions/2
Authorization: Bearer {{accessToken}}
HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists

GET {{baseUrl}}/api/merchant-transactions/transactions/api-key/API_KEY_002_IJKLMNOP789012
Authorization: Bearer {{accessToken}}
HTTP 200

[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.data" exists
