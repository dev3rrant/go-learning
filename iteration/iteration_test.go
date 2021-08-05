package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Repeat specified number times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q, received %q", expected, repeated)
		}
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleIteration(b *testing.B) {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// "aaaaa"
}
