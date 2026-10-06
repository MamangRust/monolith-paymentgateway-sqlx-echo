import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjUzOTUxfQ.pLsMRMayor5GPFVmxl_rHJyb-WuLCr5APy0mKWsABJ0';

const TRANSFER_ID = 1;

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
    `/api/transfers?page=1&page_size=10`,
    `/api/transfers/active?page=1&page_size=10`,
    `/api/transfers/trashed?page=1&page_size=10`,
    `/api/transfers/${TRANSFER_ID}`,
    `/api/transfers/monthly-success?year=2024&month=1`,
    `/api/transfers/yearly-success?year=2024`,
    `/api/transfers/monthly-failed?year=2024&month=1`,
    `/api/transfers/yearly-failed?year=2024`,
    `/api/transfers/monthly-amounts?year=2024`,
    `/api/transfers/yearly-amounts?year=2024`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
