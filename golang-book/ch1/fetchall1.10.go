package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // create a channel that received and sends strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from the channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// stole this from stackoverflow
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send the err to the channel
		return
	}
	hashedUrl := fmt.Sprint(hash(url)) // don't deal with naming
	f, err := os.Create(hashedUrl)     // create a file to write to
	if err != nil {
		ch <- fmt.Sprintf("Error creating a file for %s: %v", url, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	f.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
