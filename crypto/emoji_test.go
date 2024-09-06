package crypto_test

import (
	"testing"

	"github.com/elysiamae/sisyphus/crypto"
)

func TestEmojiEncrypt(t *testing.T) {
	str := "flag{10ve_4nd_Peace}"
	options := crypto.EmojiAESOptions{
		Rot: 0,
		Key: "GAME",
	}
	t.Log(crypto.EmojiAESEncrypt(str, options))
}

func TestEmojiDecrypt(t *testing.T) {
	str := "🙃💵🌿🎤🚪🌏🐎🥋🚫😆😍🥋🐘🍴🚰😍☀🌿😇😍😡🏎👉🛩🤣🖐💧☺🌉🏎😇😆🎈💧⏩☺🔄🌪⌨🐅🎅🙃🍌🙃🔪☂🏹🕹☃🌿🌉💵🐎🐍😇🍵😍🐅🎈🥋🚰✅🎈🎈"
	options := crypto.EmojiAESOptions{
		Rot: 0,
		Key: "GAME",
	}
	t.Log(crypto.EmojiAESDecrypt(str, options))
}
