package cipher

import (
	"fmt"
	"math"
	// "block-cipher-webapp/backend/helper"
)

func xorOperationBlock(block1, block2 []byte) []byte {
	// Perform XOR operation between two blocks
	messageBlocks := make([]byte, len(block1))
	for i := 0; i < len(block1); i++ {
		messageBlocks[i] = block1[i] ^ block2[i]
	}
	return messageBlocks
}

func SubBytesSubstitutionArr(arr []int) []string {
	// Perform Sub-Bytes Substitution for an array of values
	messageBlocks := make([]string, len(arr))
	for i, val := range arr {
		// messageBlocks[i] = SubBytesSubstitution(val)
		messageBlocks[i] = fmt.Sprintf("0x%X", SubBytesSubstitution(byte(val)))
	}
	return messageBlocks
}

func InverseSubBytesSubstitutionArr(arr []int) []string {
	// Perform Inverse Sub-Bytes Substitution for an array of values
	messageBlocks := make([]string, len(arr))
	for i, val := range arr {
		// messageBlocks[i] = InverseSubBytesSubstitution(val)
		messageBlocks[i] = fmt.Sprintf("0x%X", InverseSubBytesSubstitution(byte(val)))
	}
	return messageBlocks
}

func SubBytesSubstitution(value byte) byte {
	// Substitusi Baris x Kolom
	// Contoh: 2F (ganti dengan nilai dari tabel S-Box pada baris 2 dan kolom F)
	sbox := [][]byte{
		// 0     1     2     3     4     5     6     7     8     9     A     B     C     D     E     F
		{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76}, // 0
		{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0}, // 1
		{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15}, // 2
		{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75}, // 3
		{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84}, // 4
		{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf}, // 5
		{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8}, // 6
		{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2}, // 7
		{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73}, // 8
		{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb}, // 9
		{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79}, // A
		{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08}, // B
		{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a}, // C
		{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e}, // D
		{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf}, // E
		{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16}, // F
	}

	return sbox[value>>4][value&0x0F]
}

func InverseSubBytesSubstitution(value byte) byte {
	sboxInverse := [][]byte{
		// 0     1     2     3     4     5     6     7     8     9     A     B     C     D     E     F
		{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb}, // 0
		{0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb}, // 1
		{0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e}, // 2
		{0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25}, // 3
		{0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92}, // 4
		{0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84}, // 5
		{0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06}, // 6
		{0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b}, // 7
		{0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73}, // 8
		{0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e}, // 9
		{0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b}, // A
		{0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4}, // B
		{0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f}, // C
		{0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef}, // D
		{0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61}, // E
		{0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d}, // F
	}
	return sboxInverse[value>>4][value&0x0F]
}

func ShiftRows(state [][]int) [][]int {
	// Shift a matrix, with first row unchanged, second row shifted 1 to the left, third row shifted 2 to the left, and fourth row shifted 3 to the left, etc
	tempState := make([][]int, len(state))
	for i := range state {
		tempState[i] = make([]int, len(state[i]))
	}
	for i := range state { // Row
		for j := range state[i] { // Column
			maxShift := len(state) - i
			if j > maxShift-1 {
				fmt.Printf("IF: Filling new table %d%d with contents from old table %d%d\n", i, j, i, j-maxShift)
				tempState[i][j] = state[i][j-maxShift]
			} else {
				fmt.Printf("ELSE: Filling new table %d%d with contents from old table %d%d\n", i, j, i, j+i)
				tempState[i][j] = state[i][j+i]
			}
		}
	}
	return tempState
}

func InverseShiftRows(state [][]int) [][]int {
	// Shift a matrix, with first row unchanged, second row shifted 1 to the right, third row shifted 2 to the right, and fourth row shifted 3 to the right, etc
	tempState := make([][]int, len(state))
	for i := range state {
		tempState[i] = make([]int, len(state[i]))
	}
	for i := range state { // Row
		for j := range state[i] { // Column
			maxShift := len(state) - i
			if j < i {
				fmt.Printf("IF: Filling new table %d%d with contents from old table %d%d\n", i, j, i, j+maxShift)
				tempState[i][j] = state[i][j+maxShift]
			} else {
				fmt.Printf("ELSE: Filling new table %d%d with contents from old table %d%d\n", i, j, i, j-i)
				tempState[i][j] = state[i][j-i]
			}
		}
	}
	return tempState
}

