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

func BenchmarkMatcherWithFormat(b *testing.B) {
	p := "/posts/comments/:id.:format"
	u := "/posts/comments/new"

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

// BenchmarkMatcher                 5000000               337 ns/op
// BenchmarkMatcherWithFormat       5000000               346 ns/op
