package compression

import (
	"bytes"
	"errors"
)

func RARPseudoEncrypt(data []byte) (fixedBin []byte, err error) {
	// MARK: Check RAR Version
	h := data[:7] // HEAD
	// fmt.Printf("%x\n", h)

	rar4 := []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00}
	rar5 := []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x01}

	if bytes.Equal(h, rar5) {
		return fixedBin, errors.New("RAR Version is RAR 5")
	}

	if !bytes.Equal(h, rar4) {
		return fixedBin, errors.New("RAR MARK_HEAD Error")
	}

	// uv := uint32(data[44]) / 10 // UNP_VER
	hf := data[23] // HEAD_FLAGS

	if hf == 0x24 {
		data[23] = 0x20
	}

	return data, nil
}
