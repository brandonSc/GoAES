package main

//
// An implementation of the standard AES block cipher in Go.
// The intention of this project is to implement AES following NIST's publication FIPS-197
// Citation: http://csrc.nist.gov/publications/fips/fips197/fips-197.pdf
//

import (
	//"fmt"
	"github.com/brandonSc/GoAES/cipher"
)

func main() {
	plaintext := get_sample_text()
	key := get_sample_key()
	cipher.Cipher(plaintext, key)
}

func get_sample_key() [44][4]byte {
	var key [44][4]byte
	for i := 0; i < 44; i++ {
		for j := 0; j < 4; j++ {
			key[i][j] = byte(i * j)
		}
	}
	return key
}

func get_sample_text() [16]byte {
	// make something 128 bits long ..
	text := "Hello, world! Hello, world! Hello, world! Hello world!"
	textBytes := []byte(text)
	var message [16]byte
	for i := 0; i < 16; i++ {
		message[i] = textBytes[i]
	}
	return message
}
