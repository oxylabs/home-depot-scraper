curl --user user:pass \
'https://realtime.oxylabs.io/v1/queries' \
-H "Content-Type: application/json" \
-d '{"source": "universal_ecommerce", "url": "https://www.homedepot.com/p/DEWALT-ATOMIC-20-Volt-MAX-Cordless-Brushless-Compact-Drill-Impact-Combo-Kit-2-Tool-with-2-1-3Ah-Batteries-Charger-Bag-DCK278C2/308067489", "geo_location": "United States"}'