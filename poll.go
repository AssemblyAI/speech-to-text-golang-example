package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const API_KEY = "YOUR_API_KEY"
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"
	const POLLING_URL = TRANSCRIPT_URL + "/" + "ogs9rsf5pa-7cc9-45f1-a105-f01da771f2e4"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", POLLING_URL, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if result["status"] == "completed" {
		fmt.Println(result["text"])
	}
}
