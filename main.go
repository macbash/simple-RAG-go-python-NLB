package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

)

// Response function to provide an initial response
func Response() {
	fmt.Println("Hi! I'm your friendly chatbot. How can I assist you today?")
}

// Acknowledge function to acknowledge the user's input
func Acknowledge(userInput string) {
	fmt.Println("You mentioned:", userInput)
}

// Generate function to generate a reply based on user input and sentiment analysis
func Generate(userInput string) string {
	polarity, _, err := callPythonService(userInput)
	if err != nil {
		return "Error processing your input, please try again."
	}
	if polarity > 0 {
		return "That sounds positive! How can I assist you further?"
	} else if polarity < 0 {
		return "That seems a bit negative. Is there anything I can do to help?"
	}
	return "Thank you for your neutral input. How can I assist you?"
}

// Function to call the Python Flask service for sentiment analysis
func callPythonService(userInput string) (float64, float64, error) {
	payload := map[string]string{"text": userInput}
	jsonValue, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:5000/analyze_sentiment", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	var result map[string]float64
	json.NewDecoder(resp.Body).Decode(&result)
	return result["polarity"], result["subjectivity"], nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Start conversation
	Response()

	for {
		// Get user input
		fmt.Print("Enter your message: ")
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "exit" {
			break
		}

		// Process user input
		Acknowledge(userInput)
		response := Generate(userInput)
		fmt.Println("Chatbot:", response)
	}

	fmt.Println("Thank you for using our chatbot. Have a great day!")
}
