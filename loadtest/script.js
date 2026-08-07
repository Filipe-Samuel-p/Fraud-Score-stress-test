import http from 'k6/http';
import { check, sleep } from 'k6';


const TARGET_URL = __ENV.TARGET_URL || 'http://fraud-score.default.svc.cluster.local/transaction';


export const options = {
  stages: [
    { duration: '1m', target: 50 },  
    { duration: '2m', target: 50 },  
    { duration: '30s', target: 0 },  
  ],
  thresholds: {
    http_req_failed: ['rate<0.05'],      
    http_req_duration: ['p(95)<800'],    
  },
};


function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
    const r = (Math.random() * 16) | 0;
    const v = c === 'x' ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

const countries = ['BR', 'US', 'PT', 'AR', 'DE'];
const currencies = ['BRL', 'USD', 'EUR'];
const merchants = ['amazon', 'mercadolivre', 'shopee', 'uber', 'ifood'];

function randomFrom(arr) {
  return arr[Math.floor(Math.random() * arr.length)];
}

export default function () {
  const payload = JSON.stringify({
    transaction_id: uuidv4(),
    account_id: `acc-${Math.floor(Math.random() * 100000)}`,
    amount: (Math.random() * 5000).toFixed(2),
    currency: randomFrom(currencies),
    country: randomFrom(countries),
    merchant: randomFrom(merchants),
    ip_address: `${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}.0.1`,
    occurred_at: new Date().toISOString(),
  });

  const res = http.post(TARGET_URL, payload, {
    headers: { 'Content-Type': 'application/json' },
  });

  check(res, {
    'status 201': (r) => r.status === 201,
  });

  sleep(0.5);
}
