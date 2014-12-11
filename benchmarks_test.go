package urlp

import (
	"testing"
)

func BenchmarkMatcher(b *testing.B) {
	p := "/posts/comments/:id"
	u := "/posts/comments/new"

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

// BenchmarkMatcher         5000000               325 ns/op
