package main

import (
	"encoding/json"
	"fmt"

	grafana "github.com/AngelLozan/grafana/data"
)

func extractOrderRefs(jsonData []byte) ([]string, error) {
	// Define a struct to represent the JSON object
	type LogEntry struct {
		Line string `json:"line"`
	}

	// Define a slice to store the extracted order refs
	var orderRefs []string

	// Unmarshal the JSON array into a slice of LogEntry structs
	var logs []LogEntry
	if err := json.Unmarshal(jsonData, &logs); err != nil {
		return nil, err
	}

	// Iterate over the LogEntry objects and extract the orderRef value
	for _, log := range logs {
		var entry map[string]interface{}
		if err := json.Unmarshal([]byte(log.Line), &entry); err != nil {
			return nil, err
		}
		if orderRef, ok := entry["orderRef"].(string); ok {
			orderRefs = append(orderRefs, orderRef)
		}
	}

	return orderRefs, nil
}

func main() {
	// JSON data representing an array of objects
	jsonData := []byte(grafana.JsonData)

	// Extract order refs from the JSON data
	orderRefs, err := extractOrderRefs(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the extracted order refs
	fmt.Println("Extracted order refs: #", len(orderRefs))

	for _, ref := range orderRefs {
		fmt.Println(ref)
	}
}
