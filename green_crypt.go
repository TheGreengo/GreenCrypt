package main

import (
	"flag"
	"os"
)

func main() {
	file_ptr := flag.String("f","", "The name of the file to be encrypted")
	out_ptr := flag.String("o","", "The name of the output file to be made")
	dec_ptr := flag.Bool("d", false, `Flag to indicate if flag is being decrypted
	 default (false) means that file is to be encrypted`)	
	key_ptr := flag.String("k", "secret", "The key to be used to encrypt the file")
    chunk_ptr := flag.Int("c", 10000, "The chunk size used when encrypting/decrypting")

	flag.Parse()

    out_match := false
    if (*out_ptr == "") {
        temp := *file_ptr
        out_ptr = &temp
        *out_ptr += ".temp"
        out_match = true
    }

	file, err := os.Open(*file_ptr)
	if (err != nil) {
		panic(err)
	}
    
    f2, errr := os.Create(*out_ptr) 
	if (errr != nil) {
		panic(errr)
	}
	
    defer f2.Close()
	
    key := *key_ptr

    bits := make([]byte,*chunk_ptr)
	ind := 0
	for {
		n, err := file.Read(bits) 
		
		if (err != nil) {
			break
		}
		
        for i := 0; i < n; i++ {
		    if (*dec_ptr) {
			    bits[i] = bits[i] - key[ind]
		    } else {
			    bits[i] = bits[i] + key[ind]
		    }
		    ind = (ind + 1) % len(key)
        }

        f2.Write(bits[:n])
	}
	
    if (out_match) {
        os.Remove(*file_ptr)
        os.Rename(*out_ptr, *file_ptr)
    }
	
	file.Close()
}