func subKeyGenerator(key []byte) []byte {
	// Generates a key from the original key as seed
	// Scramble key from ABCDEF to BADCFE
	newKey := make([]byte, len(key))
	for i := 0; i < len(key); i += 2 {
		newKey[i] = key[(i+2)%len(key)]
		newKey[i+1] = key[i]
	}
	// Do left shift 1 bit for even index, 2 bits for odd index, and XOR with next index
	for i := 0; i < len(key); i++ {
		if i%2 == 0 {
			newKey[i] = newKey[i]<<1 ^ newKey[(i+1)%len(key)]
		} else {
			newKey[i] = newKey[i]<<2 ^ newKey[(i+1)%len(key)]
		}
	}

	return newKey
}

func subKeysGenerator(key []byte, numOfRounds int) [][]byte {
	// Converts a 128-bit into an array of subkeys for each round
	// Create empty array to store subkeys
	keyArr := make([][]byte, numOfRounds)

	// Generate subkeys for each round with subKeyGenerator function, the previous key becomes the seed for the next one
	for i := 0; i < numOfRounds; i++ {
		if i == 0 {
			keyArr[i] = subKeyGenerator(key)
		} else {
			keyArr[i] = subKeyGenerator(keyArr[i-1])
		}
	}

	fmt.Println("keyArr", keyArr)

	return keyArr
}

func formatMessageIntoBlocks(message string) [][]byte {
	// TODO: handle cases where padding is not needed
	// Convert message and key to hexadecimal array
	messageHex := []byte(message)

	// Add padding of nulls to message so it can be divided to 16-byte (128-bit) blocks
	num_of_padding := 16 - len(messageHex)%16
	for i := 0; i < num_of_padding; i++ {
		messageHex = append(messageHex, 0)
	}

	// Split the message into 16-byte (128-bit) blocks (spek)
	messageBlocks := make([][]byte, int(math.Ceil(float64(len(messageHex))/16)))
	for i := 0; i < len(messageHex); i += 16 {
		messageBlocks[i/16] = make([]byte, 16)
		for j := 0; j < 16; j++ {
			messageBlocks[i/16][j] = messageHex[i+j]
		}
	}
	return messageBlocks
}

func formatKeyInto128Bit(key string) []byte {
	// Only take the first 128 bits of the key. If the key is less than 128 bits, pad it with 0s
	keyHex := []byte(key)
	if len(keyHex) < 16 {
		keyHex = append(keyHex, make([]byte, 16-len(keyHex))...)
	}
	return keyHex[:16]
}

func oneRoundEncryption(messageBlock, key []byte) []byte {
	// TODO: IDK WHAT I'M DOING IN THIS ONE, DO CHANGE IT TO A BETTER ALGO

	// XOR the round key with the right half of the block
	for i := 8; i < 16; i++ {
		messageBlock[i] = messageBlock[i] ^ key[i]
	}

	// Pass the right half of the block through the S-box
	for i := 8; i < 16; i++ {
		messageBlock[i] = SubBytesSubstitution(messageBlock[i])
	}

	// Pass the right half of the block through the P-box

	// XOR the left half of the block with the right half of the block
	for i := 0; i < 8; i++ {
		messageBlock[i] = messageBlock[i] ^ messageBlock[i+8]
	}

	// Swap the left and right halves of the block
	for i := 0; i < 8; i++ {
		messageBlock[i], messageBlock[i+8] = messageBlock[i+8], messageBlock[i]
	}

	return messageBlock
}

func oneRoundDecryption(messageBlock, key []byte) []byte {
	// TODO: IDK WHAT I'M DOING IN THIS ONE, DO CHANGE IT TO A BETTER ALGO

	// Swap the left and right halves of the block
	for i := 0; i < 8; i++ {
		messageBlock[i], messageBlock[i+8] = messageBlock[i+8], messageBlock[i]
	}

	// XOR the left half of the block with the right half of the block
	for i := 0; i < 8; i++ {
		messageBlock[i] = messageBlock[i] ^ messageBlock[i+8]
	}

	// Pass the right half of the block through the P-box

	// Pass the right half of the block through the S-box
	for i := 8; i < 16; i++ {
		messageBlock[i] = InverseSubBytesSubstitution(messageBlock[i])
	}

	// XOR the round key with the right half of the block
	for i := 8; i < 16; i++ {
		messageBlock[i] = messageBlock[i] ^ key[i]
	}

	return messageBlock
}

func ecb(messageBlocks [][]byte, key []byte, mode string) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting ECB.")

	// Encrypt/Decrypt each block
	for i, block := range messageBlocks {
		if mode == "encrypt" {
			messageBlocks[i] = oneRoundEncryption(block, key)
		} else {
			messageBlocks[i] = oneRoundDecryption(block, key)
		}
	}

	return messageBlocks
}

