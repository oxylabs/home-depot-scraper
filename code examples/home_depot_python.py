import requests
from pprint import pprint

# Structure payload.
payload = {
    'source': 'universal_ecommerce',
    'url': 'https://www.homedepot.com/p/DEWALT-ATOMIC-20-Volt-MAX-Cordless-Brushless-Compact-Drill-Impact-Combo-Kit-2-Tool-with-2-1-3Ah-Batteries-Charger-Bag-DCK278C2/308067489',
    'geo_location': 'United States',
}

# Get response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('user', 'pass1'),
    json=payload,
)

# Instead of response with job status and results url, this will return the
# JSON response with the result.
pprint(response.json())