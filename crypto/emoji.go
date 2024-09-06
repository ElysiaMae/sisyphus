package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"strings"
)

// https://aghorler.github.io/emoji-aes/

type EmojiAESOptions struct {
	Rot int
	Key string
}

// MARK: Emoji AES Encrypt

// EmojiAESEncrypt Emoji AES CBC TODO:
func EmojiAESEncrypt(text string, options EmojiAESOptions) (string, error) {
	emojisInit := [65]string{"🍎", "🍌", "🏎", "🚪", "👁", "👣", "😀", "🖐", "ℹ", "😂", "🥋", "✉", "🚹", "🌉", "👌", "🍍", "👑",
		"👉", "🎤", "🚰", "☂", "🐍", "💧", "✖", "☀", "🦓", "🏹", "🎈", "😎", "🎅", "🐘", "🌿", "🌏", "🌪", "☃", "🍵", "🍴",
		"🚨", "📮", "🕹", "📂", "🛩", "⌨", "🔄", "🔬", "🐅", "🙃", "🐎", "🌊", "🚫", "❓", "⏩", "😁", "😆", "💵", "🤣", "☺",
		"😊", "😇", "😡", "🎃", "😍", "✅", "🔪", "🗒"}

	// MARK:
	if text == "" || options.Key == "" {
		return "", errors.New("Text and Key cannot be blank.")
	}
	if options.Rot < 0 || options.Rot > 65 {
		return "", errors.New("Rotation must be between 0 and 64.")
	}

	encryptedByte, err := AESCBCEncrypt([]byte(text), []byte(options.Key))
	if err != nil {
		panic(err)
	}

	encrypted := base64.StdEncoding.EncodeToString(encryptedByte)

	var emojis [65]string
	if options.Rot != 0 {
		for i := 0; i < len(emojisInit); i++ {
			if i+options.Rot < len(emojisInit) {
				emojis[i] = emojisInit[i+options.Rot]
			} else {
				emojis[i] = emojisInit[i+options.Rot-len(emojisInit)]
			}
		}
	} else {
		emojis = emojisInit
	}

	replacements := map[string]string{}

	// a - z
	for i := 0; i < 26; i++ {
		replacements[string(rune('a'+i))] = emojis[i]
	}
	// A - Z
	for i := 0; i < 26; i++ {
		replacements[string(rune('A'+i))] = emojis[i+26]
	}
	// 0 - 9
	for i := 0; i < 10; i++ {
		replacements[string(rune('0'+i))] = emojis[i+52]
	}

	replacements["+"] = emojis[62]
	replacements["/"] = emojis[63]
	replacements["="] = emojis[64]

	for old, new := range replacements {
		encrypted = strings.Replace(encrypted, old, new, -1)
	}

	return encrypted, nil
}

// MARK: Emoji AES Decrypt

// EmojiAESDecrypt Decryption Function
// This function decrypts AES-encrypted text encoded with emojis using the provided key.
func EmojiAESDecrypt(text string, options EmojiAESOptions) (string, error) {
	// Initialize the emoji array used for mapping letters, numbers, and symbols
	emojisInit := [65]string{"🍎", "🍌", "🏎", "🚪", "👁", "👣", "😀", "🖐", "ℹ", "😂", "🥋", "✉", "🚹", "🌉", "👌", "🍍", "👑",
		"👉", "🎤", "🚰", "☂", "🐍", "💧", "✖", "☀", "🦓", "🏹", "🎈", "😎", "🎅", "🐘", "🌿", "🌏", "🌪", "☃", "🍵", "🍴",
		"🚨", "📮", "🕹", "📂", "🛩", "⌨", "🔄", "🔬", "🐅", "🙃", "🐎", "🌊", "🚫", "❓", "⏩", "😁", "😆", "💵", "🤣", "☺",
		"😊", "😇", "😡", "🎃", "😍", "✅", "🔪", "🗒"}
	// Check if the input text and key are empty
	if text == "" || options.Key == "" {
		return "", errors.New("Text and Key cannot be blank.")
	}
	// Check if the rotation value is within the valid range
	if options.Rot < 0 || options.Rot > 65 {
		return "", errors.New("Rotation must be between 0 and 64.")
	}

	var emojis [65]string
	// If the rotation value is not 0, rearrange the emoji array
	if options.Rot != 0 {
		for i := 0; i < len(emojisInit); i++ {
			if i+options.Rot < len(emojisInit) {
				emojis[i] = emojisInit[i+options.Rot]
			} else {
				emojis[i] = emojisInit[i+options.Rot-len(emojisInit)]
			}
		}
	} else {
		emojis = emojisInit
	}

	// Create a map that maps letters, numbers, and symbols to emojis
	replacements := map[string]string{}

	// Map lowercase letters a-z to emojis
	for i := 0; i < 26; i++ {
		replacements[string(rune('a'+i))] = emojis[i]
	}
	// Map uppercase letters A-Z to emojis
	for i := 0; i < 26; i++ {
		replacements[string(rune('A'+i))] = emojis[i+26]
	}
	// Map digits 0-9 to emojis
	for i := 0; i < 10; i++ {
		replacements[string(rune('0'+i))] = emojis[i+52]
	}

	replacements["+"] = emojis[62]
	replacements["/"] = emojis[63]
	replacements["="] = emojis[64]

	// Reverse the mapping to map emojis back to letters, numbers, and symbols
	invertedMap := make(map[string]string)
	for letter, emoji := range replacements {
		invertedMap[emoji] = letter
	}

	// 进行替换操作，将表情符号转换回文本
	// Perform replacement operation to convert emojis back to text
	unemojified := text
	for emoji, letter := range invertedMap {
		unemojified = strings.ReplaceAll(unemojified, emoji, letter)
	}

	// 解码 Base64 编码的密文
	// Decode the Base64 encoded ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(unemojified)
	if err != nil {
		return "", err
	}

	// 解码 Base64 编码的密文
	// Decode the Base64 encoded ciphertext
	iv, calKey := getIVAndKey(ciphertext, options.Key)
	block, err := aes.NewCipher(calKey)
	if err != nil {
		return "", err
	}

	// 使用 CBC 模式进行解密
	// Use CBC mode for decryption
	mode := cipher.NewCBCDecrypter(block, iv)

	// 去除前缀和盐值
	// Remove the prefix and salt
	ciphertext = ciphertext[16:]
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充
	// Remove padding
	paddingLen := int(plaintext[len(plaintext)-1])
	if paddingLen > len(plaintext) {
		return "", errors.New("padding len error")
	}

	return string(plaintext[:len(plaintext)-paddingLen]), nil
}

// getIVAndKey
// by https://www.cnblogs.com/caiawo/p/17255857.html
func getIVAndKey(ciphertext []byte, key string) (iv []byte, calKey []byte) {
	// 从密文中提取盐值
	// Extract salt from ciphertext
	salt := ciphertext[8:16]
	// 计算第一个 MD5 哈希值
	// Calculate the first MD5 hash
	hash1 := md5.Sum([]byte(key + string(salt)))
	// 计算第二个 MD5 哈希值
	// Calculate the second MD5 hash
	hash2 := md5.Sum(append(hash1[:], []byte(key+string(salt))...))
	// 计算第三个 MD5 哈希值
	// Calculate the third MD5 hash
	hash3 := md5.Sum(append(hash2[:], []byte(key+string(salt))...))
	// 组合前两个哈希值作为密钥
	// Combine the first two hashes to form the key
	calKey = append(hash1[:], hash2[:]...)
	// 使用第三个哈希值作为 IV
	// Use the third hash as the IV
	iv = hash3[:]
	return
}
