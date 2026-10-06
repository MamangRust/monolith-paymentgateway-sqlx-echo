import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = (__ENV.BASE_URL || 'http://localhost:5000').replace(/\/$/, '');
const RAW_TOKEN = __ENV.K6_TOKEN || __ENV.TOKEN || '';
const TOKEN = RAW_TOKEN && RAW_TOKEN.startsWith('Bearer ') ? RAW_TOKEN : `Bearer ${RAW_TOKEN}`;
const API_KEY = __ENV.K6_API_KEY || '';
const USER_ID = Number(__ENV.K6_USER_ID || 1);
const MERCHANT_ID = Number(__ENV.K6_MERCHANT_ID || 0);
const CARD_NUMBER = __ENV.K6_CARD_NUMBER || '';
const RECEIVER_CARD_NUMBER = __ENV.K6_RECEIVER_CARD_NUMBER || '';
const SALDO_ID = Number(__ENV.K6_SALDO_ID || 0);
const ALLOW_FINANCIAL_DELETE = (__ENV.K6_ALLOW_FINANCIAL_DELETE || 'false') === 'true';
const DOMAINS = (__ENV.K6_DOMAINS || 'user,role,merchant,card,merchant-document').split(',').map((v) => v.trim()).filter(Boolean);

export const options = {
  scenarios: {
    lifecycle: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: __ENV.K6_RAMP_UP || '5s', target: Number(__ENV.K6_VUS || 1) },
        { duration: __ENV.K6_HOLD || '15s', target: Number(__ENV.K6_VUS || 1) },
        { duration: __ENV.K6_RAMP_DOWN || '5s', target: 0 },
      ],
      gracefulRampDown: '10s',
    },
  },
  thresholds: {
    http_req_failed: ['rate<0.05'],
    checks: ['rate>0.95'],
  },
};

function params(extraHeaders = {}) {
  return {
    headers: {
      Authorization: TOKEN,
      'Content-Type': 'application/json',
      ...extraHeaders,
    },
    tags: { suite: 'crud-lifecycle' },
  };
}

function url(path) {
  return `${BASE_URL}${path}`;
}

function json(res) {
  try {
    return res.json();
  } catch (_) {
    return {};
  }
}

function idOf(res) {
  const body = json(res);
  return body?.data?.id || body?.data?.document_id || body?.id || body?.document_id || 0;
}

function request(method, path, body = null, extraHeaders = {}) {
  const payload = body === null ? null : JSON.stringify(body);
  return http.request(method, url(path), payload, params(extraHeaders));
}

function expect(res, name, statuses = [200]) {
  return check(res, {
    [`${name}: status ${statuses.join('/')}`]: (r) => statuses.includes(r.status),
  });
}

function create(path, body, name) {
  const res = request('POST', path, body);
  expect(res, `${name} create`);
  return idOf(res);
}

function runUser() {
  const suffix = `${__VU}-${__ITER}-${Date.now()}`;
  const email = `k6.lifecycle.${suffix}@example.com`;
  const id = create('/api/user-command/create', {
    firstname: 'K6', lastname: 'Lifecycle', email, password: 'password123', confirm_password: 'password123',
  }, 'user');
  if (!id) return false;

  expect(request('GET', `/api/user-query/${id}`), 'user read');
  expect(request('POST', `/api/user-command/update/${id}`, {
    firstname: 'K6Updated', lastname: 'Lifecycle', email, password: 'password123', confirm_password: 'password123',
  }), 'user update');
  expect(request('POST', `/api/user-command/trashed/${id}`), 'user trash');
  expect(request('POST', `/api/user-command/restore/${id}`), 'user restore');
  expect(request('DELETE', `/api/user-command/permanent/${id}`), 'user permanent delete');
  return true;
}

function runRole() {
  const name = `k6-lifecycle-${__VU}-${__ITER}-${Date.now()}`;
  const id = create('/api/role', { name }, 'role');
  if (!id) return false;

  expect(request('GET', `/api/role-query/${id}`), 'role read');
  expect(request('POST', `/api/role/${id}`, { name: `${name}-updated` }), 'role update');
  expect(request('POST', `/api/role/trashed/${id}`), 'role trash');
  expect(request('PUT', `/api/role/restore/${id}`), 'role restore');
  expect(request('DELETE', `/api/role/permanent/${id}`), 'role permanent delete');
  return true;
}

function runMerchant() {
  const name = `K6 Merchant ${__VU}-${__ITER}-${Date.now()}`;
  const id = create('/api/merchant-command/create', { name, user_id: USER_ID }, 'merchant');
  if (!id) return 0;

  expect(request('GET', `/api/merchant-query/${id}`), 'merchant read');
  expect(request('POST', `/api/merchant-command/update/${id}`, {
    name: `${name} Updated`, user_id: USER_ID, status: 'active',
  }), 'merchant update');
  return id;
}

