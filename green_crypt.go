package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	file_ptr := flag.String("f","", "The name of the file to be encrypted")
	// dec_ptr := flag.Bool("d", false, `Flag to indicate if flag is being decrypted 
	// default (false) means that file is to be encrypted`)	

	flag.Parse()

	file, err := os.Open(*file_ptr)
	if (err != nil) {
		panic(err)
	}

	br := bufio.NewReader(file)

	for {
		b, err := br.ReadByte() 
		
		if (err != nil) {
			break
		}

		fmt.Printf("%c", b)
	}
}
