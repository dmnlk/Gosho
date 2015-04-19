package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	"net/url"
	"strings"

	"github.com/dmnlk/stringUtils"
	"github.com/k0kubun/pp"
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
	requestAPI("http://twitter.com", key)
}

func requestAPI(urli string, apikey string) string {
	p := url.Values{}
	p.Add("longUrl", urli)
//	p.Add("api_key", apikey)
	req, err := http.NewRequest("POST", API_URL, strings.NewReader(p.Encode()))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	pp.Print(req)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("do")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	// ioutil.readallでバイト配列にする
	val, err := ioutil.ReadAll(resp.Body)

	// バイト配列を文字列にして表示する
	pp.Print(string(val))
	return ""
}
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
