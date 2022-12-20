# Home Depot Scraper

# How it works

You can get Home Depot results by providing your own URLs to our service. We can return the HTML of any Home Depot page you like.

#### Python code example

The example below illustrates how you can get a Home Depot product page result in HTML format.

```python
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
```

Code examples for other languages can be found [**here.**](https://github.com/oxylabs/home-depot-scraper/tree/main/code%20examples)

#### Output

After executing the code example above, you can expect to get an output similar to this:

```json
{
    "results": [
        {
            "content":"<!doctype html>\n<html lang=\"en\">\n<head>
            ...
            </script></body>\n</html>\n",
            "created_at": "2022-11-17 14:53:52",
            "updated_at": "2022-11-17 14:53:54",
            "page": 1,
            "url": "https://www.homedepot.com/p/DEWALT-ATOMIC-20-Volt-MAX-Cordless-Brushless-Compact-Drill-Impact-Combo-Kit-2-Tool-with-2-1-3Ah-Batteries-Charger-Bag-DCK278C2/308067489",
            "job_id": "6999021798385787905",
            "status_code": 200
        }
    ]
}
```
