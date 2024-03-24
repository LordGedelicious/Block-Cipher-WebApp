package cipher

import (
	"block-cipher-webapp/backend/goblockc"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"time"
)

func xorOperationBlock(block1, block2 []byte) []byte {
	// Perform XOR operation between two blocks
	messageBlocks := make([]byte, len(block1))
	for i := 0; i < len(block1); i++ {
		messageBlocks[i] = block1[i] ^ block2[i]
	}
	return messageBlocks
}

func TextToHex(text string) string {
	hexString := ""
	for _, c := range text {
		hexString += string(hex.EncodeToString([]byte(string(c))))
	}
	return hexString
}

func HexToText(hexString string) string {
	text := ""
	hexString = strings.ReplaceAll(hexString, " ", "")
	for i := 0; i < len(hexString); i += 2 {
		hexByte, _ := hex.DecodeString(hexString[i : i+2])
		text += string(hexByte)
	}
	return text
}

func formatMessageIntoBlocks(messageHex []byte) [][]byte {
	// TODO: handle cases where padding is not needed
	// Convert message and key to hexadecimal array

	// Add padding of nulls to message so it can be divided to 16-byte (128-bit) blocks
	var num_of_padding int
	if len(messageHex)%16 != 0 {
		num_of_padding = 16 - len(messageHex)%16
	}
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

func ecb(messageBlocks [][]byte, key []byte, isEncrypt bool) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting ECB.")
	var error error

	// Encrypt/Decrypt each block
	for i, block := range messageBlocks {
		messageBlocks[i], error = goblockc.Parse(block, key, isEncrypt)
		if error != nil {
			fmt.Println("Error in ECB.")
		}
	}

	return messageBlocks
}

func cbc(messageBlocks [][]byte, key []byte, isEncrypt bool) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting CBC.")
	var error error

	// Encrypt/Decrypt each block
	// Random IV
	iv := []byte{98, 170, 137, 30, 192, 246, 212, 94, 116, 58, 128, 119, 118, 73, 72, 128}
	for i, block := range messageBlocks {
		if isEncrypt {
			messageBlocks[i] = xorOperationBlock(block, iv)
			messageBlocks[i], error = goblockc.Parse(messageBlocks[i], key, true)
			if error != nil {
				fmt.Println("Error in CBC.")
			}
			iv = messageBlocks[i]
		} else {
			messageBlocks[i], error = goblockc.Parse(block, key, false)
			if error != nil {
				fmt.Println("Error in CBC.")
			}
			messageBlocks[i] = xorOperationBlock(messageBlocks[i], iv)
			iv = block
		}
	}

	return messageBlocks
}

func ofb(messageBlocks [][]byte, key []byte, isEncrypt bool) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting OFB.")

	result := make([][]byte, len(messageBlocks))

	// This is an 8-bit OFB
	// Random IV
	shiftRegister := []byte{62, 52, 12, 66, 21, 82, 112, 173, 92, 216, 252, 222, 2, 82, 11, 97}
	// TODO: I think this should be better written as a nested loop but at the time of writing, my brain can only work with a single loop
	for i := 0; i < len(messageBlocks)*len(messageBlocks[0]); i++ {
		// Encrypt the shift register (with OFB, E = D)
		output, error := goblockc.Parse(shiftRegister, key, true)
		if error != nil {
			fmt.Println("Error in OFB.")
		}
		// Take the leftmost byte of the output
		leftMostByte := output[0]
		// First byte of the ciphertext is the XOR of the leftmost byte of the output and the first byte of the plaintext
		result[i/16] = append(result[i/16], messageBlocks[i/16][i%16]^leftMostByte)
		// Shift the shift register to the left by 1 bit
		shiftRegister = append(shiftRegister[1:], output[0])
	}
	return result
}

func cfb(messageBlocks [][]byte, key []byte, isEncrypt bool) [][]byte {
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
		output, error := goblockc.Parse(shiftRegister, key, true)
		if error != nil {
			fmt.Println("Error in CFB.")
		}
		// Take the leftmost byte of the output
		leftMostByte := output[0]
		// First byte of the ciphertext is the XOR of the leftmost byte of the output and the first byte of the plaintext
		result[i/16] = append(result[i/16], messageBlocks[i/16][i%16]^leftMostByte)
		// Shift the shift register to the left by 1 bit
		if isEncrypt {
			shiftRegister = append(shiftRegister[1:], result[i/16][i%16])
		} else {
			shiftRegister = append(shiftRegister[1:], messageBlocks[i/16][i%16])
		}
	}

	return result
}

func counter(messageBlocks [][]byte, key []byte, isEncrypt bool) [][]byte {
	// INPUT:
	// 	messageBlocks: plaintext or ciphertext as array of blocks
	// 	keyArr: key for encryption/decryption as array of bytes
	// 	mode: "encrypt" or "decrypt"
	// OUTPUT:
	// 	ciphertext or plaintext as array of blocks
	fmt.Println("Starting counter.")

	// Random IV
	counter := []byte{65, 23, 8, 11, 71, 115, 72, 163, 18, 82, 94, 217, 85, 37, 78, 245}

	for i := 0; i < len(messageBlocks); i++ {
		// Encrypt the counter
		output, error := goblockc.Parse(counter, key, true)
		if error != nil {
			fmt.Println("Error in counter.")
		}
		// XOR the output with the plaintext
		messageBlocks[i] = xorOperationBlock(messageBlocks[i], output)
		// Increment the counter
		counter[15]++
	}

	return messageBlocks
}

// Process encrypt/decrypt (main function)
func GoBlockC(message, key, mode string, isEncrypt bool) (string, time.Duration) {
	fmt.Println("Encrypting plaintext", message, "with key", key, "using mode", mode)

	// Split message into blocks
	var messageHex []byte
	if isEncrypt {
		messageHex = []byte(message)
	} else {
		messageHex, _ = hex.DecodeString(message)
		fmt.Println("Message hex: ", messageHex)
	}
	messageBlocks := formatMessageIntoBlocks(messageHex)
	fmt.Println("Message blocks: ", messageBlocks)

	// Convert key to hexadecimal array
	keyHex := formatKeyInto128Bit(key)

	start := time.Now()
	result := make([][]byte, len(messageBlocks))
	// TODO: Change to switch case
	if mode == "ecb" {
		result = ecb(messageBlocks, keyHex, isEncrypt)
	} else if mode == "cbc" {
		result = cbc(messageBlocks, keyHex, isEncrypt)
	} else if mode == "ofb" {
		result = ofb(messageBlocks, keyHex, isEncrypt)
	} else if mode == "cfb" {
		result = cfb(messageBlocks, keyHex, isEncrypt)
	} else if mode == "ctr" {
		result = counter(messageBlocks, keyHex, isEncrypt)
	}

	fmt.Println("Result: ", result)

	resultString := ""
	for _, block := range result {
		if isEncrypt {
			resultString += string(hex.EncodeToString(block[:]))
		} else {
			resultString += string(block[:])
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)

	return resultString, elapsed
}
