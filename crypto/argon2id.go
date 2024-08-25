package crypto

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/elysiamae/sisyphus/generate"

	"golang.org/x/crypto/argon2"
)

// Argon2id 参数
type argon2Params struct {
	memory      uint32 // 内存使用量（以字节为单位）
	iterations  uint32 // 迭代次数
	parallelism uint8  // 并行线程数
	saltLength  uint32 // salt 长度
	keyLength   uint32 // 密钥长度
}

// HashPassword 使用 Argon2id 对密码进行 hash
func HashPassword(password string) (string, error) {
	passwordHash, err := argon2Hash(password)
	return passwordHash, err
}

// VerifyPassword 验证密码的 hash 值
func VerifyPassword(password, encodedHash string) (match bool, err error) {
	// 从 password hash 中提取参数，盐值
	p, salt, hash, err := decodeArgon2Hash(encodedHash)
	if err != nil {
		return false, err
	}

	// 再次计算 hash 值
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// 使用 crypto/subtle.ConstantTimeCompare 防止定时攻击 prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

// Argon2id 的 Hash 实现
func argon2Hash(str string) (string, error) {
	// 参数设置根据 https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
	params := &argon2Params{
		memory:      19 * 1024, // 19 MiB
		iterations:  2,         // 迭代次数
		parallelism: 1,         // 并行线程数
		saltLength:  16,        // salt 长度
		keyLength:   32,        // 密钥长度
	}

	// 生成指定长度的 salt 值
	salt, err := generate.GenerateSafeRandomBytes(params.saltLength)
	if err != nil {
		return "", err
	}

	// 调用 crypto/argon2 进行 hash
	hash := argon2.IDKey([]byte(str), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// 编码信息和 hash 值
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.memory, params.iterations, params.parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

// 解码 Argon2 Hash 返回加密参数等
func decodeArgon2Hash(encodedHash string) (params *argon2Params, salt, hash []byte, err error) {
	vales := strings.Split(encodedHash, "$")
	if len(vales) != 6 {
		return nil, nil, nil, errors.New("encoded hash format error")
	}

	var version int
	_, err = fmt.Sscanf(vales[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}

	params = &argon2Params{}
	_, err = fmt.Sscanf(vales[3], "m=%d,t=%d,p=%d", &params.memory, &params.iterations, &params.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vales[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vales[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.keyLength = uint32(len(hash))

	return params, salt, hash, nil
}
