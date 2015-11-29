package main

//
// An implementation of the standard AES block cipher in Go.
// The intention of this project is to implement AES following NIST's publication FIPS-197
// Citation here: http://csrc.nist.gov/publications/fips/fips197/fips-197.pdf
//

import (
	"fmt"
	"github.com/brandonSc/GoAES/cipher"
)

func main() {
	// something 128 bits long ..
	text := "Hello, world! Hello, world! Hello, world! Hello world! Hello world! Hello world! Hello world! Hello world! Hello world! Hello world!Hello world! Hello world! Hello world!"
	textBytes := []byte(text)
	fmt.Printf("INPUT LENGTH: %d\n", len(textBytes))
	var message [128]byte
	for i := 0; i < 128; i++ {
		message[i] = textBytes[i]
	}
	cipher.Encrypt(message, nil)
}
