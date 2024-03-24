package cipher

import "testing"

func TestGoBlockC(t *testing.T)	{
	plaintext := "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
	key := "0102030405060708090a0b0c0d0e0f10"

	ciphertext, _ := GoBlockC(plaintext, key, "ecb", true)
	if ciphertext != "" {
		t.Log("GoBlockC encryption success")
	} else {
		t.Error("GoBlockC encryption failed")
	}
}