package urlp

import (
	"testing"
)

func BenchmarkMatcher(b *testing.B) {
	p := "/posts/comments/:id"
	u := "/posts/comments/new"

	i := 0
	for ; i < b.N; i++ {
		m := NewMatcher(p)
		m.Match(u)
	}
}

func BenchmarkCacheMatcher(b *testing.B) {
	p := "/posts/comments/:id"
	u := "/posts/comments/new"
	m := NewMatcher(p)

	i := 0
	for ; i < b.N; i++ {
		m.Match(u)
	}
}

// BenchmarkMatcher         5000000               412 ns/op
// BenchmarkCacheMatcher    5000000               316 ns/op
