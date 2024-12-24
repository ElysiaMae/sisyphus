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

// IsNumeric checks if a string contains only numeric characters.
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

// Combinations
// The Combinations function generates all possible combinations of a specified length n from a given charset (character set).
// It uses a closure to produce each combination on demand.
// The closure can be called repeatedly to obtain one combination at a time, returning the combination as a byte slice and a boolean value indicating whether more combinations are available.
// Combinations 函数从给定的字符集 charset 中生成指定长度 n 的所有可能组合。它使用闭包按需生成每一个组合。闭包可以被反复调用，每次返回一个新的组合，并且返回组合的字节切片和一个布尔值，表示是否还有更多的组合
func Combinations(charset []byte, n int) func() ([]byte, bool) {
	// Initialize an array `indices` to track the current index for each character in the charset.
	// `indices` will keep track of the current position in the charset for each position in the combination.
	// We also initialize `result` to store the current combination.
	// `indices` stores the index positions, and `result` stores the combination.
	indices := make([]int, n)
	result := make([]byte, n)

	// Closure: This function generates a new combination each time it is called.
	// 闭包：每次调用时生成一个新的排列组合
	return func() ([]byte, bool) {
		// Fill the result array with the current combination based on `indices` positions in the charset.
		// 使用 `indices` 中的当前索引填充 `result` 数组，形成当前的组合
		for i := 0; i < n; i++ {
			result[i] = charset[indices[i]]
		}

		// Update the `indices` array to simulate carrying over (like in counting systems).
		// 从右到左递增 `indices` 数组，模拟进位操作
		for i := n - 1; i >= 0; i-- {
			indices[i]++
			// If the index at `indices[i]` is less than the charset length, stop the carry process.
			// 如果 `indices[i]` 小于字符集长度，停止进位
			if indices[i] < len(charset) {
				break
			}
			// If the index exceeds the charset length, reset it to 0 and move to the next index.
			// 如果 `indices[i]` 超过字符集长度，将其重置为 0，并递增下一个索引
			indices[i] = 0
			// If we reach the first index and need to reset, we return the current combination and `false` indicating no more combinations are available.
			// 如果到达第一个索引并且需要进位，表示已生成所有组合，返回当前组合并返回 `false`，表示没有更多组合
			if i == 0 {
				return result, false
			}
		}

		// If there are still combinations left, return the current combination and `true`.
		// 如果还有更多组合，返回当前组合和 `true`
		return result, true
	}
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
