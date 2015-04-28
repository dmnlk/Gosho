package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"

	"encoding/json"

	"github.com/dmnlk/stringUtils"
	"strings"
	"github.com/go-lang-plugin-org/go-lang-idea-plugin/testdata/parsing/statements/go"
)

const (
	API_URL = "https://www.googleapis.com/urlshortener/v1/url"
)

type GoogleResponse struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

// require argument
func main() {
	key, err := getGoogleAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	bitlyKey, err := getBitlyAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) == 1 {
		return
	}

	googleUrl, err := requestGoogleUrlShortenerApi(os.Args[1], key)
	if err != nil {
		fmt.Println(err)
		return
	}

	bitlyUrl, err := requestBitlyApi(os.Args[1], bitlyKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Google: "+googleUrl)
	fmt.Println("bit.ly: "+bitlyUrl)
}

func requestGoogleUrlShortenerApi(originalUrl string, apikey string) (string, error) {
	var jsonStr = []byte(`{"longUrl":"` + originalUrl + `"}`)
	req, err := http.NewRequest("POST", API_URL+"?key="+apikey, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)

	var res GoogleResponse
	json.Unmarshal(val, &res)

	return res.Id, nil
}

func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("api_key not found")
	}
	return api_key, nil
}

func requestBitlyApi(originalUrl string, apikey string) (string, error) {
	return "", nil
}



func getBitlyAPIKey() (string, error) {
	api_key := os.Getenv("BITLY_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("api_key not found")
	}
	return api_key, nil
}
