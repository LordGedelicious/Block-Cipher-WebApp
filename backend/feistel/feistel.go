// --------------------------------
// put into quarantine, idk what this is
// --------------------------------
package feistel

// import (
// 	"block-cipher-webapp/backend/cipher"
// 	"block-cipher-webapp/backend/helper"
// )

// type FeistelNetwork struct {
// 	lhs  string // Left half of the block, string of 16 hexadecimal digits
// 	rhs  string // Right half of the block, string of 16 hexadecimal digits
// 	key  string // Original key inputted, key generation use loop count
// 	loop int    // Loop count for key generation and for Feistel Network, 0 if not started yet
// }

// func NewFeistelNetwork(lhs, rhs, key string, loop int) *FeistelNetwork {
// 	return &FeistelNetwork{
// 		lhs:  lhs,
// 		rhs:  rhs,
// 		key:  key,
// 		loop: loop,
// 	}
// }

// func (f *FeistelNetwork) String() string {
// 	return fmt.Sprintf("FeistelNetwork(lhs=%s, rhs=%s, key=%s, loop=%d)", f.lhs, f.rhs, f.key, f.loop)
// }

// func (f *FeistelNetwork) Encrypt() *FeistelNetwork {
// 	// Loop for 16 times
// 	for i := 0; i < 16; i++ {
// 		// Generate subkey (round key) from the round loop count
// 		subkey := SubKeyGenerator(f.key, i)
// 		// Perform Feistel Network
// 		newLHS := f.rhs
// 		newRHS := xorOperation(f.lhs, f.FeistelFunction(f.rhs, subkey))
// 		// Update the left and right half of the block
// 		f.lhs = newLHS
// 		f.rhs = newRHS
// 	}
// 	return f
// }

// func (f *FeistelNetwork) Decrypt() *FeistelNetwork {
// 	// TODO: NotImplemented
// 	return f
// }

// func (f *FeistelNetwork) FeistelFunction(rhs, subkey string) string {
// 	// Feistel function is a function that takes a right half of the block and a subkey as input
// 	// The function is defined as follows:
// 	// 1. Perform substitution using S-box
// 	// 2. Perform permutation using shift-rows
// 	// 3. Perform XOR with the subkey
// 	// 4. Return the result
// 	rhsSubstituted := SubBytesSubstitutionArr(rhs)
// 	rhsPermuted := ShiftRows(rhsSubstituted)
// 	xorResult := xorOperation(rhsPermuted, subkey)
// 	return xorResult
// }