package main

import (
	"io"
	"net/http"
	"fmt"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		streamURL(url, os.Stdout)
	}
}

func streamURL(url string, out io.Writer){
	resp, err := http.Get(url)

	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(out, resp.Body)

	resp.Body.Close()

	if err != nil{
		fmt.Fprintf(os.Stderr, "Problem writing response body to stdout %v\n", err)
	}
}
