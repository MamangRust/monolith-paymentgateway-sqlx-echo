import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjUzOTUxfQ.pLsMRMayor5GPFVmxl_rHJyb-WuLCr5APy0mKWsABJ0';

const TOPUP_ID = 1;
const CARD_NUMBER = '1234567890123456';

export const options = {
  scenarios: {
    load_test: {
      executor: 'constant-vus',
      vus: 1000,
      duration: '2m',
    },
  },
  thresholds: {
    http_req_duration: ['p(95)<200'],
    http_req_failed: ['rate<0.01'],
  },
};

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    `/api/topups?page=1&page_size=10`,
    `/api/topups/active?page=1&page_size=10`,
    `/api/topups/trashed?page=1&page_size=10`,
    `/api/topups/${TOPUP_ID}`,
    `/api/topups/card-number/${CARD_NUMBER}`,
    `/api/topups/monthly-success?year=2024&month=1`,
    `/api/topups/yearly-success?year=2024`,
    `/api/topups/monthly-failed?year=2024&month=1`,
    `/api/topups/yearly-failed?year=2024`,
    `/api/topups/monthly-methods?year=2024`,
    `/api/topups/yearly-methods?year=2024`,
    `/api/topups/monthly-amounts?year=2024`,
    `/api/topups/yearly-amounts?year=2024`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
