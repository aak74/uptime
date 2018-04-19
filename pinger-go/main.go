package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type site struct {
	ID string `json:"id"`
	// Domain string `json:"domain"`
	URL string `json:"url"`
	// Verifyssl bool   `json:"verifyssl"`
	// Timeout   int    `json:"timeout"`
}

func getUrls() (sites []site, err error) {
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
	var wg sync.WaitGroup
	sites, err := getUrls()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	start := time.Now()
	for i, url := range sites {
		wg.Add(1)
		go func(url string, i int) {
			start := time.Now()
			resp, _ := http.Get(url)
			secs := time.Since(start).Seconds()
			if resp == nil {
				fmt.Println(fmt.Sprintf("%d task fail after %.2fs: %s", i, secs, url))
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(fmt.Sprintf("%d task %.2fs elapsed with response length: %d %s", i, secs, len(body), url))
			}
			wg.Done()
		}(url.URL, i)
		wg.Wait()
	}

	// for i := range sites {
	// 	res, _ := <-ch
	// 	fmt.Println(fmt.Printf("%d/%d - %s", i+1, len(sites), res))
	// }
	// close(ch)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
