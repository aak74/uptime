package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	if resp == nil {
		fmt.Print(resp)
		ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, 0, url)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
	}
}

type Site struct {
	ID string `json:"id"`
	// Domain string `json:"domain"`
	URL string `json:"url"`
	// Verifyssl bool   `json:"verifyssl"`
	// Timeout   int    `json:"timeout"`
}

func GetUrls() (sites []Site, err error) {
	resp, err := http.Get("http://127.0.0.1/api/v1/projects")
	// fmt.Println("resp:", err)
	if err == nil {

		jsonBlob, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(jsonBlob, &sites)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	// fmt.Printf("%+v", sites)
	return
}

func main() {
	var sites, err = GetUrls()
	// fmt.Print(sites, err)
	if err == nil {
		// fmt.Print(sites)
		// return fmt.Print(urls)
		start := time.Now()
		ch := make(chan string)
		for _, url := range sites {
			go MakeRequest(url.URL, ch)
		}

		for range sites {
			fmt.Println(<-ch)
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	}

}
