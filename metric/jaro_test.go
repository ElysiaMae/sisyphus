package metric_test

import (
	"fmt"
	"testing"

	"github.com/elysiamae/sisyphus/metric"
)

func TestJaroWinkler(t *testing.T) {
	s1 := "SHACKLEFORD"
	s2 := "SHACKELFORD"

	// 计算Jaro-Winkler相似度，前缀权重通常设置为0.1
	similarity := metric.JaroWinkler(s1, s2, 0.1)

	fmt.Printf("Jaro-Winkler similarity between '%s' and '%s': %.4f\n", s1, s2, similarity)

}
