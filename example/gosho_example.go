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
		client.GoogleApiKey = gapikey
		surl, err := client.GetGoogleSUrl(originalUrl)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("google short url:" + surl)
	}

	bapikey := getBitlyAPIKey()
	if len(bapikey) != 0 {
		client.BitlyApiKey = bapikey
		surl, err := client.GetBitlySUrl(originalUrl)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("bitly short url:" + surl)
	}

	surl, err := client.GetUxnuUrl(originalUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("ux.nushort url:" + surl)

}

func getBitlyAPIKey() string {
	api_key := os.Getenv("BITLY_API_KEY")
	return api_key
}

func getGoogleAPIKey() string {
	api_key := os.Getenv("GOOGLE_API_KEY")

	return api_key
}
