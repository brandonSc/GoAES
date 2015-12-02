package main

//
// An implementation of the standard AES block cipher in Go.
// The intention of this project is to implement AES following NIST's publication FIPS-197
// Citation: http://csrc.nist.gov/publications/fips/fips197/fips-197.pdf
//

import (
	"github.com/brandonSc/GoAES/cipher"
)

func main() {
	plaintext := getSampleText()
	cipher.Encrypt(plaintext, nil)
}

func getSampleText() [128]byte {
	// make something 128 bits long ..
	text := "Hello, world! Hello, world! Hello, world! Hello world! Hello world! Hello world! Hello world! Hello world! Hello world! Hello world!Hello world! Hello world! Hello world!"
	textBytes := []byte(text)
	var message [128]byte
	for i := 0; i < 128; i++ {
		message[i] = textBytes[i]
	}
	return message
}
