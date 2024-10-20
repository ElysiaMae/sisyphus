package metric

import (
	"fmt"
	"math"
)

// 计算Jaro距离
func JaroDistance(s1, s2 string) float64 {
	len1, len2 := len(s1), len(s2)
	if len1 == 0 {
		if len2 == 0 {
			return 1.0
		}
		return 0.0
	}

	// 最大匹配距离
	matchDistance := int(math.Floor(math.Max(float64(len1), float64(len2))/2)) - 1

	// 标记匹配字符
	s1Matches := make([]bool, len1)
	s2Matches := make([]bool, len2)

	// 统计匹配字符数
	matches := 0
	for i := 0; i < len1; i++ {
		low := int(math.Max(0, float64(i-matchDistance)))
		high := int(math.Min(float64(len2-1), float64(i+matchDistance)))

		for j := low; j <= high; j++ {
			if s2Matches[j] || s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matches++
			break
		}
	}

	if matches == 0 {
		return 0.0
	}

	// 统计转置数量
	transpositions := 0
	k := 0
	for i := 0; i < len1; i++ {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++
		}
		k++
	}

	// 计算Jaro距离
	jaro := (float64(matches)/float64(len1) +
		float64(matches)/float64(len2) +
		float64(matches-transpositions/2)/float64(matches)) / 3.0
	return jaro
}

// 计算 Jaro-Winkler 相似度
func JaroWinkler(s1, s2 string, prefixScale float64) float64 {
	jaro := JaroDistance(s1, s2)

	// 找出公共前缀长度，最多4个字符
	prefixLength := 0
	for i := 0; i < int(math.Min(4, math.Min(float64(len(s1)), float64(len(s2))))); i++ {
		if s1[i] == s2[i] {
			prefixLength++
		} else {
			break
		}
	}

	// 计算Jaro-Winkler相似度
	return jaro + float64(prefixLength)*prefixScale*(1-jaro)
}

func main() {
	s1 := "MARTHA"
	s2 := "MARHTA"

	// 计算Jaro-Winkler相似度，前缀权重通常设置为0.1
	similarity := JaroWinkler(s1, s2, 0.1)

	fmt.Printf("Jaro-Winkler similarity between '%s' and '%s': %.4f\n", s1, s2, similarity)
}
