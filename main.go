package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type result struct {
	nlines   int
	nwords   int
	nbytes   int
	lastRune rune // last rune that was read
}

func main() {
	nlines := flag.Bool("l", false, "Count lines")
	nbytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()
	rs, err := count(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if *nlines {
		fmt.Println(rs.nlines)
	} else if *nbytes {
		fmt.Println(rs.nbytes)
	} else {
		fmt.Println(rs.nlines, rs.nwords, rs.nbytes)
	}
}

func count(r io.Reader) (result, error) {
	rs := result{}
	// Set up a reader to read text
	reader := bufio.NewReader(r)
	// Set the scanner to split by bytes

	for {
		b, err := reader.ReadByte()
		ch := rune(b)

		if err != nil {
			if err == io.EOF {
				// If we get to the end of the line and the last
				// character read is not a space character,
				// increment the nwords and nlines accordingly
				if !unicode.IsSpace(rs.lastRune) {
					rs.nlines++
					rs.nwords++
				}
				break
			} else {
				return rs, err
			}
		} else {
			rs.nbytes++
			// If the current rune is a newline character,
			// increment the number of lines
			if ch == '\n' {
				rs.nlines++
			}
			// If the current rune is a space char
			// and the previous character is not a space, increment
			// the number of words
			if unicode.IsSpace(ch) && !unicode.IsSpace(rs.lastRune) {
				rs.nwords++
			}
			// Maintain a reference to the last rune that was read
			rs.lastRune = ch
		}
	}
	return rs, nil
}
