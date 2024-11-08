package aes128

func keyExpansion(key []byte) [][]byte {
	const (
		nk = 4  // Número de palavras na chave (4 para AES-128)
		nr = 10 // Número de rodadas (10 para AES-128)
	)

	roundKeys := make([][]byte, nr+1)
	for i := range roundKeys {
		roundKeys[i] = make([]byte, 16)
	}

	// A chave inicial é a primeira chave de rodada
	copy(roundKeys[0], key)
	for i := 1; i <= nr; i++ {
		// Últimos 4 bytes da rodada anterior
		tmp := roundKeys[i-1][12:16]

		// RotWord, SubWord e XOR com rcon
		tmp = append(tmp[1:], tmp[0]) // RotWord

		// SubWord
		tmp[0] = sbox[tmp[0]]
		tmp[1] = sbox[tmp[1]]
		tmp[2] = sbox[tmp[2]]
		tmp[3] = sbox[tmp[3]]

		tmp[0] ^= rcon[i-1]

		// Gera os 16 bytes da chave de rodada
		for j := 0; j < 16; j++ {
			if j < 4 {
				roundKeys[i][j] = roundKeys[i-1][j] ^ tmp[j]
			} else {
				roundKeys[i][j] = roundKeys[i-1][j] ^ roundKeys[i][j-4]
			}
		}
	}
	return roundKeys
}

func addRoundKey(state, roundKey []byte) {
	for i := 0; i < len(state); i++ {
		state[i] ^= roundKey[i]
	}
}
