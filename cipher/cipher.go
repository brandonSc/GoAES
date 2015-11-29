package cipher

import (
//"fmt"
)

//
// AES encrypt a block of data
// @param plaintext - the input data to encrypt
// @param key - key schedule as 2D byte-array
// @return the AES encrypted plaintext
//
func Encrypt(plaintext [128]byte, key [][]byte) [128][128]byte {
	var state [128][128]byte

	Nb := 4 // number of columns (32 bit words) [FIPS-197]
	//Nr := (len(state) / (Nb - 1)) // number of rounds (10, 12, or 24) [FIPS-197]

	for i := 0; i < 4*Nb; i++ {
		state[i%4][i/4] = plaintext[i]
	}

	return state
}
