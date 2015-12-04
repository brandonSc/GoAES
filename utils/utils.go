package utils

func XorWord(w1 [4]byte, w2 [4]byte) [4]byte {
	var w [4]byte
	for i := range w {
		w[i] = w1[i] ^ w2[i]
	}
	return w
}
