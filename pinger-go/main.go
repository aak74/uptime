package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

func MakeRequest(url string, ch chan<- string, i int) {
	start := time.Now()
	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	if resp == nil {
		ch <- fmt.Sprintf("%d task fail after %.2fs: %s", i, secs, url)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		ch <- fmt.Sprintf("%d task %.2fs elapsed with response length: %d %s", i, secs, len(body), url)
	}
}

func main() {
	var sites, err = GetUrls()
	// fmt.Print(sites, err)
	if err == nil {
		// fmt.Print(sites)
		// return fmt.Print(urls)
		start := time.Now()
		ch := make(chan string, 20)
		for i, url := range sites {
			go MakeRequest(url.URL, ch, i)
		}

		for i := range sites {
			res, _ := <-ch
			fmt.Println(fmt.Printf("%d/%d - %s", i+1, len(sites), res))
		}
		close(ch)

		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	}

}
