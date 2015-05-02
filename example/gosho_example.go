package main

import (
	"os"

	"fmt"
	"log"

	"github.com/dmnlk/Gosho"
)

func main() {
	var originalUrl string = "http://golang-jp.org/"
	client := Gosho.NewClient()

	gapikey := getGoogleAPIKey()
	if len(gapikey) != 0 {
		surl, err := client.GetGoogleSUrl(originalUrl, gapikey)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("google short url:" + surl)
	}

	bapikey := getBitlyAPIKey()
	if len(bapikey) != 0 {
		surl, err := client.GetBitlySUrl(originalUrl, bapikey)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("bitly short url:" + surl)
	}

}

func getBitlyAPIKey() string {
	api_key := os.Getenv("BITLY_API_KEY")
	return api_key
}

func getGoogleAPIKey() string {
	api_key := os.Getenv("GOOGLE_API_KEY")

	return api_key
}
