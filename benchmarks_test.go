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

// // BenchmarkMatcher                 5000000               337 ns/op
// // BenchmarkMatcherWithFormat       5000000               574 ns/op
// // BenchmarkMatcherWithFormatNoExt  5000000               345 ns/op