function finishMerchant(id) {
  expect(request('POST', `/api/merchant-command/trashed/${id}`), 'merchant trash');
  expect(request('POST', `/api/merchant-command/restore/${id}`), 'merchant restore');
  expect(request('DELETE', `/api/merchant-command/permanent/${id}`), 'merchant permanent delete');
}

function runCard() {
  const id = create('/api/card-command/create', {
    user_id: USER_ID,
    card_type: __ENV.K6_CARD_TYPE || 'debit',
    expire_date: __ENV.K6_CARD_EXPIRE_DATE || '2030-12-31T00:00:00Z',
    cvv: __ENV.K6_CARD_CVV || '123',
    card_provider: __ENV.K6_CARD_PROVIDER || 'visa',
  }, 'card');
  if (!id) return false;

  expect(request('GET', `/api/card-query/${id}`), 'card read');
  expect(request('POST', `/api/card-command/update/${id}`, {
    card_id: id, user_id: USER_ID, card_type: 'debit',
    expire_date: __ENV.K6_CARD_EXPIRE_DATE || '2030-12-31T00:00:00Z',
    cvv: __ENV.K6_CARD_CVV || '123', card_provider: __ENV.K6_CARD_PROVIDER || 'visa',
  }), 'card update');
  expect(request('POST', `/api/card-command/trashed/${id}`), 'card trash');
  expect(request('POST', `/api/card-command/restore/${id}`), 'card restore');
  expect(request('DELETE', `/api/card-command/permanent/${id}`), 'card permanent delete');
  return true;
}

function runSaldo() {
  if (!CARD_NUMBER || !SALDO_ID) {
    check(false, { 'saldo requires K6_CARD_NUMBER and K6_SALDO_ID': () => false });
    return false;
  }
  expect(request('GET', `/api/saldo-query/${SALDO_ID}`), 'saldo read');
  expect(request('POST', `/api/saldo-command/update/${SALDO_ID}`, {
    saldo_id: SALDO_ID, card_number: CARD_NUMBER, total_balance: Number(__ENV.K6_SALDO_BALANCE || 100000),
  }), 'saldo update');
  if (ALLOW_FINANCIAL_DELETE) {
    expect(request('POST', `/api/saldo-command/trashed/${SALDO_ID}`), 'saldo trash');
    expect(request('POST', `/api/saldo-command/restore/${SALDO_ID}`), 'saldo restore');
  }
  return true;
}

function runRecordLifecycle(domain, createPath, readPath, updatePath, trashPath, restorePath, deletePath, body, updateBody, cleanup = false) {
  const id = create(createPath, body, domain);
  if (!id) return false;
  expect(request('GET', readPath.replace(':id', id)), `${domain} read`);
  expect(request('POST', updatePath.replace(':id', id), updateBody(id)), `${domain} update`);
  if (cleanup || ALLOW_FINANCIAL_DELETE) {
    expect(request('POST', trashPath.replace(':id', id)), `${domain} trash`);
    expect(request('POST', restorePath.replace(':id', id)), `${domain} restore`);
    expect(request('DELETE', deletePath.replace(':id', id)), `${domain} permanent delete`);
  }
  return true;
}

function runTopup() {
  if (!CARD_NUMBER) return false;
  const key = `topup:create:k6:${__VU}-${__ITER}-${Date.now()}`;
  return runRecordLifecycle('topup', '/api/topup-command/create', '/api/topup-query/:id', '/api/topup-command/update/:id', '/api/topup-command/trashed/:id', '/api/topup-command/restore/:id', '/api/topup-command/permanent/:id', {
    card_number: CARD_NUMBER, topup_amount: Number(__ENV.K6_AMOUNT || 50000), topup_method: __ENV.K6_TOPUP_METHOD || 'visa', idempotency_key: key,
  }, (id) => ({ topup_id: id, card_number: CARD_NUMBER, topup_amount: Number(__ENV.K6_AMOUNT || 50000), topup_method: __ENV.K6_TOPUP_METHOD || 'visa' }));
}

function runTransaction() {
  if (!CARD_NUMBER || !API_KEY || !MERCHANT_ID) return false;
  const key = `transaction:create:k6:${__VU}-${__ITER}-${Date.now()}`;
  return runRecordLifecycle('transaction', '/api/transaction-command/create', '/api/transaction-query/:id', '/api/transaction-command/update/:id', '/api/transaction-command/trashed/:id', '/api/transaction-command/restore/:id', '/api/transaction-command/permanent/:id', {
    card_number: CARD_NUMBER, amount: Number(__ENV.K6_AMOUNT || 50000), payment_method: __ENV.K6_PAYMENT_METHOD || 'visa', merchant_id: MERCHANT_ID, transaction_time: new Date().toISOString(), idempotency_key: key,
  }, (id) => ({ transaction_id: id, card_number: CARD_NUMBER, amount: Number(__ENV.K6_AMOUNT || 50000), payment_method: __ENV.K6_PAYMENT_METHOD || 'visa', merchant_id: MERCHANT_ID, transaction_time: new Date().toISOString() }));
}

