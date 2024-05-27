package main

import (
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
	
	key := *key_ptr

	br := bufio.NewReader(file)

	var new_bits []byte
	var old_bits []byte
	ind := 0
	for {
		b, err := br.ReadByte() 
		
		if (err != nil) {
			break
		}
		
		if (*dec_ptr) {
			new_bits = append(new_bits, b - key[ind])
			old_bits = append(old_bits, b)
			ind = (ind + 1) % len(key)
		} else {
			new_bits = append(new_bits, b + key[ind])
			old_bits = append(old_bits, b)
			ind = (ind + 1) % len(key)
		}
	}
	
	os.WriteFile("test_out.txt", new_bits, 0644)
	
	file.Close()
}
