import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0MSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjUzOTUxfQ.pLsMRMayor5GPFVmxl_rHJyb-WuLCr5APy0mKWsABJ0';

const WITHDRAW_ID = 1;
const CARD_NUMBER = '1234567890123456';

export const options = {
  scenarios: {
    endurance_test: {
      executor: 'constant-vus',
      vus: 150,
      duration: '30m',
    },
  },
};

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    `/api/withdraws?page=1&page_size=10`,
    `/api/withdraws/active?page=1&page_size=10`,
    `/api/withdraws/trashed?page=1&page_size=10`,
    `/api/withdraws/${WITHDRAW_ID}`,
    `/api/withdraws/card-number/${CARD_NUMBER}`,
    `/api/withdraws/monthly-success?year=2024&month=1`,
    `/api/withdraws/yearly-success?year=2024`,
    `/api/withdraws/monthly-failed?year=2024&month=1`,
    `/api/withdraws/yearly-failed?year=2024`,
    `/api/withdraws/monthly-amount?year=2024`,
    `/api/withdraws/yearly-amount?year=2024`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
