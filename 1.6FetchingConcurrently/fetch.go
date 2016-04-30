package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	results, err := os.OpenFile("results.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer results.Close()

	if err != nil{
		fmt.Fprintf(os.Stderr, "Couldn't open results.txt %v", err)
		os.Exit(1)
	}

	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}

	for range os.Args[1:]{
		_, err := results.WriteString(<-ch)

		if err != nil{
			fmt.Fprintf(os.Stderr, "Couldn't write results to results.txt %v", err)
			os.Exit(1)
		}

		fmt.Fprintf(results, "%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}