package encoding

import (
	"encoding/hex"
)

// HexToByte Hex Bytes To String Bytes
func HexToByte(input []byte) ([]byte, error) {
	o := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(o, input)
	if err != nil {
		return o, err
	}

	return o, nil
}

// ByteToHex String Bytes To Hex Bytes
func ByteToHex(input []byte) []byte {
	o := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(o, input)

	return o
}
