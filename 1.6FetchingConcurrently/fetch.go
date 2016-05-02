package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	results, err := os.OpenFile("results.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer results.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open results.txt %v", err)
		os.Exit(1)
	}

	fetchFast(os.Args[1:], results)

}

func fetchFast(urls []string, out io.Writer) {
	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprint(out, <-ch)
		fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Problem fetching %s : %v", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	sizePart := fmt.Sprintf(sizeFormat, nbytes, url)
	ch <- fmt.Sprintf("%.2fs %s", secs, sizePart)

}

const sizeFormat = "%7d %s"
