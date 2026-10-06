import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = (__ENV.BASE_URL || 'http://localhost:5000').replace(/\/$/, '');
const RAW_TOKEN = __ENV.K6_TOKEN || __ENV.TOKEN || '';
const TOKEN = RAW_TOKEN && RAW_TOKEN.startsWith('Bearer ') ? RAW_TOKEN : `Bearer ${RAW_TOKEN}`;
const API_KEY = __ENV.K6_API_KEY || '';
const CARD_NUMBER = __ENV.K6_CARD_NUMBER || '';
const RECEIVER_CARD_NUMBER = __ENV.K6_RECEIVER_CARD_NUMBER || '';
const MERCHANT_ID = Number(__ENV.K6_MERCHANT_ID || 0);
const AMOUNT = Number(__ENV.K6_AMOUNT || 50000);
const DOMAIN = (__ENV.K6_FINANCIAL_DOMAIN || 'topup').toLowerCase();
const REPLAY = (__ENV.K6_REPLAY || 'true') === 'true';
const CONCURRENT = (__ENV.K6_CONCURRENT || 'true') === 'true';

export const options = {
  scenarios: {
    financial: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: __ENV.K6_RAMP_UP || '5s', target: Number(__ENV.K6_VUS || 1) },
        { duration: __ENV.K6_HOLD || '20s', target: Number(__ENV.K6_VUS || 1) },
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

function headers() {
  const result = {
    Authorization: TOKEN,
    'Content-Type': 'application/json',
  };
  if (API_KEY) result['X-Api-Key'] = API_KEY;
  return result;
}

function request(path, body) {
  return {
    method: 'POST',
    url: `${BASE_URL}${path}`,
    body: JSON.stringify(body),
    params: { headers: headers(), tags: { suite: 'financial-concurrency', domain: DOMAIN } },
  };
}

function responseID(res) {
  try {
    return res.json('data.id') || res.json('data.transaction_id') || '';
  } catch (_) {
    return '';
  }
}

function sameReplayResponses(responses, label) {
  const firstID = responseID(responses[0]);
  return check(responses, {
    [`${label}: every replay is HTTP 200`]: (rs) => rs.length === 2 && rs.every((r) => r.status === 200),
    [`${label}: replay returns the same record`]: (rs) => firstID !== '' && rs.every((r) => responseID(r) === firstID),
  });
}

function sequentialReplay(path, body, label) {
  const first = http.post(`${BASE_URL}${path}`, JSON.stringify(body), { headers: headers(), tags: { suite: 'financial-concurrency', domain: DOMAIN } });
  check(first, { [`${label}: first request is HTTP 200`]: (r) => r.status === 200 });
  if (!REPLAY) return;

  const replay = http.post(`${BASE_URL}${path}`, JSON.stringify(body), { headers: headers(), tags: { suite: 'financial-concurrency', domain: DOMAIN } });
  check(replay, {
    [`${label}: sequential replay is HTTP 200`]: (r) => r.status === 200,
    [`${label}: sequential replay keeps ID`]: (r) => responseID(r) !== '' && responseID(r) === responseID(first),
  });
}

function concurrentReplay(path, body, label) {
  if (!CONCURRENT) return;
  const concurrentBody = { ...body, idempotency_key: `${body.idempotency_key}:race` };
  const responses = http.batch([request(path, concurrentBody), request(path, concurrentBody)]);
  sameReplayResponses(responses, label);
}

function runTopup() {
  if (!CARD_NUMBER) {
    check(false, { 'topup requires K6_CARD_NUMBER': () => false });
    return;
  }
  const body = {
    card_number: CARD_NUMBER,
    topup_amount: AMOUNT,
    topup_method: __ENV.K6_TOPUP_METHOD || 'visa',
    idempotency_key: `topup:create:k6:${__VU}-${__ITER}-${Date.now()}`,
  };
  sequentialReplay('/api/topup-command/create', body, 'topup');
  concurrentReplay('/api/topup-command/create', body, 'topup concurrent replay');
}

function runTransfer() {
  if (!CARD_NUMBER || !RECEIVER_CARD_NUMBER) {
    check(false, { 'transfer requires K6_CARD_NUMBER and K6_RECEIVER_CARD_NUMBER': () => false });
    return;
  }
  const body = {
    transfer_from: CARD_NUMBER,
    transfer_to: RECEIVER_CARD_NUMBER,
    transfer_amount: AMOUNT,
    idempotency_key: `transfer:create:k6:${__VU}-${__ITER}-${Date.now()}`,
  };
  sequentialReplay('/api/transfer-command/create', body, 'transfer');
  concurrentReplay('/api/transfer-command/create', body, 'transfer concurrent replay');
}

function runWithdraw() {
  if (!CARD_NUMBER) {
    check(false, { 'withdraw requires K6_CARD_NUMBER': () => false });
    return;
  }
  const body = {
    card_number: CARD_NUMBER,
    withdraw_amount: AMOUNT,
    withdraw_time: new Date().toISOString(),
    idempotency_key: `withdraw:create:k6:${__VU}-${__ITER}-${Date.now()}`,
  };
  sequentialReplay('/api/withdraw-command/create', body, 'withdraw');
  concurrentReplay('/api/withdraw-command/create', body, 'withdraw concurrent replay');
}

function runTransaction() {
  if (!CARD_NUMBER || !API_KEY || !MERCHANT_ID) {
    check(false, { 'transaction requires K6_CARD_NUMBER, K6_API_KEY and K6_MERCHANT_ID': () => false });
    return;
  }
  const body = {
    card_number: CARD_NUMBER,
    amount: AMOUNT,
    payment_method: __ENV.K6_PAYMENT_METHOD || 'visa',
    merchant_id: MERCHANT_ID,
    transaction_time: new Date().toISOString(),
    idempotency_key: `transaction:create:k6:${__VU}-${__ITER}-${Date.now()}`,
  };
  sequentialReplay('/api/transaction-command/create', body, 'transaction');
  concurrentReplay('/api/transaction-command/create', body, 'transaction concurrent replay');
}

export default function () {
  if (!TOKEN || TOKEN === 'Bearer ') {
    check(false, { 'K6_TOKEN is configured': () => false });
    return;
  }

  if (DOMAIN === 'topup') runTopup();
  else if (DOMAIN === 'transfer') runTransfer();
  else if (DOMAIN === 'withdraw') runWithdraw();
  else if (DOMAIN === 'transaction') runTransaction();
  else check(false, { 'K6_FINANCIAL_DOMAIN is supported': () => false });

  sleep(Number(__ENV.K6_SLEEP || 0.2));
}
