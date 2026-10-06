import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN =
  'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjU1MTU0fQ.czewquqL85w2H0jYiDNQpvN34qqOaHAOUDtQxNd0NQU';

export const options = {
  scenarios: {
    stress_test: {
      executor: 'ramping-vus',
      startVUs: 100,
      stages: [
        { duration: '30s', target: 300 },
        { duration: '30s', target: 600 },
        { duration: '30s', target: 1000 },
        { duration: '30s', target: 1500 },
      ],
    },
  },
};

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    // ===== card =====
    '/api/card?page=1&limit=10',
    '/api/card/active?page=1&limit=10',
    '/api/card/trashed?page=1&limit=10',
    '/api/card/user?user_id=11',
    '/api/card/card_number/4389195775213720',
    '/api/card/11',

    // ===== Dashboard =====
    '/api/card/dashboard',
    '/api/card/dashboard/4016531504435114',

    // ===== Balance =====
    '/api/card/monthly-balance?year=2025&month=1',
    '/api/card/yearly-balance?year=2025',
    '/api/card/monthly-balance-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-balance-by-card?year=2025&card_number=4016531504435114',

    // ===== Topup =====
    '/api/card/monthly-topup-amount?year=2025&month=1',
    '/api/card/yearly-topup-amount?year=2025',
    '/api/card/monthly-topup-amount-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-topup-amount-by-card?year=2025&card_number=4016531504435114',

    // ===== Withdraw =====
    '/api/card/monthly-withdraw-amount?year=2025&month=1',
    '/api/card/yearly-withdraw-amount?year=2025',
    '/api/card/monthly-withdraw-amount-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-withdraw-amount-by-card?year=2025&card_number=4016531504435114',

    // ===== Transaction =====
    '/api/card/monthly-transaction-amount?year=2025&month=1',
    '/api/card/yearly-transaction-amount?year=2025',
    '/api/card/monthly-transaction-amount-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-transaction-amount-by-card?year=2025&card_number=4016531504435114',

    // ===== Transfer Sender =====
    '/api/card/monthly-transfer-sender-amount?year=2025&month=1',
    '/api/card/yearly-transfer-sender-amount?year=2025',
    '/api/card/monthly-transfer-sender-amount-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-transfer-sender-amount-by-card?year=2025&card_number=4016531504435114',

    // ===== Transfer Receiver =====
    '/api/card/monthly-transfer-receiver-amount?year=2025&month=1',
    '/api/card/yearly-transfer-receiver-amount?year=2025',
    '/api/card/monthly-transfer-receiver-amount-by-card?year=2025&month=1&card_number=4016531504435114',
    '/api/card/yearly-transfer-receiver-amount-by-card?year=2025&card_number=4016531504435114',
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
