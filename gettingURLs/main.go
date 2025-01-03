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
	start := time.Now()
	ch := make(chan string)
	file, errors := os.Open(os.Args[1])
	fmt.Println("File is opened...")
	if errors != nil {
		fmt.Println(errors)
		return
	}

	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}

	for range os.Args[2:] {
		canal := <-ch
		fmt.Println("Chanal: ...", canal)
		file.WriteString(<-ch)
	}
	defer file.Close()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
