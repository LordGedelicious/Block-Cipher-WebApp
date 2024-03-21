package helper

import (
	"encoding/hex"
	"strings"
)

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

func HexToBytes(hexString string) []byte {
	hexString = strings.ReplaceAll(hexString, " ", "")
	hexBytes, _ := hex.DecodeString(hexString)
	return hexBytes
}

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func SplitIntoBlocks(text string, blockSize int) []string {
	textHex := TextToHex(text)
	blocks := make([]string, 0)
	for i := 0; i < len(textHex); i += blockSize / 4 {
		block := textHex[i : i+blockSize/4]
		block = block + strings.Repeat("0", blockSize/4-len(block))
		blocks = append(blocks, block)
	}
	return blocks
}

func StringToArr(str string) []string {
	arr := make([]string, 0)
	for i := 0; i < len(str); i += 2 {
		arr = append(arr, str[i:i+2])
	}
	return arr
}

func ArrToString(arr []string) string {
	return strings.Join(arr, "")
}

func SpliceKey(key string) string {
	keyHex := TextToHex(key)
	if len(keyHex) < 32 {
		keyHex = keyHex + strings.Repeat("0", 32-len(keyHex))
	} else if len(keyHex) > 32 {
		keyHex = keyHex[:32]
	}
	return keyHex
}
