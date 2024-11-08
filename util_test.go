package sisyphus_test

import (
	"strings"
	"testing"

	"github.com/elysiamae/sisyphus"
	"github.com/elysiamae/sisyphus/encoding"
)

func TestXor(t *testing.T) {
	s1 := []byte("1c0111001f010100061a024b53535009181c")
	s2 := []byte("686974207468652062756c6c277320657965")

	a1, err := encoding.HexToByte(s1)
	if err != nil {
		t.Fatal(err)
	}

	a2, err := encoding.HexToByte(s2)
	if err != nil {
		t.Fatal(err)
	}

	o, err := sisyphus.Xor(a1, a2)
	if err != nil {
		t.Error(err)
	}
	h := encoding.ByteToHex(o)
	if strings.Compare(string(h), "746865206b696420646f6e277420706c6179") != 0 {
		t.Error("XOR Error")
	}
	t.Log(h)
	t.Log(string(h))
}
