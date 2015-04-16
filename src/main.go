package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"io/ioutil"

	"github.com/dmnlk/stringUtils"
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
	res := Response{}
	v, _ := json.Marshal(res)
	buf := bytes.NewBuffer(v)
	resp, err := http.Post(url+"?"+apikey, "application/json", buf)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := ioutil.ReadAll(resp)
	fmt.Println(string(a))
	return ""
}
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
