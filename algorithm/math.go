package algorithm

import (
	"errors"
	"math/big"
)

// Gcd
// Euclid algorithm
// 欧几里得算法
func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// BigGcd
// Euclid algorithm
// calculates the greatest common divisor of two big integers.
// 欧几里得算法
// 计算两个大整数的最大公约数
func BigGcd(a, b *big.Int) *big.Int {
	// Create new big.Int variables to avoid modifying the input
	// 创建新的 big.Int 变量，避免修改输入
	x := new(big.Int).Set(a)
	y := new(big.Int).Set(b)

	// Temporary variable for swapping
	// 用于交换的临时变量
	temp := new(big.Int)

	// Ensure x and y are non-negative
	// 确保 x 和 y 是非负的
	x.Abs(x)
	y.Abs(y)

	// Loop until y becomes zero
	// 循环直到 y 变为零
	for y.Sign() != 0 {
		// Swap x and y
		temp.Set(x)
		x.Set(y)

		// Set y to the remainder of temp divided by y
		// 将 y 设置为 temp 除以 y 的余数
		y.Rem(temp, y)
	}

	// The GCD is the absolute value of x
	// 最大公约数是 x 的绝对值
	return x.Abs(x)
}

// SteinGcd
// calculates the greatest common divisor of two big integers using the Stein algorithm.
// 使用 Stein 算法计算最大公约数
func SteinGcd(a, b *big.Int) *big.Int {
	// Create new big.Int variables to avoid modifying the input
	// 创建新的 big.Int 变量，避免修改输入
	x := new(big.Int).Set(a)
	y := new(big.Int).Set(b)

	// Ensure x and y are non-negative
	// 确保 x 和 y 是非负的
	x.Abs(x)
	y.Abs(y)

	// If x or y is 0, return the other number
	// 如果 x 或 y 为 0，返回另一个数
	if x.Sign() == 0 {
		return y
	}
	if y.Sign() == 0 {
		return x
	}

	// Initialize k to 0
	// 初始化 k 为 0
	k := new(big.Int)

	// While both x and y are even, divide them by 2 and increase k
	// 当 x 和 y 都是偶数时，将它们除以 2 并增加 k
	for x.Bit(0) == 0 && y.Bit(0) == 0 {
		x.Rsh(x, 1)
		y.Rsh(y, 1)
		k.Add(k, big.NewInt(1))
	}

	// While x is even, divide it by 2
	// 当 x 是偶数时，将其除以 2
	for x.Bit(0) == 0 {
		x.Rsh(x, 1)
	}

	for y.Sign() != 0 {
		// While y is even, divide it by 2
		// 当 y 是偶数时，将其除以 2
		for y.Bit(0) == 0 {
			y.Rsh(y, 1)
		}

		// If x > y, swap x and y
		// 如果 x > y，交换 x 和 y
		if x.Cmp(y) > 0 {
			x, y = y, x
		}

		// Subtract x from y
		// 从 y 中减去 x
		y.Sub(y, x)
	}

	// Multiply the result by 2^k
	// 将结果乘以 2^k
	return x.Lsh(x, uint(k.Uint64()))
}

// ExGcd 扩展欧几里得算法
func ExGcd(a, b int64) (gcd, x1, y1 int64) {
	// x0 和 x1 用于存储 x1 和前一个 x1
	// y0 和 y1 用于存储 y1 和前一个 y1

	x0, x1 := int64(1), int64(0) // 1 * a + 0 * b
	y0, y1 := int64(0), int64(1) // 0 * a + 1 * b

	// 如果余数 r 不为 0
	for b != 0 {
		q := a / b           // 计算 a/b 商
		a, b = b, a-q*b      // 更新余数 a 和 b 用于存储当前和前一个余数
		x0, x1 = x1, x0-q*x1 // 更新 x1
		y0, y1 = y1, y0-q*y1 // 更新 y1
	}

	return a, x0, y0
}

// ExGcdBig 扩展欧几里得算法 Big Int 版本
func ExGcdBig(a, b *big.Int) (gcd, x, y *big.Int) {
	// 初始化 Bézout 系数
	x0, x1 := big.NewInt(1), big.NewInt(0) // 1 * a + 0 * b
	y0, y1 := big.NewInt(0), big.NewInt(1) // 0 * a + 1 * b

	// 如果余数 r 不为 0
	for b.Sign() != 0 {
		q := new(big.Int).Div(a, b) // 计算 a/b 商
		a, b = b, new(big.Int).Mod(a, b)
		x0, x1 = x1, new(big.Int).Sub(x0, new(big.Int).Mul(q, x1))
		y0, y1 = y1, new(big.Int).Sub(y0, new(big.Int).Mul(q, y1))
	}

	// 返回 gcd 和 Bézout 系数 x 和 y
	return a, x0, y0
}

// ExGcdR 扩展欧几里得算法 递归版本
func ExGcdR(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	} else {
		gcd, x1, y1 := ExGcdR(b, a%b)
		x := y1
		y := x1 - (a/b)*y1
		return gcd, x, y
	}
}

// ModInverse 计算模逆元
func ModInverse(a, m int64) (int64, error) {
	gcd, x, _ := ExGcd(a, m)
	if gcd != 1 {
		return 0, errors.New("模逆元不存在")
	}
	// 保证结果始终是一个非负整数
	return (x%m + m) % m, nil
}

// CRT 中国剩余定理
func CRT(k int, a, r []int) int64 {
	n := int64(1)
	result := int64(0)

	// 计算所有模数的乘积
	for i := 0; i < k; i++ {
		n *= int64(r[i])
	}

	// 对每个模数求解
	for i := 0; i < k; i++ {
		q := n / int64(r[i])
		_, x, _ := ExGcd(q, int64(r[i])) // 求解模反元素 x * q mod r[i] = 1
		result = (result + int64(a[i])*q*x%n) % n
	}

	// 最后对结果取模并确保为非负数
	return (result%n + n) % n
}

// BinPow 快速幂 迭代版本
func BinPow(base, exponent uint64) uint64 {
	result := uint64(1) // 初始化结果为 1
	for exponent > 0 {
		if exponent&1 != 0 { // 如果指数是奇数
			result *= base // 更新结果
		}
		base *= base   // 底数平方
		exponent >>= 1 // 右移指数
	}
	return result
}
