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
		m, ok := Match(v.p, v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, m)
	}
}

func TestWithParams(t *testing.T) {
	p := "/posts/:post_id/comments/:id"
	u := "/posts/123/comments/456"

	m, ok := Match(p, u)
	assert.True(t, ok)
	assert.Equal(t, "123", m.Get("post_id"))
	assert.Equal(t, "456", m.Get("id"))
}

func TestRoot(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/", "/"},
		{"/", ""},
		{"", "/"},
	} {
		m, ok := Match(v.p, v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, m)
	}
}

func BenchmarkJoinVsFor(b *testing.B) {
	p := "/posts/comments/new"
	u := "/posts/comments/new"

	i := 0
	for ; i < b.N; i++ {
		Match(p, u)
	}
}

// join and compares for exact matches
// BenchmarkJointsVsFor     1000000              1365 ns/op
//
// for loop comparison
// BenchmarkJointsVsFor     2000000              1007 ns/op
