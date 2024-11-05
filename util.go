package sisyphus

import (
	"errors"
	"fmt"
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

// IntToChar Int 数组转换为 Char
func IntToChar(arr []int) string {
	var result string
	for _, num := range arr {
		result += string(rune(num))
	}

	return result
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
