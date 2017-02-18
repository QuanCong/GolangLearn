package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"time"
)

func concurrentFetchUrl(url string, done chan string) error {
	start := time.Now()
	fmt.Println(url, "fetching...")

	response, err := http.Get(url)

	if err != nil {
		return err
	}
	fmt.Println(url, " response status:", response.Status)
	b, eread := ioutil.ReadAll(response.Body)
	if eread != nil {
		return eread
	}
	response.Body.Close()
	fmt.Printf("%s\n", b)

	cose := time.Since(start).Seconds()

	done <- fmt.Sprintf("fetching %s use %f second", url, cose)
	return nil
}

func fetchUrl(urls []string) error {

	for _, url := range urls {
		fmt.Println(url, "fetching...")

		response, err := http.Get(url)

		if err != nil {
			return err
		}
		fmt.Println(url, " response status:", response.Status)
		b, eread := ioutil.ReadAll(response.Body)
		if eread != nil {
			return eread
		}
		response.Body.Close()
		fmt.Printf("%s\n", b)
	}
	return nil
}

func main() {

	fetchUrl(os.Args[1:])

	urls := os.Args[1:]

	done := make(chan string, len(urls))

	for _, url := range urls {
		concurrentFetchUrl(url, done)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-done)
	}

}