package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RequestData struct {
	Cookie string `json:"cookie"`
	Body   string `json:"body"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run app.go <path_to_json_file>")
		return
	}

	filePath := os.Args[1]

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var requestData RequestData
	err = json.Unmarshal(fileContent, &requestData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.parentsquare.com/sessions", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", requestData.Cookie)
	req.Body = ioutil.NopCloser(strings.NewReader(requestData.Body))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response body:", string(responseBody))
}
