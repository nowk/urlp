package urlp

import (
	"testing"
)

func BenchmarkMatcherExact(b *testing.B) {
	p := NewPath("/posts/comments/new")
	u := "/posts/comments/new"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

func BenchmarkMatcherExactWithTrailingSlash(b *testing.B) {
	p := NewPath("/posts/comments/new")
	u := "/posts/comments/new/"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

func BenchmarkMatcher1Param(b *testing.B) {
	p := NewPath("/posts/comments/:id")
	u := "/posts/comments/123"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

func BenchmarkMatcher2Params(b *testing.B) {
	p := NewPath("/posts/:post_id/comments/:id")
	u := "/posts/123/comments/456"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

// func BenchmarkMatcherWithFormat(b *testing.B) {
// 	p := "/posts/comments/:id.:format"
// 	u := "/posts/comments/new.html"

// 	i := 0
// 	for ; i < b.N; i++ {
// 		Match(p, u)
// 	}
// }

// func BenchmarkMatcherWithFormatNoExt(b *testing.B) {
// 	p := "/posts/comments/:id.:format"
// 	u := "/posts/comments/new"

// 	i := 0
// 	for ; i < b.N; i++ {
// 		Match(p, u)
// 	}
// }

// urlp.v2
// BenchmarkMatcherExact                   200000000               8.89 ns/op               0 B/op          0 allocs/op
// BenchmarkMatcherExactWithTrailingSlash  100000000               12.7 ns/op               0 B/op          0 allocs/op
// BenchmarkMatcher1Param                   10000000                193 ns/op              32 B/op          1 allocs/op
// BenchmarkMatcher2Params                   5000000                292 ns/op              64 B/op          1 allocs/op
