package cipher

import (
	//"fmt"
	"github.com/brandonSc/GoAES/commons"
)

//
// AES decrypt a block of data - FIPS-197, Section 5.3
// @param input - the input data to decrypt
// @param key - key schedule as 2D byte-array
// @return the AES decrypted ciphertext
//
func InvCipher(input [4 * Nb]byte, word [Nb * (Nr + 1)][4]byte) [4 * Nb]byte {
	var state [4][Nb]byte

	// state = input
	for i := 0; i < 4*Nb; i++ {
		state[i%4][i/4] = input[i]
	}

	state = commons.AddRoundKey(state, word[Nr*Nb:((Nr+1)*Nb)])

	for rnd := Nr; rnd > 0; rnd-- {
		state = InvShiftRows(state)
		state = InvSubBytes(state)
		state = commons.AddRoundKey(state, word[rnd*Nb:((rnd+1)*Nb)])
		state = InvMixColumns(state)
	}

	state = InvSubBytes(state)
	state = InvShiftRows(state)
	state = commons.AddRoundKey(state, word[(Nr*Nb):((Nr+1)*Nb)])

	var output [4 * Nb]byte
	return output
}

// FIPS-197 section 5.3.2.
func InvSubBytes(state [4][Nb]byte) [4][Nb]byte {
	for r := 0; r < 4; r++ {
		for c := 0; c < Nb; c++ {
			state[r][c] = commons.SBox[state[r][c]]
		}
	}
	return state
}

// FIPS-197 section 5.3.1.
func InvShiftRows(state [4][Nb]byte) [4][Nb]byte {
	var temp [4]byte
	for r := 1; r < 4; r++ {
		for c := 0; c < 4; c++ {
			temp[c] = state[r][(c+r)%Nb]
		}
		for c := 0; c < 4; c++ {
			state[r][c] = temp[c]
		}
	}
	return state
}

// FIPS-197 section 5.3.3.
func InvMixColumns(state [4][Nb]byte) [4][Nb]byte {
	for c := 0; c < 4; c++ {
		var a [128]byte
		var b [128]byte
		for i := 0; i < 4; i++ {
			a[i] = state[i][c]
			// should further understand the significance of next line ..
			if (state[i][c] & 0x80) == 1 {
				b[i] = state[i][c] << (1 ^ 0x011b)
			} else {
				b[i] = state[i][c] << 1
			}
		}
		// FIPS-197 Section 4.3
		state[0][c] = b[0] ^ a[1] ^ b[1] ^ a[2] ^ a[3]
		state[1][c] = a[0] ^ b[1] ^ a[2] ^ b[2] ^ a[3]
		state[2][c] = a[0] ^ a[1] ^ b[2] ^ a[3] ^ b[3]
		state[3][c] = a[0] ^ b[0] ^ a[1] ^ a[2] ^ b[3]
	}
	return state
}
