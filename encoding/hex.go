package encoding

import "encoding/hex"

func HexToStr(text string) (string, error) {
	bytes, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func StrToHex(text string) string {
	return hex.EncodeToString([]byte(text))
}
