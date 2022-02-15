package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	//Get the flags decide if we encode or decode
	option := flag.Bool("d", false, "will encode by default")
	flag.Parse()

	//receive string from stdin
	scanner := bufio.NewScanner(os.Stdin)
	if *option {
		//if set run the decode function
		decode(*scanner)
	} else {
		//encode by default
		encode(*scanner)
	}

}
func encode(scanner bufio.Scanner) {

	for scanner.Scan() {
		query := scanner.Text()
		encoded := url.QueryEscape(query)
		fmt.Println("url encoded\n", encoded)
		double_encoded := url.QueryEscape(encoded)
		fmt.Println("double encoded\n", double_encoded)
	}
}
func decode(scanner bufio.Scanner) {
	for scanner.Scan() {
		query := scanner.Text()
		decoded, _ := url.QueryUnescape(query)

		fmt.Println(decoded)

	}
}
