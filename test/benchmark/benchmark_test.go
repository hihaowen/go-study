package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Println("hi")
	}

	b.StopTimer()
}
