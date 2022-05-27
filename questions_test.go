package leetcodepractice

import (
	"testing"
)

func BenchmarkLeetCodeFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueriesOnPointsInsideACircle()
	}
}

func FuzzQueriesOnPointsInsideCircle(f *testing.F) {
	f.Add(1, 2, 3, 4, 5)
	f.Fuzz(func(t *testing.T, a, b, c, d, e int) {
		point := [][]int{{a, b}}
		query := [][]int{{c, d, e}}
		out := countPoints(point, query)
		t.Log(out)
	})
}

func TestComplementOfDecimal(t *testing.T) {
	out := complimentDecimal(5)
	if out != 2 {
		t.Errorf("expected 2 but got %v", out)
	}
	out = complimentDecimal(-5)
	if out != 4 {
		t.Errorf("expected 4 but got %v", out)
	}
}
