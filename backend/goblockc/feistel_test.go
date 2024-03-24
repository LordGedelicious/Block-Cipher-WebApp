package goblockc

import (
	"testing"
)

func TestFeistel(t *testing.T) {
	forward, err := Feistel(0x7d7d7d7d7d7d7d7d, 0xffffffffffffffff, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	if forward != 0x0 {
		t.Error("Feistel is not correct")
	}
}

func TestFeistelInvertibility(t *testing.T) {
	forward, err := Feistel(0x0102030405060708, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	if backward, err := Feistel(forward, 0x090A0B0C0D0E0F10, false); err != nil || backward != 0x0102030405060708 {
		t.Error("Feistel is not invertible")
	}
}

func TestFeistelInvertibility2(t *testing.T) {
	forward, err := Feistel(0x0102030405060708, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	forward2, err := Feistel(forward, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	backward2, err := Feistel(forward2, 0x090A0B0C0D0E0F10, false)
	if err != nil {
		t.Error("Unexpected error")
	}

	if backward, err := Feistel(backward2, 0x090A0B0C0D0E0F10, false); err != nil || backward != 0x0102030405060708 {
		t.Error("Feistel is not invertible")
	}
}

func TestPermute(t *testing.T) {
	left := Permute(0x0102030405060708, true)

	if left != 0x0102030406070805 {
		t.Error("Left is not correct")
		t.Log(left)
	}
}

func TestPermuteInvertibility(t *testing.T) {
	right := Permute(0x0102030405060708, false)
	left := Permute(right, true)

	if left != 0x0102030405060708 {
		t.Error("Left is not correct")
	}
}
