package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ // autopopulated on first occurrence to be 0
	}
	fmt.Println("DUPLICATE LINES")
	for line, n := range counts {
		if n > 0 {
			fmt.Println(line, n)
		}
	}

}
