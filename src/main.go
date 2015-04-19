package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dmnlk/stringUtils"
	"github.com/k0kubun/pp"
)

const (
	API_URL = " https://www.googleapis.com/urlshortener/v1/url"
)

type Response struct {
	kind    string `kind`
	id      string `id`
	longUrl string `longUrl`
}

func main() {
	key, err := getGoogleAPIKey()
	if err != nil {
		return
	}
	requestAPI("http://twitter.com", key)
}

func requestAPI(url string, apikey string) string {

	resp, err := http.Post(url+"?"+apikey, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("a")
	pp.Print(resp)
	return ""
}
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
