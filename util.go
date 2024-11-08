package sisyphus

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 |
		~string
}

// Max 返回多个数字型中的最大值
func Max[T Numeric](n ...T) (T, error) {
	if len(n) == 0 {
		return T(0), errors.New("empty slice")
	}
	max := n[0]
	for _, i := range n {
		if max < i {
			max = i
		}
	}

	return max, nil
}

// Min 返回多个数字型中的最小值
func Min[T Numeric](n ...T) (T, error) {
	if len(n) == 0 {
		return T(0), errors.New("empty slice")
	}
	min := n[0]
	for _, i := range n {
		if i < min {
			min = i
		}
	}

	return min, nil
}

// SliceReverse 翻转
func SliceReverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

// Xor Go XOR
func Xor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("sisyphus.Xor: length mismatch")
	}

	o := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		o[i] = a[i] ^ b[i]
	}

	return o, nil
}

// IntToChar Int 数组转换为 Char
func IntToChar(arr []int) string {
	var result string
	for _, num := range arr {
		result += string(rune(num))
	}

	return result
}

// ScoreByte
// 根据英文字母频率计算 Bytes 分数
func ScoreByte(input []byte) (score float64, err error) {
	score = 0.0
	lf := getLetterFrequency()
	for _, i := range input {
		score += lf[i]
	}
	return score, nil
}

// ScoreStr
// 根据英文字母频率计算 String 分数
func ScoreStr(input string) (score float64, err error) {
	score = 0.0
	arr := []byte(strings.ToUpper(input))
	lf := getLetterFrequency()
	for _, i := range arr {
		score += lf[i]
	}
	return score, nil
}

// getLetterFrequency Statistical Distributions of English Text
// by https://web.archive.org/web/20040603075055/http://www.data-compression.com/english.html
func getLetterFrequency() map[byte]float64 {
	return map[byte]float64{
		'A': 0.0651738, 'B': 0.0124248, 'C': 0.0217339,
		'D': 0.0349835, 'E': 0.1041442, 'F': 0.0197881,
		'G': 0.0158610, 'H': 0.0492888, 'I': 0.0558094,
		'J': 0.0009033, 'K': 0.0050529, 'L': 0.0331490,
		'M': 0.0202124, 'N': 0.0564513, 'O': 0.0596302,
		'P': 0.0137645, 'Q': 0.0008606, 'R': 0.0497563,
		'S': 0.0515760, 'T': 0.0729357, 'U': 0.0225134,
		'V': 0.0082903, 'W': 0.0171272, 'X': 0.0013692,
		'Y': 0.0145984, 'Z': 0.0007836, ' ': 0.1918182,
	}
}

// isNumeric checks if a string contains only numeric characters.
// It iterates through each rune in the string and returns false
// 遍历检查字符串是否为仅数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

func PrintHexArr(data []byte) {
	for _, b := range data {
		fmt.Printf("%02X, ", b)
	}

	fmt.Println()
}

func ReversePrintArr(arr []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}

	fmt.Println()
}
