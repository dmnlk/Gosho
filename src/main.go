package main

import (
	"fmt"
	"os"
	"github.com/dmnlk/stringUtils"
)

func main() {
	fmt.Println("a")
}

func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
