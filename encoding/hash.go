package encoding

type HashType int

const (
	MD5 HashType = iota
	SHA_1
	SHA_256
	SHA_384
	SHA_512
	SHA3_224
	SHA3_256
	SHA3_512
	Blake2b
)
