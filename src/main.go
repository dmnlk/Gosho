package main

import (
	"fmt"
	"os"
	"github.com/dmnlk/stringUtils"
)

const (
	API_URL=" https://www.googleapis.com/urlshortener/v1/url"
)

func main() {
	key, err := getGoogleAPIKey()
	if err != nil {
		return
	}
	fmt.Println(key)
}

func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
