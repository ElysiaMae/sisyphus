package crypto

import (
	"crypto/rand"

	"golang.org/x/crypto/curve25519"
)

func generateSharedKey() ([32]byte, error) {
	var alicePrivate, alicePublic, bobPrivate, bobPublic [32]byte

	// Alice 生成私钥
	_, err := rand.Read(alicePrivate[:])
	if err != nil {
		return [32]byte{}, err
	}

	// Alice 生成公钥
	curve25519.ScalarBaseMult(&alicePublic, &alicePrivate)

	// Bob 生成私钥
	_, err = rand.Read(bobPrivate[:])
	if err != nil {
		return [32]byte{}, err
	}

	// Bob 生成公钥
	curve25519.ScalarBaseMult(&bobPublic, &bobPrivate)

	// Alice 和 Bob 交换公钥并生成共享密钥
	var sharedKey [32]byte
	curve25519.ScalarMult(&sharedKey, &alicePrivate, &bobPublic)

	return sharedKey, nil
}
