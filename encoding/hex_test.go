package encoding_test

import (
	"strings"
	"testing"

	"github.com/elysiamae/sisyphus/encoding"
)

func TestHexToByte(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	a := []byte(s)
	t.Log(a)
	o, err := encoding.HexToByte([]byte(s))
	if err != nil {
		t.Error(err)
	}
	t.Log(o)
	t.Log(string(o))
}

func TestByteToHex(t *testing.T) {
	s := "I'm killing your brain like a poisonous mushroom"
	o := encoding.ByteToHex([]byte(s))
	t.Log(o)
	t.Log(string(o))
	s = string(o)
	if strings.Compare(s, "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d") != 0 {
		t.Error("ByteToHex Func Error")
	}
}