function runTransfer() {
  if (!CARD_NUMBER || !RECEIVER_CARD_NUMBER) return false;
  const key = `transfer:create:k6:${__VU}-${__ITER}-${Date.now()}`;
  return runRecordLifecycle('transfer', '/api/transfer-command/create', '/api/transfer-query/:id', '/api/transfer-command/update/:id', '/api/transfer-command/trashed/:id', '/api/transfer-command/restore/:id', '/api/transfer-command/permanent/:id', {
    transfer_from: CARD_NUMBER, transfer_to: RECEIVER_CARD_NUMBER, transfer_amount: Number(__ENV.K6_AMOUNT || 50000), idempotency_key: key,
  }, (id) => ({ transfer_id: id, transfer_from: CARD_NUMBER, transfer_to: RECEIVER_CARD_NUMBER, transfer_amount: Number(__ENV.K6_AMOUNT || 50000) }));
}

function runWithdraw() {
  if (!CARD_NUMBER) return false;
  const key = `withdraw:create:k6:${__VU}-${__ITER}-${Date.now()}`;
  return runRecordLifecycle('withdraw', '/api/withdraw-command/create', '/api/withdraw-query/:id', '/api/withdraw-command/update/:id', '/api/withdraw-command/trashed/:id', '/api/withdraw-command/restore/:id', '/api/withdraw-command/permanent/:id', {
    card_number: CARD_NUMBER, withdraw_amount: Number(__ENV.K6_AMOUNT || 50000), withdraw_time: new Date().toISOString(), idempotency_key: key,
  }, (id) => ({ withdraw_id: id, card_number: CARD_NUMBER, withdraw_amount: Number(__ENV.K6_AMOUNT || 50000), withdraw_time: new Date().toISOString() }));
}

function runMerchantDocument(merchantID) {
  const id = merchantID || MERCHANT_ID;
  if (!id) return false;
  return runRecordLifecycle('merchant-document', '/api/merchant-document-command/create', '/api/merchant-document-query/:id', '/api/merchant-document-command/update/:id', '/api/merchant-document-command/trashed/:id', '/api/merchant-document-command/restore/:id', '/api/merchant-document-command/permanent/:id', {
    merchant_id: id, document_type: 'license', document_url: `https://example.com/k6/${__VU}-${__ITER}`,
  }, (docID) => ({ document_id: docID, merchant_id: id, document_type: 'license', document_url: `https://example.com/k6/${__VU}-${__ITER}`, status: 'pending', note: 'k6 lifecycle update' }), true);
}

export function setup() {
  if (!TOKEN || TOKEN === 'Bearer ') throw new Error('K6_TOKEN is required');

  const selected = new Set(DOMAINS);
  if (selected.has('saldo') && (!CARD_NUMBER || !SALDO_ID)) throw new Error('saldo requires K6_CARD_NUMBER and K6_SALDO_ID');
  if (selected.has('topup') && !CARD_NUMBER) throw new Error('topup requires K6_CARD_NUMBER');
  if (selected.has('withdraw') && !CARD_NUMBER) throw new Error('withdraw requires K6_CARD_NUMBER');
  if (selected.has('transfer') && (!CARD_NUMBER || !RECEIVER_CARD_NUMBER)) throw new Error('transfer requires K6_CARD_NUMBER and K6_RECEIVER_CARD_NUMBER');
  if (selected.has('transaction') && (!CARD_NUMBER || !API_KEY || !MERCHANT_ID)) throw new Error('transaction requires K6_CARD_NUMBER, K6_API_KEY and K6_MERCHANT_ID');
  if (selected.has('merchant-document') && !selected.has('merchant') && !MERCHANT_ID) throw new Error('merchant-document requires K6_MERCHANT_ID when merchant is not selected');
  return {};
}

export default function () {
  let merchantID = 0;
  for (const domain of DOMAINS) {
    if (domain === 'user') runUser();
    else if (domain === 'role') runRole();
    else if (domain === 'merchant') merchantID = runMerchant();
    else if (domain === 'card') runCard();
    else if (domain === 'saldo') runSaldo();
    else if (domain === 'topup') runTopup();
    else if (domain === 'transaction') runTransaction();
    else if (domain === 'transfer') runTransfer();
    else if (domain === 'withdraw') runWithdraw();
    else if (domain === 'merchant-document') runMerchantDocument(merchantID);
  }

  if (merchantID) finishMerchant(merchantID);
  sleep(Number(__ENV.K6_SLEEP || 0.2));
}
