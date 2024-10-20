package compression

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"github.com/elysiamae/sisyphus"
	"github.com/elysiamae/sisyphus/encoding"
)

// TODO: 伪加密

// func ZIPPseudoEncrypt(data []byte) (fixedBin []byte, err error) {
// 	return nil, nil
// }

// Unzip 解压缩文件
func Unzip(filePath string, out string) error {
	color := slog.New(sisyphus.NewColorLogHandler())

	// 存储待处理的 ZIP 文件路径
	files := []string{filePath}

	for len(files) > 0 {
		file := files[len(files)-1]
		files = files[:len(files)-1]

		// color.Info("Open File:", slog.String("File", file))

		// Open ZIP File
		r, err := zip.OpenReader(file)
		if err != nil {
			return err
		}
		defer r.Close()

		// UnZIP File
		for _, f := range r.File {
			// Join File Path
			p := filepath.Join(out, f.Name)

			// Make dir
			if f.FileInfo().IsDir() {
				err := os.MkdirAll(p, os.ModePerm)
				if err != nil {
					return err
				}
				continue
			}

			// Make File
			outFile, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			rc, err := f.Open()
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, rc); err != nil {
				return err
			}
			outFile.Close()
			rc.Close()

			color.Info("Unzip File:", slog.String("File Name", filepath.Base(outFile.Name())))

			if filepath.Ext(p) == ".zip" {
				files = append(files, p)
			}
		}

	}
	return nil
}

// ZIPHash Calculate ZIP File Hash 并发计算压缩包内每个文件的 Hash
func ZIPHashC(f *zip.File, t encoding.HashType, wg *sync.WaitGroup, r chan<- sisyphus.Pair[string, []byte]) {
	defer wg.Done()

	// Open File
	file, err := f.Open()
	if err != nil {
		r <- sisyphus.Pair[string, []byte]{First: f.Name}
		return
	}
	defer file.Close()

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

	// Read File
	if _, err := io.Copy(hasher, file); err != nil {
		r <- sisyphus.Pair[string, []byte]{First: f.Name}
		return
	}

	// Calculate Hash
	r <- sisyphus.Pair[string, []byte]{First: f.Name,Second: hasher.Sum(nil)}
}
