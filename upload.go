package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const API_KEY = "YOUR_API_KEY"
	const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"

	// Load file
	data, err := ioutil.ReadFile("english.m4a")
	if err != nil {
		log.Fatalln(err)
	}

	// Setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
	req.Header.Set("authorization", API_KEY)
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	// decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// print the upload_url
	fmt.Println(result["upload_url"])
}
