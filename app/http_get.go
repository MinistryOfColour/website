package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GETJson(url string, data interface{}) bool {
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("Authorization", "Basic "+basicAuth(u, p))
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("%s: did not get acceptable status code: %v body: %q",
			url, resp.Status, string(body))
		return false
	}
	err = json.NewDecoder(resp.Body).Decode(data)
	// Session.Items = []shared.Item{}
	// err = json.NewDecoder(resp.Body).Decode(&Session.Items)
	// print("items appears to be", Session.Items)
	if err != nil {
		return false
	}

	return true
}
