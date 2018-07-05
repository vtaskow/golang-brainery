package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Args)

	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("%.6fs\n", secs)

	start = time.Now()
	for i, val := range os.Args {
		fmt.Println(i, val)
	}
	secs = time.Since(start).Seconds()
	fmt.Printf("%.6fs\n", secs)
}
