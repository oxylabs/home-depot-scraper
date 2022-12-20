import fetch from 'node-fetch';

const username = 'YOUR_USERNAME';
const password = 'YOUR_PASSWORD';
const body = {
  'source': 'universal_ecommerce',
  'url': 'https://www.homedepot.com/p/DEWALT-ATOMIC-20-Volt-MAX-Cordless-Brushless-Compact-Drill-Impact-Combo-Kit-2-Tool-with-2-1-3Ah-Batteries-Charger-Bag-DCK278C2/308067489',
  'geo_location': 'United States'
};
const response = await fetch('https://realtime.oxylabs.io/v1/queries', {
  method: 'post',
  body: JSON.stringify(body),
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Basic ' + Buffer.from(`${username}:${password}`).toString('base64'),
  }
});

console.log(await response.json());