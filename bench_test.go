package urlp

import (
	"testing"
)

func BenchmarkMatcherRoot(b *testing.B) {
	p := NewPattern("/")
	u := ""

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		p.Match(u)
	}
}

func BenchmarkMatcherExact(b *testing.B) {
	p := NewPattern("/posts/comments/new")
	u := "/posts/comments/new"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		p.Match(u)
	}
}

func BenchmarkMatcherExactWithTrailingSlash(b *testing.B) {
	p := NewPattern("/posts/comments/new")
	u := "/posts/comments/new/"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		p.Match(u)
	}
}

func BenchmarkMatcher1Param(b *testing.B) {
	p := NewPattern("/posts/comments/:id")
	u := "/posts/comments/123"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		p.Match(u)
	}
}

func BenchmarkMatcher2Params(b *testing.B) {
	p := NewPattern("/posts/:post_id/comments/:id")
	u := "/posts/123/comments/456"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		p.Match(u)
	}
}

func BenchmarkMatcher2ParamsWithNewPattern(b *testing.B) {
	u := "/posts/123/comments/456"

	b.ReportAllocs()
	b.ResetTimer()

	i := 0
	for ; i < b.N; i++ {
		NewPattern("/posts/:post_id/comments/:id").Match(u)
	}
}

// BenchmarkMatcherRoot                     200000000                8.38 ns/op            0 B/op          0 allocs/op
// BenchmarkMatcherExact                    200000000                8.11 ns/op            0 B/op          0 allocs/op
// BenchmarkMatcherExactWithTrailingSlash   100000000                11.8 ns/op            0 B/op          0 allocs/op
// BenchmarkMatcher1Param                    10000000                 199 ns/op           32 B/op          1 allocs/op
// BenchmarkMatcher2Params                    3000000                 416 ns/op           96 B/op          2 allocs/op
// BenchmarkMatcher2ParamsWithNewPattern      1000000                1183 ns/op          272 B/op          6 allocs/op
// ok      github.com/nowk/urlp    11.291s
