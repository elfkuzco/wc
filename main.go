package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	nw, err := count(os.Stdin, *lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nw)
}

func count(r io.Reader, countLines bool) (int, error) {
	// Set up a scanner to read text
	scanner := bufio.NewScanner(r)
	if !countLines {
		// Define the scanner split type to words
		// (default is split by lines)
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wc, nil
}
