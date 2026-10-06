import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjUzOTUxfQ.pLsMRMayor5GPFVmxl_rHJyb-WuLCr5APy0mKWsABJ0';

const TRANSACTION_ID = 1;
const CARD_NUMBER = '1234567890123456';

export const options = {
  scenarios: {
    spike_test: {
      executor: 'ramping-vus',
      stages: [
        { duration: '10s', target: 50 },
        { duration: '10s', target: 1000 },
        { duration: '30s', target: 1000 },
        { duration: '10s', target: 50 },
      ],
    },
  },
};

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    `/api/transactions?page=1&page_size=10`,
    `/api/transactions/active?page=1&page_size=10`,
    `/api/transactions/trashed?page=1&page_size=10`,
    `/api/transactions/${TRANSACTION_ID}`,
    `/api/transactions/card-number/${CARD_NUMBER}`,
    `/api/transactions/monthly-success?year=2024&month=1`,
    `/api/transactions/yearly-success?year=2024`,
    `/api/transactions/monthly-failed?year=2024&month=1`,
    `/api/transactions/yearly-failed?year=2024`,
    `/api/transactions/monthly-methods?year=2024`,
    `/api/transactions/yearly-methods?year=2024`,
    `/api/transactions/monthly-amounts?year=2024`,
    `/api/transactions/yearly-amounts?year=2024`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
