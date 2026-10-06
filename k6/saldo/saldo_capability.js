import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjUzOTUxfQ.pLsMRMayor5GPFVmxl_rHJyb-WuLCr5APy0mKWsABJ0';

const SALDO_ID = 1;
const USER_ID = 1;
const CARD_NUMBER = '1234567890123456';

export const options = {
  scenarios: {
    scalability_test: {
      executor: 'ramping-arrival-rate',
      startRate: 50,
      timeUnit: '1s',
      stages: [
        { duration: '1m', target: 100 },
        { duration: '1m', target: 300 },
        { duration: '1m', target: 600 },
      ],
      preAllocatedVUs: 100,
      maxVUs: 900,
    },
  },
};

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    `/api/saldos?page=1&page_size=10`,
    `/api/saldos/active?page=1&page_size=10`,
    `/api/saldos/trashed?page=1&page_size=10`,
    `/api/saldos/${SALDO_ID}`,
    `/api/saldos/card-number/${CARD_NUMBER}`,
    `/api/saldos/user/${USER_ID}`,
    `/api/saldos/monthly-total-balance?year=2024&month=1`,
    `/api/saldos/yearly-total-balance?year=2024`,
    `/api/saldos/monthly-balances?year=2024`,
    `/api/saldos/yearly-balances?year=2024`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
