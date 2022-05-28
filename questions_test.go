package leetcodepractice

import (
	"testing"
)

func assert(t *testing.T, expected interface{}, output interface{}) {
	if output != expected {
		t.Errorf("expected %v but got %v", expected, output)
	}
}

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

func TestCalculateNumberOfStepToReduceANumberToZero(t *testing.T) {
	out := calculateNumberOfStepToReduceANumberToZero(14)
	expected := 6
	assert(t, expected, out)
	out = calculateNumberOfStepToReduceANumberToZero(8)
	expected = 4
	assert(t, expected, out)
	out = calculateNumberOfStepToReduceANumberToZero(123)
	expected = 12
	assert(t, expected, out)
}

func TestGetMissingNumber(t *testing.T) {
	out := getMissingNumber([]int{3, 0, 1})
	expected := 2
	assert(t, expected, out)
	out = getMissingNumber([]int{0, 1})
	expected = 2
	assert(t, expected, out)
	out = getMissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	expected = 8
	assert(t, expected, out)
}
