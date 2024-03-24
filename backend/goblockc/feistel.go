package goblockc

import (
	"encoding/binary"
	"slices"

	"block-cipher-webapp/backend/utils"
)

func Feistel(right uint64, key uint64, forward bool) (uint64, error) {
	if forward {
		sbox := utils.GetSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, right)

		for i := range 8 {
			bytes[i] = sbox[bytes[i]]
		}

		substitutedRight, err := utils.BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}

		permutedRight := Permute(substitutedRight, forward)

		return permutedRight ^ key, nil
	} else {

		unXORedRight := right ^ key

		unpermutedRight := Permute(unXORedRight, forward)

		invsbox := utils.GetInvSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, unpermutedRight)

		for i := range 8 {
			bytes[i] = invsbox[bytes[i]]
		}

		unsubstitutedRight, err := utils.BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}


		return unsubstitutedRight, nil
	}
}

func Permute(right uint64, forward bool) uint64 {
	rightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(rightBytes, right)

	rightBytesOne := make([]byte, 4)
	copy(rightBytesOne, rightBytes[:4])
	rightBytesTwo := make([]byte, 4)
	copy(rightBytesTwo, rightBytes[4:])

	result := slices.Concat(rightBytesOne, utils.Rotate(rightBytesTwo, 1, forward))
	return binary.BigEndian.Uint64(result)
}
