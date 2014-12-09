package urlp

import (
	"github.com/nowk/assert"
	"testing"
)

func TestExactMatches(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/posts", "/posts"},
		{"/posts", "posts/"},
		{"/posts", "/posts/"},
		{"posts", "/posts"},
		{"posts", "posts/"},
		{"posts", "/posts/"},
		{"/posts/new", "/posts/new"},
	} {
		m := NewMatcher(v.p)
		p, ok := m.Match(v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, p)
	}
}

func TestWithParams(t *testing.T) {
	p := "/posts/:post_id/comments/:id"
	u := "/posts/123/comments/456"

	m := NewMatcher(p)
	v, ok := m.Match(u)
	assert.True(t, ok)
	assert.Equal(t, "123", v["post_id"])
	assert.Equal(t, "456", v["id"])
}

func TestRoot(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/", "/"},
		{"/", ""},
		{"", "/"},
	} {
		m := NewMatcher(v.p)
		p, ok := m.Match(v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, p)
	}
}

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

// BenchmarkMatcher         1000000              1575 ns/op
// BenchmarkCacheMatcher    1000000              1145 ns/op
