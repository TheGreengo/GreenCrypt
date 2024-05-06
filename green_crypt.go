package main

import (
	"fmt"
	"flag"
)

func main() {
	file_ptr := flag.String("f","", "The name of the file to be encrypted")
	dec_ptr := flag.Bool("d", false, `Flag to indicate if flag is being decrypted 
	default (false) means that file is to be encrypted`)	

	flag.Parse()

	fmt.Println("File name:", *file_ptr, "\nDecrypt:", *dec_ptr)	
}
