package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	"net/url"

	"encoding/json"

	"github.com/dmnlk/stringUtils"
	"github.com/k0kubun/pp"
)

const (
	API_URL = "https://www.googleapis.com/urlshortener/v1/url"
)

type GoogleResponse struct {
	kind    string `json:"kind"`
	id      string `json:"id"`
	longUrl string `json:"longUrl"`
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
	var jsonStr = []byte(`{"longUrl":"` + urli + `"}`)
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
	// ioutil.readallでバイト配列にする
	val, err := ioutil.ReadAll(resp.Body)

	var res GoogleResponse
	json.Unmarshal(val, &res)
	pp.Print(res)
	pp.Print(string(val))

	var result interface{}
	json.Unmarshal(val, &result)
	pp.Print(result)

	b := json.NewDecoder(val)
	b.Decode(&res)

	pp.Print(res)
	// バイト配列を文字列にして表示する
	return res.id, nil
}
func getGoogleAPIKey() (string, error) {
	api_key := os.Getenv("GOOGLE_API_KEY")
	if stringUtils.IsEmpty(api_key) {
		return "", fmt.Errorf("not found api_key")
	}
	return api_key, nil
}
