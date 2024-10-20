package metric_test

import (
	"fmt"
	"testing"

	"github.com/elysiamae/sisyphus/metric"
)

func TestLevenshtein(t *testing.T) {
	a := "abcd"
	b := "abdc"
	distance := metric.DamerauLevenshteinDistance(a, b)
	fmt.Printf("Damerau-Levenshtein 距离: %d\n", distance)
}
