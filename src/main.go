package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	"net/url"

	"github.com/dmnlk/stringUtils"
	"github.com/k0kubun/pp"
	"strings"
)

const (
	API_URL = "https://www.googleapis.com/urlshortener/v1/url"
)

type Response struct {
	kind    string `kind`
	id      string `id`
	longUrl string `longUrl`
}

func main() {
	key, err := getGoogleAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	st, err := requestAPI("https://developers.google.com/url-shortener/v1/getting_started", key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(st)
}

func requestAPI(urli string, apikey string) (string, error) {
	p := url.Values{}
	p.Add("longUrl", urli)
	//p.Add("key", apikey)

	//resp, err := http.PostForm(API_URL+"?key="+apikey, p)

	req, err := http.NewRequest("POST", API_URL+"?key="+apikey, strings.NewReader(p.Encode()))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	pp.Print(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// ioutil.readallでバイト配列にする
	val, err := ioutil.ReadAll(resp.Body)

	// バイト配列を文字列にして表示する
	pp.Print(string(val))
	return string(val), nil
}
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
