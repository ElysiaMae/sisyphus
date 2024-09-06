package image

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"strings"
)

// CalcCRC32 Calculate CRC32 Value
// 计算 CRC32 的值 返回 CRC32 是否正确 计算的值 提取的值
func CheckPNGCRC32(fileType string, data []byte) (flag bool, calCRC32 uint32, fileCRC32 uint32, err error) {
	switch strings.ToLower(fileType) {
	case ".png":
		ihdrData := data[12:29]                          // IHDR Data Block
		fileCRC32 = binary.BigEndian.Uint32(data[29:33]) // 提取 CRC32 的值
		calCRC32 = crc32.ChecksumIEEE(ihdrData)          // 计算 CRC32 的值
	default:
		return false, uint32(0), uint32(0), errors.New("file type error")
	}

	if calCRC32 == fileCRC32 {
		flag = true
	}

	return flag, calCRC32, fileCRC32, nil
}

// FixPNGIHDR
// 修复 PNG Image 的宽高
func FixPNGIHDR(bin []byte, maxDim int) (fixedBin []byte, width []byte, height []byte, err error) {
	ihdr := bin[12:29]                              // IHDR Data Block
	crc32Key := binary.BigEndian.Uint32(bin[29:33]) // Extract the CRC32 value

	for w := 0; w < maxDim; w++ {
		width = make([]byte, 4)
		binary.BigEndian.PutUint32(width, uint32(w))

		for h := 0; h < maxDim; h++ {
			height = make([]byte, 4)
			binary.BigEndian.PutUint32(height, uint32(h))

			// Write width and height into the IHDR data block
			copy(ihdr[4:8], width)
			copy(ihdr[8:12], height)

			crc32Result := crc32.ChecksumIEEE(ihdr)

			if crc32Result == crc32Key {
				copy(bin[12:29], ihdr) // 将 IHDR Copy 到原位置
				return bin, width, height, nil
			}
		}

	}

	return bin, []byte{00}, []byte{00}, errors.New("no matching width and height found")
}
