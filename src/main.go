package main

import (
	"fmt"
	"os"
	"github.com/dmnlk/stringUtils"
	"net/http"
	"io"
	"bytes"
)

const (
	API_URL=" https://www.googleapis.com/urlshortener/v1/url"
)
type Response struct {
	kind string `kind`
	id string `id`
	longUrl string `longUrl`

}

func main() {
	key, err := getGoogleAPIKey()
	if err != nil {
		return
	}
	fmt.Println(key)
}

func requestAPI(url string) string {
	buf := bytes.NewBuffer()
	http.Post(url, "application/json", buf)
	return ""
} 
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
