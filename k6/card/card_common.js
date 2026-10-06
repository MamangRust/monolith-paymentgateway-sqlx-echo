import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN =
  'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjU1MTU0fQ.czewquqL85w2H0jYiDNQpvN34qqOaHAOUDtQxNd0NQU';

export default function () {
  const params = {
    headers: { Authorization: TOKEN, 'Content-Type': 'application/json' },
  };

  const basicEndpoints = [
    // Cards
    '/api/cards',
    '/api/cards/active?page=1&limit=10',
    '/api/cards/by-user/1',

    // Stats: Balance
    '/api/cards/stats/balance/monthly?year=2025',
    '/api/cards/stats/balance/yearly?year=2025',
    '/api/cards/stats/balance/monthly/by-card?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/balance/yearly/by-card?year=2025&card_number=4415599957419074',

    // Stats: Topup
    '/api/cards/stats/topup/monthly?year=2025',
    '/api/cards/stats/topup/yearly?year=2025',
    '/api/cards/stats/topup/monthly/by-card?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/topup/yearly/by-card?year=2025&card_number=4415599957419074',

    // Stats: Transaction
    '/api/cards/stats/transaction/monthly?year=2025',
    '/api/cards/stats/transaction/yearly?year=2025',
    '/api/cards/stats/transaction/monthly/by-card?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/transaction/yearly/by-card?year=2025&card_number=4415599957419074',

    // Stats: Transfer
    '/api/cards/stats/transfer/monthly/sender?year=2025',
    '/api/cards/stats/transfer/monthly/receiver?year=2025',
    '/api/cards/stats/transfer/yearly/sender?year=2025',
    '/api/cards/stats/transfer/yearly/receiver?year=2025',
    '/api/cards/stats/transfer/monthly/by-card/sender?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/transfer/monthly/by-card/receiver?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/transfer/yearly/by-card/sender?year=2025&card_number=4415599957419074',
    '/api/cards/stats/transfer/yearly/by-card/receiver?year=2025&card_number=4415599957419074',

    // Stats: Withdraw
    '/api/cards/stats/withdraw/monthly?year=2025',
    '/api/cards/stats/withdraw/yearly?year=2025',
    '/api/cards/stats/withdraw/monthly/by-card?year=2025&month=1&card_number=4415599957419074',
    '/api/cards/stats/withdraw/yearly/by-card?year=2025&card_number=4415599957419074',
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
