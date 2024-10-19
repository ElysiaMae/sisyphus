package compression

import (
	"archive/tar"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"

	"github.com/elysiamae/sisyphus"
	"github.com/elysiamae/sisyphus/encoding"
)

func TGZCount(tr *tar.Reader) int {
	count := 0
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // 到达文件末尾
		}
		if err != nil {
			log.Fatal(err)
		}

		if header.Typeflag != tar.TypeDir {
			count += 1
		}
	}
	fmt.Println()
	return count
}

func TGZHash(tf *tar.Header, f *tar.Reader, t encoding.HashType) (sisyphus.Pair[string, []byte], error) {
	var hasher hash.Hash
	switch t {
	case encoding.MD5:
		hasher = md5.New()
	case encoding.SHA_1:
		hasher = sha1.New()
	case encoding.SHA_256:
		hasher = sha256.New()
	case encoding.SHA_512:
		hasher = sha512.New()
	default:
		hasher = md5.New()
	}

	result := sisyphus.Pair[string, []byte]{
		First: tf.Name,
	}
	// Read File
	if _, err := io.Copy(hasher, f); err != nil {
		return result, errors.New(fmt.Sprintf("Error calculating hashes for file %s: %v", tf.Name, err))
	}

	// Calculate Hash
	result.First, result.Second = tf.Name, hasher.Sum(nil)
	return result, nil

}
