package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Request struct {
	ID       uint   `json:"id"`
	Endpoint string `json:"endpoint"`
}

type Result struct {
	RequestID uint   `json:"requestId"`
	Result    string `json:"result"`
	Operator  string `json:"operator"`
}

// fetchRequests fetches pending requests from a server or contract
func fetchRequests(serverURL string) ([]Request, error) {
	resp, err := http.Get(serverURL + "/requests")
	if err != nil {
		return nil, fmt.Errorf("error fetching requests: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var requests []Request
	err = json.Unmarshal(body, &requests)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling requests: %v", err)
	}

	return requests, nil
}

// performPing checks if an endpoint is reachable using ICMP ping
func performPing(endpoint string) bool {
	cmd := exec.Command("ping", "-c", "1", endpoint)
	err := cmd.Run()
	return err == nil
}

// performTraceroute performs a traceroute to analyze the network path
func performTraceroute(endpoint string) string {
	cmd := exec.Command("traceroute", endpoint)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("error performing traceroute: %v", err)
	}
	return string(output)
}

// submitResult submits the operator's result back to the server or contract
func submitResult(serverURL string, result Result) error {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error marshalling result: %v", err)
	}

	resp, err := http.Post(serverURL+"/results", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error submitting result: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error from server: %s", resp.Status)
	}

	return nil
}

func main() {
	serverURL := "http://localhost:8080" // Replace with your server or contract's URL

	for {
		fmt.Println("Fetching pending requests...")
		requests, err := fetchRequests(serverURL)
		if err != nil {
			fmt.Printf("Error fetching requests: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, request := range requests {
			fmt.Printf("Processing request ID: %d, Endpoint: %s\n", request.ID, request.Endpoint)

			isAlive := performPing(request.Endpoint)
			var resultText string

			if isAlive {
				fmt.Printf("Endpoint %s is reachable.\n", request.Endpoint)
				resultText = "reachable"
			} else {
				fmt.Printf("Endpoint %s is not reachable.\n", request.Endpoint)
				resultText = "unreachable"
			}

			tracerouteOutput := performTraceroute(request.Endpoint)
			fmt.Println("Traceroute Result:")
			fmt.Println(tracerouteOutput)

			// Combine result text with traceroute for submission
			finalResult := fmt.Sprintf("%s\n\nTraceroute:\n%s", resultText, tracerouteOutput)

			// Submit result
			result := Result{
				RequestID: request.ID,
				Result:    finalResult,
				Operator:  "operator_address_here", // Replace with the actual operator's address
			}

			err := submitResult(serverURL, result)
			if err != nil {
				fmt.Printf("Error submitting result for request ID %d: %v\n", request.ID, err)
			} else {
				fmt.Printf("Result submitted successfully for request ID %d.\n", request.ID)
			}
		}

		fmt.Println("Waiting for the next cycle...")
		time.Sleep(30 * time.Second) // Adjust the polling interval as needed
	}
}
