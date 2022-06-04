package pkg

import (
	"fmt"
	"testing"
)

func BenchmarkSumSeries(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SumSeries(1, size)
			}
		})
	}
}

func TestSumSeries(t *testing.T) {
	if result := SumSeries(1, 3); result != 6 {
		t.Errorf("SumSeries(1, 3) is %d, not 6", result)
	}
}
