package metric

import "unicode/utf8"

// LevenshteinDistance 计算两个字符串之间的Levenshtein距离
func LevenshteinDistance(a, b string) int {
	la := utf8.RuneCountInString(a)
	lb := utf8.RuneCountInString(b)

	// 初始化距离矩阵
	matrix := make([][]int, la+1)
	for i := range matrix {
		matrix[i] = make([]int, lb+1)
	}

	// 初始化矩阵边界值
	for i := 0; i <= la; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		matrix[0][j] = j
	}

	// 动态规划填充矩阵
	for i, ra := range a {
		for j, rb := range b {
			cost := 0
			if ra != rb {
				cost = 1
			}
			matrix[i+1][j+1] = min(
				matrix[i][j+1]+1,  // 删除
				matrix[i+1][j]+1,  // 插入
				matrix[i][j]+cost, // 替换
			)
		}
	}

	return matrix[la][lb]
}

// DamerauLevenshteinDistance 计算两个字符串之间的Damerau-Levenshtein距离
func DamerauLevenshteinDistance(a, b string) int {
	la := utf8.RuneCountInString(a)
	lb := utf8.RuneCountInString(b)

	// 初始化距离矩阵
	matrix := make([][]int, la+1)
	for i := range matrix {
		matrix[i] = make([]int, lb+1)
	}

	// 初始化矩阵边界值
	for i := 0; i <= la; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		matrix[0][j] = j
	}

	// 动态规划填充矩阵
	for i, ra := range a {
		for j, rb := range b {
			cost := 0
			if ra != rb {
				cost = 1
			}
			matrix[i+1][j+1] = min(
				matrix[i][j+1]+1,  // 删除
				matrix[i+1][j]+1,  // 插入
				matrix[i][j]+cost, // 替换
			)

			// 如果可以交换字符
			if i > 0 && j > 0 && ra == rune(b[j-1]) && rune(a[i-1]) == rb {
				matrix[i+1][j+1] = min(
					matrix[i+1][j+1],
					matrix[i-1][j-1]+1, // 交换
				)
			}
		}
	}

	return matrix[la][lb]
}
