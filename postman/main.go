package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
		"github.com/joho/godotenv"
		"os"
)

func main() {

		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error loading .env file")
		}

		token := os.Getenv("BEARER")
		retailer := os.Getenv("RETAILER")
		
		ids := []string{}

		for _, id := range ids {
			postIt(id, token, retailer)
		}

}

func postIt(id string, envVar string, retailer string) {
	url := "https://emmasleep.api.fluentretail.com/api/v4.1/event/async"

	requestBody := map[string]interface{}{
			"name":       "ConsignmentDelivered",
			"retailerId": retailer,
			"entityType": "CONSIGNMENT",
			"entityId":   id,
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
			fmt.Println("Error creating request:", err)
			return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + envVar)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			fmt.Println("Error sending request:", err)
			return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
			fmt.Println("Error reading response body:", err)
			return
	}

	fmt.Println("Response Status:", resp.Status)
	if resp.Status == "200" {
		fmt.Println("Successfully marked consigment as delivered!")
	} else {
		fmt.Println("Failed to update consignment.")
		fmt.Println("Response Body:", string(body))
	}

}
