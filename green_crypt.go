package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	file_ptr := flag.String("f","", "The name of the file to be encrypted")
	dec_ptr := flag.Bool("d", false, `Flag to indicate if flag is being decrypted
	 default (false) means that file is to be encrypted`)	
	key_ptr := flag.String("k", "secret", "The key to be used to encrypt the file")

	flag.Parse()

	file, err := os.Open(*file_ptr)
	if (err != nil) {
		panic(err)
	}

	br := bufio.NewReader(file)

	var new_bits []byte
	for {
		b, err := br.ReadByte() 
		
		if (err != nil) {
			break
		}
		
		append(new_bits, b + ind)
		ind = (ind + 1) % len(*key_ptr)
		fmt.Printf("%c", b)
		fmt.Printf("%d", b + ind)
	}
	
	os.WriteFile(*file_ptr, new_bits, 0644)
}
