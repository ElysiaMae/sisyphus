package generate

import (
	"crypto/rand"
	"encoding/base64"
)

// 健壮的生成随机数字节数组
func GenerateSafeRandomBytes(length uint32) ([]byte, error) {
	random := make([]byte, length)
	_, err := rand.Read(random) // 使用 crypto/rand 来生成
	if err != nil {
		return nil, err
	}
	return random, nil
}

// 健壮的生成随机数字节数组并 base64 字符串
func GenerateSafeRandomData(length uint32) (string, error) {
	random := make([]byte, length)
	_, err := rand.Read(random)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(random), nil
}