func cbc(messageBlocks [][]byte, key []byte, mode string) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting CBC.")

	// Encrypt/Decrypt each block
	// Random IV
	iv := []byte{98, 170, 137, 30, 192, 246, 212, 94, 116, 58, 128, 119, 118, 73, 72, 128}
	for i, block := range messageBlocks {
		if mode == "encrypt" {
			messageBlocks[i] = xorOperationBlock(block, iv)
			messageBlocks[i] = oneRoundEncryption(messageBlocks[i], key)
			iv = messageBlocks[i]
		} else {
			messageBlocks[i] = oneRoundDecryption(block, key)
			messageBlocks[i] = xorOperationBlock(block, iv)
			iv = block
		}
	}

	return messageBlocks
}

func ofb(messageBlocks [][]byte, key []byte, mode string) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting CFB.")

	result := make([][]byte, len(messageBlocks))

	// This is an 8-bit OFB
	// Random IV
	shiftRegister := []byte{62, 52, 12, 66, 21, 82, 112, 173, 92, 216, 252, 222, 2, 82, 11, 97}
	// TODO: I think this should be better written as a nested loop but at the time of writing, my brain can only work with a single loop
	for i := 0; i < len(messageBlocks)*len(messageBlocks[0]); i++ {
		// Encrypt the shift register (with CFB, E = D)
		output := oneRoundEncryption(shiftRegister, key)
		// Take the leftmost byte of the output
		leftMostByte := output[0]
		// First byte of the ciphertext is the XOR of the leftmost byte of the output and the first byte of the plaintext
		result[i/16] = append(result[i/16], messageBlocks[i/16][i%16]^leftMostByte)
		// Shift the shift register to the left by 1 bit
		shiftRegister = append(shiftRegister[1:], output[0])
	}
	return result
}

func cfb(messageBlocks [][]byte, key []byte, mode string) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting CFB.")

	result := make([][]byte, len(messageBlocks))

	// This is an 8-bit CFB
	// Random IV
	shiftRegister := []byte{12, 214, 112, 31, 45, 61, 83, 72, 115, 221, 75, 53, 251, 79, 27, 183}
	// TODO: I think this should be better written as a nested loop but at the time of writing, my brain can only work with a single loop
	for i := 0; i < len(messageBlocks)*len(messageBlocks[0]); i++ {
		// Encrypt the shift register (with CFB, E = D)
		output := oneRoundEncryption(shiftRegister, key)
		// Take the leftmost byte of the output
		leftMostByte := output[0]
		// First byte of the ciphertext is the XOR of the leftmost byte of the output and the first byte of the plaintext
		result[i/16] = append(result[i/16], messageBlocks[i/16][i%16]^leftMostByte)
		// Shift the shift register to the left by 1 bit
		if mode == "encrypt" {
			shiftRegister = append(shiftRegister[1:], result[i/16][i%16])
		} else {
			shiftRegister = append(shiftRegister[1:], messageBlocks[i/16][i%16])
		}
	}

	return result
}

func counter(messageBlocks [][]byte, key []byte, mode string) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks

	// TODO: NotImplemented

	return messageBlocks
}

// Process encrypt/decrypt (main function)
func GoBlockC(message, key, encryptOrDecrypt, mode string) string {
	fmt.Println("Encrypting plaintext", message, "with key", key, "using mode", mode)

	// Split message into blocks
	messageBlocks := formatMessageIntoBlocks(message)

	// Convert key to hexadecimal array
	keyHex := formatKeyInto128Bit(key)

	result := make([][]byte, len(messageBlocks))
	// TODO: Change to switch case
	if mode == "ecb" {
		result = ecb(messageBlocks, keyHex, encryptOrDecrypt)
	} else if mode == "cbc" {
		result = cbc(messageBlocks, keyHex, encryptOrDecrypt)
	} else if mode == "ofb" {
		result = ofb(messageBlocks, keyHex, encryptOrDecrypt)
	} else if mode == "cfb" {
		result = cfb(messageBlocks, keyHex, encryptOrDecrypt)
	} else if mode == "ctr" {
		result = counter(messageBlocks, keyHex, encryptOrDecrypt)
	}

	fmt.Println("Result: ", result)

	resultString := ""
	for _, block := range result {
		for _, b := range block {
			resultString += fmt.Sprintf("%02X", b)
		}
	}
	return resultString
}
