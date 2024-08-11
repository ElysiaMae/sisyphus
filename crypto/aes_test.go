package crypto_test

import (
	"fmt"
	"testing"

	"github.com/elysiamae/sisyphus/crypto"
	"github.com/elysiamae/sisyphus/generate"
)

func TestAes(t *testing.T) {
	// 密钥
	key, err := generate.GenerateSafeRandomData(32)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	t.Log("key:")
	fmt.Printf("%s\n", key)
	// 明文
	plainText := string("Hello, world!")

	// 加密
	cipherText, err := crypto.AESGCMEncryptData(plainText, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	t.Log("Cipher text:")
	fmt.Printf("%s\n", cipherText)

	// 解密
	decryptedText, err := crypto.AESGCMDecryptData(cipherText, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	t.Log("Decrypted text:")
	fmt.Printf("%s\n", decryptedText)
}
