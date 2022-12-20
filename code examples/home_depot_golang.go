package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const Username = "YOUR_USERNAME"
	const Password = "YOUR_PASSWORD"

	payload := map[string]string{
		"source": "universal_ecommerce",
		"url": "https://www.homedepot.com/p/DEWALT-ATOMIC-20-Volt-MAX-Cordless-Brushless-Compact-Drill-Impact-Combo-Kit-2-Tool-with-2-1-3Ah-Batteries-Charger-Bag-DCK278C2/308067489",
        	"geo_location": "United States",
	}

	jsonValue, _ := json.Marshal(payload)

	client := &http.Client{}
	request, _ := http.NewRequest("POST",
		"https://realtime.oxylabs.io/v1/queries",
		bytes.NewBuffer(jsonValue),
	)

	request.SetBasicAuth(Username, Password)
	response, _ := client.Do(request)

	responseText, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseText))
}