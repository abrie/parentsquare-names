package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestData struct {
	Cookie string `json:"cookie"`
	Body   string `json:"body"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the JSON file as an argument")
	}

	filePath := os.Args[1]
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var requestData RequestData
	err = json.Unmarshal(fileContent, &requestData)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.parentsquare.com/sessions", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", requestData.Cookie)
	req.Body = ioutil.NopCloser(strings.NewReader(requestData.Body))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(respBody))
}
