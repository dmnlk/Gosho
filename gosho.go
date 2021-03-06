package Gosho

import (
	"bytes"
	"io/ioutil"

	"net/http"

	"encoding/json"
)

const (
	API_URL   = "https://www.googleapis.com/urlshortener/v1/url"
	BITLY_URL = "https://api-ssl.bitly.com/v3/shorten"
	UXNU_URL  = "http://ux.nu/api/short"
	NAZR_URL = "http://nazr.in/api/short_links"
)

type Client struct {
	GoogleApiKey string
	BitlyApiKey  string
	c* http.Client
}

func NewClient() Client {
	c := Client{}
	c.c = http.DefaultClient
	return c
}

type GoogleResponse struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type BitlyResponse struct {
	StatusCode int64  `json:status_code`
	StatusTxt  string `json:status_txt`
	Data       Data   `json:data`
}

type Data struct {
	LongUrl    string `json:"longUrl"`
	Url        string `json:"url"`
	Hash       string `json:"hash"`
	GlobalHash string `json:"global_hash"`
	NewHash    int64  `json:"new_hash"`
}

type UxnuResponse struct {
	StatusCode int64 `json:status_code`
	Data       Data  `json:data`
}

type NazrResponse struct {
	D62 string `json:"d62"`
	OriginalUrl string `json:"original_url"`
	CreatedAt string `json:"created_at"`
	Url string `json:"url"`

}

func (c Client) GetGoogleSUrl(originalUrl string) (string, error) {
	var jsonStr = []byte(`{"longUrl":"` + originalUrl + `"}`)
	req, err := http.NewRequest("POST", API_URL+"?key="+c.GoogleApiKey, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", err
	}


	resp, err := c.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)

	var res GoogleResponse
	json.Unmarshal(val, &res)

	return res.Id, nil
}

func (c Client) GetBitlySUrl(originalUrl string) (string, error) {
	req, err := http.NewRequest("GET", BITLY_URL+"?access_token="+c.BitlyApiKey+"&longUrl="+originalUrl, nil)
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
	var res BitlyResponse
	json.Unmarshal(val, &res)

	return res.Data.Url, nil
}

func (c Client) GetUxnuUrl(originalUrl string) (string, error) {
	req, err := http.NewRequest("GET", UXNU_URL+"?url="+originalUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	var res UxnuResponse
	json.Unmarshal(val, &res)

	return res.Data.Url, nil
}

func (c Client) GetNazrUrl(originalUrl string) (string, error) {
	req, err := http.NewRequest("POST", NAZR_URL+"?url="+originalUrl, nil)
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
	var res NazrResponse
	json.Unmarshal(val, &res)

	return res.Url, nil
}
