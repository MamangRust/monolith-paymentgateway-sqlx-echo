import http from 'k6/http';
import { check, sleep } from 'k6';

const baseURL = __ENV.BASE_URL || 'http://localhost:5000';
const token = __ENV.K6_TOKEN || '';
const vus = Number(__ENV.K6_VUS || 1);
const rampUp = __ENV.K6_RAMP_UP || '30s';
const hold = __ENV.K6_HOLD || '60s';
const rampDown = __ENV.K6_RAMP_DOWN || '30s';
const p95 = Number(__ENV.K6_P95_MS || 500);
const endpoints = (__ENV.K6_ENDPOINTS || '/health')
  .split(',')
  .map((endpoint) => endpoint.trim())
  .filter((endpoint) => endpoint.length > 0);

export const options = {
  scenarios: {
    read_baseline: {
      executor: 'ramping-vus',
      startVUs: 1,
      stages: [
        { duration: rampUp, target: vus },
        { duration: hold, target: vus },
        { duration: rampDown, target: 0 },
      ],
      gracefulRampDown: '10s',
    },
  },
  thresholds: {
    http_req_duration: [`p(95)<${p95}`],
    http_req_failed: ['rate<0.01'],
    checks: ['rate>0.99'],
  },
};

export default function () {
  const params = {
    headers: token ? { Authorization: `Bearer ${token}` } : {},
    tags: { workload: 'p2-read-baseline' },
  };

  for (const endpoint of endpoints) {
    const response = http.get(`${baseURL}${endpoint}`, params);
    check(response, {
      [`${endpoint} returns 2xx`]: (result) => result.status >= 200 && result.status < 300,
    });
  }

  sleep(1);
}
