package image

import (
	"image"
)

func LSB(img image.Image) ([]byte, error) {
	bounds := img.Bounds()
	var bits []uint8

	// 遍历图像的每个像素
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取像素的 RGBA 值
			r, g, b, _ := img.At(x, y).RGBA()

			// 提取每个颜色通道的最低有效位
			bits = append(bits, uint8(r&1))
			bits = append(bits, uint8(g&1))
			bits = append(bits, uint8(b&1))
		}
	}

	// 将位组合成字节
	var bytes []byte
	for i := 0; i+7 < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			b |= bits[i+j] << (7 - j)
		}
		if b == 0 { // 使用空字节作为结束标志
			break
		}
		bytes = append(bytes, b)
	}

	return bytes, nil
}
